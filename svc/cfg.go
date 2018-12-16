// Package svc provides definitions for services that run on the center.
package svc

import (
	"github.com/kostiamol/centerms/proto"
	"golang.org/x/net/context"

	"math/rand"
	"time"

	"encoding/json"

	"github.com/Sirupsen/logrus"
	gproto "github.com/golang/protobuf/proto"
	consul "github.com/hashicorp/consul/api"
	"github.com/kostiamol/centerms/entity"
	"github.com/nats-io/go-nats"
	"github.com/satori/go.uuid"
)

const (
	aggregate = "cfg_svc"
	event     = "cfg_patched"
)

// Cfg is used to deal with device configurations.
type Cfg struct {
	addr      entity.Addr
	store     entity.Storer
	ctrl      Ctrl
	log       *logrus.Entry
	retry     time.Duration
	sub       entity.Subscription
	agent     *consul.Agent
	agentName string
	ttl       time.Duration
}

// NewCfgService creates and initializes a new instance of Cfg service.
func NewCfgService(a entity.Addr, s entity.Storer, c Ctrl, l *logrus.Entry, retry time.Duration, subj string,
	agentName string, ttl time.Duration) *Cfg {

	return &Cfg{
		addr:  a,
		store: s,
		ctrl:  c,
		log:   l.WithFields(logrus.Fields{"component": "svc", "name": "cfg"}),
		retry: retry,
		sub: entity.Subscription{
			ChanName: subj,
			Chan:     make(chan []byte),
		},
		agentName: agentName,
		ttl:       ttl,
	}
}

// Run launches the service by running goroutines for listening the service termination and config patches.
func (c *Cfg) Run() {
	c.log.WithFields(logrus.Fields{
		"func":  "Run",
		"event": entity.EventSVCStarted,
	}).Infof("running on host: [%s], port: [%d]", c.addr.Host, c.addr.Port)

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		if r := recover(); r != nil {
			c.log.WithFields(logrus.Fields{
				"func":  "Run",
				"event": entity.EventPanic,
			}).Errorf("%s", r)

			cancel()
			c.terminate()
		}
	}()
	go c.listenTermination()
	go c.listenCfgPatches(ctx)
	c.runConsulAgent()
}

// GetAddr returns address of the service.
func (c *Cfg) GetAddr() entity.Addr {
	return c.addr
}

func (c *Cfg) runConsulAgent() {
	cl, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		c.log.WithFields(logrus.Fields{
			"func":  "Run",
			"event": entity.EventPanic,
		}).Errorf("%s", err)
		panic("consul init error")
	}
	a := &consul.AgentServiceRegistration{
		Name: c.agentName,
		Port: c.addr.Port,
		Check: &consul.AgentServiceCheck{
			TTL: c.ttl.String(),
		},
	}
	c.agent = cl.Agent()
	if err := c.agent.ServiceRegister(a); err != nil {
		c.log.WithFields(logrus.Fields{
			"func":  "Run",
			"event": entity.EventPanic,
		}).Errorf("%s", err)
		panic("consul init error")
	}
	go c.updateTTL(c.check)
}

// todo: substitute bool with byte
func (c *Cfg) check() (bool, error) {
	// while the service is alive - everything is ok
	return true, nil
}

func (c *Cfg) updateTTL(check func() (bool, error)) {
	t := time.NewTicker(c.ttl / 2)
	for range t.C {
		c.update(check)
	}
}

func (c *Cfg) update(check func() (bool, error)) {
	var health string
	ok, err := check()
	if !ok {
		c.log.WithFields(logrus.Fields{
			"func":  "update",
			"event": entity.EventUpdConsulStatus,
		}).Errorf("check has failed: %s", err)

		// failed check will remove a service instance from DNS and HTTP query
		// to avoid returning errors or invalid data.
		health = consul.HealthCritical
	} else {
		health = consul.HealthPassing
	}

	if err := c.agent.UpdateTTL("svc:"+c.agentName, "", health); err != nil {
		c.log.WithFields(logrus.Fields{
			"func":  "update",
			"event": entity.EventUpdConsulStatus,
		}).Error(err)
	}
}

func (c *Cfg) listenTermination() {
	for {
		select {
		case <-c.ctrl.StopChan:
			c.terminate()
			return
		}
	}
}

func (c *Cfg) terminate() {
	defer func() {
		if r := recover(); r != nil {
			c.log.WithFields(logrus.Fields{
				"func":  "terminate",
				"event": entity.EventPanic,
			}).Errorf("%s", r)
			c.ctrl.Terminate()
		}
	}()

	if err := c.store.CloseConn(); err != nil {
		c.log.WithFields(logrus.Fields{
			"func": "terminate",
		}).Errorf("%s", err)
	}

	c.log.WithFields(logrus.Fields{
		"func":  "terminate",
		"event": entity.EventSVCShutdown,
	}).Infoln("svc is down")
	c.ctrl.Terminate()
}

// SetDevInitCfg check's whether device is already registered in the system. If it's already registered,
// the func returns actual configuration. Otherwise it returns default config for that type of device.
func (c *Cfg) SetDevInitCfg(m *entity.DevMeta) (*entity.DevCfg, error) {
	conn, err := c.store.CreateConn()
	if err != nil {
		c.log.WithFields(logrus.Fields{
			"func": "SetDevInitCfg",
		}).Errorf("%s", err)
		return nil, err
	}
	defer conn.CloseConn()

	if err = conn.SetDevMeta(m); err != nil {
		c.log.WithFields(logrus.Fields{
			"func": "SetDevInitCfg",
		}).Errorf("%s", err)
		return nil, err
	}

	var d *entity.DevCfg
	id := entity.DevID(m.MAC)

	if ok, err := conn.DevIsRegistered(m); ok {
		if err != nil {
			c.log.WithFields(logrus.Fields{
				"func": "SetDevInitCfg",
			}).Errorf("%s", err)
			return nil, err
		}

		d, err = conn.GetDevCfg(id)
		if err != nil {
			c.log.WithFields(logrus.Fields{
				"func": "SetDevInitCfg",
			}).Errorf("%s", err)
			return nil, err
		}
	} else {
		if err != nil {
			c.log.WithFields(logrus.Fields{
				"func": "SetDevInitCfg",
			}).Errorf("%s", err)
			return nil, err
		}

		d, err = conn.GetDevDefaultCfg(m)
		if err != nil {
			c.log.WithFields(logrus.Fields{
				"func": "SetDevInitCfg",
			}).Errorf("%s", err)
			return nil, err
		}

		if err = conn.SetDevCfg(id, d); err != nil {
			c.log.WithFields(logrus.Fields{
				"func": "SetDevInitCfg",
			}).Errorf("%s", err)
			return nil, err
		}
		c.log.WithFields(logrus.Fields{
			"func":  "SetDevInitCfg",
			"event": entity.EventDevRegistered,
		}).Infof("devices' meta: %+v", m)
	}
	return d, err
}

func (c *Cfg) listenCfgPatches(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.log.WithFields(logrus.Fields{
				"func":  "listenCfgPatches",
				"event": entity.EventPanic,
			}).Errorf("%s", r)
			c.terminate()
		}
	}()

	conn, err := c.store.CreateConn()
	if err != nil {
		c.log.WithFields(logrus.Fields{
			"func": "listenCfgPatches",
		}).Errorf("%s", err)
		return
	}
	defer conn.CloseConn()

	go conn.Subscribe(c.sub.Chan, c.sub.ChanName)

	var d entity.DevCfg
	for {
		select {
		case msg := <-c.sub.Chan:
			if err := json.Unmarshal(msg, &d); err != nil {
				c.log.WithFields(logrus.Fields{
					"func": "listenCfgPatches",
				}).Errorf("%s", err)
			} else {
				go c.pubNewCfgPatchEvent(&d)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (c *Cfg) pubNewCfgPatchEvent(cfg *entity.DevCfg) {
	defer func() {
		if r := recover(); r != nil {
			c.log.WithFields(logrus.Fields{
				"func":  "pubNewCfgPatchEvent",
				"event": entity.EventPanic,
			}).Errorf("%s", r)
			c.terminate()
		}
	}()

	conn, err := nats.Connect(nats.DefaultURL)
	for err != nil {
		c.log.WithFields(logrus.Fields{
			"func": "pubNewCfgPatchEvent",
		}).Error("nats connectivity status: DISCONNECTED")
		duration := time.Duration(rand.Intn(int(c.retry.Seconds())))
		time.Sleep(time.Second*duration + 1)
		conn, err = nats.Connect(nats.DefaultURL)
	}
	defer conn.Close()

	s := proto.EventStore{
		AggregateId:   cfg.MAC,
		AggregateType: aggregate,
		EventId:       uuid.NewV4().String(),
		EventType:     event,
		EventData:     string(cfg.Data),
	}
	subj := "Cfg.Patch." + cfg.MAC
	b, err := gproto.Marshal(&s)
	if err != nil {
		c.log.WithFields(logrus.Fields{
			"func": "pubNewCfgPatchEvent",
		}).Errorf("marshal has failed: %s", err)
	}

	if err := conn.Publish(subj, b); err != nil {

	}
	c.log.WithFields(logrus.Fields{
		"func":  "pubNewCfgPatchEvent",
		"event": entity.EventCfgPatchCreated,
	}).Infof("cfg patch [%s] for device with ID [%s]", cfg.Data, cfg.MAC)
}

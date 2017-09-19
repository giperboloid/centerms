package servers

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"os"

	"context"

	"sync"

	"github.com/Sirupsen/logrus"
	"github.com/giperboloid/centerms/entities"
)

type ConnPool struct {
	sync.Mutex
	conn map[string]net.Conn
}

func (p *ConnPool) Init() {
	p.Lock()
	defer p.Unlock()
	p.conn = make(map[string]net.Conn)
}

func (p *ConnPool) AddConn(conn net.Conn, key string) {
	p.Lock()
	p.conn[key] = conn
	defer p.Unlock()
}

func (p *ConnPool) GetConn(key string) net.Conn {
	p.Lock()
	defer p.Unlock()
	return p.conn[key]
}

func (p *ConnPool) RemoveConn(key string) {
	p.Lock()
	defer p.Unlock()
	delete(p.conn, key)
}

type DevConfigServer struct {
	Server     entities.Server
	DevStorage entities.DevStorage
	Controller entities.ServersController
	Log        *logrus.Logger
	Reconnect  *time.Ticker
	ConnPool   ConnPool
	Messages   chan []string
}

func NewDevConfigServer(s entities.Server, ds entities.DevStorage, c entities.ServersController, l *logrus.Logger,
	r *time.Ticker, msgs chan []string) *DevConfigServer {
	l.Out = os.Stdout

	return &DevConfigServer{
		Server:     s,
		DevStorage: ds,
		Controller: c,
		Log:        l,
		Reconnect:  r,
		Messages:   msgs,
	}
}

func (s *DevConfigServer) Run() {
	s.Log.Infof("DevConfigServer has started on host: %s, port: %d", s.Server.Host, s.Server.Port)
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		if r := recover(); r != nil {
			s.Log.Errorf("DevConfigServer: Run(): panic: %s", r)
			cancel()
			s.gracefulHalt()
		}
	}()

	go s.handleTermination()

	s.ConnPool.Init()
	go s.configSubscribe(ctx, entities.DevConfigChan, s.Messages)

	go s.listenForConnections(ctx)
}

func (s *DevConfigServer) handleTermination() {
	for {
		select {
		case <-s.Controller.StopChan:
			s.gracefulHalt()
			return
		}
	}
}

func (s *DevConfigServer) gracefulHalt() {
	s.DevStorage.CloseConn()
	s.Log.Infoln("DevConfigServer has shut down")
	s.Controller.Terminate()
}

func (s *DevConfigServer) listenForConnections(ctx context.Context) {
	ln, err := net.Listen("tcp", s.Server.Host+":"+fmt.Sprint(s.Server.Port))
	for err != nil {
		for range s.Reconnect.C {
			ln, err = net.Listen("tcp", s.Server.Host+":"+fmt.Sprint(s.Server.Port))
			if err != nil {
				s.Log.Errorf("DevConfigServer: Run(): Listen() has failed: %s", err)
			}
		}
		s.Reconnect.Stop()

		select {
		case <-ctx.Done():
			return
		}
	}

	for {
		cn, err := ln.Accept()
		if err != nil {
			s.Log.Errorf("DevConfigServer: Run(): Accept() has failed: %s", err)
		}
		go s.sendDefaultConfig(ctx, cn, &s.ConnPool)

		select {
		case <-ctx.Done():
			return
		}
	}
}

func (s *DevConfigServer) sendDefaultConfig(ctx context.Context, c net.Conn, cp *ConnPool) {
	var req entities.Request
	if err := json.NewDecoder(c).Decode(&req); err != nil {
		s.Log.Errorf("DevConfigServer: sendDefaultConfig(): Request marshalling has failed: %s", err)
		return
	}
	cp.AddConn(c, req.Meta.MAC)

	cn, err := s.DevStorage.CreateConn()
	if err != nil {
		s.Log.Errorf("DevConfigServer: sendDefaultConfig(): storage connection hasn't been established: %s", err)
		return
	}
	defer cn.CloseConn()

	var dc *entities.DevConfig
	if ok, err := cn.DevIsRegistered(&req.Meta); ok {
		if err != nil {
			s.Log.Errorf("DevConfigServer: sendDefaultConfig(): DevIsRegistered() has failed: %s", err)
			return
		}

		dc, err = cn.GetDevConfig(&req.Meta)
		if err != nil {
			s.Log.Errorf("DevConfigServer: sendDefaultConfig(): GetDevConfig() has failed: %s", err)
			return
		}
	} else {
		dc, err = cn.GetDevDefaultConfig(&req.Meta)
		if err != nil {
			s.Log.Errorf("DevConfigServer: sendDefaultConfig(): GetDevDefaultConfig() has failed: %s", err)
			return
		}

		if err = cn.SetDevConfig(&req.Meta, dc); err != nil {
			s.Log.Errorf("DevConfigServer: sendDefaultConfig(): SetDevConfig() has failed: %s", err)
			return
		}
	}

	if _, err = c.Write(dc.Data); err != nil {
		s.Log.Errorf("DevConfigServer: sendDefaultConfig(): DevConfig.Data writing has failed: %s", err)
		return
	}

	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}

func (s *DevConfigServer) configSubscribe(ctx context.Context, channel string, msg chan []string) {
	cn, err := s.DevStorage.CreateConn()
	if err != nil {
		s.Log.Errorf("DevConfigServer: configSubscribe(): storage connection hasn't been established: ", err)
		return
	}
	defer cn.CloseConn()

	cn.Subscribe(msg, channel)
	for {
		var dc entities.DevConfig
		select {
		case msg := <-msg:
			if msg[0] == "message" {
				if err := json.Unmarshal([]byte(msg[2]), &dc); err != nil {
					s.Log.Errorf("DevConfigServer: configSubscribe(): DevConfig unmarshalling has failed: ", err)
					return
				}
				go s.sendNewConfig(ctx, &dc)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (s *DevConfigServer) sendNewConfig(ctx context.Context, c *entities.DevConfig) {
	cn := s.ConnPool.GetConn(c.MAC)
	if cn == nil {
		s.Log.Error("DevConfigServer: sendNewConfig(): there isn't such a connection in pool")
		return
	}

	if _, err := cn.Write(c.Data); err != nil {
		s.Log.Errorf("DevConfigServer: sendNewConfig(): DevConfig.Data writing has failed: %s", err)
		s.ConnPool.RemoveConn(c.MAC)
		return
	}

	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}

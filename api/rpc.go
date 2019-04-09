package api

import (
	"github.com/kostiamol/centerms/svc"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"

	"github.com/kostiamol/centerms/cfg"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/kostiamol/centerms/proto"

	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func (a *API) runRPCServer() {
	defer func() {
		if r := recover(); r != nil {
			a.log.WithFields(logrus.Fields{
				"func":  "listenCfg",
				"event": cfg.EventPanic,
			}).Errorf("%s", r)
		}
	}()
	l, err := net.Listen("tcp", ":"+fmt.Sprint(a.portRPC))
	if err != nil {
		a.log.WithFields(logrus.Fields{
			"func": "runRPCServer",
		}).Fatalf("Listen() failed: %s", err)
	}

	grpc_logrus.ReplaceGrpcLogger(a.log)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_logrus.UnaryServerInterceptor(a.log),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	proto.RegisterCenterServiceServer(s, a)
	if err := s.Serve(l); err != nil {
		a.log.WithFields(logrus.Fields{
			"func": "runRPCServer",
		}).Fatalf("Serve() failed: %s", err)
	}
}

// SetDevInitCfg sets device's initial configuration when it connects to the center for the first time using CfgService
// and returns that configuration to the device.
func (a *API) SetDevInitCfg(ctx context.Context, r *proto.SetDevInitCfgRequest) (*proto.SetDevInitCfgResponse, error) {
	m := svc.DevMeta{
		Type: r.Meta.Type,
		Name: r.Meta.Name,
		MAC:  r.Meta.Mac,
	}
	c, err := a.cfgProvider.SetDevInitCfg(&m)
	if err != nil {
		a.log.WithFields(logrus.Fields{
			"func": "SetDevInitCfg",
		}).Errorf("SetDevInitCfg() failed: %s", err)
	}
	return &proto.SetDevInitCfgResponse{
		Cfg: c.Data,
	}, nil
}

// SaveDevData saves data from device using DataService.
func (a *API) SaveDevData(ctx context.Context, r *proto.SaveDevDataRequest) (*proto.SaveDevDataResponse, error) {
	d := svc.DevData{
		Time: r.Time,
		Meta: svc.DevMeta{
			Type: r.Meta.Type,
			Name: r.Meta.Name,
			MAC:  r.Meta.Mac,
		},
		Data: r.Data,
	}
	if err := a.dataProvider.SaveDevData(&d); err != nil {
		a.log.WithFields(logrus.Fields{
			"func": "SaveDevData",
		}).Errorf("SaveDevData() failed: %s", err)
	}
	return &proto.SaveDevDataResponse{
		Status: "OK",
	}, nil
}

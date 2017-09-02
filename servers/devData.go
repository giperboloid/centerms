package servers

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/Sirupsen/logrus"

	"github.com/giperboloid/centerms/entities"
	"github.com/giperboloid/centerms/storages/redis"
	"github.com/pkg/errors"
)

type DevDataServer struct {
	Server     entities.Server
	Storage    entities.Storage
	Controller entities.RoutinesController
	Log        *logrus.Logger
	Reconnect  *time.Ticker
}

func NewDevDataServer(s entities.Server, st entities.Storage, c entities.RoutinesController,
	l *logrus.Logger, r *time.Ticker) *DevDataServer {
	return &DevDataServer{
		Server:     s,
		Storage:    st,
		Controller: c,
		Log:        l,
		Reconnect:  r,
	}
}

func (s *DevDataServer) Run() {
	defer func() {
		if r := recover(); r != nil {
			errors.New("DevConfigServer: Run(): panic leads to halt")
			s.gracefulHalt()
			s.Controller.Close()
		}
	}()

	ln, err := net.Listen("tcp", s.Server.Host+":"+fmt.Sprint(s.Server.Port))
	if err != nil {
		errors.Wrap(err, "DevConfigServer: Run(): Listen() has failed")
	}

	for err != nil {
		for range s.Reconnect.C {
			ln, err = net.Listen("tcp", s.Server.Host+":"+fmt.Sprint(s.Server.Port))
			if err != nil {
				errors.Wrap(err, "DevConfigServer: Run(): Listen() has failed")
			}
		}
		s.Reconnect.Stop()
	}

	for {
		conn, err := ln.Accept()
		if err == nil {
			go s.devDataHandler(conn)
		}
	}
}

func (s *DevDataServer) gracefulHalt() {
	s.Storage.CloseConnection()
}

func (s *DevDataServer) devDataHandler(conn net.Conn) {
	var req entities.Request
	var resp entities.Response
	for {
		err := json.NewDecoder(conn).Decode(&req)
		if err != nil {
			errors.Wrap(err, "DevConfigServer: devDataHandler(): Request decoding has failed")
			return
		}

		go s.devTypeHandler(&req)

		resp = entities.Response{
			Status: 200,
			Descr:  "OK",
		}
		err = json.NewEncoder(conn).Encode(&resp)
		if err != nil {
			errors.Wrap(err, "DevConfigServer: devDataHandler(): Response encoding has failed")
		}
	}
}

func (s *DevDataServer) devTypeHandler(r *entities.Request) {
	conn, err := s.Storage.CreateConnection()
	if err != nil {
		errors.Wrap(err, "DevConfigServer: devTypeHandler(): storage connection hasn't been established")
	}
	defer conn.CloseConnection()

	switch r.Action {
	case "update":
		conn.SetDevData(r)
		go storages.PublishWS(r, "devWS", conn)

	default:
		errors.Wrap(err, "DevConfigServer: devTypeHandler(): device Request - unknown action")
	}
}

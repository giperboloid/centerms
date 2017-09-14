package servers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"os"

	"github.com/giperboloid/centerms/entities"
	"github.com/pkg/errors"
)

type WebServer struct {
	Server     entities.Server
	DevStore   entities.DevStorage
	Controller entities.RoutinesController
	Log        *logrus.Logger
}

func NewWebServer(s entities.Server, ds entities.DevStorage, c entities.RoutinesController, l *logrus.Logger) *WebServer {
	l.Out = os.Stdout
	return &WebServer{
		Server:     s,
		DevStore:   ds,
		Controller: c,
		Log:        l,
	}
}

func (s *WebServer) Run() {
	defer func() {
		if r := recover(); r != nil {
			errors.New("WebServer: Run(): panic leads to halt")
			s.gracefulHalt()
			s.Controller.Close()
		}
	}()

	r := mux.NewRouter()
	r.HandleFunc("/devices", s.getDevsDataHandler).Methods(http.MethodGet)
	r.HandleFunc("/devices/{id}/data", s.getDevDataHandler).Methods(http.MethodGet)
	r.HandleFunc("/devices/{id}/config", s.getDevConfigHandler).Methods(http.MethodGet)
	r.HandleFunc("/devices/{id}/config", s.patchDevConfigHandler).Methods(http.MethodPatch)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./view/")))

	port := fmt.Sprint(s.Server.Port)
	srv := &http.Server{
		Handler:      r,
		Addr:         s.Server.Host + ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	http.ListenAndServe(s.Server.Host+":"+port, handlers.CORS()(r))

	go s.Log.Fatal(srv.ListenAndServe())
}

func (s *WebServer) gracefulHalt() {
	s.DevStore.CloseConn()
}

func (s *WebServer) getDevsDataHandler(w http.ResponseWriter, r *http.Request) {
	cn, err := s.DevStore.CreateConn()
	if err != nil {
		errors.Wrap(err, "WebServer: getDevicesHandler(): storage connection hasn't been established")
		return
	}
	defer cn.CloseConn()

	ds, err := cn.GetDevsData()
	if err != nil {
		errors.Wrap(err, "WebServer: getDevicesHandler(): devices info extraction has failed")
	}

	err = json.NewEncoder(w).Encode(ds)
	if err != nil {
		errors.Wrap(err, "WebServer: getDevicesHandler(): []DevData encoding has failed")
	}
}

func (s *WebServer) getDevDataHandler(w http.ResponseWriter, r *http.Request) {
	cn, err := s.DevStore.CreateConn()
	if err != nil {
		errors.Wrap(err, "WebServer: getDevDataHandler(): storage connection hasn't been established")
		return
	}
	defer cn.CloseConn()

	m := entities.DevMeta{
		Type: r.FormValue("type"),
		Name: r.FormValue("name"),
		MAC:  r.FormValue("mac"),
	}

	d, err := cn.GetDevData(&m)
	if err != nil {
		errors.Wrap(err, "WebServer: getDevDataHandler(): DevData extraction has failed")
	}

	err = json.NewEncoder(w).Encode(d)
	if err != nil {
		errors.Wrap(err, "WebServer: getDevDataHandler(): DevData encoding has failed")
	}
}

func (s *WebServer) getDevConfigHandler(w http.ResponseWriter, r *http.Request) {
	cn, err := s.DevStore.CreateConn()
	if err != nil {
		errors.Wrap(err, "WebServer: getDevConfigHandler(): storage connection hasn't been established")
		return
	}
	defer cn.CloseConn()

	m := entities.DevMeta{
		Type: r.FormValue("type"),
		Name: r.FormValue("name"),
		MAC:  r.FormValue("mac"),
	}

	c, err := cn.GetDevConfig(&m)
	if err != nil {
		errors.Wrap(err, "WebServer: getDevConfigHandler(): DevConfig extraction has failed")
	}
	w.Write(c.Data)
}

func (s *WebServer) patchDevConfigHandler(w http.ResponseWriter, r *http.Request) {
	cn, err := s.DevStore.CreateConn()
	if err != nil {
		errors.Wrap(err, "WebServer: patchDevConfigHandler(): storage connection hasn't been established")
		return
	}
	defer cn.CloseConn()

	var c entities.DevConfig
	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		errors.Wrap(err, "WebServer: patchDevConfigHandler(): DevConfig decoding has failed")
	}

	m := entities.DevMeta{
		Type: r.FormValue("type"),
		Name: r.FormValue("name"),
		MAC:  r.FormValue("mac"),
	}

	err = cn.SetDevConfig(&m, &c)
	if err != nil {
		errors.Wrap(err, "WebServer: patchDevConfigHandler(): DevConfig setting has failed")
	}

	jc, err := json.Marshal(c)
	if err != nil {
		errors.Wrap(err, "WebServer: patchDevConfigHandler(): DevConfig marshalling has failed")
	}
	cn.Publish("configChan", jc)
}

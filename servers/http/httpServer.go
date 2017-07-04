package http

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"github.com/gorilla/handlers"
	log "github.com/Sirupsen/logrus"
	"encoding/json"
	"strings"

	. "github.com/KharkivGophers/center-smart-house/models"
	"github.com/KharkivGophers/center-smart-house/drivers"
	. "github.com/KharkivGophers/center-smart-house/dao"
	"strconv"
)

type HTTPServer struct{
	LocalServer Server
	DbServer Server
}

func NewHTTPServer (local , db Server) *HTTPServer{
	return &HTTPServer{LocalServer:local, DbServer: db}
}


func (server *HTTPServer)RunDynamicServer() {
	r := mux.NewRouter()
	r.HandleFunc("/devices", server.getDevicesHandler).Methods("GET")
	r.HandleFunc("/devices/{id}/data", server.getDevDataHandler).Methods("GET")
	r.HandleFunc("/devices/{id}/config", server.getDevConfigHandler).Methods("GET")

	r.HandleFunc("/devices/{id}/config", server.patchDevConfigHandler).Methods("PATCH")

	//provide static html pages
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("../view/")))

	var port string =strconv.FormatUint(uint64(server.LocalServer.Port),10)

	srv := &http.Server{
		Handler:      r,
		Addr:         server.LocalServer.IP + ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	//CORS provides Cross-Origin Resource Sharing middleware
	http.ListenAndServe( server.LocalServer.IP + ":" + port, handlers.CORS()(r))

	go log.Fatal(srv.ListenAndServe())
}

//----------------------http Dynamic Connection----------------------------------------------------------------------------------

func (server *HTTPServer)getDevicesHandler(w http.ResponseWriter, r *http.Request) {

	dbClient, err := GetDBConnection(server.DbServer)
	defer dbClient.Close()
	if CheckError("HTTP Dynamic Connection: getDevicesHandler", err)!= nil{
		return
	}

	devices := dbClient.GetAllDevices()
	err = json.NewEncoder(w).Encode(devices)
	CheckError("getDevicesHandler JSON enc", err)
}

func (server *HTTPServer) getDevDataHandler(w http.ResponseWriter, r *http.Request) {

	dbClient, err := GetDBConnection(server.DbServer)
	defer dbClient.Close()
	if CheckError("HTTP Dynamic Connection: getDevDataHandler", err)!= nil{
		return
	}

	vars := mux.Vars(r)
	devID := "device:" + vars["id"]

	devParamsKeysTokens := []string{}
	devParamsKeysTokens = strings.Split(devID, ":")
	devParamsKey := devID + ":" + "params"

	device := dbClient.GetDevice(devParamsKey, devParamsKeysTokens)
	err = json.NewEncoder(w).Encode(device)
	CheckError("getDevDataHandler JSON enc", err)
}

func (server *HTTPServer) getDevConfigHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"] // type + mac
	macAndType := strings.Split(id, ":")

	dbClient, err := GetDBConnection(server.DbServer)
	defer dbClient.Close()
	if CheckError("HTTP Dynamic Connection: getDevConfigHandler", err)!= nil{
		return
	}

	configInfo := macAndType[2] + ":" + "config" // key

	var device drivers.ConfigDevDriver = *drivers.IdentDevString(macAndType[0])
	log.Info(device)
	if device==nil{
		http.Error(w, "This type is not found", 400)
		return
	}
	config := device.GetDevConfig(configInfo, macAndType[2], dbClient.GetClient())
	json.NewEncoder(w).Encode(config)
}

func (server *HTTPServer) patchDevConfigHandler(w http.ResponseWriter, r *http.Request) {

	var config *DevConfig
	vars := mux.Vars(r)
	id := vars["id"] // warning!! type : name : mac
	typeAndMac := strings.Split(id, ":")

	dbClient, err := GetDBConnection(server.DbServer)
	defer dbClient.Close()
	if CheckError("HTTP Dynamic Connection: patchDevConfigHandler", err)!= nil{
		return
	}

	configInfo := typeAndMac[2] + ":" + "config" // key

	var device drivers.ConfigDevDriver = *drivers.IdentDevString(typeAndMac[0])
	if device==nil{
		http.Error(w, "This type is not found", 400)
		return
	}
	config = device.GetDevConfig(configInfo, typeAndMac[2], dbClient.GetClient())


	// log.Warnln("Config Before", config)
	err = json.NewDecoder(r.Body).Decode(&config)

	if err != nil {
		http.Error(w, err.Error(), 400)
		log.Errorln("NewDec: ", err)
	}

	log.Info(config)
	CheckError("Encode error", err) // log.Warnln("Config After: ", config)


	if !ValidateMAC(config.MAC) {
		log.Error("Invalid MAC")
		http.Error(w, "Invalid MAC", 400)
	} else if !ValidateStreamOn(config.StreamOn) {
		log.Error("Invalid Stream Value")
		http.Error(w, "Invalid Stream Value", 400)
	} else if !ValidateTurnedOn(config.TurnedOn) {
		log.Error("Invalid Turned Value")
		http.Error(w, "Invalid Turned Value", 400)
	} else if !ValidateCollectFreq(config.CollectFreq) {
		log.Error("Invalid Collect Frequency Value")
		http.Error(w, "Collect Frequency should be more than 150!", 400)
	} else if !ValidateSendFreq(config.SendFreq) {
		log.Error("Invalid Send Frequency Value")
		http.Error(w, "Send Frequency should be more than 150!", 400)
	} else {
		// Save New Configuration to DB
		device.SetDevConfig(configInfo,config, dbClient.GetClient())
		log.Println("New Config was added to DB: ", config)
		JSONСonfig, _ := json.Marshal(config)
		dbClient.Publish("configChan",JSONСonfig)
	}
}
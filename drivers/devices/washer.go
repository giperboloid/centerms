package devices

import (
	. "github.com/KharkivGophers/center-smart-house/dao"
	. "github.com/KharkivGophers/center-smart-house/models"

)

type Washer struct {
	Data   WasherData
	Config DevConfig
	Meta   DevMeta
}

type WasherData struct {
	Mode   string
	Drying string
	Temp   map[int64]float32
}


func(washer *Washer)GetDevConfig(configInfo, mac string, worker DbRedisDriver) (*DevConfig){
	return nil
}
func(washer *Washer)SetDevConfig(configInfo string, config *DevConfig, worker DbRedisDriver){

}
func(washer *Washer)ValidateDevData(config DevConfig) (bool, string){
	return true,""
}


func(washer *Washer)GetDevData(devParamsKey string, devParamsKeysTokens []string, worker DbRedisDriver) DevData{
	return DevData{}
}
func(washer *Washer)SetDevData(req *Request, worker DbRedisDriver) *ServerError{
	return nil
}
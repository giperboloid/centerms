package cfg

import (
	"encoding/base64"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"
)

// Inner log events.
const (
	DevCfgChan  = "dev_cfg"
	DevDataChan = "dev_data"

	EventCfgPatchCreated = "cfg_patch_created"
	EventDevRegistered   = "dev_registered"
	EventMSTerminated    = "ms_terminated"
	EventPanic           = "panic"
	EventStoreInit       = "store_init"
	EventSVCStarted      = "svc_started"
	EventSVCShutdown     = "svc_shutdown"
	EventWSConnAdded     = "ws_conn_added"
	EventWSConnRemoved   = "ws_conn_removed"
)

type (
	// Config holds the app config.
	Config struct {
		Service Service
		Store   Store
		Token   Token
	}

	сonfiger interface {
		validate() error
	}
)

// NewConfig initializes and returns configuration structure with environment variables.
func NewConfig() (*Config, error) {
	publicKey, err := decodeEnv("PUBLIC_KEY")
	if err != nil {
		return nil, err
	}

	privateKey, err := decodeEnv("PRIVATE_KEY")
	if err != nil {
		return nil, err
	}

	c := &Config{
		Service: Service{
			AppID:         os.Getenv("APP_ID"),
			LogLevel:      os.Getenv("LOG_LEVEL"),
			RetryNumber:   uintEnv("RETRY_NUMBER"),
			RetryTimeout:  time.Duration(uintEnv("RETRY_TIMEOUT")),
			PortRPC:       uintEnv("PORT_RPC"),
			PortREST:      uintEnv("PORT_REST"),
			PortWebSocket: uintEnv("PORT_WEBSOCKET"),
		},
		Store: Store{
			Host:     os.Getenv("STORE_HOST"),
			Port:     uintEnv("STORE_PORT"),
			Password: os.Getenv("STORE_PASSWORD"),
		},
		Token: Token{
			PublicKey:  publicKey,
			PrivateKey: privateKey,
		},
	}

	err = c.validate()
	return c, err
}

func (c *Config) validate() error {
	v := reflect.ValueOf(c).Elem()
	baseConfigType := reflect.TypeOf((*сonfiger)(nil)).Elem()

	for i := 0; i < v.NumField(); i++ {
		if v.Type().Field(i).Type.Implements(baseConfigType) {
			if err := v.Field(i).Interface().(сonfiger).validate(); err != nil {
				return err
			}
		}
	}
	return nil
}

func uintEnv(env string) uint64 {
	v := os.Getenv(env)
	if v == "" {
		return 0
	}

	u, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return 0
	}
	return u
}

func decodeEnv(env string) (string, error) {
	v := os.Getenv(env)
	if v == "" {
		return "", nil
	}

	dec, err := base64.StdEncoding.DecodeString(v)
	if err != nil {
		return "", fmt.Errorf("DecodeString(): %s", err)
	}
	return string(dec), nil
}

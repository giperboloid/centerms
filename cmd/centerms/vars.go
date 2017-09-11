package main

import (
	"github.com/giperboloid/centerms/entities"
	"os"
	"strconv"
	"github.com/pkg/errors"
)

var (
	StorageServer = entities.Server{
		Host: getEnvStorageHost("REDIS_PORT_6379_TCP_ADDR"),
		Port: getEnvStoragePort("REDIS_PORT_6379_TCP_PORT"),
	}

	localhost     = "0.0.0.0"
	streamPort    = uint(2540)
	devDataPort   = uint(3030)
	devConfigPort = uint(3000)
	webPort       = uint(8100)
)

func getEnvStoragePort(key string) uint {
	parsed, err := strconv.ParseUint(os.Getenv(key), 10, 64)
	if err != nil {
		errors.New("main: getEnvStoragePort(): ParseUint() has failed")
	}

	port := uint(parsed)
	if port == 0 {
		return uint(6379)
	}
	return port
}

func getEnvStorageHost(key string) string {
	host := os.Getenv(key)
	if len(host) == 0 {
		return "127.0.0.1"
	}
	return host
}

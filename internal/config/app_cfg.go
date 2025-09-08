package config

import (
	"github.com/dart-chain/hello-actions/internal/consts"
)

type AppCfg struct {
	Host    string
	Port    string
	IsDebug bool
}

func LoadConfig(host, port string, isDebug bool) (*AppCfg, error) {
	if host == "" {
		host = consts.DEFAULT_HOST
	}
	if port == "" {
		port = consts.DEFAULT_PORT
	}

	return &AppCfg{
		Host:    host,
		Port:    port,
		IsDebug: isDebug,
	}, nil
}

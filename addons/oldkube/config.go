package main

import (
	"dtstack.com/dtstack/easymatrix/addons/oldkube/pkg/api"
	"dtstack.com/dtstack/easymatrix/addons/oldkube/pkg/base"
	"dtstack.com/dtstack/easymatrix/go-common/log"
	"fmt"
	"github.com/elastic/go-ucfg"
	"github.com/elastic/go-ucfg/yaml"
	"net"
	"os"
	"strconv"
)

const LOG_PREFIX = "oldkube"

type LogConfig struct {
	Dir        string `config:"dir" validate:"required"`
	MaxSize    int    `config:"max-logger-size"`
	MaxBackups int    `config:"max-logger-backups"`
	MaxAge     int    `config:"days-to-keep"`
}

type ApiConfig struct {
	Host     string `config:"host"`
	Port     int    `config:"port" validate:"required"`
	Restrict bool   `config:"restrict-api-check"`
}

type Config struct {
	Log LogConfig `config:"log" validate:"required"`
	Api ApiConfig `config:"api" validate:"required"`
}

func ParseConfig(configFile string) error {
	configContent, err := yaml.NewConfigWithFile(configFile, ucfg.PathSep("."))
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("config file not found!")
		}
		return err
	}

	config := Config{}
	if err := configContent.Unpack(&config); err != nil {
		return err
	}
	clog := config.Log
	if err := log.ConfigureLogger(LOG_PREFIX, clog.Dir, clog.MaxSize, clog.MaxBackups, clog.MaxAge); err != nil {
		return err
	} else {
		fmt.Printf("Saving logs at %s\n", clog.Dir)
	}

	apiconf := &config.Api

	if err := base.ConfigureApiServer(apiconf.Host, apiconf.Port, &api.ApiV2Schema, apiconf.Restrict); err != nil {
		return err
	} else {
		fmt.Printf("Running API service at %v\n", net.JoinHostPort(apiconf.Host, strconv.Itoa(apiconf.Port)))
	}

	return nil
}

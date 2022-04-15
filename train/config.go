package main

import (
	"dtstack.com/dtstack/easymatrix/matrix/log"
	"dtstack.com/dtstack/easymatrix/matrix/model"
	"dtstack.com/dtstack/easymatrix/train/define"
	"dtstack.com/dtstack/easymatrix/train/man"
	"fmt"
	"github.com/elastic/go-ucfg"
	"github.com/elastic/go-ucfg/yaml"
	"os"
)

func ParseConfig(configFile string) error {
	configContent, err := yaml.NewConfigWithFile(configFile, ucfg.PathSep("."))
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("config file not found!")
		}
		return err
	}

	config := define.Config{}
	if err := configContent.Unpack(&config); err != nil {
		return err
	}

	clog := &config.Log
	if err := log.ConfigureLogger(clog.Dir, clog.MaxSize, clog.MaxBackups, clog.MaxAge); err != nil {
		return err
	} else {
		fmt.Printf("Saving logs at %s\n", clog.Dir)
	}

	agent := &config.Agent
	em := &config.Em
	operatpors := config.OperatorList

	man.InitDockerCompose(em.DockerCompose, operatpors)

	db := &config.MysqlDb
	if err := model.ConfigureMysqlDatabase(db.Host, db.Port, db.User, db.Password, db.DbName); err != nil {
		fmt.Printf("%v\n", err)
	}
	if err := man.InitAgentClient(agent.Host, em.DockerCompose); err != nil {
		fmt.Printf("%v\n", err)
	}
	return nil
}

package main

import (
	"dtstack.com/dtstack/easymatrix/addons/easyfiler/prod/server/base"
	"dtstack.com/dtstack/easymatrix/go-common/log"
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	fmt.Println("Easyfiler " + VERSION)
	fmt.Println("Copyright (c) 2019 DTStack Inc.")
	base.ConfigureProductVersion(VERSION)

	app := cli.NewApp()
	app.Version = VERSION
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config,c",
			Usage: "config path",
		},
		cli.BoolFlag{
			Name:  "debug",
			Usage: "debug info",
		},
	}

	app.Action = func(ctx *cli.Context) error {
		log.SetDebug(ctx.Bool("debug"))
		srv, err := ParseConfig(ctx.String("config"))
		if err != nil {
			fmt.Printf("failed to parse config file ,err:%v\n", err)
			os.Exit(1)
		}
		return base.Run(srv)
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("exit with failure: %v\n", err)
		os.Exit(1)
	}
}

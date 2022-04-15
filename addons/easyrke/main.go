package main

import (
	"fmt"
	"os"

	"dtstack.com/dtstack/easymatrix/addons/easyrke/base"
	"dtstack.com/dtstack/easymatrix/addons/easyrke/rke"
	"dtstack.com/dtstack/easymatrix/go-common/log"
	"github.com/urfave/cli"
)

func main() {
	fmt.Println(base.VERSION)
	fmt.Println("Copyright (c) 2017 DTStack Inc.")
	base.ConfigureProductVersion(base.VERSION)

	app := cli.NewApp()
	app.Version = base.VERSION
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config,c",
			Usage: "config path",
		},
		cli.BoolFlag{
			Name:  "debug",
			Usage: "debug info",
		},
		cli.StringFlag{
			Name:  "k8s",
			Usage: "k8s  version ",
		},
	}
	//
	app.Action = func(ctx *cli.Context) error {
		log.SetDebug(ctx.Bool("debug"))
		if err := ParseConfig(ctx.String("config")); err != nil {
			return err
		}
		rke.SaveImage(ctx.String("k8s"))
		return base.Run()
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "exit with failure: %v\n", err)
	}
}

package main

import (
	"fmt"
	"os"

	"dtstack.com/dtstack/easymatrix/matrix/base"
	"dtstack.com/dtstack/easymatrix/matrix/log"
	"dtstack.com/dtstack/easymatrix/train/man"
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
			Value: "config.yml",
		},
		cli.BoolFlag{
			Name:  "debug",
			Usage: "debug info",
		},
	}
	//
	app.Action = func(ctx *cli.Context) error {
		log.SetDebug(ctx.Bool("debug"))
		if err := ParseConfig(ctx.String("config")); err != nil {
			return err
		}
		man.TrainMan.Exec(ctx.Args().First())
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "exit with failure: %v\n", err)
	}
}

package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

const (
	VERSION = "1.0.0"
)

func main() {
	app := cli.NewApp()
	app.Name = "Easyfiler"
	app.Version = VERSION
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "e.g: --config ./config.yml",
		},
	}
	app.Commands = []cli.Command{
		listCommand(),
		downloadCommand(),
		uploadCommand(),
		previewCommand(),
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("exit with failure: %v\n", err)
		os.Exit(1)
	}

}

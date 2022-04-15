package main

import (
	"fmt"
	"os"

	"dtstack.com/dtstack/easymatrix/matrix/base"
	"dtstack.com/dtstack/easymatrix/matrix/log"
	"github.com/urfave/cli"
)

// @title           maxtrix
// @version         2.0
// @description     matrtix swagger api doc
// @termsOfService  http://swagger.io/terms/
// @contact.name    API Support
// @contact.url     http://www.swagger.io/support
// @contact.email   support@swagger.io
// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html
// @host            localhost:8864
// @BasePath        /api/v2
// @securityDefinitions.basic  BasicAuth
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
	}
	// 解析配置文件并启动iris web服务
	app.Action = func(ctx *cli.Context) error {
		log.SetDebug(ctx.Bool("debug"))
		config := ctx.String("config")
		if err := ParseConfig(config); err != nil {
			return err
		}
		return base.Run()
	}
	//
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "exit with failure: %v\n", err)
	}
}

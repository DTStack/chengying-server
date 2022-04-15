package base

import (
	serve "dtstack.com/dtstack/easymatrix/addons/easyfiler/pkg/rpc-server"
	"dtstack.com/dtstack/easymatrix/go-common/log"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(srv *serve.Server) error {
	healthChecker := time.NewTimer(time.Second * 10)
	signals := make(chan os.Signal, 1)
	log.Infof("easyfiler start")
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go srv.Start()
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "FATAL %v\n", r)
		}
		healthChecker.Stop()
		exitGracefully()
	}()

LOOP:
	for {
		select {
		case sig := <-signals:
			fmt.Printf("Quit according to signal '%s'\n", sig.String())
			break LOOP
		case systemFailure := <-_SYSTEM_FAIL:
			if systemFailure.ExitCode > 0 {
				return fmt.Errorf("SYSTEM FAILURE: %d\nREASON: %s", systemFailure.ExitCode, systemFailure.Reason)
			}
		case <-healthChecker.C:
			checkSystemHealth()
		}
	}
	return nil
}

func exitGracefully() {

}

func checkSystemHealth() {

}

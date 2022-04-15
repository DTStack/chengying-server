package base

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kataras/iris"
)

func Run() error {
	healthChecker := time.NewTimer(time.Second * 10)
	signals := make(chan os.Signal, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "FATAL!!! %v\n", r)
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
	if app := iris.Default(); app != nil && app.Shutdown != nil {
		fmt.Println("Stopping API server...")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		if err := app.Shutdown(ctx); err != nil {
			fmt.Fprintf(os.Stderr, "Unable to stop API server: %v\n", err)
		}
	}
}

func checkSystemHealth() {

}

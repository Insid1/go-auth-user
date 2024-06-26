package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Insid1/go-auth-user/auth-service/internal/app"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {

	ctx := context.Background()
	application, err := app.NewApp(ctx)

	if err != nil {
		panic(err)
	}

	go func() {
		err = application.Run()
		if err != nil {
			application.Logger.Errorf("Failed to run application: %v", err)
		}
	}()
	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sign := <-stop

	application.Logger.Info("Stopping application...", zap.String("signal", sign.String()))
	err = application.Stop()
	if err != nil {
		application.Logger.Error("Failed to stop application", zap.Error(err))
	} else {
		application.Logger.Info("Application stopped")
	}
}

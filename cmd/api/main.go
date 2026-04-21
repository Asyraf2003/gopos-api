package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"pos-go/internal/app/bootstrap"
	"pos-go/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	app, err := bootstrap.New(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer app.DB.Close()

	go func() {
		if err := app.Echo.Start(cfg.HTTPAddr()); err != nil {
			log.Printf("http server stopped: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.Echo.Shutdown(shutdownCtx); err != nil {
		log.Fatal(err)
	}
}

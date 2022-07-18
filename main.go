package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/andyfen/oauth-server/server"
	"github.com/andyfen/oauth-server/server/config"
)

func main() {

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	conf := config.New()

	srv, err := server.New(conf)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-stopChan
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)

	defer cancel()
}

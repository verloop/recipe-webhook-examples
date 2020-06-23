package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	logger "github.com/sirupsen/logrus"
)

const (
	DefaultHTTPPort = "8080"
)

func Serve() {
	api := NewRechargeAPI()

	server := http.Server{
		Addr:         fmt.Sprintf(":%s", GetEnvOrDefault("PORT", DefaultHTTPPort)),
		Handler:      api,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	serverErrors := make(chan error, 1)
	go func() {
		logger.Info("Starting HTTP Server")
		serverErrors <- server.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		//Possible mem leak need to check
		case err := <-serverErrors:
			logger.WithError(err).Fatalln("Error starting http server")
			return
		case <-shutdown:
			logger.Warningln("Received a signal to shutdown http server")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			err := server.Shutdown(ctx)
			if err != nil {
				logger.WithError(err).Error("Graceful shutdown of http server failed")
				err = server.Close()
			}

			if err != nil {
				logger.WithError(err).Error("Could not stop http server")
			}
			return
		}
	}
}

package main

import (
	"fmt"
	"net/http"

	"go-cloud-native/app/app"
	"go-cloud-native/app/router"
	"go-cloud-native/config"
	lr "go-cloud-native/util/logger"
)

func main() {
	appConf := config.AppConfig()

	logger := lr.New(appConf.Debug)

	application := app.New(logger)

	appRouter := router.New(application)

	address := fmt.Sprintf(":%d", appConf.Server.Port)

	logger.Info().Msgf("Starting server %v", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server startup failed")
	}
}

// Greet ...
func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

package main

import (
	"fmt"

	"go.uber.org/zap"

	"example.com/learn/pkg/publishPlatformMusic/api"
	"example.com/learn/pkg/publishPlatformMusic/config"
	"example.com/learn/pkg/publishPlatformMusic/server"
	migrations "example.com/learn/pkg/publishPlatformMusic/storage/migrations"
)

func main() {
	cnf := config.Load()

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	if err := migrations.Run(cnf.App.PostgresURI); err != nil {
		logger.Fatal("postgres connection error", zap.Error(err))
	}

	handler, err := api.NewHandler(logger, api.HandlerConfig{
		AppPostgresURI: cnf.App.PostgresURI,
		JwtSecret:      []byte(cnf.JwtSecret),
	})
	if err != nil {
		logger.Fatal("api.NewHandler() failed", zap.Error(err))
	}

	srv, err := server.NewServer(&server.ServerConfig{
		Handler:   handler,
		JwtSecret: []byte(cnf.JwtSecret),
		Address:   fmt.Sprintf(":%d", cnf.API.Port),
	})
	if err != nil {
		logger.Fatal("server.NewServer() failed", zap.Error(err))
	}

	fmt.Printf("running server :%d\n", cnf.API.Port)
	srv.Start()
}

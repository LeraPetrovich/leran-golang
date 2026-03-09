package main

import (
	"fmt"

	"example.com/learn/pkg/publishPlatformMusic/core/migrations"
	"example.com/learn/pkg/publishPlatformMusic/service/music"
	"go.uber.org/zap"

	"example.com/learn/pkg/publishPlatformMusic/api"
	"example.com/learn/pkg/publishPlatformMusic/config"
	"example.com/learn/pkg/publishPlatformMusic/db"
	"example.com/learn/pkg/publishPlatformMusic/server"
	"example.com/learn/pkg/publishPlatformMusic/storage"
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
	postgres, err := db.NewPostgresDb(cnf.App.PostgresURI, cnf.App.PostgresConnections)
	if err != nil {
		panic(err)
	}
	appStorage := storage.New(logger, postgres)
	appService := music.NewService(logger, appStorage)

	handler := api.NewHandler(logger, appService, api.HandlerConfig{
		JwtSecret: []byte(cnf.JwtSecret),
	})

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

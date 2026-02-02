package main

import (
	"fmt"

	"go.uber.org/zap"

	"example.com/learn/pkg/learn/api"
	homepageConfig "example.com/learn/pkg/learn/config"
	"example.com/learn/pkg/learn/server"
)

func main() {
	cnf := homepageConfig.Load()

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	handlerConfig := api.HandlerConfig{
		PostgresURI: cnf.App.PostgresURI,
		JwtSecret:   []byte(cnf.JwtSecret),
	}

	handler, error := api.New(logger, handlerConfig)

	if error != nil {
		panic(error)
	}

	serverConfig := server.ServerConfig{
		Handler:   handler,
		Address:   fmt.Sprintf(":%v", cnf.API.Port),
		JwtSecret: []byte(cnf.JwtSecret),
	}

	server, err := server.New(logger, serverConfig)

	if err != nil {
		logger.Fatal("api.NewServer() failed", zap.Error(err))
	}

	fmt.Printf("running server :%v\n", cnf.API.Port)
	server.Run()
}

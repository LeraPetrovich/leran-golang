package main

import (
	"log"

	"github.com/LeraPetrovich/learn-golang/pkg/handler"
	"github.com/LeraPetrovich/learn-golang/pkg/repository"
	"github.com/LeraPetrovich/learn-golang/pkg/service"
	"github.com/LeraPetrovich/learn-golang/server"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init configs: %s", err.Error())
	}
	repository := repository.NewRepository()
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	svr := new(server.Server)

	if err := svr.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

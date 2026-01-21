package main

import (
	"log"

	"github.com/LeraPetrovich/learn-golang/pkg/handler"
	"github.com/LeraPetrovich/learn-golang/pkg/repository"
	"github.com/LeraPetrovich/learn-golang/pkg/service"
	"github.com/LeraPetrovich/learn-golang/server"
)

func main() {
	repository := repository.NewRepository()
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	svr := new(server.Server)

	if err := svr.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

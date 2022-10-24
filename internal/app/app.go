package app

import (
	"log"
	"main/internal/delivery"
	"main/internal/service"
)

func Run() {
	service := new(service.Service)
	handler := delivery.NewHandler(service)
	server := new(delivery.Server)

	if err := server.ServerRun(":8080", handler.InitRouter()); err != nil {
		log.Printf("Server doesn`t run :%s", err.Error())
		return
	}
}

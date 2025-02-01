package main

import (
	"github.com/TimmyTurner98/consult/pkg/handler"
	"github.com/TimmyTurner98/consult/pkg/service"
)

func main() {
	//services := &service.Service{}

	services := service.NewService()

	h := handler.NewHandler(services)

	router := h.InitRoutes()
	router.Run(":8080")

}

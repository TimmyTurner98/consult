package main

import (
	"github.com/TimmyTurner98/consult/pkg/handler"
	"github.com/TimmyTurner98/consult/pkg/repository"
	"github.com/TimmyTurner98/consult/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	repos := repository.NewRepository()

	services := service.NewService(repos)

	h := handler.NewHandler(services)

	router := h.InitRoutes()
	if err := router.Run(":" + viper.GetString("port")); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

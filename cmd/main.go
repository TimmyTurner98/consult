package main

import (
	"github.com/TimmyTurner98/consult/pkg/handler"
	"github.com/TimmyTurner98/consult/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfigs(); err != nil {
		log.Fatalf("Error initializing configs: %s", err.Error())
	}

	services := service.NewService()

	h := handler.NewHandler(services)

	router := h.InitRoutes()
	if err := router.Run(":" + viper.GetString("port")); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

package main

import (
	"github.com/TimmyTurner98/consult/pkg/dbs/postgres"
	"github.com/TimmyTurner98/consult/pkg/handler"
	"github.com/TimmyTurner98/consult/pkg/repository"
	"github.com/TimmyTurner98/consult/pkg/service"
	"github.com/sirupsen/logrus"
)

func main() {
	// Загружаем конфигурацию
	if err := postgres.InitDBConfig(); err != nil {
		logrus.Fatalf("❌ Ошибка загрузки конфигурации: %v", err)
	}
	// Подключаемся к БД
	postgres.ConnectDB()
	repos := repository.NewRepository(postgres.DB)

	services := service.NewService(repos)

	h := handler.NewHandler(services)

	port := postgres.DBConfig().Server.Port

	router := h.InitRoutes()

	// Запускаем сервер
	logrus.Infof("🚀 Сервер запущен на порту %s", port)
	if err := router.Run(":" + port); err != nil {
		logrus.Fatalf("❌ Ошибка при запуске сервера: %s", err.Error())
	}
}

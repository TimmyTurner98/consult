package handler

import (
	"github.com/TimmyTurner98/consult/pkg/service"
	"github.com/gin-gonic/gin"
)

//type Handler struct {
//	services *service.Service
//}

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/clients", h.createClient)

	return router
}

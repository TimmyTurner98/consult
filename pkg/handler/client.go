package handler

import (
	"github.com/TimmyTurner98/consult/pkg/modules"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createClient(c *gin.Context) {
	var client modules.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Передаем клиентский объект в сервис для создания
	createdClient, err := h.services.CreateClient(client)
	if err != nil {
		// Если произошла ошибка при создании клиента
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create client"})
		return
	}

	// Если клиент успешно создан, возвращаем статус 201 и данные клиента
	c.JSON(http.StatusCreated, createdClient)
}

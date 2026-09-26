package handlers

import (
	"net/http"
	interfaces "to-do-server/internal/Application/Interfaces/services"

	"github.com/gin-gonic/gin"
)

type PrioridadHandler struct {
	service interfaces.IPrioridadService
}

func NewPrioridadHandler(
	service interfaces.IPrioridadService,
) *PrioridadHandler {
	return &PrioridadHandler{
		service: service,
	}
}

func (h *PrioridadHandler) ObtenerTodos(c *gin.Context) {

	prioridades, err := h.service.ObtenerPrioridades(
		c.Request.Context(),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, prioridades)
}

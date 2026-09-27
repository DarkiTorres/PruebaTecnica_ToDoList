package handlers

import (
	"net/http"
	"strconv"
	"time"

	tareas "to-do-server/internal/Application/DTOs/Tareas"
	services "to-do-server/internal/Application/Interfaces/services"
	entities "to-do-server/internal/Domain/Entities"

	"github.com/gin-gonic/gin"
)

type SubTareaHandler struct {
	service services.ISubTareaService
}

func NewSubTareaHandler(
	service services.ISubTareaService,
) *SubTareaHandler {
	return &SubTareaHandler{
		service: service,
	}
}

func (h *SubTareaHandler) ObtenerPorTareaId(c *gin.Context) {

	tareaId, err := strconv.ParseInt(
		c.Param("id"),
		10,
		64,
	)

	if err != nil || tareaId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "El id de la tarea no es válido",
		})
		return
	}

	subTareas, err := h.service.ObtenerPorTareaId(
		c.Request.Context(),
		tareaId,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, subTareas)
}

func (h *SubTareaHandler) Crear(c *gin.Context) {

	tareaId, err := strconv.ParseInt(
		c.Param("id"),
		10,
		64,
	)

	if err != nil || tareaId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "El id de la tarea no es válido",
		})
		return
	}

	var request tareas.CrearSubTareaRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "El cuerpo de la solicitud no es válido",
		})
		return
	}

	subTarea := &entities.SubTarea{
		TareaId:   tareaId,
		Titulo:    request.Titulo,
		CreadoEl:  time.Now(),
		CreadoPor: request.CreadoPor,
	}

	if err := h.service.CrearSubTarea(
		c.Request.Context(),
		subTarea,
	); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, subTarea)
}

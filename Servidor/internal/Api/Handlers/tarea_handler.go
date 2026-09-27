package handlers

import (
	"net/http"
	"strconv"
	"time"
	tareas "to-do-server/internal/Application/DTOs/Tareas"
	"to-do-server/internal/Application/Interfaces/repositories"
	interfaces "to-do-server/internal/Application/Interfaces/services"
	entities "to-do-server/internal/Domain/Entities"

	"github.com/gin-gonic/gin"
)

type TareaHandler struct {
	service            interfaces.ITareaService
	subTareaService    interfaces.ISubTareaService
	subTareaRepository repositories.ISubTareaRepository
}

func NewTareaHandler(
	service interfaces.ITareaService,
	subTareaService interfaces.ISubTareaService,
	subTareaRepository repositories.ISubTareaRepository,
) *TareaHandler {
	return &TareaHandler{
		service:            service,
		subTareaService:    subTareaService,
		subTareaRepository: subTareaRepository,
	}
}

func (h *TareaHandler) Crear(c *gin.Context) {

	var request tareas.CrearTareaRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "El cuerpo de la solicitud no es válido",
		})
		return
	}

	var fechaEntrega *time.Time

	if request.FechaEntrega != nil {

		fecha, err := time.Parse(
			time.RFC3339,
			*request.FechaEntrega,
		)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "La fecha de entrega no tiene un formato válido",
			})
			return
		}

		fechaEntrega = &fecha
	}

	tarea := &entities.Tarea{
		Titulo:       request.Titulo,
		Descripcion:  request.Descripcion,
		PrioridadId:  request.PrioridadId,
		FechaEntrega: fechaEntrega,
		CreadoEl:     time.Now(),
		CreadoPor:    request.CreadoPor,
	}

	err := h.service.CrearTarea(
		c.Request.Context(),
		tarea,
		request.AsignadoA,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := h.construirTareaResponse(c, *tarea)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (h *TareaHandler) ObtenerTodas(c *gin.Context) {
	tareasObtenidas, err := h.service.ObtenerTareas(
		c.Request.Context(),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	respuestas := make(
		[]tareas.TareaResponse,
		0,
		len(tareasObtenidas),
	)

	for _, tarea := range tareasObtenidas {

		response, err := h.construirTareaResponse(
			c,
			tarea,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		respuestas = append(respuestas, response)
	}

	c.JSON(http.StatusOK, respuestas)
}

func (h *TareaHandler) ObtenerPorId(c *gin.Context) {
	id, err := strconv.ParseInt(
		c.Param("id"),
		10,
		64,
	)

	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "El id de la tarea no es válido",
		})
		return
	}

	tarea, err := h.service.ObtenerTareaPorId(
		c.Request.Context(),
		id,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := h.construirTareaResponse(
		c,
		*tarea,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *TareaHandler) Actualizar(c *gin.Context) {

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

	var request tareas.ActualizarTareaRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "El cuerpo de la solicitud no es válido",
		})
		return
	}

	var fechaEntrega *time.Time

	if request.FechaEntrega != nil {

		fecha, err := time.Parse(
			time.RFC3339,
			*request.FechaEntrega,
		)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "La fecha de entrega no tiene un formato válido",
			})
			return
		}

		fechaEntrega = &fecha
	}

	tarea := &entities.Tarea{
		Id:            tareaId,
		Titulo:        request.Titulo,
		Descripcion:   request.Descripcion,
		PrioridadId:   request.PrioridadId,
		FechaEntrega:  fechaEntrega,
		ModificadoEl:  func() *time.Time { ahora := time.Now(); return &ahora }(),
		ModificadoPor: &request.ModificadoPor,
	}

	if err := h.service.ActualizarTarea(
		c.Request.Context(),
		tarea,
	); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, subTareaRequest := range request.SubTareas {

		if subTareaRequest.Id == 0 {

			subTarea := &entities.SubTarea{
				TareaId:       tareaId,
				Titulo:        subTareaRequest.Titulo,
				EstaTerminada: subTareaRequest.EstaTerminada,
				CreadoPor:     subTareaRequest.CreadoPor,
			}

			if err := h.subTareaService.CrearSubTarea(
				c.Request.Context(),
				subTarea,
			); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			continue
		}

		if subTareaRequest.EstaEliminada {

			if err := h.subTareaService.EliminarSubTarea(
				c.Request.Context(),
				subTareaRequest.Id,
				subTareaRequest.ModificadoPor,
			); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			continue
		}

		subTarea := &entities.SubTarea{
			Id:            subTareaRequest.Id,
			TareaId:       tareaId,
			Titulo:        subTareaRequest.Titulo,
			EstaTerminada: subTareaRequest.EstaTerminada,
			CreadoPor:     subTareaRequest.CreadoPor,
			ModificadoPor: &subTareaRequest.ModificadoPor,
		}

		if err := h.subTareaService.ActualizarSubTarea(
			c.Request.Context(),
			subTarea,
		); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, tarea)
}

func (h *TareaHandler) Eliminar(c *gin.Context) {

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

	var request tareas.EliminarTareaRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "El cuerpo de la solicitud no es válido",
		})
		return
	}

	if request.ModificadoPor <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "El usuario modificador no es válido",
		})
		return
	}

	if err := h.service.EliminarTarea(
		c.Request.Context(),
		tareaId,
		request.ModificadoPor,
	); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}
func (h *TareaHandler) Completar(c *gin.Context) {

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

	var request tareas.CompletarTareaRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "El cuerpo de la solicitud no es válido",
		})
		return
	}

	if err := h.service.CompletarTarea(
		c.Request.Context(),
		tareaId,
		request.EstaTerminada,
	); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	println("ENTRO A COMPLETAR - ESCRIBIENDO 204")
	c.Status(http.StatusNoContent)
	println("GIN STATUS:", c.Writer.Status())
	println("RECORDED STATUS:", c.Writer.Status())
}
func (h *TareaHandler) ObtenerPorUsuario(c *gin.Context) {
	usuarioId, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil || usuarioId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "El id del usuario no es válido",
		})
		return
	}

	tareasObtenidas, err := h.service.ObtenerTareasPorUsuario(
		c.Request.Context(),
		usuarioId,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	respuestas := make(
		[]tareas.TareaResponse,
		0,
		len(tareasObtenidas),
	)

	for _, tarea := range tareasObtenidas {

		response, err := h.construirTareaResponse(
			c,
			tarea,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		respuestas = append(respuestas, response)
	}

	c.JSON(http.StatusOK, respuestas)

}

func (h *TareaHandler) construirTareaResponse(
	c *gin.Context,
	tarea entities.Tarea,
) (tareas.TareaResponse, error) {

	subTareas, err := h.subTareaRepository.ObtenerPorTareaId(
		c.Request.Context(),
		tarea.Id,
	)

	if err != nil {
		return tareas.TareaResponse{}, err
	}

	subTareasResponse := make(
		[]tareas.SubTareaResponse,
		0,
		len(subTareas),
	)

	for _, subTarea := range subTareas {

		subTareasResponse = append(
			subTareasResponse,
			tareas.SubTareaResponse{
				Id:            subTarea.Id,
				TareaId:       subTarea.TareaId,
				Titulo:        subTarea.Titulo,
				EstaTerminada: subTarea.EstaTerminada,
				EstaEliminada: subTarea.EstaEliminada,
				CreadoEl:      subTarea.CreadoEl,
				CreadoPor:     subTarea.CreadoPor,
				ModificadoEl:  subTarea.ModificadoEl,
				ModificadoPor: subTarea.ModificadoPor,
			},
		)
	}

	return tareas.TareaResponse{
		Id:            tarea.Id,
		Titulo:        tarea.Titulo,
		Descripcion:   tarea.Descripcion,
		PrioridadId:   tarea.PrioridadId,
		FechaEntrega:  tarea.FechaEntrega,
		EstaTerminada: tarea.EstaTerminada,
		EstaEliminada: tarea.EstaEliminada,
		CreadoEl:      tarea.CreadoEl,
		CreadoPor:     tarea.CreadoPor,
		ModificadoEl:  tarea.ModificadoEl,
		ModificadoPor: tarea.ModificadoPor,
		SubTareas:     subTareasResponse,
	}, nil
}

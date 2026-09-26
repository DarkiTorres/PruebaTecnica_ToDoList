package handlers

import (
	"net/http"
	"strconv"
	interfaces "to-do-server/internal/Application/Interfaces/services"

	"github.com/gin-gonic/gin"
)

type UsuarioHandler struct {
	service interfaces.IUsuarioService
}

func NewUsuarioHandler(
	service interfaces.IUsuarioService,
) *UsuarioHandler {
	return &UsuarioHandler{
		service: service,
	}
}

func (h *UsuarioHandler) ObtenerTodos(c *gin.Context) {

	usuarios, err := h.service.ObtenerUsuarios(
		c.Request.Context(),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, usuarios)
}

func (h *UsuarioHandler) ObtenerPorId(c *gin.Context) {

	usuarioId, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil || usuarioId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "El id del usuario no es válido",
		})
		return
	}

	usuario, err := h.service.ObtenerUsuarioPorId(
		c.Request.Context(),
		usuarioId,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

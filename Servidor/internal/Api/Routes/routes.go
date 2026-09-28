package routes

import (
	handlers "to-do-server/internal/Api/Handlers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConfigurarRutas(
	router *gin.Engine,
	tareaHandler *handlers.TareaHandler,
	subtareaHandler *handlers.SubTareaHandler,
	usuarioHandler *handlers.UsuarioHandler,
	prioridadHandler *handlers.PrioridadHandler,
	db *pgxpool.Pool,

) {
	router.GET("/health", func(c *gin.Context) {
		if err := db.Ping(c.Request.Context()); err != nil {
			c.JSON(503, gin.H{
				"status":   "error",
				"database": "disconnected",
			})
			return
		}

		c.JSON(200, gin.H{
			"status":   "ok",
			"database": "connected",
		})
	})

	router.POST("/tareas", tareaHandler.Crear)
	router.GET("/tareas", tareaHandler.ObtenerTodas)
	router.GET("/tareas/:id", tareaHandler.ObtenerPorId)
	router.PUT("/tareas/:id", tareaHandler.Actualizar)
	router.PATCH("/tareas/:id/eliminar", tareaHandler.Eliminar)
	router.PATCH("/tareas/:id/completar", tareaHandler.Completar)
	router.GET("/usuarios/:id/tareas", tareaHandler.ObtenerPorUsuario)

	router.POST("/tareas/:id/subtareas", subtareaHandler.Crear)
	router.GET("/tareas/:id/subtareas", subtareaHandler.ObtenerPorTareaId)

	router.GET("/usuarios", usuarioHandler.ObtenerTodos)
	router.GET("/usuarios/:id", usuarioHandler.ObtenerPorId)
	router.PATCH("/subtareas/:id/completar", subtareaHandler.Completar)

	router.GET("/prioridades", prioridadHandler.ObtenerTodos)
}

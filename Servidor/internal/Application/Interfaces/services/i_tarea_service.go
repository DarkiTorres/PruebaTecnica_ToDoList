package services

import (
	"context"
	entities "to-do-server/internal/Domain/Entities"
)

type ITareaService interface {
	CrearTarea(context.Context, *entities.Tarea, int) error
	ObtenerTareas(context.Context) ([]entities.Tarea, error)
	ObtenerTareaPorId(context.Context, int64) (*entities.Tarea, error)
	ActualizarTarea(context.Context, *entities.Tarea, int) error
	EliminarTarea(context.Context, int64, int) error
	EliminarFisicamente(context.Context, int64) error
	CompletarTarea(context.Context, int64, bool) error
	ObtenerTareasPorUsuario(context.Context, int) ([]entities.Tarea, error)
	ObtenerBitacora(context.Context) ([]entities.Tarea, error)
	RestaurarTarea(context.Context, int64, int) error
}

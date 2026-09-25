package repositories

import (
	"context"
	entities "to-do-server/internal/Domain/Entities"
)

type ITareaRepository interface {
	Crear(context context.Context, tarea *entities.Tarea) error

	ObtenerPorId(context context.Context, id int64) (*entities.Tarea, error)

	ObtenerTodos(context context.Context) ([]entities.Tarea, error)

	Actualizar(context context.Context, tarea *entities.Tarea) error

	Eliminar(context context.Context, id int64, modificadoPor int) error
}

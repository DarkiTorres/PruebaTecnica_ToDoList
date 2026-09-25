package repositories

import (
	"context"
	entities "to-do-server/internal/Domain/Entities"
)

type ISubTareaRepository interface {
	Crear(context context.Context, subTarea *entities.SubTarea) error

	ObtenerPorId(context context.Context, id int64) (*entities.SubTarea, error)

	ObtenerPorTareaId(context context.Context, tareaId int64) ([]entities.SubTarea, error)

	Actualizar(context context.Context, subTarea *entities.SubTarea) error

	Eliminar(context context.Context, id int64, modificadoPor int) error
}

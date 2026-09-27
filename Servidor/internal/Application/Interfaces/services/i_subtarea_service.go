package services

import (
	"context"

	entities "to-do-server/internal/Domain/Entities"
)

type ISubTareaService interface {
	CrearSubTarea(
		context.Context,
		*entities.SubTarea,
	) error

	ActualizarSubTarea(
		context.Context,
		*entities.SubTarea,
	) error

	EliminarSubTarea(
		context.Context,
		int64,
		int,
	) error

	ObtenerPorTareaId(
		context context.Context,
		tareaId int64,
	) ([]entities.SubTarea, error)
}

package services

import (
	"context"
	entities "to-do-server/internal/Domain/Entities"
)

type ITareaUsuarioService interface {
	Crear(context.Context, *entities.TareaUsuario) error

	ObtenerPorTareaId(
		context.Context,
		int64,
	) ([]entities.TareaUsuario, error)

	ObtenerPorUsuarioId(
		context.Context,
		int,
	) ([]entities.TareaUsuario, error)

	Eliminar(
		context.Context,
		int64,
		int,
	) error
}

package repositories

import (
	"context"
	entities "to-do-server/internal/Domain/Entities"
)

type ITareaUsuarioRepository interface {
	Crear(context context.Context, tareaUsuario *entities.TareaUsuario) error

	ObtenerPorTareaId(context context.Context, tareaId int64) ([]entities.TareaUsuario, error)

	ObtenerPorUsuarioId(context context.Context, usuarioId int) ([]entities.TareaUsuario, error)

	Eliminar(context context.Context, tareaId int64, usuarioId int) error
}

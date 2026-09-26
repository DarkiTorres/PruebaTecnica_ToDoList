package services

import (
	"context"
	entities "to-do-server/internal/Domain/Entities"
)

type IUsuarioService interface {
	ObtenerUsuarioPorId(
		context.Context,
		int,
	) (*entities.Usuario, error)

	ObtenerUsuarios(
		context.Context,
	) ([]entities.Usuario, error)
}

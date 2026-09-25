package repositories

import (
	"context"
	entities "to-do-server/internal/Domain/Entities"
)

type IUsuarioRepository interface {
	ObtenerPorId(context context.Context, id int) (*entities.Usuario, error)
	ObtenerTodos(context context.Context) ([]entities.Usuario, error)
}

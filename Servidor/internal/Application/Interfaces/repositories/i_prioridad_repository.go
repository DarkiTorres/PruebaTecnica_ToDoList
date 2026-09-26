package repositories

import (
	"context"
	entities "to-do-server/internal/Domain/Entities"
)

type IPrioridadRepository interface {
	ObtenerTodos(context.Context) ([]entities.Prioridad, error)
}

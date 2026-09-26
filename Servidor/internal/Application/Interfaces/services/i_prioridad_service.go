package services

import (
	"context"

	entities "to-do-server/internal/Domain/Entities"
)

type IPrioridadService interface {
	ObtenerPrioridades(context.Context) ([]entities.Prioridad, error)
}

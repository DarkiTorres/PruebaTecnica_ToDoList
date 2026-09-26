package services

import (
	"context"

	"to-do-server/internal/Application/Interfaces/repositories"
	interfaces "to-do-server/internal/Application/Interfaces/services"
	entities "to-do-server/internal/Domain/Entities"
)

var _ interfaces.IPrioridadService = (*PrioridadService)(nil)

type PrioridadService struct {
	prioridadRepository repositories.IPrioridadRepository
}

func NewPrioridadService(
	prioridadRepository repositories.IPrioridadRepository,
) *PrioridadService {
	return &PrioridadService{
		prioridadRepository: prioridadRepository,
	}
}

func (s *PrioridadService) ObtenerPrioridades(
	context context.Context,
) ([]entities.Prioridad, error) {

	return s.prioridadRepository.ObtenerTodos(context)
}

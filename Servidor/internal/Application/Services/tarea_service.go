package services

import (
	"context"
	"errors"
	"to-do-server/internal/Application/Interfaces/repositories"
)

type TareaService struct {
	tareaRepository    repositories.ITareaRepository
	subTareaRepository repositories.ISubTareaRepository
}

func NewTareaService(
	tareaRepository repositories.ITareaRepository,
	subTareaRepository repositories.ISubTareaRepository,
) *TareaService {
	return &TareaService{
		tareaRepository:    tareaRepository,
		subTareaRepository: subTareaRepository,
	}
}

func (s *TareaService) CompletarTarea(context context.Context, tareaId int64) error {
	tarea, err := s.tareaRepository.ObtenerPorId(context, tareaId)
	if err != nil {
		return err
	}

	subTareas, err := s.subTareaRepository.ObtenerPorTareaId(context, tareaId)
	if err != nil {
		return err
	}

	for _, subTarea := range subTareas {
		if !subTarea.EstaTerminada {
			return errors.New("No se puede completar la tarea porque tiene subtareas pendientes.")
		}
	}

	tarea.EstaTerminada = true

	return s.tareaRepository.Actualizar(
		context, tarea,
	)
}

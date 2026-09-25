package services

import (
	"context"
	"errors"
	"time"
	"to-do-server/internal/Application/Interfaces/repositories"
	validators "to-do-server/internal/Application/Validators"
	entities "to-do-server/internal/Domain/Entities"
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

func (s *TareaService) CrearTarea(context context.Context, tarea *entities.Tarea) error {
	if err := validators.ValidarTarea(tarea); err != nil {
		return err
	}

	return s.tareaRepository.Crear(context, tarea)
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
		if subTarea.EstaEliminada {
			continue
		}

		if !subTarea.EstaTerminada {
			return errors.New("No se puede completar la tarea porque tiene subtareas pendientes.")
		}
	}

	tarea.EstaTerminada = true

	return s.tareaRepository.Actualizar(
		context, tarea,
	)
}

func (s *TareaService) ActualizarTarea(context context.Context, tarea *entities.Tarea) error {
	if tarea == nil {
		return errors.New("la tarea no puede ser nula")
	}

	if tarea.Id <= 0 {
		return errors.New("el id de la tarea no es válido")
	}

	if err := validators.ValidarTarea(tarea); err != nil {
		return err
	}

	return s.tareaRepository.Actualizar(
		context,
		tarea,
	)
}

func (s *TareaService) EliminarTarea(context context.Context, tareaId int64, modificadoPor int) error {
	if tareaId <= 0 {
		return errors.New("el id de la tarea no es válido")
	}

	if modificadoPor <= 0 {
		return errors.New("el usuario modificador no es válido")
	}

	tarea, err := s.tareaRepository.ObtenerPorId(
		context,
		tareaId,
	)
	if err != nil {
		return err
	}

	ahora := time.Now()

	tarea.EstaEliminada = true
	tarea.ModificadoEl = &ahora
	tarea.ModificadoPor = &modificadoPor

	return s.tareaRepository.Actualizar(
		context,
		tarea,
	)
}

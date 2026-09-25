package services

import (
	"context"
	"errors"
	"time"
	"to-do-server/internal/Application/Interfaces/repositories"
	validators "to-do-server/internal/Application/Validators"
	entities "to-do-server/internal/Domain/Entities"
)

type SubTareaService struct {
	subTareaRepository repositories.ISubTareaRepository
	tareaRepository    repositories.ITareaRepository
}

func NewSubTareaService(subTareaRepository repositories.ISubTareaRepository, tareaRepository repositories.ITareaRepository) *SubTareaService {
	return &SubTareaService{
		subTareaRepository: subTareaRepository,
		tareaRepository:    tareaRepository,
	}
}

func (s *SubTareaService) CrearSubTarea(context context.Context, subTarea *entities.SubTarea) error {

	if err := validators.ValidarSubTarea(subTarea); err != nil {
		return err
	}

	tarea, err := s.tareaRepository.ObtenerPorId(
		context,
		subTarea.TareaId,
	)
	if err != nil {
		return err
	}

	if tarea.EstaEliminada {
		return errors.New(
			"No se puede crear una subtarea para una tarea eliminada.",
		)
	}

	return s.subTareaRepository.Crear(
		context,
		subTarea,
	)
}

func (s *SubTareaService) ActualizarSubTarea(context context.Context, subTarea *entities.SubTarea) error {

	if subTarea == nil {
		return errors.New("la subtarea no puede ser nula")
	}

	if subTarea.Id <= 0 {
		return errors.New("el id de la subtarea no es válido")
	}

	if err := validators.ValidarSubTarea(subTarea); err != nil {
		return err
	}

	subTareaActual, err := s.subTareaRepository.ObtenerPorId(
		context,
		subTarea.Id,
	)
	if err != nil {
		return err
	}

	if subTareaActual.EstaEliminada {
		return errors.New(
			"No se puede actualizar una subtarea eliminada.",
		)
	}

	return s.subTareaRepository.Actualizar(
		context,
		subTarea,
	)
}

func (s *SubTareaService) EliminarSubTarea(context context.Context, subTareaId int64, modificadoPor int) error {
	if subTareaId <= 0 {
		return errors.New("el id de la subtarea no es válido")
	}

	if modificadoPor <= 0 {
		return errors.New("el usuario modificador no es válido")
	}

	subTarea, err := s.subTareaRepository.ObtenerPorId(
		context,
		subTareaId,
	)
	if err != nil {
		return err
	}

	if subTarea.EstaEliminada {
		return errors.New(
			"La subtarea ya está eliminada.",
		)
	}

	ahora := time.Now()

	subTarea.EstaEliminada = true
	subTarea.ModificadoEl = &ahora
	subTarea.ModificadoPor = &modificadoPor

	return s.subTareaRepository.Actualizar(
		context,
		subTarea,
	)
}

func (s *SubTareaService) CompletarSubTarea(context context.Context, subTareaId int64) error {
	if subTareaId <= 0 {
		return errors.New("el id de la subtarea no es válido")
	}

	subTarea, err := s.subTareaRepository.ObtenerPorId(
		context,
		subTareaId,
	)
	if err != nil {
		return err
	}

	if subTarea.EstaEliminada {
		return errors.New(
			"No se puede completar una subtarea eliminada.",
		)
	}

	subTarea.EstaTerminada = true

	return s.subTareaRepository.Actualizar(
		context,
		subTarea,
	)
}

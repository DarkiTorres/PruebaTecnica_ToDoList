package mocks

import (
	"context"
	"to-do-server/internal/Application/Interfaces/repositories"
	entities "to-do-server/internal/Domain/Entities"
)

type SubTareaRepositoryMock struct {
	SubTareas []entities.SubTarea
	SubTarea  *entities.SubTarea

	CrearLlamado      bool
	ActualizarLlamado bool

	ErrorCrear             error
	ErrorObtenerPorId      error
	ErrorActualizar        error
	ErrorObtenerPorTareaId error
}

var _ repositories.ISubTareaRepository = (*SubTareaRepositoryMock)(nil)

func (m *SubTareaRepositoryMock) Crear(
	context context.Context,
	subTarea *entities.SubTarea,
) error {
	m.CrearLlamado = true

	if m.ErrorCrear != nil {
		return m.ErrorCrear
	}

	return nil
}

func (m *SubTareaRepositoryMock) ObtenerPorId(
	context context.Context,
	id int64,
) (*entities.SubTarea, error) {
	if m.ErrorObtenerPorId != nil {
		return nil, m.ErrorObtenerPorId
	}

	return m.SubTarea, nil
}

func (m *SubTareaRepositoryMock) ObtenerPorTareaId(
	context context.Context,
	tareaId int64,
) ([]entities.SubTarea, error) {

	if m.ErrorObtenerPorTareaId != nil {
		return nil, m.ErrorObtenerPorTareaId
	}

	return m.SubTareas, nil
}

func (m *SubTareaRepositoryMock) Actualizar(
	context context.Context,
	subTarea *entities.SubTarea,
) error {
	m.ActualizarLlamado = true

	if m.ErrorActualizar != nil {
		return m.ErrorActualizar
	}

	m.SubTarea = subTarea

	return nil
}

func (m *SubTareaRepositoryMock) Eliminar(
	context context.Context,
	id int64,
	modificadoPor int,
) error {
	return nil
}

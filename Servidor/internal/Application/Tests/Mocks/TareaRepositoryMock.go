package mocks

import (
	"context"
	"to-do-server/internal/Application/Interfaces/repositories"
	entities "to-do-server/internal/Domain/Entities"
)

type TareaRepositoryMock struct {
	Tarea             *entities.Tarea
	ErrorObtenerPorId error
	ActualizarLlamado bool
}

var _ repositories.ITareaRepository = (*TareaRepositoryMock)(nil)

func (m *TareaRepositoryMock) Crear(
	context context.Context,
	tarea *entities.Tarea,
) error {
	return nil
}

func (m *TareaRepositoryMock) ObtenerPorId(
	context context.Context,
	id int64,
) (*entities.Tarea, error) {

	if m.ErrorObtenerPorId != nil {
		return nil, m.ErrorObtenerPorId
	}

	return m.Tarea, nil
}

func (m *TareaRepositoryMock) ObtenerTodos(
	context context.Context,
) ([]entities.Tarea, error) {
	return nil, nil
}

func (m *TareaRepositoryMock) Actualizar(
	context context.Context,
	tarea *entities.Tarea,
) error {
	m.ActualizarLlamado = true

	m.Tarea = tarea

	return nil
}

func (m *TareaRepositoryMock) Eliminar(
	context context.Context,
	id int64,
	modificadoPor int,
) error {
	return nil
}

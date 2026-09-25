package mocks

import (
	"context"
	"to-do-server/internal/Application/Interfaces/repositories"
	entities "to-do-server/internal/Domain/Entities"
)

type SubTareaRepositoryMock struct {
	SubTareas []entities.SubTarea
}

var _ repositories.ISubTareaRepository = (*SubTareaRepositoryMock)(nil)

func (m *SubTareaRepositoryMock) Crear(
	context context.Context,
	subTarea *entities.SubTarea,
) error {
	return nil
}

func (m *SubTareaRepositoryMock) ObtenerPorId(
	context context.Context,
	id int64,
) (*entities.SubTarea, error) {
	return nil, nil
}

func (m *SubTareaRepositoryMock) ObtenerPorTareaId(
	context context.Context,
	tareaId int64,
) ([]entities.SubTarea, error) {
	return m.SubTareas, nil
}

func (m *SubTareaRepositoryMock) Actualizar(
	context context.Context,
	subTarea *entities.SubTarea,
) error {
	return nil
}

func (m *SubTareaRepositoryMock) Eliminar(
	context context.Context,
	id int64,
	modificadoPor int,
) error {
	return nil
}

package mocks

import (
	"context"
	"to-do-server/internal/Application/Interfaces/repositories"
	entities "to-do-server/internal/Domain/Entities"
)

type TareaRepositoryMock struct {
	Tarea  *entities.Tarea
	Tareas []entities.Tarea

	CrearLlamado          bool
	ActualizarLlamado     bool
	ObtenerPorIdLlamado   bool
	EliminarFisicoLlamado bool

	ErrorCrear          error
	ErrorObtenerPorId   error
	ErrorActualizar     error
	ErrorEliminarFisico error
}

var _ repositories.ITareaRepository = (*TareaRepositoryMock)(nil)

func (m *TareaRepositoryMock) Crear(
	context context.Context,
	tarea *entities.Tarea,
) error {
	m.CrearLlamado = true

	if m.ErrorCrear != nil {
		return m.ErrorCrear
	}

	m.Tarea = tarea

	return nil
}

func (m *TareaRepositoryMock) ObtenerPorId(
	context context.Context,
	id int64,
) (*entities.Tarea, error) {

	m.ObtenerPorIdLlamado = true

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

	if m.ErrorActualizar != nil {
		return m.ErrorActualizar
	}

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

func (m *TareaRepositoryMock) EliminarFisicoPorId(
	context context.Context,
	id int64,
) error {
	m.EliminarFisicoLlamado = true

	if m.ErrorEliminarFisico != nil {
		return m.ErrorEliminarFisico
	}

	return nil
}

func (m *TareaRepositoryMock) ObtenerPorUsuarioId(
	context context.Context,
	usuarioId int,
) ([]entities.Tarea, error) {
	return m.Tareas, nil
}

func (r *TareaRepositoryMock) ObtenerEliminadas(
	context context.Context,
) ([]entities.Tarea, error) {
	return nil, nil
}

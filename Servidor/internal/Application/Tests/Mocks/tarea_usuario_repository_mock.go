package mocks

import (
	"context"
	"to-do-server/internal/Application/Interfaces/repositories"
	entities "to-do-server/internal/Domain/Entities"
)

type TareaUsuarioRepositoryMock struct {
	TareasUsuario []entities.TareaUsuario

	CrearLlamado    bool
	EliminarLlamado bool

	ErrorCrear error
}

var _ repositories.ITareaUsuarioRepository = (*TareaUsuarioRepositoryMock)(nil)

func (m *TareaUsuarioRepositoryMock) Crear(
	context context.Context,
	tareaUsuario *entities.TareaUsuario,
) error {

	m.CrearLlamado = true

	if m.ErrorCrear != nil {
		return m.ErrorCrear
	}

	m.TareasUsuario = append(
		m.TareasUsuario,
		*tareaUsuario,
	)

	return nil
}

func (m *TareaUsuarioRepositoryMock) ObtenerPorTareaId(
	context context.Context,
	tareaId int64,
) ([]entities.TareaUsuario, error) {

	var resultado []entities.TareaUsuario

	for _, tareaUsuario := range m.TareasUsuario {
		if tareaUsuario.TareaId == tareaId {
			resultado = append(resultado, tareaUsuario)
		}
	}

	return resultado, nil
}

func (m *TareaUsuarioRepositoryMock) ObtenerPorUsuarioId(
	context context.Context,
	usuarioId int,
) ([]entities.TareaUsuario, error) {

	var resultado []entities.TareaUsuario

	for _, tareaUsuario := range m.TareasUsuario {
		if tareaUsuario.UsuarioId == usuarioId {
			resultado = append(resultado, tareaUsuario)
		}
	}

	return resultado, nil
}

func (m *TareaUsuarioRepositoryMock) Eliminar(
	context context.Context,
	tareaId int64,
	usuarioId int,
) error {

	m.EliminarLlamado = true

	return nil
}

package mocks

import (
	"context"
	"to-do-server/internal/Application/Interfaces/repositories"
	entities "to-do-server/internal/Domain/Entities"
)

type UsuarioRepositoryMock struct {
	Usuarios []entities.Usuario
	Usuario  *entities.Usuario

	ObtenerPorIdLlamado bool
	ObtenerTodosLlamado bool

	ErrorObtenerPorId error
	ErrorObtenerTodos error
}

var _ repositories.IUsuarioRepository = (*UsuarioRepositoryMock)(nil)

func (m *UsuarioRepositoryMock) ObtenerPorId(
	context context.Context,
	id int,
) (*entities.Usuario, error) {

	m.ObtenerPorIdLlamado = true

	if m.ErrorObtenerPorId != nil {
		return nil, m.ErrorObtenerPorId
	}

	if m.Usuario != nil {
		return m.Usuario, nil
	}

	for _, usuario := range m.Usuarios {
		if usuario.Id == id {
			return &usuario, nil
		}
	}

	return nil, nil
}

func (m *UsuarioRepositoryMock) ObtenerTodos(
	context context.Context,
) ([]entities.Usuario, error) {

	m.ObtenerTodosLlamado = true

	if m.ErrorObtenerTodos != nil {
		return nil, m.ErrorObtenerTodos
	}

	return m.Usuarios, nil
}

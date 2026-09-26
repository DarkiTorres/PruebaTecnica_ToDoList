package services

import (
	"context"
	"errors"
	"to-do-server/internal/Application/Interfaces/repositories"
	interfaces "to-do-server/internal/Application/Interfaces/services"
	entities "to-do-server/internal/Domain/Entities"
)

var _ interfaces.IUsuarioService = (*UsuarioService)(nil)

type UsuarioService struct {
	usuarioRepository repositories.IUsuarioRepository
}

func NewUsuarioService(
	usuarioRepository repositories.IUsuarioRepository,
) *UsuarioService {
	return &UsuarioService{
		usuarioRepository: usuarioRepository,
	}
}

func (s *UsuarioService) ObtenerUsuarioPorId(
	context context.Context,
	usuarioId int,
) (*entities.Usuario, error) {

	if usuarioId <= 0 {
		return nil, errors.New(
			"el id del usuario no es válido",
		)
	}

	usuario, err := s.usuarioRepository.ObtenerPorId(
		context,
		usuarioId,
	)

	if err != nil {
		return nil, err
	}

	if usuario == nil {
		return nil, errors.New(
			"el usuario no existe",
		)
	}

	return usuario, nil
}

func (s *UsuarioService) ObtenerUsuarios(
	context context.Context,
) ([]entities.Usuario, error) {

	return s.usuarioRepository.ObtenerTodos(context)
}

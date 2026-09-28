package services

import (
	"context"
	"to-do-server/internal/Application/Interfaces/repositories"
	serv "to-do-server/internal/Application/Interfaces/services"
	entities "to-do-server/internal/Domain/Entities"
)

type TareaUsuarioService struct {
	repository repositories.ITareaUsuarioRepository
}

var _ serv.ITareaUsuarioService = (*TareaUsuarioService)(nil)

func NewTareaUsuarioService(
	repository repositories.ITareaUsuarioRepository,
) *TareaUsuarioService {
	return &TareaUsuarioService{
		repository: repository,
	}
}

func (s *TareaUsuarioService) Crear(
	context context.Context,
	tareaUsuario *entities.TareaUsuario,
) error {
	return s.repository.Crear(context, tareaUsuario)
}

func (s *TareaUsuarioService) ObtenerPorTareaId(
	context context.Context,
	tareaId int64,
) ([]entities.TareaUsuario, error) {
	return s.repository.ObtenerPorTareaId(context, tareaId)
}

func (s *TareaUsuarioService) ObtenerPorUsuarioId(
	context context.Context,
	usuarioId int,
) ([]entities.TareaUsuario, error) {
	return s.repository.ObtenerPorUsuarioId(context, usuarioId)
}

func (s *TareaUsuarioService) Eliminar(
	context context.Context,
	tareaId int64,
	usuarioId int,
) error {
	return s.repository.Eliminar(context, tareaId, usuarioId)
}

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
	tareaRepository        repositories.ITareaRepository
	subTareaRepository     repositories.ISubTareaRepository
	usuarioRepository      repositories.IUsuarioRepository
	tareaUsuarioRepository repositories.ITareaUsuarioRepository
}

func NewTareaService(
	tareaRepository repositories.ITareaRepository,
	subTareaRepository repositories.ISubTareaRepository,
	usuarioRepository repositories.IUsuarioRepository,
	tareaUsuarioRepository repositories.ITareaUsuarioRepository,
) *TareaService {
	return &TareaService{
		tareaRepository:        tareaRepository,
		subTareaRepository:     subTareaRepository,
		usuarioRepository:      usuarioRepository,
		tareaUsuarioRepository: tareaUsuarioRepository,
	}
}

func (s *TareaService) CrearTarea(
	context context.Context,
	tarea *entities.Tarea,
	asignadoA int,
) error {

	if err := validators.ValidarTarea(tarea); err != nil {
		return err
	}

	if asignadoA <= 0 {
		return errors.New("El usuario asignado no es valido.")
	}

	creador, err := s.usuarioRepository.ObtenerPorId(
		context,
		tarea.CreadoPor,
	)

	if err != nil {
		return err
	}

	if creador == nil {
		return errors.New("El usuario creador no existe")
	}

	if creador.EstaDesactivado {
		return errors.New("El usuario creador esta desactivado.")
	}

	asignado, err := s.usuarioRepository.ObtenerPorId(
		context,
		asignadoA,
	)

	if err != nil {
		return err
	}

	if asignado == nil {
		return errors.New("El usuario asignado no existe")
	}

	if asignado.EstaDesactivado {
		return errors.New("El usuario asignado esta desactivado.")
	}

	if creador.RolId == 2 && creador.Id != asignado.Id {
		return errors.New(
			"Un contribuidor solamente puede asignarse tareas a si mismo",
		)
	}

	if err := s.tareaRepository.Crear(
		context,
		tarea,
	); err != nil {
		return err
	}

	tareaUsuario := &entities.TareaUsuario{
		TareaId:   tarea.Id,
		UsuarioId: asignado.Id,
	}

	return s.tareaUsuarioRepository.Crear(
		context,
		tareaUsuario,
	)
}

func (s *TareaService) CompletarTarea(context context.Context, tareaId int64, estaTerminada bool) error {

	if tareaId <= 0 {
		return errors.New("el id de la tarea no es válido")
	}

	tarea, err := s.tareaRepository.ObtenerPorId(context, tareaId)
	if err != nil {
		return err
	}

	if tarea == nil {
		return errors.New("la tarea no existe")
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

	tarea.EstaTerminada = estaTerminada

	if estaTerminada && tarea.FechaEntrega != nil && tarea.FechaEntrega.Before(time.Now()) {
		tarea.EstaEliminada = true
	}

	return s.tareaRepository.Actualizar(
		context, tarea,
	)
}

func (s *TareaService) ActualizarTarea(context context.Context, tarea *entities.Tarea, asignadoA int) error {
	if tarea == nil {
		return errors.New("la tarea no puede ser nula")
	}

	if tarea.Id <= 0 {
		return errors.New("el id de la tarea no es válido")
	}

	if asignadoA <= 0 {
		return errors.New("El usuario asignado no es valido.")
	}

	if err := validators.ValidarTarea(tarea); err != nil {
		return err
	}

	asignado, err := s.usuarioRepository.ObtenerPorId(
		context, asignadoA,
	)

	if err != nil {
		return err
	}

	if asignado == nil {
		return errors.New("El usuario asignado no existe.")
	}

	if asignado.EstaDesactivado {
		return errors.New("El usuario asignado esta desactivado.")
	}

	if err := s.tareaRepository.Actualizar(context, tarea); err != nil {
		return err
	}

	relaciones, err := s.tareaUsuarioRepository.ObtenerPorTareaId(
		context,
		tarea.Id,
	)

	if err != nil {
		return err
	}

	for _, relacion := range relaciones {
		if err := s.tareaUsuarioRepository.Eliminar(
			context,
			tarea.Id,
			relacion.UsuarioId,
		); err != nil {
			return err
		}
	}

	return s.tareaUsuarioRepository.Crear(
		context,
		&entities.TareaUsuario{
			TareaId:   tarea.Id,
			UsuarioId: asignadoA,
		},
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

	if tarea == nil {
		return errors.New("La tarea no existe")
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

func (s *TareaService) ObtenerTareaPorId(
	context context.Context,
	tareaId int64,
) (*entities.Tarea, error) {

	if tareaId <= 0 {
		return nil, errors.New("el id de la tarea no es válido")
	}

	tarea, err := s.tareaRepository.ObtenerPorId(
		context,
		tareaId,
	)

	if err != nil {
		return nil, err
	}

	if tarea == nil {
		return nil, errors.New("la tarea no existe")
	}

	return tarea, nil
}

func (s *TareaService) ObtenerTareas(
	context context.Context,
) ([]entities.Tarea, error) {

	return s.tareaRepository.ObtenerTodos(context)
}

func (s *TareaService) ObtenerTareasPorUsuario(
	context context.Context,
	usuarioId int,
) ([]entities.Tarea, error) {

	if usuarioId <= 0 {
		return nil, errors.New("el id del usuario no es válido")
	}

	usuario, err := s.usuarioRepository.ObtenerPorId(
		context,
		usuarioId,
	)

	if err != nil {
		return nil, err
	}

	if usuario == nil {
		return nil, errors.New("el usuario no existe")
	}

	if usuario.EstaDesactivado {
		return nil, errors.New("el usuario está desactivado")
	}

	return s.tareaRepository.ObtenerPorUsuarioId(
		context,
		usuarioId,
	)
}

func (s *TareaService) ObtenerBitacora(context context.Context) ([]entities.Tarea, error) {

	tareasActivas, err := s.tareaRepository.ObtenerTodos(context)

	if err != nil {
		return nil, err
	}

	tareasEliminadas, err := s.tareaRepository.ObtenerEliminadas(context)

	if err != nil {
		return nil, err
	}

	var bitacora []entities.Tarea

	for _, tarea := range tareasActivas {
		if tarea.EstaTerminada {
			bitacora = append(bitacora, tarea)
		}
	}

	bitacora = append(bitacora, tareasEliminadas...)

	return bitacora, nil
}

func (s *TareaService) RestaurarTarea(
	context context.Context,
	tareaId int64,
	modificadoPor int,
) error {

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

	if tarea == nil {
		return errors.New("la tarea no existe")
	}

	if !tarea.EstaEliminada {
		return errors.New("la tarea no está eliminada")
	}

	usuario, err := s.usuarioRepository.ObtenerPorId(
		context,
		modificadoPor,
	)

	if err != nil {
		return err
	}

	if usuario == nil {
		return errors.New("el usuario no existe")
	}

	if usuario.EstaDesactivado {
		return errors.New("el usuario está desactivado")
	}

	ahora := time.Now()

	tarea.EstaTerminada = false
	tarea.EstaEliminada = false
	tarea.ModificadoEl = &ahora
	tarea.ModificadoPor = &modificadoPor

	return s.tareaRepository.Actualizar(
		context,
		tarea,
	)
}

func (s *TareaService) EliminarFisicamente(
	context context.Context,
	tareaId int64,
) error {

	if tareaId <= 0 {
		return errors.New("el id de la tarea no es válido")
	}

	tarea, err := s.tareaRepository.ObtenerPorId(
		context,
		tareaId,
	)

	if err != nil {
		return err
	}

	if tarea == nil {
		return errors.New("la tarea no existe")
	}

	if !tarea.EstaEliminada {
		return errors.New(
			"No se puede eliminar físicamente una tarea activa.",
		)
	}

	return s.tareaRepository.EliminarFisicoPorId(
		context,
		tareaId,
	)
}

package postgres

import (
	"context"
	"to-do-server/internal/Application/Interfaces/repositories"
	entities "to-do-server/internal/Domain/Entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TareaRepositoryPostgres struct {
	db *pgxpool.Pool
}

var _ repositories.ITareaRepository = (*TareaRepositoryPostgres)(nil)

func NewTareaRepository(db *pgxpool.Pool) *TareaRepositoryPostgres {
	return &TareaRepositoryPostgres{
		db: db,
	}
}

func (r *TareaRepositoryPostgres) Crear(context context.Context, tarea *entities.Tarea) error {
	query := `
		INSERT INTO Tareas (
			Titulo,
			Descripcion,
			PrioridadId,
			FechaEntrega,
			CreadoEl,
			CreadoPor
		)
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING Id
	`

	return r.db.QueryRow(
		context,
		query,
		tarea.Titulo,
		tarea.Descripcion,
		tarea.PrioridadId,
		tarea.FechaEntrega,
		tarea.CreadoEl,
		tarea.CreadoPor,
	).Scan(&tarea.Id)
}

func (r *TareaRepositoryPostgres) ObtenerPorId(context context.Context, id int64) (*entities.Tarea, error) {
	query := `
		SELECT
			Id,
			Titulo,
			Descripcion,
			PrioridadId,
			FechaEntrega,
			EstaTerminado,
			EstaEliminado,
			CreadoEl,
			CreadoPor,
			ModificadoEl,
			ModificadoPor
		From Tareas
		WHERE Id = $1
	`

	tarea := &entities.Tarea{}

	err := r.db.QueryRow(
		context, query, id,
	).Scan(
		&tarea.Id,
		&tarea.Titulo,
		&tarea.Descripcion,
		&tarea.PrioridadId,
		&tarea.FechaEntrega,
		&tarea.EstaTerminada,
		&tarea.EstaEliminada,
		&tarea.CreadoEl,
		&tarea.CreadoPor,
		&tarea.ModificadoEl,
		&tarea.ModificadoPor,
	)

	if err != nil {
		return nil, err
	}

	return tarea, nil
}

func (r *TareaRepositoryPostgres) ObtenerTodos(context context.Context) ([]entities.Tarea, error) {
	query := `
		SELECT
			Id,
			Titulo,
			Descripcion,
			PrioridadId,
			FechaEntrega,
			EstaTerminado,
			EstaEliminado,
			CreadoEl,
			CreadoPor,
			ModificadoEl,
			ModificadoPor
		FROM Tareas
		WHERE EstaEliminado = FALSE
		ORDER BY Id
	`

	rows, err := r.db.Query(context, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tareas []entities.Tarea

	for rows.Next() {
		var tarea entities.Tarea

		err := rows.Scan(
			&tarea.Id,
			&tarea.Titulo,
			&tarea.Descripcion,
			&tarea.PrioridadId,
			&tarea.FechaEntrega,
			&tarea.EstaTerminada,
			&tarea.EstaEliminada,
			&tarea.CreadoEl,
			&tarea.CreadoPor,
			&tarea.ModificadoEl,
			&tarea.ModificadoPor,
		)

		if err != nil {
			return nil, err
		}

		tareas = append(tareas, tarea)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tareas, nil
}

func (r *TareaRepositoryPostgres) Actualizar(context context.Context, tarea *entities.Tarea) error {
	query := `
		UPDATE Tareas
		SET
			Titulo = $1,
			Descripcion = $2,
			PrioridadId = $3,
			FechaEntrega = $4,
			EstaTerminado = $5,
			EstaEliminado = $6,
			ModificadoEl = $7,
			ModificadoPor = $8
		WHERE Id = $9
	`

	_, err := r.db.Exec(
		context,
		query,
		tarea.Titulo,
		tarea.Descripcion,
		tarea.PrioridadId,
		tarea.FechaEntrega,
		tarea.EstaTerminada,
		tarea.EstaEliminada,
		tarea.ModificadoEl,
		tarea.ModificadoPor,
		tarea.Id,
	)

	return err
}

func (r *TareaRepositoryPostgres) Eliminar(context context.Context, id int64, modificadoPor int) error {
	query := `
		UPDATE Tareas
		SET
			EstaEliminado = True,
			ModificadoEl = CURRENT_TIMESTAMP,
			ModificadoPor = $2
		WHERE Id = $1
	`

	_, err := r.db.Exec(
		context, query, id, modificadoPor,
	)

	return err
}

func (r *TareaRepositoryPostgres) EliminarFisicoPorId(
	context context.Context,
	id int64,
) error {

	query := `
		DELETE FROM Tareas
		WHERE Id = $1
	`

	_, err := r.db.Exec(
		context,
		query,
		id,
	)

	return err
}

func (r *TareaRepositoryPostgres) ObtenerPorUsuarioId(
	context context.Context,
	usuarioId int,
) ([]entities.Tarea, error) {

	query := `
		SELECT
			t.Id,
			t.Titulo,
			t.Descripcion,
			t.PrioridadId,
			t.FechaEntrega,
			t.EstaTerminado,
			t.EstaEliminado,
			t.CreadoEl,
			t.CreadoPor,
			t.ModificadoEl,
			t.ModificadoPor
		FROM Tareas t
		INNER JOIN TareaXUsuario txu
			ON txu.TareaId = t.Id
		WHERE txu.UsuarioId = $1
			AND t.EstaEliminado = FALSE
		ORDER BY t.Id
	`

	rows, err := r.db.Query(
		context,
		query,
		usuarioId,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tareas []entities.Tarea

	for rows.Next() {
		var tarea entities.Tarea

		err := rows.Scan(
			&tarea.Id,
			&tarea.Titulo,
			&tarea.Descripcion,
			&tarea.PrioridadId,
			&tarea.FechaEntrega,
			&tarea.EstaTerminada,
			&tarea.EstaEliminada,
			&tarea.CreadoEl,
			&tarea.CreadoPor,
			&tarea.ModificadoEl,
			&tarea.ModificadoPor,
		)

		if err != nil {
			return nil, err
		}

		tareas = append(tareas, tarea)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tareas, nil
}

func (r *TareaRepositoryPostgres) ObtenerEliminadas(
	context context.Context,
) ([]entities.Tarea, error) {

	query := `
		SELECT
			Id,
			Titulo,
			Descripcion,
			PrioridadId,
			FechaEntrega,
			EstaTerminado,
			EstaEliminado,
			CreadoEl,
			CreadoPor,
			ModificadoEl,
			ModificadoPor
		FROM Tareas
		WHERE EstaEliminado = TRUE
		ORDER BY Id
	`

	rows, err := r.db.Query(context, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tareas []entities.Tarea

	for rows.Next() {
		var tarea entities.Tarea

		err := rows.Scan(
			&tarea.Id,
			&tarea.Titulo,
			&tarea.Descripcion,
			&tarea.PrioridadId,
			&tarea.FechaEntrega,
			&tarea.EstaTerminada,
			&tarea.EstaEliminada,
			&tarea.CreadoEl,
			&tarea.CreadoPor,
			&tarea.ModificadoEl,
			&tarea.ModificadoPor,
		)

		if err != nil {
			return nil, err
		}

		tareas = append(tareas, tarea)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tareas, nil
}

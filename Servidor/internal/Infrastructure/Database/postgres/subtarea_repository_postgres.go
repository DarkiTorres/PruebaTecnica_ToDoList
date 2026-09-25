package postgres

import (
	"context"
	"to-do-server/internal/Application/Interfaces/repositories"
	entities "to-do-server/internal/Domain/Entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SubTareaRepositoryPostgres struct {
	db *pgxpool.Pool
}

func NewSubTareaRepository(db *pgxpool.Pool) *SubTareaRepositoryPostgres {
	return &SubTareaRepositoryPostgres{
		db: db,
	}
}

var _ repositories.ISubTareaRepository = (*SubTareaRepositoryPostgres)(nil)

func (r *SubTareaRepositoryPostgres) Crear(context context.Context, subTarea *entities.SubTarea) error {
	query := `
		INSERT INTO SubTareas (
			TareaId,
			Titulo,
			CreadoEl,
			CreadoPor
		)
		VALUES ($1, $2, $3, $4)
		RETURNING Id
	`

	return r.db.QueryRow(
		context,
		query,
		subTarea.TareaId,
		subTarea.Titulo,
		subTarea.CreadoEl,
		subTarea.CreadoPor,
	).Scan(&subTarea.Id)
}

func (r *SubTareaRepositoryPostgres) ObtenerPorId(context context.Context, id int64) (*entities.SubTarea, error) {
	query := `
		SELECT
			Id,
			TareaId,
			Titulo,
			EstaTerminado,
			EstaEliminado,
			CreadoEl,
			CreadoPor,
			ModificadoEl,
			ModificadoPor
		FROM SubTareas
		WHERE Id = $1
	`

	subTarea := &entities.SubTarea{}

	err := r.db.QueryRow(
		context,
		query,
		id,
	).Scan(
		&subTarea.Id,
		&subTarea.TareaId,
		&subTarea.Titulo,
		&subTarea.EstaTerminada,
		&subTarea.EstaEliminada,
		&subTarea.CreadoEl,
		&subTarea.CreadoPor,
		&subTarea.ModificadoEl,
		&subTarea.ModificadoPor,
	)

	if err != nil {
		return nil, err
	}

	return subTarea, nil
}

func (r *SubTareaRepositoryPostgres) ObtenerPorTareaId(context context.Context, tareaId int64) ([]entities.SubTarea, error) {
	query := `
		SELECT
			Id,
			TareaId,
			Titulo,
			EstaTerminado,
			EstaEliminado,
			CreadoEl,
			CreadoPor,
			ModificadoEl,
			ModificadoPor
		FROM SubTareas
		WHERE TareaId = $1
			AND EstaEliminado = FALSE
		ORDER BY Id
	`

	rows, err := r.db.Query(
		context,
		query,
		tareaId,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var subTareas []entities.SubTarea

	for rows.Next() {
		var subTarea entities.SubTarea

		err := rows.Scan(
			&subTarea.Id,
			&subTarea.TareaId,
			&subTarea.Titulo,
			&subTarea.EstaTerminada,
			&subTarea.EstaEliminada,
			&subTarea.CreadoEl,
			&subTarea.CreadoPor,
			&subTarea.ModificadoEl,
			&subTarea.ModificadoPor,
		)

		if err != nil {
			return nil, err
		}

		subTareas = append(subTareas, subTarea)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return subTareas, nil
}

func (r *SubTareaRepositoryPostgres) Actualizar(context context.Context, subTarea *entities.SubTarea) error {
	query := `
		UPDATE SubTareas
		SET
			Titulo = $1,
			EstaTerminado = $2,
			EstaEliminado = $3,
			ModificadoEl = $4,
			ModificadoPor = $5
		WHERE Id = $6
	`

	_, err := r.db.Exec(
		context,
		query,
		subTarea.Titulo,
		subTarea.EstaTerminada,
		subTarea.EstaEliminada,
		subTarea.ModificadoEl,
		subTarea.ModificadoPor,
		subTarea.Id,
	)

	return err
}

func (r *SubTareaRepositoryPostgres) Eliminar(context context.Context, id int64, modificadoPor int) error {
	query := `
		UPDATE SubTareas
		SET
			EstaEliminado = TRUE,
			ModificadoEl = CURRENT_TIMESTAMP,
			ModificadoPor = $2
		WHERE Id = $1
	`

	_, err := r.db.Exec(
		context,
		query,
		id,
		modificadoPor,
	)

	return err
}

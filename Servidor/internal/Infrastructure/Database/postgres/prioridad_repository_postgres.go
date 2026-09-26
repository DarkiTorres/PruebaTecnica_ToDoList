package postgres

import (
	"context"

	"to-do-server/internal/Application/Interfaces/repositories"
	entities "to-do-server/internal/Domain/Entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PrioridadRepositoryPostgres struct {
	db *pgxpool.Pool
}

var _ repositories.IPrioridadRepository = (*PrioridadRepositoryPostgres)(nil)

func NewPrioridadRepository(
	db *pgxpool.Pool,
) *PrioridadRepositoryPostgres {
	return &PrioridadRepositoryPostgres{
		db: db,
	}
}

func (r *PrioridadRepositoryPostgres) ObtenerTodos(
	context context.Context,
) ([]entities.Prioridad, error) {

	query := `
		SELECT
			Id,
			Descripcion
		FROM Prioridades
		ORDER BY Id
	`

	rows, err := r.db.Query(
		context,
		query,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var prioridades []entities.Prioridad

	for rows.Next() {

		var prioridad entities.Prioridad

		err := rows.Scan(
			&prioridad.Id,
			&prioridad.Descripcion,
		)

		if err != nil {
			return nil, err
		}

		prioridades = append(
			prioridades,
			prioridad,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return prioridades, nil
}

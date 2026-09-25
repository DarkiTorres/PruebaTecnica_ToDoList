package postgres

import (
	"context"
	"to-do-server/internal/Application/Interfaces/repositories"
	entities "to-do-server/internal/Domain/Entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TareaUsuarioRepositoryPostgres struct {
	db *pgxpool.Pool
}

var _ repositories.ITareaUsuarioRepository = (*TareaUsuarioRepositoryPostgres)(nil)

func NewTareaUsuarioRepository(
	db *pgxpool.Pool,
) *TareaUsuarioRepositoryPostgres {
	return &TareaUsuarioRepositoryPostgres{
		db: db,
	}
}

func (r *TareaUsuarioRepositoryPostgres) Crear(
	context context.Context,
	tareaUsuario *entities.TareaUsuario,
) error {

	query := `
		INSERT INTO TareaXUsuario (
			TareaId,
			UsuarioId
		)
		VALUES ($1, $2)
	`

	_, err := r.db.Exec(
		context,
		query,
		tareaUsuario.TareaId,
		tareaUsuario.UsuarioId,
	)

	return err
}

func (r *TareaUsuarioRepositoryPostgres) ObtenerPorTareaId(
	context context.Context,
	tareaId int64,
) ([]entities.TareaUsuario, error) {

	query := `
		SELECT
			TareaId,
			UsuarioId
		FROM TareaXUsuario
		WHERE TareaId = $1
		ORDER BY UsuarioId
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

	var tareasUsuario []entities.TareaUsuario

	for rows.Next() {
		var tareaUsuario entities.TareaUsuario

		err := rows.Scan(
			&tareaUsuario.TareaId,
			&tareaUsuario.UsuarioId,
		)

		if err != nil {
			return nil, err
		}

		tareasUsuario = append(
			tareasUsuario,
			tareaUsuario,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tareasUsuario, nil
}

func (r *TareaUsuarioRepositoryPostgres) ObtenerPorUsuarioId(
	context context.Context,
	usuarioId int,
) ([]entities.TareaUsuario, error) {

	query := `
		SELECT
			TareaId,
			UsuarioId
		FROM TareaXUsuario
		WHERE UsuarioId = $1
		ORDER BY TareaId
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

	var tareasUsuario []entities.TareaUsuario

	for rows.Next() {
		var tareaUsuario entities.TareaUsuario

		err := rows.Scan(
			&tareaUsuario.TareaId,
			&tareaUsuario.UsuarioId,
		)

		if err != nil {
			return nil, err
		}

		tareasUsuario = append(
			tareasUsuario,
			tareaUsuario,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tareasUsuario, nil
}

func (r *TareaUsuarioRepositoryPostgres) Eliminar(
	context context.Context,
	tareaId int64,
	usuarioId int,
) error {

	query := `
		DELETE FROM TareaXUsuario
		WHERE TareaId = $1
			AND UsuarioId = $2
	`

	_, err := r.db.Exec(
		context,
		query,
		tareaId,
		usuarioId,
	)

	return err
}

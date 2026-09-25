package postgres

import (
	"context"
	"to-do-server/internal/Application/Interfaces/repositories"
	entities "to-do-server/internal/Domain/Entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UsuarioRepositoryPostgres struct {
	db *pgxpool.Pool
}

var _ repositories.IUsuarioRepository = (*UsuarioRepositoryPostgres)(nil)

func NewUsuarioRepository(db *pgxpool.Pool) *UsuarioRepositoryPostgres {
	return &UsuarioRepositoryPostgres{
		db: db,
	}
}

func (r *UsuarioRepositoryPostgres) ObtenerTodos(context context.Context) ([]entities.Usuario, error) {
	query := `
		SELECT
			Id,
			Nombre,
			RolId,
			EstaDesactivado
		FROM Usuarios
		ORDER BY Id
	`

	rows, err := r.db.Query(context, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var usuarios []entities.Usuario

	for rows.Next() {
		var usuario entities.Usuario

		err := rows.Scan(
			&usuario.Id,
			&usuario.Nombre,
			&usuario.RolId,
			&usuario.EstaDesactivado,
		)

		if err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return usuarios, nil
}

func (r *UsuarioRepositoryPostgres) ObtenerPorId(
	context context.Context,
	id int,
) (*entities.Usuario, error) {

	query := `
		SELECT
			Id,
			Nombre,
			RolId,
			EstaDesactivado
		FROM Usuarios
		WHERE Id = $1
	`

	usuario := &entities.Usuario{}

	err := r.db.QueryRow(
		context,
		query,
		id,
	).Scan(
		&usuario.Id,
		&usuario.Nombre,
		&usuario.RolId,
		&usuario.EstaDesactivado,
	)

	if err != nil {
		return nil, err
	}

	return usuario, nil
}

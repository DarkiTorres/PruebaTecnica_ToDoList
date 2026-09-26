package initialization

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var databaseSQL string

type DatabaseInitializer struct {
	db *pgxpool.Pool
}

func NewDatabaseInitializer(db *pgxpool.Pool) *DatabaseInitializer {
	return &DatabaseInitializer{
		db: db,
	}
}

func (d *DatabaseInitializer) Initialize(
	context context.Context,
) error {

	_, err := d.db.Exec(context, databaseSQL)

	if err != nil {
		return fmt.Errorf(
			"no se pudo inicializar la base de datos: %w",
			err,
		)
	}

	return nil
}

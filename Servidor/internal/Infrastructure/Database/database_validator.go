package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RequiredTable struct {
	Name    string
	Columns []string
}

var requiredTables = []RequiredTable{
	{
		Name: "prioridades",
		Columns: []string{
			"id",
			"descripcion",
		},
	},
	{
		Name: "roles",
		Columns: []string{
			"id",
			"descripcion",
		},
	},
	{
		Name: "usuarios",
		Columns: []string{
			"id",
			"nombre",
			"rolid",
			"estadesactivado",
		},
	},
	{
		Name: "tareas",
		Columns: []string{
			"id",
			"titulo",
			"descripcion",
			"prioridadid",
			"fechaentrega",
			"estaterminado",
			"estaeliminado",
			"creadoel",
			"creadopor",
			"modificadoel",
			"modificadopor",
		},
	},
	{
		Name: "subtareas",
		Columns: []string{
			"id",
			"tareaid",
			"titulo",
			"estaterminado",
			"estaeliminado",
			"creadoel",
			"creadopor",
			"modificadoel",
			"modificadopor",
		},
	},
	{
		Name: "tareaXusuario",
		Columns: []string{
			"tareaid",
			"usuarioid",
		},
	},
}

func ValidarBaseDeDatos(
	ctx context.Context,
	db *pgxpool.Pool,
) error {

	if err := db.Ping(ctx); err != nil {
		return fmt.Errorf(
			"no se pudo conectar con PostgreSQL: %w",
			err,
		)
	}

	for _, table := range requiredTables {

		existe, err := tableExists(
			ctx,
			db,
			table.Name,
		)

		if err != nil {
			return fmt.Errorf(
				"error verificando la tabla %s: %w",
				table.Name,
				err,
			)
		}

		if !existe {
			return fmt.Errorf(
				"la tabla requerida '%s' no existe",
				table.Name,
			)
		}

		for _, column := range table.Columns {

			existe, err := columnExists(
				ctx,
				db,
				table.Name,
				column,
			)

			if err != nil {
				return fmt.Errorf(
					"error verificando la columna '%s.%s': %w",
					table.Name,
					column,
					err,
				)
			}

			if !existe {
				return fmt.Errorf(
					"la columna requerida '%s.%s' no existe",
					table.Name,
					column,
				)
			}
		}
	}

	return nil
}

func tableExists(
	ctx context.Context,
	db *pgxpool.Pool,
	tableName string,
) (bool, error) {

	var exists bool

	err := db.QueryRow(
		ctx,
		`
		SELECT EXISTS (
			SELECT 1
			FROM information_schema.tables
			WHERE table_schema = 'public'
			AND LOWER(table_name) = LOWER($1)
		)
		`,
		tableName,
	).Scan(&exists)

	return exists, err
}

func columnExists(
	ctx context.Context,
	db *pgxpool.Pool,
	tableName string,
	columnName string,
) (bool, error) {

	var exists bool

	err := db.QueryRow(
		ctx,
		`
		SELECT EXISTS (
			SELECT 1
			FROM information_schema.columns
			WHERE table_schema = 'public'
			AND LOWER(table_name) = LOWER($1)
			AND LOWER(column_name) = LOWER($2)
		)
		`,
		tableName,
		columnName,
	).Scan(&exists)

	return exists, err
}

package postgres

import (
	"context"
	"testing"
)

func TestUsuarioRepositoryPostgres_ObtenerPorId(t *testing.T) {
	pool := crearPoolPrueba(t)

	repository := NewUsuarioRepository(pool)

	usuario, err := repository.ObtenerPorId(
		context.Background(),
		1,
	)

	if err != nil {
		t.Fatalf(
			"No se pudo obtener el usuario: %v",
			err,
		)
	}

	if usuario == nil {
		t.Fatal("Se esperaba un usuario")
	}

	if usuario.Id != 1 {
		t.Fatalf(
			"Se esperaba Id 1, se obtuvo %d",
			usuario.Id,
		)
	}

	if usuario.RolId != 1 {
		t.Fatalf(
			"Se esperaba RolId 1, se obtuvo %d",
			usuario.RolId,
		)
	}
}

func TestUsuarioRepositoryPostgres_ObtenerTodos(t *testing.T) {
	pool := crearPoolPrueba(t)

	repository := NewUsuarioRepository(pool)

	usuarios, err := repository.ObtenerTodos(
		context.Background(),
	)

	if err != nil {
		t.Fatalf(
			"No se pudieron obtener los usuarios: %v",
			err,
		)
	}

	if len(usuarios) != 2 {
		t.Fatalf(
			"Se esperaban 2 usuarios, se obtuvieron %d",
			len(usuarios),
		)
	}

	if usuarios[0].Id != 1 {
		t.Fatalf(
			"Se esperaba que el primer usuario tuviera Id 1, se obtuvo %d",
			usuarios[0].Id,
		)
	}

	if usuarios[1].Id != 2 {
		t.Fatalf(
			"Se esperaba que el segundo usuario tuviera Id 2, se obtuvo %d",
			usuarios[1].Id,
		)
	}
}

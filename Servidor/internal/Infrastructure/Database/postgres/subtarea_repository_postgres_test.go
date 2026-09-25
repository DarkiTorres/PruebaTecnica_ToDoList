package postgres

import (
	"context"
	"testing"
	"time"
	entities "to-do-server/internal/Domain/Entities"
)

func TestSubTareaRepositoryPostgres_Crear(t *testing.T) {
	pool := crearPoolPrueba(t)

	repository := NewSubTareaRepository(pool)

	subTarea := &entities.SubTarea{
		TareaId:   1,
		Titulo:    "Subtarea de prueba",
		CreadoEl:  time.Now(),
		CreadoPor: 1,
	}

	err := repository.Crear(
		context.Background(),
		subTarea,
	)

	if err != nil {
		t.Fatalf(
			"No se pudo crear la subtarea: %v",
			err,
		)
	}

	if subTarea.Id == 0 {
		t.Fatal("Se esperaba que PostgreSQL asignara un Id")
	}

	t.Logf(
		"Subtarea creada correctamente con Id: %d",
		subTarea.Id,
	)
}

func TestSubTareaRepositoryPostgres_ObtenerPorId(t *testing.T) {
	pool := crearPoolPrueba(t)

	repository := NewSubTareaRepository(pool)

	subTarea := &entities.SubTarea{
		TareaId:   1,
		Titulo:    "Subtarea para obtener",
		CreadoEl:  time.Now(),
		CreadoPor: 1,
	}

	err := repository.Crear(
		context.Background(),
		subTarea,
	)

	if err != nil {
		t.Fatalf(
			"No se pudo crear la subtarea: %v",
			err,
		)
	}

	obtenida, err := repository.ObtenerPorId(
		context.Background(),
		subTarea.Id,
	)

	if err != nil {
		t.Fatalf(
			"No se pudo obtener la subtarea: %v",
			err,
		)
	}

	if obtenida == nil {
		t.Fatal("Se esperaba una subtarea")
	}

	if obtenida.Id != subTarea.Id {
		t.Fatalf(
			"Se esperaba Id %d, se obtuvo %d",
			subTarea.Id,
			obtenida.Id,
		)
	}

	if obtenida.TareaId != subTarea.TareaId {
		t.Fatalf(
			"Se esperaba TareaId %d, se obtuvo %d",
			subTarea.TareaId,
			obtenida.TareaId,
		)
	}

	if obtenida.Titulo != subTarea.Titulo {
		t.Fatalf(
			"Se esperaba título %q, se obtuvo %q",
			subTarea.Titulo,
			obtenida.Titulo,
		)
	}
}

func TestSubTareaRepositoryPostgres_ObtenerPorTareaId(t *testing.T) {
	pool := crearPoolPrueba(t)

	repository := NewSubTareaRepository(pool)

	subTareaActiva := &entities.SubTarea{
		TareaId:   1,
		Titulo:    "Subtarea activa",
		CreadoEl:  time.Now(),
		CreadoPor: 1,
	}

	err := repository.Crear(
		context.Background(),
		subTareaActiva,
	)

	if err != nil {
		t.Fatalf(
			"No se pudo crear la subtarea activa: %v",
			err,
		)
	}

	subTareaEliminada := &entities.SubTarea{
		TareaId:   1,
		Titulo:    "Subtarea eliminada",
		CreadoEl:  time.Now(),
		CreadoPor: 1,
	}

	err = repository.Crear(
		context.Background(),
		subTareaEliminada,
	)

	if err != nil {
		t.Fatalf(
			"No se pudo crear la subtarea eliminada: %v",
			err,
		)
	}

	err = repository.Eliminar(
		context.Background(),
		subTareaEliminada.Id,
		1,
	)

	if err != nil {
		t.Fatalf(
			"No se pudo eliminar la subtarea: %v",
			err,
		)
	}

	subTareas, err := repository.ObtenerPorTareaId(
		context.Background(),
		1,
	)

	if err != nil {
		t.Fatalf(
			"No se pudieron obtener las subtareas: %v",
			err,
		)
	}

	encontradaActiva := false
	encontradaEliminada := false

	for _, subTarea := range subTareas {

		if subTarea.Id == subTareaActiva.Id {
			encontradaActiva = true
		}

		if subTarea.Id == subTareaEliminada.Id {
			encontradaEliminada = true
		}
	}

	if !encontradaActiva {
		t.Fatal(
			"La subtarea activa debería aparecer",
		)
	}

	if encontradaEliminada {
		t.Fatal(
			"La subtarea eliminada no debería aparecer",
		)
	}
}

func TestSubTareaRepositoryPostgres_Actualizar(t *testing.T) {
	pool := crearPoolPrueba(t)

	repository := NewSubTareaRepository(pool)

	subTarea := &entities.SubTarea{
		TareaId:   1,
		Titulo:    "Título original",
		CreadoEl:  time.Now(),
		CreadoPor: 1,
	}

	err := repository.Crear(
		context.Background(),
		subTarea,
	)

	if err != nil {
		t.Fatalf(
			"No se pudo crear la subtarea: %v",
			err,
		)
	}

	tituloNuevo := "Título modificado"

	subTarea.Titulo = tituloNuevo
	subTarea.EstaTerminada = true

	ahora := time.Now()
	subTarea.ModificadoEl = &ahora

	modificadoPor := 1
	subTarea.ModificadoPor = &modificadoPor

	err = repository.Actualizar(
		context.Background(),
		subTarea,
	)

	if err != nil {
		t.Fatalf(
			"No se pudo actualizar la subtarea: %v",
			err,
		)
	}

	actualizada, err := repository.ObtenerPorId(
		context.Background(),
		subTarea.Id,
	)

	if err != nil {
		t.Fatalf(
			"No se pudo obtener la subtarea actualizada: %v",
			err,
		)
	}

	if actualizada.Titulo != tituloNuevo {
		t.Fatalf(
			"Se esperaba título %q, se obtuvo %q",
			tituloNuevo,
			actualizada.Titulo,
		)
	}

	if !actualizada.EstaTerminada {
		t.Fatal(
			"Se esperaba que la subtarea estuviera terminada",
		)
	}
}

func TestSubTareaRepositoryPostgres_Eliminar(t *testing.T) {
	pool := crearPoolPrueba(t)

	repository := NewSubTareaRepository(pool)

	subTarea := &entities.SubTarea{
		TareaId:   1,
		Titulo:    "Subtarea para eliminar",
		CreadoEl:  time.Now(),
		CreadoPor: 1,
	}

	err := repository.Crear(
		context.Background(),
		subTarea,
	)

	if err != nil {
		t.Fatalf(
			"No se pudo crear la subtarea: %v",
			err,
		)
	}

	err = repository.Eliminar(
		context.Background(),
		subTarea.Id,
		1,
	)

	if err != nil {
		t.Fatalf(
			"No se pudo eliminar la subtarea: %v",
			err,
		)
	}

	eliminada, err := repository.ObtenerPorId(
		context.Background(),
		subTarea.Id,
	)

	if err != nil {
		t.Fatalf(
			"No se pudo obtener la subtarea eliminada: %v",
			err,
		)
	}

	if !eliminada.EstaEliminada {
		t.Fatal(
			"Se esperaba que EstaEliminada fuera true",
		)
	}

	subTareas, err := repository.ObtenerPorTareaId(
		context.Background(),
		subTarea.TareaId,
	)

	if err != nil {
		t.Fatalf(
			"No se pudieron obtener las subtareas: %v",
			err,
		)
	}

	for _, item := range subTareas {
		if item.Id == subTarea.Id {
			t.Fatal(
				"La subtarea eliminada no debería aparecer en ObtenerPorTareaId",
			)
		}
	}
}

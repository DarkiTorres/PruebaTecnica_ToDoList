package postgres

import (
	"context"
	"testing"
	"time"
	entities "to-do-server/internal/Domain/Entities"
	"to-do-server/internal/Infrastructure/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func crearPoolPrueba(t *testing.T) *pgxpool.Pool {
	t.Helper()

	err := godotenv.Load("../../../../.env")
	if err != nil {
		t.Fatalf(
			"No se pudo cargar el .env: %v",
			err,
		)
	}

	cfg := config.FromEnvironment()

	pool, err := pgxpool.New(
		context.Background(),
		cfg.GenerateConnectionString(),
	)

	if err != nil {
		t.Fatalf(
			"No se pudo crear el pool: %v",
			err,
		)
	}

	t.Cleanup(func() {
		pool.Close()
	})

	return pool
}

func TestTareaRepositoryPostgres_ObtenerPorId(t *testing.T) {
	pool := crearPoolPrueba(t)

	repository := NewTareaRepository(pool)

	tarea := &entities.Tarea{
		Titulo:      "Tarea para obtener",
		PrioridadId: 1,
		CreadoEl:    time.Now(),
		CreadoPor:   1,
	}

	err := repository.Crear(
		context.Background(),
		tarea,
	)

	if err != nil {
		t.Fatalf("No se pudo crear la tarea: %v", err)
	}

	tareaObtenida, err := repository.ObtenerPorId(
		context.Background(),
		tarea.Id,
	)

	if err != nil {
		t.Fatalf(
			"No se pudo obtener la tarea: %v",
			err,
		)
	}

	if tareaObtenida == nil {
		t.Fatal("Se esperaba una tarea")
	}

	if tareaObtenida.Id != tarea.Id {
		t.Fatalf(
			"Se esperaba Id %d, se obtuvo %d",
			tarea.Id,
			tareaObtenida.Id,
		)
	}

	if tareaObtenida.Titulo != tarea.Titulo {
		t.Fatalf(
			"Se esperaba título %q, se obtuvo %q",
			tarea.Titulo,
			tareaObtenida.Titulo,
		)
	}

	if tareaObtenida.PrioridadId != tarea.PrioridadId {
		t.Fatalf(
			"Se esperaba prioridad %d, se obtuvo %d",
			tarea.PrioridadId,
			tareaObtenida.PrioridadId,
		)
	}
}

func TestTareaRepositoryPostgres_Crear(t *testing.T) {
	pool := crearPoolPrueba(t)

	repository := NewTareaRepository(pool)

	tarea := &entities.Tarea{
		Titulo:      "Tarea de prueba de integración",
		PrioridadId: 1,
		CreadoEl:    time.Now(),
		CreadoPor:   1,
	}

	err := repository.Crear(
		context.Background(),
		tarea,
	)

	if err != nil {
		t.Fatalf(
			"No se pudo crear la tarea: %v",
			err,
		)
	}

	if tarea.Id == 0 {
		t.Fatal("Se esperaba que PostgreSQL asignara un Id")
	}

	t.Logf(
		"Tarea creada correctamente con Id: %d",
		tarea.Id,
	)
}

func TestTareaRepositoryPostgres_ObtenerTodos(t *testing.T) {
	pool := crearPoolPrueba(t)

	repository := NewTareaRepository(pool)

	tarea := &entities.Tarea{
		Titulo:      "Tarea para ObtenerTodos",
		PrioridadId: 2,
		CreadoEl:    time.Now(),
		CreadoPor:   1,
	}

	err := repository.Crear(
		context.Background(),
		tarea,
	)

	if err != nil {
		t.Fatalf("No se pudo crear la tarea: %v", err)
	}

	tareas, err := repository.ObtenerTodos(
		context.Background(),
	)

	if err != nil {
		t.Fatalf(
			"No se pudieron obtener las tareas: %v",
			err,
		)
	}

	encontrada := false

	for _, item := range tareas {
		if item.Id == tarea.Id {
			encontrada = true
			break
		}
	}

	if !encontrada {
		t.Fatalf(
			"La tarea creada con Id %d no apareció en ObtenerTodos",
			tarea.Id,
		)
	}
}

func TestTareaRepositoryPostgres_Actualizar(t *testing.T) {
	pool := crearPoolPrueba(t)

	repository := NewTareaRepository(pool)

	tarea := &entities.Tarea{
		Titulo:      "Título original",
		PrioridadId: 1,
		CreadoEl:    time.Now(),
		CreadoPor:   1,
	}

	err := repository.Crear(
		context.Background(),
		tarea,
	)

	if err != nil {
		t.Fatalf(
			"No se pudo crear la tarea: %v",
			err,
		)
	}

	tituloNuevo := "Título modificado"

	tarea.Titulo = tituloNuevo
	tarea.PrioridadId = 3

	ahora := time.Now()
	tarea.ModificadoEl = &ahora

	modificadoPor := 1
	tarea.ModificadoPor = &modificadoPor

	err = repository.Actualizar(
		context.Background(),
		tarea,
	)

	if err != nil {
		t.Fatalf(
			"No se pudo actualizar la tarea: %v",
			err,
		)
	}

	tareaActualizada, err := repository.ObtenerPorId(
		context.Background(),
		tarea.Id,
	)

	if err != nil {
		t.Fatalf(
			"No se pudo obtener la tarea actualizada: %v",
			err,
		)
	}

	if tareaActualizada.Titulo != tituloNuevo {
		t.Fatalf(
			"Se esperaba título %q, se obtuvo %q",
			tituloNuevo,
			tareaActualizada.Titulo,
		)
	}

	if tareaActualizada.PrioridadId != 3 {
		t.Fatalf(
			"Se esperaba prioridad 3, se obtuvo %d",
			tareaActualizada.PrioridadId,
		)
	}
}

func TestTareaRepositoryPostgres_Eliminar(t *testing.T) {
	pool := crearPoolPrueba(t)

	repository := NewTareaRepository(pool)

	tarea := &entities.Tarea{
		Titulo:      "Tarea para eliminar",
		PrioridadId: 4,
		CreadoEl:    time.Now(),
		CreadoPor:   1,
	}

	err := repository.Crear(
		context.Background(),
		tarea,
	)

	if err != nil {
		t.Fatalf(
			"No se pudo crear la tarea: %v",
			err,
		)
	}

	err = repository.Eliminar(
		context.Background(),
		tarea.Id,
		1,
	)

	if err != nil {
		t.Fatalf(
			"No se pudo eliminar la tarea: %v",
			err,
		)
	}

	tareaEliminada, err := repository.ObtenerPorId(
		context.Background(),
		tarea.Id,
	)

	if err != nil {
		t.Fatalf(
			"No se pudo obtener la tarea eliminada: %v",
			err,
		)
	}

	if !tareaEliminada.EstaEliminada {
		t.Fatal(
			"Se esperaba que EstaEliminada fuera true",
		)
	}

	tareas, err := repository.ObtenerTodos(
		context.Background(),
	)

	if err != nil {
		t.Fatalf(
			"No se pudieron obtener las tareas: %v",
			err,
		)
	}

	for _, item := range tareas {
		if item.Id == tarea.Id {
			t.Fatal(
				"La tarea eliminada no debería aparecer en ObtenerTodos",
			)
		}
	}
}

func TestTareaRepositoryPostgres_EliminarFisicoPorId(t *testing.T) {
	pool := crearPoolPrueba(t)

	repository := NewTareaRepository(pool)
	tarea := &entities.Tarea{
		Titulo:      "Tarea para eliminación física",
		PrioridadId: 1,
		CreadoEl:    time.Now(),
		CreadoPor:   1,
	}

	err := repository.Crear(
		context.Background(),
		tarea,
	)

	if err != nil {
		t.Fatalf("No se pudo crear la tarea de prueba: %v", err)
	}

	err = repository.EliminarFisicoPorId(
		context.Background(),
		tarea.Id,
	)

	if err != nil {
		t.Fatalf(
			"No se pudo eliminar físicamente la tarea: %v",
			err,
		)
	}

	tareaEliminada, err := repository.ObtenerPorId(
		context.Background(),
		tarea.Id,
	)

	if err == nil {
		t.Fatal("Se esperaba un error al consultar una tarea eliminada físicamente")
	}

	if tareaEliminada != nil {
		t.Fatal("La tarea no debería existir después de eliminarla físicamente")
	}
}

func TestTareaRepositoryPostgres_EliminarFisicoPorId_Cascada(
	t *testing.T,
) {
	pool := crearPoolPrueba(t)

	repository := NewTareaRepository(pool)
	subTareaRepository := NewSubTareaRepository(pool)
	tareaUsuarioRepository := NewTareaUsuarioRepository(pool)

	ctx := context.Background()

	// 1. Crear tarea
	tarea := &entities.Tarea{
		Titulo:      "Tarea para prueba de cascada",
		PrioridadId: 1,
		CreadoEl:    time.Now(),
		CreadoPor:   1,
	}

	err := repository.Crear(ctx, tarea)

	if err != nil {
		t.Fatalf(
			"No se pudo crear la tarea de prueba: %v",
			err,
		)
	}

	// 2. Crear tres subtareas
	subTareas := []*entities.SubTarea{
		{
			TareaId:   tarea.Id,
			Titulo:    "Subtarea 1",
			CreadoEl:  time.Now(),
			CreadoPor: 1,
		},
		{
			TareaId:   tarea.Id,
			Titulo:    "Subtarea 2",
			CreadoEl:  time.Now(),
			CreadoPor: 1,
		},
		{
			TareaId:   tarea.Id,
			Titulo:    "Subtarea 3",
			CreadoEl:  time.Now(),
			CreadoPor: 1,
		},
	}

	for _, subTarea := range subTareas {
		err := subTareaRepository.Crear(ctx, subTarea)

		if err != nil {
			t.Fatalf(
				"No se pudo crear la subtarea de prueba: %v",
				err,
			)
		}
	}

	// 3. Asignar la tarea al usuario 1
	tareaUsuario := &entities.TareaUsuario{
		TareaId:   tarea.Id,
		UsuarioId: 1,
	}

	err = tareaUsuarioRepository.Crear(
		ctx,
		tareaUsuario,
	)

	if err != nil {
		t.Fatalf(
			"No se pudo asignar la tarea al usuario: %v",
			err,
		)
	}

	// 4. Comprobar que existen las subtareas
	subTareasAntes, err := subTareaRepository.ObtenerPorTareaId(
		ctx,
		tarea.Id,
	)

	if err != nil {
		t.Fatalf(
			"No se pudieron consultar las subtareas: %v",
			err,
		)
	}

	if len(subTareasAntes) != 3 {
		t.Fatalf(
			"Se esperaban 3 subtareas, se encontraron %d",
			len(subTareasAntes),
		)
	}

	// 5. Comprobar que existe la asignación
	asignacionesAntes, err := tareaUsuarioRepository.ObtenerPorTareaId(
		ctx,
		tarea.Id,
	)

	if err != nil {
		t.Fatalf(
			"No se pudieron consultar las asignaciones: %v",
			err,
		)
	}

	if len(asignacionesAntes) != 1 {
		t.Fatalf(
			"Se esperaba 1 asignación, se encontraron %d",
			len(asignacionesAntes),
		)
	}

	// 6. Eliminar físicamente la tarea
	err = repository.EliminarFisicoPorId(
		ctx,
		tarea.Id,
	)

	if err != nil {
		t.Fatalf(
			"No se pudo eliminar físicamente la tarea: %v",
			err,
		)
	}

	// 7. Comprobar que la tarea ya no existe
	tareaEliminada, err := repository.ObtenerPorId(
		ctx,
		tarea.Id,
	)

	if err == nil {
		t.Fatal(
			"Se esperaba un error al consultar la tarea eliminada",
		)
	}

	if tareaEliminada != nil {
		t.Fatal(
			"La tarea debería haber sido eliminada físicamente",
		)
	}

	// 8. Comprobar que las subtareas fueron eliminadas por CASCADE
	subTareasDespues, err := subTareaRepository.ObtenerPorTareaId(
		ctx,
		tarea.Id,
	)

	if err != nil {
		t.Fatalf(
			"No se pudieron consultar las subtareas después de eliminar la tarea: %v",
			err,
		)
	}

	if len(subTareasDespues) != 0 {
		t.Fatalf(
			"Se esperaban 0 subtareas después de la cascada, se encontraron %d",
			len(subTareasDespues),
		)
	}

	// 9. Comprobar que las asignaciones fueron eliminadas por CASCADE
	asignacionesDespues, err := tareaUsuarioRepository.ObtenerPorTareaId(
		ctx,
		tarea.Id,
	)

	if err != nil {
		t.Fatalf(
			"No se pudieron consultar las asignaciones después de eliminar la tarea: %v",
			err,
		)
	}

	if len(asignacionesDespues) != 0 {
		t.Fatalf(
			"Se esperaban 0 asignaciones después de la cascada, se encontraron %d",
			len(asignacionesDespues),
		)
	}
}

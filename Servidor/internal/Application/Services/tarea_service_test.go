package services

import (
	"context"
	"errors"
	"strings"
	"testing"
	mocks "to-do-server/internal/Application/Tests/Mocks"
	validators "to-do-server/internal/Application/Validators"
	entities "to-do-server/internal/Domain/Entities"
)

func TestCompletarTarea_SinSubTareas_CompletaTarea(t *testing.T) {
	tareaRepository := &mocks.TareaRepositoryMock{
		Tarea: &entities.Tarea{
			Id:            1,
			Titulo:        "Test",
			EstaTerminada: false,
		},
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{
		SubTareas: []entities.SubTarea{},
	}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
	)

	err := service.CompletarTarea(
		context.Background(),
		1,
	)

	if err != nil {
		t.Fatalf(
			"Se esperaba que no hubiera error, se obtuvo: %v",
			err,
		)
	}

	if !tareaRepository.Tarea.EstaTerminada {
		t.Fatal("Se esperaba que la tarea estuviera terminada")
	}

	if !tareaRepository.ActualizarLlamado {
		t.Fatal("Se esperaba que Actualizar fuera llamado")
	}
}

func TestCompletarTarea_ConSubTareaPendiente_NoCompleta(t *testing.T) {
	tareaRepository := &mocks.TareaRepositoryMock{
		Tarea: &entities.Tarea{
			Id:            1,
			Titulo:        "Test",
			EstaTerminada: false,
		},
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{
		SubTareas: []entities.SubTarea{
			{
				Id:            1,
				TareaId:       1,
				Titulo:        "Subtarea pendiente",
				EstaTerminada: false,
				EstaEliminada: false,
			},
		},
	}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
	)

	err := service.CompletarTarea(
		context.Background(),
		1,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareaRepository.Tarea.EstaTerminada {
		t.Fatal("La tarea no debería haberse completado")
	}

	if tareaRepository.ActualizarLlamado {
		t.Fatal("Actualizar no debería haberse llamado")
	}
}

func TestCompletarTarea_ConTodasLasSubTareasTerminadas_CompletaTarea(t *testing.T) {
	tareaRepository := &mocks.TareaRepositoryMock{
		Tarea: &entities.Tarea{
			Id:            1,
			Titulo:        "Test",
			EstaTerminada: false,
		},
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{
		SubTareas: []entities.SubTarea{
			{
				Id:            1,
				TareaId:       1,
				EstaTerminada: true,
			},
			{
				Id:            2,
				TareaId:       1,
				EstaTerminada: true,
			},
		},
	}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
	)

	err := service.CompletarTarea(
		context.Background(),
		1,
	)

	if err != nil {
		t.Fatalf("No se esperaba error: %v", err)
	}

	if !tareaRepository.Tarea.EstaTerminada {
		t.Fatal("La tarea debería estar terminada")
	}
}

func TestCompletarTarea_ConSubTareaEliminada_NoBloquea(t *testing.T) {
	tareaRepository := &mocks.TareaRepositoryMock{
		Tarea: &entities.Tarea{
			Id:            1,
			Titulo:        "Test",
			EstaTerminada: false,
		},
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{
		SubTareas: []entities.SubTarea{
			{
				Id:            1,
				TareaId:       1,
				EstaTerminada: false,
				EstaEliminada: true,
			},
		},
	}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
	)

	err := service.CompletarTarea(
		context.Background(),
		1,
	)

	if err != nil {
		t.Fatalf("No se esperaba error: %v", err)
	}

	if !tareaRepository.Tarea.EstaTerminada {
		t.Fatal("La tarea debería estar terminada")
	}
}

func TestCompletarTarea_ErrorAlObtenerTarea_RegresaError(t *testing.T) {

	errorEsperado := errors.New("error de base de datos")

	tareaRepository := &mocks.TareaRepositoryMock{
		ErrorObtenerPorId: errorEsperado,
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
	)

	err := service.CompletarTarea(
		context.Background(),
		1,
	)

	if !errors.Is(err, errorEsperado) {
		t.Fatalf(
			"Se esperaba el error %v, se obtuvo %v",
			errorEsperado,
			err,
		)
	}
}

func TestCrearTarea_TareaValida_CreaTarea(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
	)

	tarea := &entities.Tarea{
		Titulo:      "Estudiar Go",
		PrioridadId: 1,
		CreadoPor:   1,
	}

	err := service.CrearTarea(
		context.Background(),
		tarea,
	)

	if err != nil {
		t.Fatalf(
			"No se esperaba error, se obtuvo: %v",
			err,
		)
	}

	if !tareaRepository.CrearLlamado {
		t.Fatal("Se esperaba que Crear fuera llamado")
	}

	if tareaRepository.Tarea != tarea {
		t.Fatal("Se esperaba que se enviara la misma tarea al repository")
	}
}

func TestCrearTarea_TituloVacio_NoCreaTarea(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
	)

	tarea := &entities.Tarea{
		Titulo:      "",
		PrioridadId: 1,
		CreadoPor:   1,
	}

	err := service.CrearTarea(
		context.Background(),
		tarea,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareaRepository.CrearLlamado {
		t.Fatal("Crear no debería haberse llamado")
	}
}

func TestCrearTarea_TituloDemasiadoLargo_NoCreaTarea(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
	)

	titulo := strings.Repeat("A", validators.MaxTituloTarea+1)

	tarea := &entities.Tarea{
		Titulo:      titulo,
		PrioridadId: 1,
		CreadoPor:   1,
	}

	err := service.CrearTarea(
		context.Background(),
		tarea,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareaRepository.CrearLlamado {
		t.Fatal("Crear no debería haberse llamado")
	}
}

func TestCrearTarea_ErrorAlCrear_RegresaError(t *testing.T) {

	errorEsperado := errors.New("error de base de datos")

	tareaRepository := &mocks.TareaRepositoryMock{
		ErrorCrear: errorEsperado,
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
	)

	tarea := &entities.Tarea{
		Titulo:      "Estudiar Go",
		PrioridadId: 1,
		CreadoPor:   1,
	}

	err := service.CrearTarea(
		context.Background(),
		tarea,
	)

	if !errors.Is(err, errorEsperado) {
		t.Fatalf(
			"Se esperaba el error %v, se obtuvo %v",
			errorEsperado,
			err,
		)
	}
}

func TestActualizarTarea_TareaValida_ActualizaTarea(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
	)

	tarea := &entities.Tarea{
		Id:          1,
		Titulo:      "Estudiar Go",
		PrioridadId: 1,
		CreadoPor:   1,
	}

	err := service.ActualizarTarea(
		context.Background(),
		tarea,
	)

	if err != nil {
		t.Fatalf(
			"No se esperaba error, se obtuvo: %v",
			err,
		)
	}

	if !tareaRepository.ActualizarLlamado {
		t.Fatal("Se esperaba que Actualizar fuera llamado")
	}

	if tareaRepository.Tarea != tarea {
		t.Fatal("Se esperaba que se enviara la misma tarea al repository")
	}
}

func TestActualizarTarea_TituloVacio_NoActualiza(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
	)

	tarea := &entities.Tarea{
		Id:          1,
		Titulo:      "",
		PrioridadId: 1,
		CreadoPor:   1,
	}

	err := service.ActualizarTarea(
		context.Background(),
		tarea,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareaRepository.ActualizarLlamado {
		t.Fatal("Actualizar no debería haberse llamado")
	}
}

func TestActualizarTarea_ErrorAlActualizar_RegresaError(t *testing.T) {

	errorEsperado := errors.New("error de base de datos")

	tareaRepository := &mocks.TareaRepositoryMock{
		ErrorActualizar: errorEsperado,
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
	)

	tarea := &entities.Tarea{
		Id:          1,
		Titulo:      "Estudiar Go",
		PrioridadId: 1,
		CreadoPor:   1,
	}

	err := service.ActualizarTarea(
		context.Background(),
		tarea,
	)

	if !errors.Is(err, errorEsperado) {
		t.Fatalf(
			"Se esperaba el error %v, se obtuvo %v",
			errorEsperado,
			err,
		)
	}
}

func TestEliminarTarea_IdInvalido_NoElimina(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}
	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
	)

	err := service.EliminarTarea(
		context.Background(),
		0,
		1,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareaRepository.ActualizarLlamado {
		t.Fatal("Actualizar no debería haberse llamado")
	}
}

func TestEliminarTarea_UsuarioInvalido_NoElimina(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}
	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
	)

	err := service.EliminarTarea(
		context.Background(),
		1,
		0,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareaRepository.ActualizarLlamado {
		t.Fatal("Actualizar no debería haberse llamado")
	}
}

func TestEliminarTarea_TareaValida_RealizaSoftDelete(t *testing.T) {

	tarea := &entities.Tarea{
		Id:            1,
		Titulo:        "Estudiar Go",
		EstaEliminada: false,
		CreadoPor:     1,
	}

	tareaRepository := &mocks.TareaRepositoryMock{
		Tarea: tarea,
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
	)

	err := service.EliminarTarea(
		context.Background(),
		1,
		2,
	)

	if err != nil {
		t.Fatalf(
			"No se esperaba error, se obtuvo: %v",
			err,
		)
	}

	if !tareaRepository.ActualizarLlamado {
		t.Fatal("Se esperaba que Actualizar fuera llamado")
	}

	if !tareaRepository.Tarea.EstaEliminada {
		t.Fatal("La tarea debería estar marcada como eliminada")
	}

	if tareaRepository.Tarea.ModificadoEl == nil {
		t.Fatal("ModificadoEl debería tener un valor")
	}

	if tareaRepository.Tarea.ModificadoPor == nil {
		t.Fatal("ModificadoPor debería tener un valor")
	}

	if *tareaRepository.Tarea.ModificadoPor != 2 {
		t.Fatalf(
			"Se esperaba ModificadoPor = 2, se obtuvo %d",
			*tareaRepository.Tarea.ModificadoPor,
		)
	}
}

func TestEliminarTarea_ErrorAlObtenerTarea_RegresaError(t *testing.T) {

	errorEsperado := errors.New("error de base de datos")

	tareaRepository := &mocks.TareaRepositoryMock{
		ErrorObtenerPorId: errorEsperado,
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
	)

	err := service.EliminarTarea(
		context.Background(),
		1,
		2,
	)

	if !errors.Is(err, errorEsperado) {
		t.Fatalf(
			"Se esperaba el error %v, se obtuvo %v",
			errorEsperado,
			err,
		)
	}

	if tareaRepository.ActualizarLlamado {
		t.Fatal("Actualizar no debería haberse llamado")
	}
}

func TestEliminarTarea_ErrorAlActualizar_RegresaError(t *testing.T) {

	errorEsperado := errors.New("error de base de datos")

	tareaRepository := &mocks.TareaRepositoryMock{
		Tarea: &entities.Tarea{
			Id:            1,
			Titulo:        "Estudiar Go",
			EstaEliminada: false,
			CreadoPor:     1,
		},
		ErrorActualizar: errorEsperado,
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
	)

	err := service.EliminarTarea(
		context.Background(),
		1,
		2,
	)

	if !errors.Is(err, errorEsperado) {
		t.Fatalf(
			"Se esperaba el error %v, se obtuvo %v",
			errorEsperado,
			err,
		)
	}
}

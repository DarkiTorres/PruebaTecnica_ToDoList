package services

import (
	"context"
	"errors"
	"testing"
	mocks "to-do-server/internal/Application/Tests/Mocks"
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

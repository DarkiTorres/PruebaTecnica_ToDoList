package services

import (
	"context"
	"errors"
	"testing"
	mocks "to-do-server/internal/Application/Tests/Mocks"
	entities "to-do-server/internal/Domain/Entities"
)

func TestCrearSubTarea_SubTareaValida_CreaSubTarea(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{
		Tarea: &entities.Tarea{
			Id:            1,
			Titulo:        "Estudiar Go",
			EstaEliminada: false,
		},
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	service := NewSubTareaService(
		subTareaRepository,
		tareaRepository,
	)

	subTarea := &entities.SubTarea{
		TareaId:   1,
		Titulo:    "Estudiar interfaces",
		CreadoPor: 2,
	}

	err := service.CrearSubTarea(
		context.Background(),
		subTarea,
	)

	if err != nil {
		t.Fatalf(
			"No se esperaba error, se obtuvo: %v",
			err,
		)
	}

	if !subTareaRepository.CrearLlamado {
		t.Fatal("Se esperaba que Crear fuera llamado")
	}
}

func TestCrearSubTarea_SubTareaInvalida_NoCreaSubTarea(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	service := NewSubTareaService(
		subTareaRepository,
		tareaRepository,
	)

	subTarea := &entities.SubTarea{
		TareaId:   1,
		Titulo:    "",
		CreadoPor: 2,
	}

	err := service.CrearSubTarea(
		context.Background(),
		subTarea,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if subTareaRepository.CrearLlamado {
		t.Fatal("Crear no debería haberse llamado")
	}
}

func TestCrearSubTarea_ErrorAlObtenerTarea_NoCreaSubTarea(t *testing.T) {

	errorEsperado := errors.New("error de base de datos")

	tareaRepository := &mocks.TareaRepositoryMock{
		ErrorObtenerPorId: errorEsperado,
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	service := NewSubTareaService(
		subTareaRepository,
		tareaRepository,
	)

	subTarea := &entities.SubTarea{
		TareaId:   1,
		Titulo:    "Estudiar interfaces",
		CreadoPor: 2,
	}

	err := service.CrearSubTarea(
		context.Background(),
		subTarea,
	)

	if !errors.Is(err, errorEsperado) {
		t.Fatalf(
			"Se esperaba el error %v, se obtuvo %v",
			errorEsperado,
			err,
		)
	}

	if subTareaRepository.CrearLlamado {
		t.Fatal("Crear no debería haberse llamado")
	}
}

func TestCrearSubTarea_TareaEliminada_NoCreaSubTarea(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{
		Tarea: &entities.Tarea{
			Id:            1,
			Titulo:        "Estudiar Go",
			EstaEliminada: true,
		},
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	service := NewSubTareaService(
		subTareaRepository,
		tareaRepository,
	)

	subTarea := &entities.SubTarea{
		TareaId:   1,
		Titulo:    "Estudiar interfaces",
		CreadoPor: 2,
	}

	err := service.CrearSubTarea(
		context.Background(),
		subTarea,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if subTareaRepository.CrearLlamado {
		t.Fatal("Crear no debería haberse llamado")
	}
}

func TestCrearSubTarea_ErrorAlCrear_RegresaError(t *testing.T) {

	errorEsperado := errors.New("error de base de datos")

	tareaRepository := &mocks.TareaRepositoryMock{
		Tarea: &entities.Tarea{
			Id:            1,
			Titulo:        "Estudiar Go",
			EstaEliminada: false,
		},
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{
		ErrorCrear: errorEsperado,
	}

	service := NewSubTareaService(
		subTareaRepository,
		tareaRepository,
	)

	subTarea := &entities.SubTarea{
		TareaId:   1,
		Titulo:    "Estudiar interfaces",
		CreadoPor: 2,
	}

	err := service.CrearSubTarea(
		context.Background(),
		subTarea,
	)

	if !errors.Is(err, errorEsperado) {
		t.Fatalf(
			"Se esperaba el error %v, se obtuvo %v",
			errorEsperado,
			err,
		)
	}
}

func TestCompletarSubTarea_IdInvalido_NoCompleta(t *testing.T) {

	subTareaRepository := &mocks.SubTareaRepositoryMock{}
	tareaRepository := &mocks.TareaRepositoryMock{}

	service := NewSubTareaService(
		subTareaRepository,
		tareaRepository,
	)

	err := service.CompletarSubTarea(
		context.Background(),
		0,
		true,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if subTareaRepository.ActualizarLlamado {
		t.Fatal("Actualizar no debería haberse llamado")
	}
}

func TestCompletarSubTarea_ErrorAlObtener_RegresaError(t *testing.T) {

	errorEsperado := errors.New("error de base de datos")

	subTareaRepository := &mocks.SubTareaRepositoryMock{
		ErrorObtenerPorId: errorEsperado,
	}

	tareaRepository := &mocks.TareaRepositoryMock{}

	service := NewSubTareaService(
		subTareaRepository,
		tareaRepository,
	)

	err := service.CompletarSubTarea(
		context.Background(),
		1,
		true,
	)

	if !errors.Is(err, errorEsperado) {
		t.Fatalf(
			"Se esperaba el error %v, se obtuvo %v",
			errorEsperado,
			err,
		)
	}

	if subTareaRepository.ActualizarLlamado {
		t.Fatal("Actualizar no debería haberse llamado")
	}
}

func TestCompletarSubTarea_SubTareaEliminada_NoCompleta(t *testing.T) {

	subTareaRepository := &mocks.SubTareaRepositoryMock{
		SubTarea: &entities.SubTarea{
			Id:            1,
			TareaId:       1,
			Titulo:        "Subtarea",
			EstaTerminada: false,
			EstaEliminada: true,
		},
	}

	tareaRepository := &mocks.TareaRepositoryMock{}

	service := NewSubTareaService(
		subTareaRepository,
		tareaRepository,
	)

	err := service.CompletarSubTarea(
		context.Background(),
		1,
		true,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if subTareaRepository.SubTarea.EstaTerminada {
		t.Fatal("La subtarea no debería haberse completado")
	}

	if subTareaRepository.ActualizarLlamado {
		t.Fatal("Actualizar no debería haberse llamado")
	}
}

func TestCompletarSubTarea_SubTareaValida_CompletaSubTarea(t *testing.T) {

	subTareaRepository := &mocks.SubTareaRepositoryMock{
		SubTarea: &entities.SubTarea{
			Id:            1,
			TareaId:       1,
			Titulo:        "Subtarea",
			EstaTerminada: false,
			EstaEliminada: false,
		},
	}

	tareaRepository := &mocks.TareaRepositoryMock{}

	service := NewSubTareaService(
		subTareaRepository,
		tareaRepository,
	)

	err := service.CompletarSubTarea(
		context.Background(),
		1,
		true,
	)

	if err != nil {
		t.Fatalf(
			"No se esperaba error, se obtuvo: %v",
			err,
		)
	}

	if !subTareaRepository.SubTarea.EstaTerminada {
		t.Fatal("La subtarea debería estar terminada")
	}

	if !subTareaRepository.ActualizarLlamado {
		t.Fatal("Se esperaba que Actualizar fuera llamado")
	}
}

func TestCompletarSubTarea_ErrorAlActualizar_RegresaError(t *testing.T) {

	errorEsperado := errors.New("error de base de datos")

	subTareaRepository := &mocks.SubTareaRepositoryMock{
		SubTarea: &entities.SubTarea{
			Id:            1,
			TareaId:       1,
			Titulo:        "Subtarea",
			EstaTerminada: false,
			EstaEliminada: false,
		},
		ErrorActualizar: errorEsperado,
	}

	tareaRepository := &mocks.TareaRepositoryMock{}

	service := NewSubTareaService(
		subTareaRepository,
		tareaRepository,
	)

	err := service.CompletarSubTarea(
		context.Background(),
		1,
		true,
	)

	if !errors.Is(err, errorEsperado) {
		t.Fatalf(
			"Se esperaba el error %v, se obtuvo %v",
			errorEsperado,
			err,
		)
	}
}

package services

import (
	"context"
	"errors"
	"testing"

	"to-do-server/internal/Application/Interfaces/repositories"
	entities "to-do-server/internal/Domain/Entities"
)

type PrioridadRepositoryMock struct {
	Prioridades []entities.Prioridad

	ObtenerTodosLlamado bool

	ErrorObtenerTodos error
}

var _ repositories.IPrioridadRepository = (*PrioridadRepositoryMock)(nil)

func (m *PrioridadRepositoryMock) ObtenerTodos(
	context context.Context,
) ([]entities.Prioridad, error) {

	m.ObtenerTodosLlamado = true

	if m.ErrorObtenerTodos != nil {
		return nil, m.ErrorObtenerTodos
	}

	return m.Prioridades, nil
}

func TestPrioridadService_ObtenerPrioridades_RetornaPrioridades(t *testing.T) {

	prioridades := []entities.Prioridad{
		{
			Id:          1,
			Descripcion: "Baja",
		},
		{
			Id:          2,
			Descripcion: "Media",
		},
		{
			Id:          3,
			Descripcion: "Alta",
		},
	}

	prioridadRepository := &PrioridadRepositoryMock{
		Prioridades: prioridades,
	}

	service := NewPrioridadService(
		prioridadRepository,
	)

	resultado, err := service.ObtenerPrioridades(
		context.Background(),
	)

	if err != nil {
		t.Fatalf(
			"No se esperaba un error, se obtuvo: %v",
			err,
		)
	}

	if !prioridadRepository.ObtenerTodosLlamado {
		t.Fatal(
			"Se esperaba que ObtenerTodos fuera llamado",
		)
	}

	if len(resultado) != 3 {
		t.Fatalf(
			"Se esperaban 3 prioridades, se obtuvieron %d",
			len(resultado),
		)
	}

	if resultado[0].Descripcion != "Baja" {
		t.Fatalf(
			"Se esperaba 'Baja', se obtuvo %s",
			resultado[0].Descripcion,
		)
	}

	if resultado[2].Descripcion != "Alta" {
		t.Fatalf(
			"Se esperaba 'Alta', se obtuvo %s",
			resultado[2].Descripcion,
		)
	}
}

func TestPrioridadService_ObtenerPrioridades_ErrorRepositorio_RetornaError(t *testing.T) {

	esperado := errors.New("error al obtener prioridades")

	prioridadRepository := &PrioridadRepositoryMock{
		ErrorObtenerTodos: esperado,
	}

	service := NewPrioridadService(
		prioridadRepository,
	)

	prioridades, err := service.ObtenerPrioridades(
		context.Background(),
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if !errors.Is(err, esperado) {
		t.Fatalf(
			"Se esperaba el error %v, se obtuvo %v",
			esperado,
			err,
		)
	}

	if prioridades != nil {
		t.Fatal("No se esperaban prioridades")
	}

	if !prioridadRepository.ObtenerTodosLlamado {
		t.Fatal(
			"Se esperaba que ObtenerTodos fuera llamado",
		)
	}
}

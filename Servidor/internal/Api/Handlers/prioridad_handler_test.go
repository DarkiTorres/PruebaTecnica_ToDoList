package handlers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	services "to-do-server/internal/Application/Interfaces/services"
	entities "to-do-server/internal/Domain/Entities"

	"github.com/gin-gonic/gin"
)

type PrioridadServiceMock struct {
	ObtenerPrioridadesLlamado bool
	Prioridades               []entities.Prioridad
	ErrorObtenerPrioridades   error
}

var _ services.IPrioridadService = (*PrioridadServiceMock)(nil)

func (m *PrioridadServiceMock) ObtenerPrioridades(
	context context.Context,
) ([]entities.Prioridad, error) {

	m.ObtenerPrioridadesLlamado = true

	if m.ErrorObtenerPrioridades != nil {
		return nil, m.ErrorObtenerPrioridades
	}

	return m.Prioridades, nil
}

func TestPrioridadHandler_ObtenerTodos_RetornaPrioridades(t *testing.T) {

	service := &PrioridadServiceMock{
		Prioridades: []entities.Prioridad{
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
		},
	}

	handler := NewPrioridadHandler(service)

	request := httptest.NewRequest(
		http.MethodGet,
		"/prioridades",
		nil,
	)

	response := httptest.NewRecorder()

	context, _ := gin.CreateTestContext(response)
	context.Request = request

	handler.ObtenerTodos(context)

	if response.Code != http.StatusOK {
		t.Fatalf(
			"Se esperaba status 200, se obtuvo %d",
			response.Code,
		)
	}

	if !service.ObtenerPrioridadesLlamado {
		t.Fatal(
			"Se esperaba que ObtenerPrioridades fuera llamado",
		)
	}

	if len(service.Prioridades) != 3 {
		t.Fatalf(
			"Se esperaban 3 prioridades, se obtuvieron %d",
			len(service.Prioridades),
		)
	}
}

func TestPrioridadHandler_ObtenerTodos_ErrorService_RetornaInternalServerError(t *testing.T) {

	service := &PrioridadServiceMock{
		ErrorObtenerPrioridades: errors.New(
			"error al obtener prioridades",
		),
	}

	handler := NewPrioridadHandler(service)

	request := httptest.NewRequest(
		http.MethodGet,
		"/prioridades",
		nil,
	)

	response := httptest.NewRecorder()

	context, _ := gin.CreateTestContext(response)
	context.Request = request

	handler.ObtenerTodos(context)

	if response.Code != http.StatusInternalServerError {
		t.Fatalf(
			"Se esperaba status 500, se obtuvo %d",
			response.Code,
		)
	}

	if !service.ObtenerPrioridadesLlamado {
		t.Fatal(
			"Se esperaba que ObtenerPrioridades fuera llamado",
		)
	}
}

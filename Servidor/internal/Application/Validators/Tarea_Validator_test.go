package validators

import (
	"strings"
	"testing"
	entities "to-do-server/internal/Domain/Entities"
)

func TestValidarTarea_TituloVacio_RegresaError(t *testing.T) {
	tarea := &entities.Tarea{
		Titulo:      "",
		PrioridadId: 1,
		CreadoPor:   1,
	}

	err := ValidarTarea(tarea)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}
}

func TestValidarTarea_TituloValido_NoRegresaError(t *testing.T) {

	descripcion := "Estudiar arquitectura limpia"

	tarea := &entities.Tarea{
		Titulo:      "Estudiar Go",
		Descripcion: &descripcion,
		PrioridadId: 1,
		CreadoPor:   1,
	}

	err := ValidarTarea(tarea)

	if err != nil {
		t.Fatalf("No se esperaba error: %v", err)
	}
}

func TestValidarTarea_TituloExcedeLongitud_RegresaError(t *testing.T) {

	titulo := strings.Repeat("A", MaxTituloTarea+1)

	tarea := &entities.Tarea{
		Titulo:      titulo,
		PrioridadId: 1,
		CreadoPor:   1,
	}

	err := ValidarTarea(tarea)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}
}

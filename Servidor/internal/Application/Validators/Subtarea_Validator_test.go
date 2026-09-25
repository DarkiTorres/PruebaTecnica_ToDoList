package validators

import (
	"strings"
	"testing"
	entities "to-do-server/internal/Domain/Entities"
)

func TestValidarSubTarea_Valida_NoRegresaError(t *testing.T) {

	subTarea := &entities.SubTarea{
		TareaId:   1,
		Titulo:    "Estudiar Go",
		CreadoPor: 1,
	}

	err := ValidarSubTarea(subTarea)

	if err != nil {
		t.Fatalf(
			"No se esperaba error, se obtuvo: %v",
			err,
		)
	}
}

func TestValidarSubTarea_Nil_RegresaError(t *testing.T) {

	err := ValidarSubTarea(nil)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}
}

func TestValidarSubTarea_TareaIdInvalido_RegresaError(t *testing.T) {

	subTarea := &entities.SubTarea{
		TareaId:   0,
		Titulo:    "Estudiar Go",
		CreadoPor: 1,
	}

	err := ValidarSubTarea(subTarea)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}
}

func TestValidarSubTarea_CreadoPorInvalido_RegresaError(t *testing.T) {

	subTarea := &entities.SubTarea{
		TareaId:   1,
		Titulo:    "Estudiar Go",
		CreadoPor: 0,
	}

	err := ValidarSubTarea(subTarea)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}
}

func TestValidarSubTarea_TituloVacio_RegresaError(t *testing.T) {

	subTarea := &entities.SubTarea{
		TareaId:   1,
		Titulo:    "   ",
		CreadoPor: 1,
	}

	err := ValidarSubTarea(subTarea)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}
}

func TestValidarSubTarea_TituloDemasiadoLargo_RegresaError(t *testing.T) {

	titulo := strings.Repeat(
		"A",
		MaxTituloSubTarea+1,
	)

	subTarea := &entities.SubTarea{
		TareaId:   1,
		Titulo:    titulo,
		CreadoPor: 1,
	}

	err := ValidarSubTarea(subTarea)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}
}

func TestValidarSubTarea_CaracterControl_RegresaError(t *testing.T) {

	subTarea := &entities.SubTarea{
		TareaId:   1,
		Titulo:    "Estudiar\nGo",
		CreadoPor: 1,
	}

	err := ValidarSubTarea(subTarea)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}
}

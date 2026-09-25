package validators

import (
	"errors"
	"strings"
	entities "to-do-server/internal/Domain/Entities"
	"unicode"
)

const MaxTituloSubTarea = 200

func ValidarSubTarea(subTarea *entities.SubTarea) error {
	if subTarea == nil {
		return errors.New("La subtarea no puede ser nula")
	}

	if subTarea.TareaId <= 0 {
		return errors.New("El id de la tarea no es válido")
	}

	if subTarea.CreadoPor <= 0 {
		return errors.New("El usuario creador no es válido")
	}

	titulo := strings.TrimSpace(subTarea.Titulo)

	if titulo == "" {
		return errors.New("El título de la subtarea es obligatorio")
	}

	if len([]rune(titulo)) > MaxTituloSubTarea {
		return errors.New(
			"El título de la subtarea no puede superar los 200 caracteres",
		)
	}

	for _, caracter := range titulo {
		if unicode.IsControl(caracter) {
			return errors.New(
				"El título contiene caracteres no permitidos",
			)
		}
	}

	return nil
}

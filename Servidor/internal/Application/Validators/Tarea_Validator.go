package validators

import (
	"errors"
	"strings"
	entities "to-do-server/internal/Domain/Entities"
)

const (
	MaxTituloTarea      = 200
	MaxDescripcionTarea = 500
)

func ValidarTarea(tarea *entities.Tarea) error {
	if tarea == nil {
		return errors.New("La tarea no puede ser nula.")
	}

	titulo := strings.TrimSpace(tarea.Titulo)

	if titulo == "" {
		return errors.New("El título es obligatorio.")
	}

	if len([]rune(titulo)) > MaxTituloTarea {
		return errors.New("El título no puede superar los 500 caracteres.")
	}

	if tarea.Descripcion != nil {
		if len([]rune(*tarea.Descripcion)) > MaxDescripcionTarea {
			return errors.New("La descripción no puede superar los 500 caracteres.")
		}
	}

	if tarea.PrioridadId <= 0 {
		return errors.New("La prioridad no es válida.")
	}

	if tarea.CreadoPor <= 0 {
		return errors.New("El usuario no es válida.")
	}

	return nil
}

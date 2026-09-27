package entities

import "time"

type SubTarea struct {
	Id            int64      `json:"id"`
	TareaId       int64      `json:"tareaId"`
	Titulo        string     `json:"titulo"`
	EstaTerminada bool       `json:"estaTerminada"`
	EstaEliminada bool       `json:"estaEliminada"`
	CreadoEl      time.Time  `json:"creadoEl"`
	CreadoPor     int        `json:"creadoPor"`
	ModificadoEl  *time.Time `json:"modificadoEl,omitempty"`
	ModificadoPor *int       `json:"modificadoPor,omitempty"`
}

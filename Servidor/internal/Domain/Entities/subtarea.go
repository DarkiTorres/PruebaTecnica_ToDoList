package entities

import "time"

type SubTarea struct {
	Id            int64
	TareaId       int64
	Titulo        string
	EstaTerminada bool
	EstaEliminada bool
	CreadoEl      time.Time
	CreadoPor     int
	ModificadoEl  *time.Time
	ModificadoPor *int
}

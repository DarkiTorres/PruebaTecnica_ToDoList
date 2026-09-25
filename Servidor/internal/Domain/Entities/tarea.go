package entities

import "time"

type Tarea struct {
	Id            int64
	Titulo        string
	Descripcion   *string
	PrioridadId   int16
	FechaEntrega  *time.Time
	EstaTerminada bool
	EstaEliminada bool
	CreadoEl      time.Time
	CreadoPor     int
	ModificadoEl  *time.Time
	ModificadoPor *int
}

package tareas

import "time"

type TareaResponse struct {
	Id            int64      `json:"id"`
	Titulo        string     `json:"titulo"`
	Descripcion   *string    `json:"descripcion"`
	PrioridadId   int16      `json:"prioridadId"`
	FechaEntrega  *time.Time `json:"fechaEntrega"`
	EstaTerminada bool       `json:"estaTerminada"`
	EstaEliminada bool       `json:"estaEliminada"`
	CreadoEl      time.Time  `json:"creadoEl"`
	CreadoPor     int        `json:"creadoPor"`
	ModificadoEl  *time.Time `json:"modificadoEl"`
	ModificadoPor *int       `json:"modificadoPor"`

	AsignadoAId *int    `json:"asignadoAId"`
	AsignadoA   *string `json:"asignadoA"`

	SubTareas []SubTareaResponse `json:"subTareas"`
}

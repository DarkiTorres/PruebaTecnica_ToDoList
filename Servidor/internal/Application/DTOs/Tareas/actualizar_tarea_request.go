package tareas

type ActualizarTareaRequest struct {
	Titulo        string                      `json:"titulo"`
	Descripcion   *string                     `json:"descripcion"`
	PrioridadId   int16                       `json:"prioridadId"`
	FechaEntrega  *string                     `json:"fechaEntrega"`
	ModificadoPor int                         `json:"modificadoPor"`
	AsignadoA     int                         `json:"asignadoA"`
	SubTareas     []ActualizarSubTareaRequest `json:"subTareas"`
}

type ActualizarSubTareaRequest struct {
	Id            int64  `json:"id"`
	Titulo        string `json:"titulo"`
	EstaTerminada bool   `json:"estaTerminada"`
	EstaEliminada bool   `json:"estaEliminada"`
	CreadoPor     int    `json:"creadoPor"`
	ModificadoPor int    `json:"modificadoPor"`
}

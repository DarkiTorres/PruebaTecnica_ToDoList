package tareas

type CrearTareaRequest struct {
	Titulo       string  `json:"titulo"`
	Descripcion  *string `json:"descripcion"`
	PrioridadId  int16   `json:"prioridadId"`
	FechaEntrega *string `json:"fechaEntrega"`
	CreadoPor    int     `json:"creadoPor"`
	AsignadoA    int     `json:"asignadoA"`
}

package tareas

type CrearSubTareaRequest struct {
	Titulo    string `json:"titulo"`
	TareaId   int64  `json:"tareaId"`
	CreadoPor int    `json:"creadoPor"`
}

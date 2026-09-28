package handlers

import (
	"context"
	"to-do-server/internal/Application/Interfaces/repositories"
	interfaces "to-do-server/internal/Application/Interfaces/services"
	entities "to-do-server/internal/Domain/Entities"
)

type TareaServiceMock struct {
	CrearTareaLlamado      bool
	ActualizarTareaLlamado bool

	Tarea         *entities.Tarea
	AsignadoA     int
	Tareas        []entities.Tarea
	TareaObtenida *entities.Tarea

	TareaActualizada *entities.Tarea

	ErrorCrearTarea      error
	ErrorObtenerTarea    error
	ErrorObtenerTareas   error
	ErrorActualizarTarea error

	EliminarTareaLlamado bool
	TareaEliminadaId     int64
	TareaEliminadaPor    int
	ErrorEliminarTarea   error

	TareasPorUsuario      []entities.Tarea
	CompletarTareaLlamado bool
	TareaCompletadaId     int64
	ErrorCompletarTarea   error

	ObtenerTareasPorUsuarioLlamado bool
	UsuarioIdTareas                int
	UsuarioIdObtenerTareas         int
	ErrorObtenerTareasPorUsuario   error
}

var _ interfaces.ITareaService = (*TareaServiceMock)(nil)

type SubTareaRepositoryMock struct {
	SubTareas []entities.SubTarea
	Error     error
}

var _ repositories.ISubTareaRepository = (*SubTareaRepositoryMock)(nil)

func (m *TareaServiceMock) CompletarTarea(
	context context.Context,
	tareaId int64,
	estaTerminada bool,
) error {
	m.CompletarTareaLlamado = true
	m.TareaCompletadaId = tareaId
	return m.ErrorCompletarTarea
}

func (m *TareaServiceMock) ObtenerTareasPorUsuario(
	context context.Context,
	usuarioId int,
) ([]entities.Tarea, error) {
	m.ObtenerTareasPorUsuarioLlamado = true
	m.UsuarioIdTareas = usuarioId

	if m.ErrorObtenerTareasPorUsuario != nil {
		return nil, m.ErrorObtenerTareasPorUsuario
	}

	return m.TareasPorUsuario, nil
}
func (m *TareaServiceMock) ActualizarTarea(
	context context.Context,
	tarea *entities.Tarea,
	asignadoA int,
) error {

	m.ActualizarTareaLlamado = true
	m.TareaActualizada = tarea

	return m.ErrorActualizarTarea
}

func (m *TareaServiceMock) EliminarTarea(
	context context.Context,
	tareaId int64,
	modificadoPor int,
) error {

	m.EliminarTareaLlamado = true
	m.TareaEliminadaId = tareaId
	m.TareaEliminadaPor = modificadoPor

	return m.ErrorEliminarTarea
}

func (m *TareaServiceMock) ObtenerTareaPorId(
	context context.Context,
	tareaId int64,
) (*entities.Tarea, error) {
	return m.TareaObtenida, m.ErrorObtenerTarea
}

func (m *TareaServiceMock) ObtenerTareas(
	context context.Context,
) ([]entities.Tarea, error) {
	return m.Tareas, m.ErrorObtenerTareas
}

func (m *TareaServiceMock) CrearTarea(
	context context.Context,
	tarea *entities.Tarea,
	asignadoA int,
) error {

	m.CrearTareaLlamado = true
	m.Tarea = tarea
	m.AsignadoA = asignadoA

	return m.ErrorCrearTarea
}

func (m *SubTareaRepositoryMock) Crear(
	context context.Context,
	subTarea *entities.SubTarea,
) error {
	return nil
}

func (m *SubTareaRepositoryMock) ObtenerPorId(
	context context.Context,
	id int64,
) (*entities.SubTarea, error) {
	return nil, nil
}

func (m *SubTareaRepositoryMock) ObtenerPorTareaId(
	context context.Context,
	tareaId int64,
) ([]entities.SubTarea, error) {
	return m.SubTareas, m.Error
}

func (m *SubTareaRepositoryMock) Actualizar(
	context context.Context,
	subTarea *entities.SubTarea,
) error {
	return nil
}

func (m *SubTareaRepositoryMock) Eliminar(
	context context.Context,
	id int64,
	modificadoPor int,
) error {
	return nil
}

// func TestTareaHandler_Crear_TareaValida_CreaTarea(t *testing.T) {

// 	service := &TareaServiceMock{}

// 	subTareaService := &SubTareaServiceMock{}

// 	subTareaRepository := &SubTareaRepositoryMock{}

// 	handler := NewTareaHandler(
// 		service,
// 		subTareaService,
// 		subTareaRepository,
// 	)

// 	body := `{
// 		"titulo": "Estudiar Go",
// 		"descripcion": "Repasar interfaces",
// 		"prioridadId": 2,
// 		"fechaEntrega": "2026-09-30T18:00:00Z",
// 		"creadoPor": 1,
// 		"asignadoA": 2
// 	}`

// 	request := httptest.NewRequest(
// 		http.MethodPost,
// 		"/tareas",
// 		strings.NewReader(body),
// 	)

// 	request.Header.Set(
// 		"Content-Type",
// 		"application/json",
// 	)

// 	response := httptest.NewRecorder()

// 	context, _ := gin.CreateTestContext(response)

// 	context.Request = request

// 	handler.Crear(context)

// 	if response.Code != http.StatusCreated {
// 		t.Fatalf(
// 			"Se esperaba status 201, se obtuvo %d",
// 			response.Code,
// 		)
// 	}

// 	if !service.CrearTareaLlamado {
// 		t.Fatal(
// 			"Se esperaba que CrearTarea fuera llamado",
// 		)
// 	}

// 	if service.Tarea == nil {
// 		t.Fatal(
// 			"Se esperaba que se enviara una tarea al service",
// 		)
// 	}

// 	if service.Tarea.Titulo != "Estudiar Go" {
// 		t.Fatalf(
// 			"Se esperaba título %q, se obtuvo %q",
// 			"Estudiar Go",
// 			service.Tarea.Titulo,
// 		)
// 	}

// 	if service.Tarea.PrioridadId != 2 {
// 		t.Fatalf(
// 			"Se esperaba prioridad 2, se obtuvo %d",
// 			service.Tarea.PrioridadId,
// 		)
// 	}

// 	if service.AsignadoA != 2 {
// 		t.Fatalf(
// 			"Se esperaba usuario asignado 2, se obtuvo %d",
// 			service.AsignadoA,
// 		)
// 	}
// }

// func TestTareaHandler_Crear_JSONInvalido_RetornaBadRequest(t *testing.T) {

// 	service := &TareaServiceMock{}

// 	subTareaService := &SubTareaServiceMock{}

// 	subTareaRepository := &SubTareaRepositoryMock{}

// 	handler := NewTareaHandler(
// 		service,
// 		subTareaService,
// 		subTareaRepository,
// 	)

// 	body := `{
// 		"titulo": "Tarea de prueba",
// 		"prioridadId":
// 	}`

// 	request := httptest.NewRequest(
// 		http.MethodPost,
// 		"/tareas",
// 		strings.NewReader(body),
// 	)

// 	request.Header.Set(
// 		"Content-Type",
// 		"application/json",
// 	)

// 	response := httptest.NewRecorder()

// 	context, _ := gin.CreateTestContext(response)
// 	context.Request = request

// 	handler.Crear(context)

// 	if response.Code != http.StatusBadRequest {
// 		t.Fatalf(
// 			"Se esperaba status 400, se obtuvo %d",
// 			response.Code,
// 		)
// 	}

// 	if service.CrearTareaLlamado {
// 		t.Fatal(
// 			"El service no debería haberse llamado",
// 		)
// 	}
// }

// func TestTareaHandler_Crear_FechaInvalida_RetornaBadRequest(t *testing.T) {

// 	service := &TareaServiceMock{}

// 	subTareaService := &SubTareaServiceMock{}

// 	subTareaRepository := &SubTareaRepositoryMock{}

// 	handler := NewTareaHandler(
// 		service,
// 		subTareaService,
// 		subTareaRepository,
// 	)

// 	body := `{
// 		"titulo": "Estudiar Go",
// 		"prioridadId": 2,
// 		"fechaEntrega": "fecha-invalida",
// 		"creadoPor": 1,
// 		"asignadoA": 2
// 	}`

// 	request := httptest.NewRequest(
// 		http.MethodPost,
// 		"/tareas",
// 		strings.NewReader(body),
// 	)

// 	request.Header.Set(
// 		"Content-Type",
// 		"application/json",
// 	)

// 	response := httptest.NewRecorder()

// 	context, _ := gin.CreateTestContext(response)
// 	context.Request = request

// 	handler.Crear(context)

// 	if response.Code != http.StatusBadRequest {
// 		t.Fatalf(
// 			"Se esperaba status 400, se obtuvo %d",
// 			response.Code,
// 		)
// 	}

// 	if service.CrearTareaLlamado {
// 		t.Fatal(
// 			"El service no debería haberse llamado",
// 		)
// 	}
// }

// func TestTareaHandler_Crear_ServiceDevuelveError_RetornaBadRequest(t *testing.T) {

// 	service := &TareaServiceMock{
// 		ErrorCrearTarea: errors.New("el usuario asignado no existe"),
// 	}

// 	subTareaService := &SubTareaServiceMock{}

// 	subTareaRepository := &SubTareaRepositoryMock{}

// 	handler := NewTareaHandler(
// 		service,
// 		subTareaService,
// 		subTareaRepository,
// 	)

// 	body := `{
// 		"titulo": "Estudiar Go",
// 		"prioridadId": 2,
// 		"creadoPor": 1,
// 		"asignadoA": 99
// 	}`

// 	request := httptest.NewRequest(
// 		http.MethodPost,
// 		"/tareas",
// 		strings.NewReader(body),
// 	)

// 	request.Header.Set(
// 		"Content-Type",
// 		"application/json",
// 	)

// 	response := httptest.NewRecorder()

// 	context, _ := gin.CreateTestContext(response)
// 	context.Request = request

// 	handler.Crear(context)

// 	if response.Code != http.StatusBadRequest {
// 		t.Fatalf(
// 			"Se esperaba status 400, se obtuvo %d",
// 			response.Code,
// 		)
// 	}

// 	if !service.CrearTareaLlamado {
// 		t.Fatal(
// 			"Se esperaba que CrearTarea fuera llamado",
// 		)
// 	}
// }

// func TestTareaHandler_ObtenerTodas_DevuelveTareasConSubTareas(t *testing.T) {

// 	descripcion := "Estudiar interfaces"

// 	service := &TareaServiceMock{
// 		Tareas: []entities.Tarea{
// 			{
// 				Id:          1,
// 				Titulo:      "Estudiar Go",
// 				Descripcion: &descripcion,
// 				PrioridadId: 2,
// 				CreadoPor:   1,
// 			},
// 		},
// 	}

// 	subTareaRepository := &SubTareaRepositoryMock{
// 		SubTareas: []entities.SubTarea{
// 			{
// 				Id:            1,
// 				TareaId:       1,
// 				Titulo:        "Repasar interfaces",
// 				EstaTerminada: false,
// 				CreadoPor:     1,
// 			},
// 			{
// 				Id:            2,
// 				TareaId:       1,
// 				Titulo:        "Repasar structs",
// 				EstaTerminada: true,
// 				CreadoPor:     1,
// 			},
// 		},
// 	}

// 	subTareaService := &SubTareaServiceMock{}

// 	handler := NewTareaHandler(
// 		service,
// 		subTareaService,
// 		subTareaRepository,
// 	)

// 	request := httptest.NewRequest(
// 		http.MethodGet,
// 		"/tareas",
// 		nil,
// 	)

// 	response := httptest.NewRecorder()

// 	context, _ := gin.CreateTestContext(response)
// 	context.Request = request

// 	handler.ObtenerTodas(context)

// 	if response.Code != http.StatusOK {
// 		t.Fatalf(
// 			"Se esperaba status 200, se obtuvo %d",
// 			response.Code,
// 		)
// 	}

// 	body := response.Body.String()

// 	if !strings.Contains(body, `"titulo":"Estudiar Go"`) {
// 		t.Fatal(
// 			"Se esperaba encontrar la tarea en la respuesta",
// 		)
// 	}

// 	if !strings.Contains(body, `"titulo":"Repasar interfaces"`) {
// 		t.Fatal(
// 			"Se esperaba encontrar la primera subtarea en la respuesta",
// 		)
// 	}

// 	if !strings.Contains(body, `"titulo":"Repasar structs"`) {
// 		t.Fatal(
// 			"Se esperaba encontrar la segunda subtarea en la respuesta",
// 		)
// 	}
// }

// func TestTareaHandler_ObtenerTodas_TareaSinSubTareas(t *testing.T) {

// 	service := &TareaServiceMock{
// 		Tareas: []entities.Tarea{
// 			{
// 				Id:          1,
// 				Titulo:      "Tarea sin subtareas",
// 				PrioridadId: 2,
// 				CreadoPor:   1,
// 			},
// 		},
// 	}

// 	subTareaService := &SubTareaServiceMock{}

// 	subTareaRepository := &SubTareaRepositoryMock{
// 		SubTareas: []entities.SubTarea{},
// 	}

// 	handler := NewTareaHandler(
// 		service,
// 		subTareaService,
// 		subTareaRepository,
// 	)

// 	request := httptest.NewRequest(
// 		http.MethodGet,
// 		"/tareas",
// 		nil,
// 	)

// 	response := httptest.NewRecorder()

// 	context, _ := gin.CreateTestContext(response)
// 	context.Request = request

// 	handler.ObtenerTodas(context)

// 	if response.Code != http.StatusOK {
// 		t.Fatalf(
// 			"Se esperaba status 200, se obtuvo %d",
// 			response.Code,
// 		)
// 	}

// 	body := response.Body.String()

// 	if !strings.Contains(body, `"titulo":"Tarea sin subtareas"`) {
// 		t.Fatal(
// 			"Se esperaba encontrar la tarea",
// 		)
// 	}

// 	if !strings.Contains(body, `"subTareas":[]`) {
// 		t.Fatal(
// 			"Se esperaba que la tarea tuviera una lista de subtareas vacía",
// 		)
// 	}
// }

// func TestTareaHandler_ObtenerTodas_ServiceDevuelveError_RetornaInternalServerError(t *testing.T) {

// 	service := &TareaServiceMock{
// 		ErrorObtenerTareas: errors.New("error al obtener las tareas"),
// 	}

// 	subTareaService := &SubTareaServiceMock{}

// 	subTareaRepository := &SubTareaRepositoryMock{}

// 	handler := NewTareaHandler(
// 		service,
// 		subTareaService,
// 		subTareaRepository,
// 	)

// 	request := httptest.NewRequest(
// 		http.MethodGet,
// 		"/tareas",
// 		nil,
// 	)

// 	response := httptest.NewRecorder()

// 	context, _ := gin.CreateTestContext(response)
// 	context.Request = request

// 	handler.ObtenerTodas(context)

// 	if response.Code != http.StatusInternalServerError {
// 		t.Fatalf(
// 			"Se esperaba status 500, se obtuvo %d",
// 			response.Code,
// 		)
// 	}
// }

// func TestTareaHandler_ObtenerPorId_DevuelveTareaConSubTareas(t *testing.T) {

// 	service := &TareaServiceMock{
// 		TareaObtenida: &entities.Tarea{
// 			Id:          1,
// 			Titulo:      "Estudiar Go",
// 			PrioridadId: 2,
// 			CreadoPor:   1,
// 		},
// 	}

// 	subTareaRepository := &SubTareaRepositoryMock{
// 		SubTareas: []entities.SubTarea{
// 			{
// 				Id:        1,
// 				TareaId:   1,
// 				Titulo:    "Repasar interfaces",
// 				CreadoPor: 1,
// 			},
// 			{
// 				Id:        2,
// 				TareaId:   1,
// 				Titulo:    "Repasar structs",
// 				CreadoPor: 1,
// 			},
// 		},
// 	}

// 	subTareaService := &SubTareaServiceMock{}

// 	handler := NewTareaHandler(
// 		service,
// 		subTareaService,
// 		subTareaRepository,
// 	)

// 	request := httptest.NewRequest(
// 		http.MethodGet,
// 		"/tareas/1",
// 		nil,
// 	)

// 	response := httptest.NewRecorder()

// 	context, _ := gin.CreateTestContext(response)
// 	context.Request = request

// 	// Como estamos usando CreateTestContext, debemos establecer el parámetro.
// 	context.Params = gin.Params{
// 		{
// 			Key:   "id",
// 			Value: "1",
// 		},
// 	}

// 	handler.ObtenerPorId(context)

// 	if response.Code != http.StatusOK {
// 		t.Fatalf(
// 			"Se esperaba status 200, se obtuvo %d",
// 			response.Code,
// 		)
// 	}

// 	body := response.Body.String()

// 	if !strings.Contains(body, `"titulo":"Estudiar Go"`) {
// 		t.Fatal("Se esperaba encontrar la tarea")
// 	}

// 	if !strings.Contains(body, `"titulo":"Repasar interfaces"`) {
// 		t.Fatal("Se esperaba encontrar la primera subtarea")
// 	}

// 	if !strings.Contains(body, `"titulo":"Repasar structs"`) {
// 		t.Fatal("Se esperaba encontrar la segunda subtarea")
// 	}
// }

// func TestTareaHandler_ObtenerPorId_IdInvalido_RetornaBadRequest(t *testing.T) {

// 	service := &TareaServiceMock{}
// 	subTareaService := &SubTareaServiceMock{}

// 	subTareaRepository := &SubTareaRepositoryMock{}

// 	handler := NewTareaHandler(
// 		service,
// 		subTareaService,
// 		subTareaRepository,
// 	)

// 	request := httptest.NewRequest(
// 		http.MethodGet,
// 		"/tareas/abc",
// 		nil,
// 	)

// 	response := httptest.NewRecorder()

// 	context, _ := gin.CreateTestContext(response)
// 	context.Request = request

// 	context.Params = gin.Params{
// 		{
// 			Key:   "id",
// 			Value: "abc",
// 		},
// 	}

// 	handler.ObtenerPorId(context)

// 	if response.Code != http.StatusBadRequest {
// 		t.Fatalf(
// 			"Se esperaba status 400, se obtuvo %d",
// 			response.Code,
// 		)
// 	}
// }

// func TestTareaHandler_ObtenerPorId_TareaNoExiste_RetornaNotFound(t *testing.T) {

// 	service := &TareaServiceMock{
// 		ErrorObtenerTarea: errors.New("la tarea no existe"),
// 	}

// 	subTareaService := &SubTareaServiceMock{}

// 	subTareaRepository := &SubTareaRepositoryMock{}

// 	handler := NewTareaHandler(
// 		service,
// 		subTareaService,
// 		subTareaRepository,
// 	)

// 	request := httptest.NewRequest(
// 		http.MethodGet,
// 		"/tareas/999",
// 		nil,
// 	)

// 	response := httptest.NewRecorder()

// 	context, _ := gin.CreateTestContext(response)
// 	context.Request = request

// 	context.Params = gin.Params{
// 		{
// 			Key:   "id",
// 			Value: "999",
// 		},
// 	}

// 	handler.ObtenerPorId(context)

// 	if response.Code != http.StatusNotFound {
// 		t.Fatalf(
// 			"Se esperaba status 404, se obtuvo %d",
// 			response.Code,
// 		)
// 	}
// }

// func TestTareaHandler_ObtenerPorId_ErrorSubTareas_RetornaInternalServerError(t *testing.T) {

// 	service := &TareaServiceMock{
// 		TareaObtenida: &entities.Tarea{
// 			Id:          1,
// 			Titulo:      "Estudiar Go",
// 			PrioridadId: 2,
// 			CreadoPor:   1,
// 		},
// 	}

// 	subTareaService := &SubTareaServiceMock{}

// 	subTareaRepository := &SubTareaRepositoryMock{
// 		Error: errors.New("error al obtener las subtareas"),
// 	}

// 	handler := NewTareaHandler(
// 		service,
// 		subTareaService,
// 		subTareaRepository,
// 	)

// 	request := httptest.NewRequest(
// 		http.MethodGet,
// 		"/tareas/1",
// 		nil,
// 	)

// 	response := httptest.NewRecorder()

// 	context, _ := gin.CreateTestContext(response)
// 	context.Request = request

// 	context.Params = gin.Params{
// 		{
// 			Key:   "id",
// 			Value: "1",
// 		},
// 	}

// 	handler.ObtenerPorId(context)

// 	if response.Code != http.StatusInternalServerError {
// 		t.Fatalf(
// 			"Se esperaba status 500, se obtuvo %d",
// 			response.Code,
// 		)
// 	}
// }

// func TestTareaHandler_Actualizar_TareaValida_ActualizaTarea(t *testing.T) {

// 	service := &TareaServiceMock{}

// 	subTareaService := &SubTareaServiceMock{}

// 	subTareaRepository := &SubTareaRepositoryMock{}

// 	handler := NewTareaHandler(
// 		service,
// 		subTareaService,
// 		subTareaRepository,
// 	)

// 	body := `{
// 		"titulo": "Estudiar Go avanzado",
// 		"descripcion": "Repasar interfaces y arquitectura",
// 		"prioridadId": 3,
// 		"fechaEntrega": "2026-10-01T18:00:00Z",
// 		"modificadoPor": 1
// 	}`

// 	request := httptest.NewRequest(
// 		http.MethodPut,
// 		"/tareas/5",
// 		strings.NewReader(body),
// 	)

// 	request.Header.Set(
// 		"Content-Type",
// 		"application/json",
// 	)

// 	response := httptest.NewRecorder()

// 	context, _ := gin.CreateTestContext(response)
// 	context.Request = request

// 	context.Params = gin.Params{
// 		{
// 			Key:   "id",
// 			Value: "5",
// 		},
// 	}

// 	handler.Actualizar(context)

// 	if response.Code != http.StatusOK {
// 		t.Fatalf(
// 			"Se esperaba status 200, se obtuvo %d",
// 			response.Code,
// 		)
// 	}

// 	if !service.ActualizarTareaLlamado {
// 		t.Fatal(
// 			"Se esperaba que ActualizarTarea fuera llamado",
// 		)
// 	}

// 	if service.TareaActualizada == nil {
// 		t.Fatal(
// 			"Se esperaba que se enviara una tarea al service",
// 		)
// 	}

// 	if service.TareaActualizada.Id != 5 {
// 		t.Fatalf(
// 			"Se esperaba id 5, se obtuvo %d",
// 			service.TareaActualizada.Id,
// 		)
// 	}

// 	if service.TareaActualizada.Titulo != "Estudiar Go avanzado" {
// 		t.Fatalf(
// 			"Se esperaba título %q, se obtuvo %q",
// 			"Estudiar Go avanzado",
// 			service.TareaActualizada.Titulo,
// 		)
// 	}

// 	if service.TareaActualizada.PrioridadId != 3 {
// 		t.Fatalf(
// 			"Se esperaba prioridad 3, se obtuvo %d",
// 			service.TareaActualizada.PrioridadId,
// 		)
// 	}

// 	if service.TareaActualizada.ModificadoPor == nil {
// 		t.Fatal(
// 			"Se esperaba ModificadoPor",
// 		)
// 	}

// 	if *service.TareaActualizada.ModificadoPor != 1 {
// 		t.Fatalf(
// 			"Se esperaba ModificadoPor 1, se obtuvo %d",
// 			*service.TareaActualizada.ModificadoPor,
// 		)
// 	}
// }

// func TestTareaHandler_Actualizar_IdInvalido_RetornaBadRequest(t *testing.T) {

// 	service := &TareaServiceMock{}
// 	subTareaService := &SubTareaServiceMock{}
// 	subTareaRepository := &SubTareaRepositoryMock{}

// 	handler := NewTareaHandler(
// 		service,
// 		subTareaService,
// 		subTareaRepository,
// 	)

// 	body := `{
// 		"titulo": "Tarea modificada",
// 		"prioridadId": 2,
// 		"modificadoPor": 1
// 	}`

// 	request := httptest.NewRequest(
// 		http.MethodPut,
// 		"/tareas/abc",
// 		strings.NewReader(body),
// 	)

// 	request.Header.Set(
// 		"Content-Type", "application/json",
// 	)

// 	response := httptest.NewRecorder()

// 	context, _ := gin.CreateTestContext(response)
// 	context.Request = request

// 	context.Params = gin.Params{
// 		{
// 			Key:   "id",
// 			Value: "abc",
// 		},
// 	}

// 	handler.Actualizar(context)

// 	if response.Code != http.StatusBadRequest {
// 		t.Fatalf(
// 			"Se esperaba status 400, se obtuvo %d",
// 			response.Code,
// 		)
// 	}

// 	if service.ActualizarTareaLlamado {
// 		t.Fatal(
// 			"El service no debería haberse llamado",
// 		)
// 	}
// }

// func TestTareaHandler_Actualizar_JSONInvalido_RetornaBadRequest(t *testing.T) {

// 	service := &TareaServiceMock{}
// 	subTareaService := &SubTareaServiceMock{}

// 	subTareaRepository := &SubTareaRepositoryMock{}

// 	handler := NewTareaHandler(
// 		service,
// 		subTareaService,
// 		subTareaRepository,
// 	)

// 	body := `{
// 		"titulo":
// 	}`

// 	request := httptest.NewRequest(
// 		http.MethodPut,
// 		"/tareas/1",
// 		strings.NewReader(body),
// 	)

// 	request.Header.Set(
// 		"Content-Type", "application/json",
// 	)

// 	response := httptest.NewRecorder()

// 	context, _ := gin.CreateTestContext(response)
// 	context.Request = request

// 	context.Params = gin.Params{
// 		{
// 			Key:   "id",
// 			Value: "1",
// 		},
// 	}

// 	handler.Actualizar(context)

// 	if response.Code != http.StatusBadRequest {
// 		t.Fatalf(
// 			"Se esperaba status 400, se obtuvo %d",
// 			response.Code,
// 		)
// 	}

// 	if service.ActualizarTareaLlamado {
// 		t.Fatal(
// 			"El service no debería haberse llamado",
// 		)
// 	}
// }

// func TestTareaHandler_Actualizar_FechaInvalida_RetornaBadRequest(t *testing.T) {

// 	service := &TareaServiceMock{}
// 	subTareaService := &SubTareaServiceMock{}

// 	subTareaRepository := &SubTareaRepositoryMock{}

// 	handler := NewTareaHandler(
// 		service,
// 		subTareaService,
// 		subTareaRepository,
// 	)

// 	body := `{
// 		"titulo": "Tarea modificada",
// 		"prioridadId": 2,
// 		"fechaEntrega": "fecha-invalida",
// 		"modificadoPor": 1
// 	}`

// 	request := httptest.NewRequest(
// 		http.MethodPut,
// 		"/tareas/1",
// 		strings.NewReader(body),
// 	)

// 	request.Header.Set(
// 		"Content-Type", "application/json",
// 	)

// 	response := httptest.NewRecorder()

// 	context, _ := gin.CreateTestContext(response)
// 	context.Request = request

// 	context.Params = gin.Params{
// 		{
// 			Key:   "id",
// 			Value: "1",
// 		},
// 	}

// 	handler.Actualizar(context)

// 	if response.Code != http.StatusBadRequest {
// 		t.Fatalf(
// 			"Se esperaba status 400, se obtuvo %d",
// 			response.Code,
// 		)
// 	}

// 	if service.ActualizarTareaLlamado {
// 		t.Fatal(
// 			"El service no debería haberse llamado",
// 		)
// 	}
// }

// func TestTareaHandler_Actualizar_ServiceDevuelveError_RetornaBadRequest(t *testing.T) {

// 	service := &TareaServiceMock{
// 		ErrorActualizarTarea: errors.New(
// 			"la tarea no existe",
// 		),
// 	}

// 	subTareaRepository := &SubTareaRepositoryMock{}

// 	subTareaService := &SubTareaServiceMock{}

// 	handler := NewTareaHandler(
// 		service,
// 		subTareaService,
// 		subTareaRepository,
// 	)

// 	body := `{
// 		"titulo": "Tarea modificada",
// 		"prioridadId": 2,
// 		"modificadoPor": 1
// 	}`

// 	request := httptest.NewRequest(
// 		http.MethodPut,
// 		"/tareas/1",
// 		strings.NewReader(body),
// 	)

// 	request.Header.Set(
// 		"Content-Type", "application/json",
// 	)

// 	response := httptest.NewRecorder()

// 	context, _ := gin.CreateTestContext(response)
// 	context.Request = request

// 	context.Params = gin.Params{
// 		{
// 			Key:   "id",
// 			Value: "1",
// 		},
// 	}

// 	handler.Actualizar(context)

// 	if response.Code != http.StatusBadRequest {
// 		t.Fatalf(
// 			"Se esperaba status 400, se obtuvo %d",
// 			response.Code,
// 		)
// 	}

// 	if !service.ActualizarTareaLlamado {
// 		t.Fatal(
// 			"Se esperaba que ActualizarTarea fuera llamado",
// 		)
// 	}
// }

// func TestTareaHandler_Actualizar_AgregaSubTarea(t *testing.T) {

// 	service := &TareaServiceMock{}

// 	subTareaService := &SubTareaServiceMock{}

// 	subTareaRepository := &SubTareaRepositoryMock{}

// 	handler := NewTareaHandler(
// 		service,
// 		subTareaService,
// 		subTareaRepository,
// 	)

// 	body := `{
// 		"titulo": "Estudiar Go",
// 		"prioridadId": 2,
// 		"modificadoPor": 1,
// 		"subTareas": [
// 			{
// 				"titulo": "Repasar interfaces",
// 				"creadoPor": 1
// 			}
// 		]
// 	}`

// 	request := httptest.NewRequest(
// 		http.MethodPut,
// 		"/tareas/5",
// 		strings.NewReader(body),
// 	)

// 	request.Header.Set(
// 		"Content-Type",
// 		"application/json",
// 	)

// 	response := httptest.NewRecorder()

// 	context, _ := gin.CreateTestContext(response)

// 	context.Request = request

// 	context.Params = gin.Params{
// 		{
// 			Key:   "id",
// 			Value: "5",
// 		},
// 	}

// 	handler.Actualizar(context)

// 	if response.Code != http.StatusOK {
// 		t.Fatalf(
// 			"Se esperaba status 200, se obtuvo %d",
// 			response.Code,
// 		)
// 	}

// 	if !service.ActualizarTareaLlamado {
// 		t.Fatal(
// 			"Se esperaba que ActualizarTarea fuera llamado",
// 		)
// 	}

// 	if !subTareaService.CrearSubTareaLlamado {
// 		t.Fatal(
// 			"Se esperaba que CrearSubTarea fuera llamado",
// 		)
// 	}

// 	if subTareaService.SubTarea == nil {
// 		t.Fatal(
// 			"Se esperaba que se enviara una subtarea al service",
// 		)
// 	}

// 	if subTareaService.SubTarea.TareaId != 5 {
// 		t.Fatalf(
// 			"Se esperaba TareaId 5, se obtuvo %d",
// 			subTareaService.SubTarea.TareaId,
// 		)
// 	}

// 	if subTareaService.SubTarea.Titulo != "Repasar interfaces" {
// 		t.Fatalf(
// 			"Se esperaba título %q, se obtuvo %q",
// 			"Repasar interfaces",
// 			subTareaService.SubTarea.Titulo,
// 		)
// 	}

// 	if subTareaService.SubTarea.CreadoPor != 1 {
// 		t.Fatalf(
// 			"Se esperaba CreadoPor 1, se obtuvo %d",
// 			1,
// 		)
// 	}
// }

// func TestTareaHandler_Actualizar_ModificaSubTarea(t *testing.T) {

// 	service := &TareaServiceMock{}

// 	subTareaService := &SubTareaServiceMock{}

// 	subTareaRepository := &SubTareaRepositoryMock{}

// 	handler := NewTareaHandler(
// 		service,
// 		subTareaService,
// 		subTareaRepository,
// 	)

// 	body := `{
// 		"titulo": "Estudiar Go",
// 		"prioridadId": 2,
// 		"modificadoPor": 1,
// 		"subTareas": [
// 			{
// 				"id": 10,
// 				"titulo": "Repasar interfaces avanzado",
// 				"estaTerminada": true,
// 				"creadoPor": 1,
// 				"modificadoPor": 1
// 			}
// 		]
// 	}`

// 	request := httptest.NewRequest(
// 		http.MethodPut,
// 		"/tareas/5",
// 		strings.NewReader(body),
// 	)

// 	request.Header.Set(
// 		"Content-Type",
// 		"application/json",
// 	)

// 	response := httptest.NewRecorder()

// 	context, _ := gin.CreateTestContext(response)
// 	context.Request = request

// 	context.Params = gin.Params{
// 		{
// 			Key:   "id",
// 			Value: "5",
// 		},
// 	}

// 	handler.Actualizar(context)

// 	if response.Code != http.StatusOK {
// 		t.Fatalf(
// 			"Se esperaba status 200, se obtuvo %d",
// 			response.Code,
// 		)
// 	}

// 	if !subTareaService.ActualizarSubTareaLlamado {
// 		t.Fatal(
// 			"Se esperaba que ActualizarSubTarea fuera llamado",
// 		)
// 	}

// 	if subTareaService.SubTareaActualizada == nil {
// 		t.Fatal(
// 			"Se esperaba que se enviara una subtarea al service",
// 		)
// 	}

// 	if subTareaService.SubTareaActualizada.Id != 10 {
// 		t.Fatalf(
// 			"Se esperaba Id 10, se obtuvo %d",
// 			subTareaService.SubTareaActualizada.Id,
// 		)
// 	}

// 	if subTareaService.SubTareaActualizada.TareaId != 5 {
// 		t.Fatalf(
// 			"Se esperaba TareaId 5, se obtuvo %d",
// 			subTareaService.SubTareaActualizada.TareaId,
// 		)
// 	}

// 	if subTareaService.SubTareaActualizada.Titulo != "Repasar interfaces avanzado" {
// 		t.Fatalf(
// 			"Se esperaba el título actualizado, se obtuvo %q",
// 			subTareaService.SubTareaActualizada.Titulo,
// 		)
// 	}

// 	if !subTareaService.SubTareaActualizada.EstaTerminada {
// 		t.Fatal(
// 			"Se esperaba que la subtarea estuviera terminada",
// 		)
// 	}
// }

// func TestTareaHandler_Actualizar_EliminaSubTarea(t *testing.T) {

// 	service := &TareaServiceMock{}

// 	subTareaService := &SubTareaServiceMock{}

// 	subTareaRepository := &SubTareaRepositoryMock{}

// 	handler := NewTareaHandler(
// 		service,
// 		subTareaService,
// 		subTareaRepository,
// 	)

// 	body := `{
// 		"titulo": "Estudiar Go",
// 		"prioridadId": 2,
// 		"modificadoPor": 1,
// 		"subTareas": [
// 			{
// 				"id": 10,
// 				"titulo": "Repasar interfaces",
// 				"estaEliminada": true,
// 				"creadoPor": 1,
// 				"modificadoPor": 1
// 			}
// 		]
// 	}`

// 	request := httptest.NewRequest(
// 		http.MethodPut,
// 		"/tareas/5",
// 		strings.NewReader(body),
// 	)

// 	request.Header.Set(
// 		"Content-Type",
// 		"application/json",
// 	)

// 	response := httptest.NewRecorder()

// 	context, _ := gin.CreateTestContext(response)
// 	context.Request = request

// 	context.Params = gin.Params{
// 		{
// 			Key:   "id",
// 			Value: "5",
// 		},
// 	}

// 	handler.Actualizar(context)

// 	if response.Code != http.StatusOK {
// 		t.Fatalf(
// 			"Se esperaba status 200, se obtuvo %d",
// 			response.Code,
// 		)
// 	}

// 	if !subTareaService.EliminarSubTareaLlamado {
// 		t.Fatal(
// 			"Se esperaba que EliminarSubTarea fuera llamado",
// 		)
// 	}

// 	if subTareaService.SubTareaEliminadaId != 10 {
// 		t.Fatalf(
// 			"Se esperaba eliminar la subtarea 10, se obtuvo %d",
// 			subTareaService.SubTareaEliminadaId,
// 		)
// 	}

// 	if subTareaService.ModificadoPorEliminado != 1 {
// 		t.Fatalf(
// 			"Se esperaba ModificadoPor 1, se obtuvo %d",
// 			subTareaService.ModificadoPorEliminado,
// 		)
// 	}
// }

// func TestTareaHandler_Eliminar_EliminaTarea(t *testing.T) {

// 	service := &TareaServiceMock{}

// 	subTareaService := &SubTareaServiceMock{}

// 	subTareaRepository := &SubTareaRepositoryMock{}

// 	handler := NewTareaHandler(
// 		service,
// 		subTareaService,
// 		subTareaRepository,
// 	)

// 	body := `{
// 		"modificadoPor": 1
// 	}`

// 	request := httptest.NewRequest(
// 		http.MethodPatch,
// 		"/tareas/5/eliminar",
// 		strings.NewReader(body),
// 	)

// 	request.Header.Set(
// 		"Content-Type",
// 		"application/json",
// 	)

// 	response := httptest.NewRecorder()

// 	context, _ := gin.CreateTestContext(response)

// 	context.Request = request

// 	context.Params = gin.Params{
// 		{
// 			Key:   "id",
// 			Value: "5",
// 		},
// 	}

// 	handler.Eliminar(context)

// 	if response.Code != http.StatusNoContent {
// 		t.Fatalf(
// 			"Se esperaba status 204, se obtuvo %d",
// 			response.Code,
// 		)
// 	}

// 	if !service.EliminarTareaLlamado {
// 		t.Fatal(
// 			"Se esperaba que EliminarTarea fuera llamado",
// 		)
// 	}

// 	if service.TareaEliminadaId != 5 {
// 		t.Fatalf(
// 			"Se esperaba tarea 5, se obtuvo %d",
// 			service.TareaEliminadaId,
// 		)
// 	}

// 	if service.TareaEliminadaPor != 1 {
// 		t.Fatalf(
// 			"Se esperaba ModificadoPor 1, se obtuvo %d",
// 			service.TareaEliminadaPor,
// 		)
// 	}
// }

// func TestTareaService_EliminarTarea_TareaValida_EliminaTarea(t *testing.T) {

// 	tareaRepository := &mocks.TareaRepositoryMock{
// 		Tareas: []entities.Tarea{
// 			{
// 				Id:            5,
// 				Titulo:        "Estudiar Go",
// 				EstaEliminada: false,
// 				CreadoPor:     1,
// 			},
// 		},
// 	}

// 	subTareaRepository := &mocks.SubTareaRepositoryMock{}

// 	usuarioRepository := &mocks.UsuarioRepositoryMock{}

// 	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

// 	service := services.NewTareaService(
// 		tareaRepository,
// 		subTareaRepository,
// 		usuarioRepository,
// 		tareaUsuarioRepository,
// 	)

// 	err := service.EliminarTarea(
// 		context.Background(),
// 		5,
// 		2,
// 	)

// 	if err != nil {
// 		t.Fatalf(
// 			"No se esperaba un error, se obtuvo: %v",
// 			err,
// 		)
// 	}

// 	if !tareaRepository.ActualizarLlamado {
// 		t.Fatal(
// 			"Se esperaba que Actualizar fuera llamado",
// 		)
// 	}

// 	if tareaRepository.Tarea == nil {
// 		t.Fatal(
// 			"Se esperaba que se enviara una tarea al repositorio",
// 		)
// 	}

// 	if !tareaRepository.Tarea.EstaEliminada {
// 		t.Fatal(
// 			"Se esperaba que la tarea quedara eliminada",
// 		)
// 	}

// 	if tareaRepository.Tarea.Id != 5 {
// 		t.Fatalf(
// 			"Se esperaba tarea 5, se obtuvo %d",
// 			tareaRepository.Tarea.Id,
// 		)
// 	}

// 	if tareaRepository.Tarea.ModificadoPor == nil {
// 		t.Fatal(
// 			"Se esperaba ModificadoPor",
// 		)
// 	}

// 	if *tareaRepository.Tarea.ModificadoPor != 2 {
// 		t.Fatalf(
// 			"Se esperaba ModificadoPor 2, se obtuvo %d",
// 			*tareaRepository.Tarea.ModificadoPor,
// 		)
// 	}

// 	if tareaRepository.Tarea.ModificadoEl == nil {
// 		t.Fatal(
// 			"Se esperaba ModificadoEl",
// 		)
// 	}
// }

// func TestTareaHandler_Completar_CompletaTarea(t *testing.T) {

// 	service := &TareaServiceMock{}

// 	handler := NewTareaHandler(
// 		service,
// 		nil,
// 		nil,
// 	)

// 	request := httptest.NewRequest(
// 		http.MethodPatch,
// 		"/tareas/5/completar",
// 		nil,
// 	)

// 	response := httptest.NewRecorder()

// 	context, _ := gin.CreateTestContext(response)
// 	context.Request = request

// 	context.Params = gin.Params{
// 		{
// 			Key:   "id",
// 			Value: "5",
// 		},
// 	}

// 	handler.Completar(context)

// 	t.Logf("STATUS GIN: %d", context.Writer.Status())
// 	t.Logf("STATUS RESPONSE: %d", response.Code)

// 	if context.Writer.Status() != http.StatusNoContent {
// 		t.Fatalf(
// 			"Se esperaba status 204, se obtuvo %d",
// 			context.Writer.Status(),
// 		)
// 	}

// 	if !service.CompletarTareaLlamado {
// 		t.Fatal(
// 			"Se esperaba que CompletarTarea fuera llamado",
// 		)
// 	}

// 	if service.TareaCompletadaId != 5 {
// 		t.Fatalf(
// 			"Se esperaba tareaId 5, se obtuvo %d",
// 			service.TareaCompletadaId,
// 		)
// 	}
// }

// func TestTareaHandler_ObtenerPorUsuario_ObtieneTareas(t *testing.T) {

// 	service := &TareaServiceMock{
// 		TareasPorUsuario: []entities.Tarea{
// 			{
// 				Id:        1,
// 				Titulo:    "Estudiar Go",
// 				CreadoPor: 5,
// 			},
// 			{
// 				Id:        2,
// 				Titulo:    "Hacer pruebas",
// 				CreadoPor: 5,
// 			},
// 		},
// 	}

// 	handler := NewTareaHandler(
// 		service,
// 		nil,
// 		nil,
// 	)

// 	request := httptest.NewRequest(
// 		http.MethodGet,
// 		"/tareas/usuario/5",
// 		nil,
// 	)

// 	response := httptest.NewRecorder()

// 	context, _ := gin.CreateTestContext(response)
// 	context.Request = request

// 	context.Params = gin.Params{
// 		{
// 			Key:   "usuarioId",
// 			Value: "5",
// 		},
// 	}

// 	handler.ObtenerPorUsuario(context)

// 	if response.Code != http.StatusOK {
// 		t.Fatalf(
// 			"Se esperaba status 200, se obtuvo %d",
// 			response.Code,
// 		)
// 	}

// 	if !service.ObtenerTareasPorUsuarioLlamado {
// 		t.Fatal(
// 			"Se esperaba que ObtenerTareasPorUsuario fuera llamado",
// 		)
// 	}

// 	if service.UsuarioIdTareas != 5 {
// 		t.Fatalf(
// 			"Se esperaba usuarioId 5, se obtuvo %d",
// 			service.UsuarioIdTareas,
// 		)
// 	}

// 	if len(service.TareasPorUsuario) != 2 {
// 		t.Fatalf(
// 			"Se esperaban 2 tareas, se obtuvieron %d",
// 			len(service.TareasPorUsuario),
// 		)
// 	}
// }

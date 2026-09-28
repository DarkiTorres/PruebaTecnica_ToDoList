package handlers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	services "to-do-server/internal/Application/Interfaces/services"
	serv "to-do-server/internal/Application/Services"
	mocks "to-do-server/internal/Application/Tests/Mocks"
	entities "to-do-server/internal/Domain/Entities"

	"github.com/gin-gonic/gin"
)

type SubTareaServiceMock struct {
	CrearSubTareaLlamado bool
	SubTarea             *entities.SubTarea
	ErrorCrearSubTarea   error

	ActualizarSubTareaLlamado bool
	SubTareaActualizada       *entities.SubTarea
	ErrorActualizarSubTarea   error

	EliminarSubTareaLlamado bool
	SubTareaEliminadaId     int64
	ModificadoPorEliminado  int
	ErrorEliminarSubTarea   error

	ObtenerPorTareaIdLlamado bool
	SubTareasPorTarea        []entities.SubTarea
	ErrorObtenerPorTareaId   error

	CompletarSubTareaLlamado bool
	SubTareaCompletadaId     int64
	EstadoSubTarea           bool
	ErrorCompletarSubTarea   error
}

var _ services.ISubTareaService = (*SubTareaServiceMock)(nil)

func (m *SubTareaServiceMock) CompletarSubTarea(
	context context.Context,
	subTareaId int64,
	estaTerminada bool,
) error {

	m.CompletarSubTareaLlamado = true
	m.SubTareaCompletadaId = subTareaId
	m.EstadoSubTarea = estaTerminada

	return m.ErrorCompletarSubTarea
}
func (m *SubTareaServiceMock) ObtenerPorTareaId(
	context context.Context,
	tareaId int64,
) ([]entities.SubTarea, error) {

	m.ObtenerPorTareaIdLlamado = true

	return m.SubTareasPorTarea, m.ErrorObtenerPorTareaId
}

func (m *SubTareaServiceMock) CrearSubTarea(
	context context.Context,
	subTarea *entities.SubTarea,
) error {

	m.CrearSubTareaLlamado = true
	m.SubTarea = subTarea

	return m.ErrorCrearSubTarea
}

func (m *SubTareaServiceMock) ActualizarSubTarea(
	context context.Context,
	subTarea *entities.SubTarea,
) error {

	m.ActualizarSubTareaLlamado = true
	m.SubTareaActualizada = subTarea

	return m.ErrorActualizarSubTarea
}

func (m *SubTareaServiceMock) EliminarSubTarea(
	context context.Context,
	subTareaId int64,
	modificadoPor int,
) error {

	m.EliminarSubTareaLlamado = true
	m.SubTareaEliminadaId = subTareaId
	m.ModificadoPorEliminado = modificadoPor

	return m.ErrorEliminarSubTarea
}

func TestSubTareaHandler_Crear_CreaSubTarea(t *testing.T) {

	service := &SubTareaServiceMock{}

	handler := NewSubTareaHandler(service)

	body := `{
		"titulo": "Repasar interfaces",
		"creadoPor": 1
	}`

	request := httptest.NewRequest(
		http.MethodPost,
		"/tareas/1/subtareas",
		strings.NewReader(body),
	)

	request.Header.Set(
		"Content-Type",
		"application/json",
	)

	response := httptest.NewRecorder()

	context, _ := gin.CreateTestContext(response)
	context.Request = request

	context.Params = gin.Params{
		{
			Key:   "id",
			Value: "1",
		},
	}

	handler.Crear(context)

	if response.Code != http.StatusCreated {
		t.Fatalf(
			"Se esperaba status 201, se obtuvo %d",
			response.Code,
		)
	}

	if !service.CrearSubTareaLlamado {
		t.Fatal(
			"Se esperaba que CrearSubTarea fuera llamado",
		)
	}

	if service.SubTarea == nil {
		t.Fatal(
			"Se esperaba que se enviara una subtarea al service",
		)
	}

	if service.SubTarea.TareaId != 1 {
		t.Fatalf(
			"Se esperaba TareaId 1, se obtuvo %d",
			service.SubTarea.TareaId,
		)
	}

	if service.SubTarea.Titulo != "Repasar interfaces" {
		t.Fatalf(
			"Se esperaba título %q, se obtuvo %q",
			"Repasar interfaces",
			service.SubTarea.Titulo,
		)
	}

	if service.SubTarea.CreadoPor != 1 {
		t.Fatalf(
			"Se esperaba CreadoPor 1, se obtuvo %d",
			service.SubTarea.CreadoPor,
		)
	}
}

func TestSubTareaHandler_Crear_IdInvalido_RetornaBadRequest(t *testing.T) {

	service := &SubTareaServiceMock{}

	handler := NewSubTareaHandler(service)

	body := `{
		"titulo": "Repasar interfaces",
		"creadoPor": 1
	}`

	request := httptest.NewRequest(
		http.MethodPost,
		"/tareas/abc/subtareas",
		strings.NewReader(body),
	)

	request.Header.Set(
		"Content-Type",
		"application/json",
	)

	response := httptest.NewRecorder()

	context, _ := gin.CreateTestContext(response)
	context.Request = request

	context.Params = gin.Params{
		{
			Key:   "id",
			Value: "abc",
		},
	}

	handler.Crear(context)

	if response.Code != http.StatusBadRequest {
		t.Fatalf(
			"Se esperaba status 400, se obtuvo %d",
			response.Code,
		)
	}

	if service.CrearSubTareaLlamado {
		t.Fatal(
			"El service no debería haberse llamado",
		)
	}
}

func TestSubTareaHandler_Crear_JSONInvalido_RetornaBadRequest(t *testing.T) {

	service := &SubTareaServiceMock{}

	handler := NewSubTareaHandler(service)

	body := `{
		"titulo":
	}`

	request := httptest.NewRequest(
		http.MethodPost,
		"/tareas/1/subtareas",
		strings.NewReader(body),
	)

	request.Header.Set(
		"Content-Type",
		"application/json",
	)

	response := httptest.NewRecorder()

	context, _ := gin.CreateTestContext(response)
	context.Request = request

	context.Params = gin.Params{
		{
			Key:   "id",
			Value: "1",
		},
	}

	handler.Crear(context)

	if response.Code != http.StatusBadRequest {
		t.Fatalf(
			"Se esperaba status 400, se obtuvo %d",
			response.Code,
		)
	}

	if service.CrearSubTareaLlamado {
		t.Fatal(
			"El service no debería haberse llamado",
		)
	}
}

func TestSubTareaHandler_Crear_ServiceDevuelveError_RetornaBadRequest(t *testing.T) {

	service := &SubTareaServiceMock{
		ErrorCrearSubTarea: errors.New(
			"No se puede crear una subtarea para una tarea eliminada.",
		),
	}

	handler := NewSubTareaHandler(service)

	body := `{
		"titulo": "Repasar interfaces",
		"creadoPor": 1
	}`

	request := httptest.NewRequest(
		http.MethodPost,
		"/tareas/1/subtareas",
		strings.NewReader(body),
	)

	request.Header.Set(
		"Content-Type",
		"application/json",
	)

	response := httptest.NewRecorder()

	context, _ := gin.CreateTestContext(response)
	context.Request = request

	context.Params = gin.Params{
		{
			Key:   "id",
			Value: "1",
		},
	}

	handler.Crear(context)

	if response.Code != http.StatusBadRequest {
		t.Fatalf(
			"Se esperaba status 400, se obtuvo %d",
			response.Code,
		)
	}

	if !service.CrearSubTareaLlamado {
		t.Fatal(
			"Se esperaba que CrearSubTarea fuera llamado",
		)
	}
}

func TestSubTareaService_CrearSubTarea_TareaValida_CreaSubTarea(t *testing.T) {

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	tareaRepository := &mocks.TareaRepositoryMock{
		Tareas: []entities.Tarea{
			{
				Id:            5,
				Titulo:        "Estudiar Go",
				EstaEliminada: false,
			},
		},
	}

	service := serv.NewSubTareaService(
		subTareaRepository,
		tareaRepository,
	)

	subTarea := &entities.SubTarea{
		TareaId:       5,
		Titulo:        "Estudiar interfaces",
		CreadoPor:     1,
		EstaTerminada: false,
	}

	err := service.CrearSubTarea(
		context.Background(),
		subTarea,
	)

	if err != nil {
		t.Fatalf(
			"No se esperaba un error, se obtuvo: %v",
			err,
		)
	}

	if !subTareaRepository.CrearLlamado {
		t.Fatal(
			"Se esperaba que Crear fuera llamado",
		)
	}

	if subTareaRepository.SubTarea == nil {
		t.Fatal(
			"Se esperaba que se enviara una subtarea al repositorio",
		)
	}

	if subTareaRepository.SubTarea.TareaId != 5 {
		t.Fatalf(
			"Se esperaba TareaId 5, se obtuvo %d",
			subTareaRepository.SubTarea.TareaId,
		)
	}

	if subTareaRepository.SubTarea.Titulo != "Estudiar interfaces" {
		t.Fatalf(
			"Se esperaba título 'Estudiar interfaces', se obtuvo %s",
			subTareaRepository.SubTarea.Titulo,
		)
	}
}

func TestSubTareaService_CrearSubTarea_TareaEliminada_NoCreaSubTarea(t *testing.T) {

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	tareaRepository := &mocks.TareaRepositoryMock{
		Tareas: []entities.Tarea{
			{
				Id:            5,
				Titulo:        "Estudiar Go",
				EstaEliminada: true,
			},
		},
	}

	service := serv.NewSubTareaService(
		subTareaRepository,
		tareaRepository,
	)

	subTarea := &entities.SubTarea{
		TareaId:       5,
		Titulo:        "Estudiar interfaces",
		CreadoPor:     1,
		EstaTerminada: false,
	}

	err := service.CrearSubTarea(
		context.Background(),
		subTarea,
	)

	if err == nil {
		t.Fatal(
			"Se esperaba un error porque la tarea está eliminada",
		)
	}

	if subTareaRepository.CrearLlamado {
		t.Fatal(
			"No se esperaba que Crear fuera llamado",
		)
	}
}

func TestSubTareaService_CrearSubTarea_TareaNoExiste_RetornaError(t *testing.T) {

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	tareaRepository := &mocks.TareaRepositoryMock{
		Tareas: []entities.Tarea{},
	}

	service := serv.NewSubTareaService(
		subTareaRepository,
		tareaRepository,
	)

	subTarea := &entities.SubTarea{
		TareaId:       99,
		Titulo:        "Estudiar interfaces",
		CreadoPor:     1,
		EstaTerminada: false,
	}

	err := service.CrearSubTarea(
		context.Background(),
		subTarea,
	)

	if err == nil {
		t.Fatal(
			"Se esperaba un error porque la tarea no existe",
		)
	}

	if err.Error() != "la tarea no existe" {
		t.Fatalf(
			"Se esperaba error 'la tarea no existe', se obtuvo %q",
			err.Error(),
		)
	}

	if subTareaRepository.CrearLlamado {
		t.Fatal(
			"No se esperaba que Crear fuera llamado",
		)
	}
}

func TestSubTareaService_ActualizarSubTarea_SubTareaValida_ActualizaSubTarea(t *testing.T) {

	subTareaRepository := &mocks.SubTareaRepositoryMock{
		SubTarea: &entities.SubTarea{
			Id:            5,
			TareaId:       1,
			Titulo:        "Estudiar Go",
			EstaEliminada: false,
			CreadoPor:     1,
		},
	}

	tareaRepository := &mocks.TareaRepositoryMock{}

	service := serv.NewSubTareaService(
		subTareaRepository,
		tareaRepository,
	)

	subTarea := &entities.SubTarea{
		Id:            5,
		TareaId:       1,
		Titulo:        "Estudiar interfaces",
		CreadoPor:     1,
		EstaTerminada: true,
	}

	err := service.ActualizarSubTarea(
		context.Background(),
		subTarea,
	)

	if err != nil {
		t.Fatalf(
			"No se esperaba un error, se obtuvo: %v",
			err,
		)
	}

	if !subTareaRepository.ActualizarLlamado {
		t.Fatal(
			"Se esperaba que Actualizar fuera llamado",
		)
	}

	if subTareaRepository.SubTarea == nil {
		t.Fatal(
			"Se esperaba que se enviara una subtarea al repositorio",
		)
	}

	if subTareaRepository.SubTarea.Id != 5 {
		t.Fatalf(
			"Se esperaba Id 5, se obtuvo %d",
			subTareaRepository.SubTarea.Id,
		)
	}

	if subTareaRepository.SubTarea.Titulo != "Estudiar interfaces" {
		t.Fatalf(
			"Se esperaba título 'Estudiar interfaces', se obtuvo %s",
			subTareaRepository.SubTarea.Titulo,
		)
	}

	if !subTareaRepository.SubTarea.EstaTerminada {
		t.Fatal(
			"Se esperaba que la subtarea quedara terminada",
		)
	}
}

func TestSubTareaService_ActualizarSubTarea_SubTareaEliminada_RetornaError(t *testing.T) {

	subTareaRepository := &mocks.SubTareaRepositoryMock{
		SubTarea: &entities.SubTarea{
			Id:            10,
			TareaId:       5,
			Titulo:        "Subtarea eliminada",
			EstaEliminada: true,
		},
	}

	tareaRepository := &mocks.TareaRepositoryMock{}

	service := serv.NewSubTareaService(
		subTareaRepository,
		tareaRepository,
	)

	subTarea := &entities.SubTarea{
		Id:        10,
		TareaId:   5,
		Titulo:    "Intentar actualizar",
		CreadoPor: 1,
	}

	err := service.ActualizarSubTarea(
		context.Background(),
		subTarea,
	)

	if err == nil {
		t.Fatal(
			"Se esperaba un error porque la subtarea está eliminada",
		)
	}

	if err.Error() != "No se puede actualizar una subtarea eliminada." {
		t.Fatalf(
			"Se esperaba error de subtarea eliminada, se obtuvo %q",
			err.Error(),
		)
	}

	if subTareaRepository.ActualizarLlamado {
		t.Fatal(
			"Actualizar no debería haberse llamado",
		)
	}
}

func TestSubTareaService_ActualizarSubTarea_IdInvalido_RetornaError(t *testing.T) {

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	tareaRepository := &mocks.TareaRepositoryMock{}

	service := serv.NewSubTareaService(
		subTareaRepository,
		tareaRepository,
	)

	subTarea := &entities.SubTarea{
		Id:        0,
		TareaId:   5,
		Titulo:    "Estudiar interfaces",
		CreadoPor: 1,
	}

	err := service.ActualizarSubTarea(
		context.Background(),
		subTarea,
	)

	if err == nil {
		t.Fatal(
			"Se esperaba un error porque el id de la subtarea no es válido",
		)
	}

	if err.Error() != "el id de la subtarea no es válido" {
		t.Fatalf(
			"Se esperaba error de id inválido, se obtuvo %q",
			err.Error(),
		)
	}

	if subTareaRepository.ActualizarLlamado {
		t.Fatal(
			"Actualizar no debería haberse llamado",
		)
	}

	if subTareaRepository.SubTarea != nil {
		t.Fatal(
			"El repositorio no debería haber recibido una subtarea",
		)
	}
}

func TestSubTareaService_ActualizarSubTarea_ErrorObtenerPorId_RetornaError(t *testing.T) {

	errRepositorio := errors.New("error al consultar la subtarea")

	subTareaRepository := &mocks.SubTareaRepositoryMock{
		ErrorObtenerPorId: errRepositorio,
	}

	tareaRepository := &mocks.TareaRepositoryMock{}

	service := serv.NewSubTareaService(
		subTareaRepository,
		tareaRepository,
	)

	subTarea := &entities.SubTarea{
		Id:        10,
		TareaId:   5,
		Titulo:    "Estudiar interfaces",
		CreadoPor: 1,
	}

	err := service.ActualizarSubTarea(
		context.Background(),
		subTarea,
	)

	if err == nil {
		t.Fatal(
			"Se esperaba un error del repositorio",
		)
	}

	if err.Error() != "error al consultar la subtarea" {
		t.Fatalf(
			"Se esperaba el error del repositorio, se obtuvo %q",
			err.Error(),
		)
	}

	if subTareaRepository.ActualizarLlamado {
		t.Fatal(
			"Actualizar no debería haberse llamado",
		)
	}
}

func TestSubTareaService_ActualizarSubTarea_ErrorActualizar_RetornaError(t *testing.T) {

	errRepositorio := errors.New("error al actualizar la subtarea")

	subTareaRepository := &mocks.SubTareaRepositoryMock{
		SubTarea: &entities.SubTarea{
			Id:            10,
			TareaId:       5,
			Titulo:        "Título anterior",
			EstaEliminada: false,
		},
		ErrorActualizar: errRepositorio,
	}

	tareaRepository := &mocks.TareaRepositoryMock{}

	service := serv.NewSubTareaService(
		subTareaRepository,
		tareaRepository,
	)

	subTarea := &entities.SubTarea{
		Id:        10,
		TareaId:   5,
		Titulo:    "Título actualizado",
		CreadoPor: 1,
	}

	err := service.ActualizarSubTarea(
		context.Background(),
		subTarea,
	)

	if err == nil {
		t.Fatal(
			"Se esperaba un error al actualizar la subtarea",
		)
	}

	if err.Error() != "error al actualizar la subtarea" {
		t.Fatalf(
			"Se esperaba el error del repositorio, se obtuvo %q",
			err.Error(),
		)
	}

	if !subTareaRepository.ActualizarLlamado {
		t.Fatal(
			"Se esperaba que Actualizar fuera llamado",
		)
	}
}

func TestSubTareaService_EliminarSubTarea_SubTareaValida_EliminaSubTarea(t *testing.T) {

	subTareaRepository := &mocks.SubTareaRepositoryMock{
		SubTarea: &entities.SubTarea{
			Id:            10,
			TareaId:       5,
			Titulo:        "Estudiar interfaces",
			EstaEliminada: false,
			CreadoPor:     1,
		},
	}

	tareaRepository := &mocks.TareaRepositoryMock{}

	service := serv.NewSubTareaService(
		subTareaRepository,
		tareaRepository,
	)

	err := service.EliminarSubTarea(
		context.Background(),
		10,
		2,
	)

	if err != nil {
		t.Fatalf(
			"No se esperaba un error, se obtuvo: %v",
			err,
		)
	}

	if !subTareaRepository.ActualizarLlamado {
		t.Fatal(
			"Se esperaba que Actualizar fuera llamado",
		)
	}

	if subTareaRepository.SubTarea == nil {
		t.Fatal(
			"Se esperaba que se enviara una subtarea al repositorio",
		)
	}

	if !subTareaRepository.SubTarea.EstaEliminada {
		t.Fatal(
			"Se esperaba que la subtarea quedara eliminada",
		)
	}

	if subTareaRepository.SubTarea.Id != 10 {
		t.Fatalf(
			"Se esperaba Id 10, se obtuvo %d",
			subTareaRepository.SubTarea.Id,
		)
	}

	if subTareaRepository.SubTarea.ModificadoPor == nil {
		t.Fatal(
			"Se esperaba ModificadoPor",
		)
	}

	if *subTareaRepository.SubTarea.ModificadoPor != 2 {
		t.Fatalf(
			"Se esperaba ModificadoPor 2, se obtuvo %d",
			*subTareaRepository.SubTarea.ModificadoPor,
		)
	}

	if subTareaRepository.SubTarea.ModificadoEl == nil {
		t.Fatal(
			"Se esperaba ModificadoEl",
		)
	}
}

func TestSubTareaService_EliminarSubTarea_SubTareaYaEliminada_RetornaError(t *testing.T) {

	subTareaRepository := &mocks.SubTareaRepositoryMock{
		SubTarea: &entities.SubTarea{
			Id:            10,
			TareaId:       5,
			Titulo:        "Subtarea eliminada",
			EstaEliminada: true,
		},
	}

	tareaRepository := &mocks.TareaRepositoryMock{}

	service := serv.NewSubTareaService(
		subTareaRepository,
		tareaRepository,
	)

	err := service.EliminarSubTarea(
		context.Background(),
		10,
		2,
	)

	if err == nil {
		t.Fatal(
			"Se esperaba un error porque la subtarea ya está eliminada",
		)
	}

	if err.Error() != "La subtarea ya está eliminada." {
		t.Fatalf(
			"Se esperaba error de subtarea ya eliminada, se obtuvo %q",
			err.Error(),
		)
	}

	if subTareaRepository.ActualizarLlamado {
		t.Fatal(
			"Actualizar no debería haberse llamado",
		)
	}
}

func TestSubTareaService_EliminarSubTarea_SubTareaNoExiste_RetornaError(t *testing.T) {

	subTareaRepository := &mocks.SubTareaRepositoryMock{
		SubTarea: nil,
	}

	tareaRepository := &mocks.TareaRepositoryMock{}

	service := serv.NewSubTareaService(
		subTareaRepository,
		tareaRepository,
	)

	err := service.EliminarSubTarea(
		context.Background(),
		10,
		2,
	)

	if err == nil {
		t.Fatal(
			"Se esperaba un error porque la subtarea no existe",
		)
	}

	if err.Error() != "la subtarea no existe" {
		t.Fatalf(
			"Se esperaba error 'la subtarea no existe', se obtuvo %q",
			err.Error(),
		)
	}

	if subTareaRepository.ActualizarLlamado {
		t.Fatal(
			"Actualizar no debería haberse llamado",
		)
	}
}

func TestSubTareaService_EliminarSubTarea_IdInvalido_RetornaError(t *testing.T) {

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	tareaRepository := &mocks.TareaRepositoryMock{}

	service := serv.NewSubTareaService(
		subTareaRepository,
		tareaRepository,
	)

	err := service.EliminarSubTarea(
		context.Background(),
		0,
		2,
	)

	if err == nil {
		t.Fatal(
			"Se esperaba un error porque el id de la subtarea no es válido",
		)
	}

	if err.Error() != "el id de la subtarea no es válido" {
		t.Fatalf(
			"Se esperaba error de id inválido, se obtuvo %q",
			err.Error(),
		)
	}

	if subTareaRepository.ActualizarLlamado {
		t.Fatal(
			"Actualizar no debería haberse llamado",
		)
	}
}

func TestSubTareaService_EliminarSubTarea_ModificadoPorInvalido_RetornaError(t *testing.T) {

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	tareaRepository := &mocks.TareaRepositoryMock{}

	service := serv.NewSubTareaService(
		subTareaRepository,
		tareaRepository,
	)

	err := service.EliminarSubTarea(
		context.Background(),
		10,
		0,
	)

	if err == nil {
		t.Fatal(
			"Se esperaba un error porque el usuario modificador no es válido",
		)
	}

	if err.Error() != "el usuario modificador no es válido" {
		t.Fatalf(
			"Se esperaba error de usuario modificador inválido, se obtuvo %q",
			err.Error(),
		)
	}

	if subTareaRepository.ActualizarLlamado {
		t.Fatal(
			"Actualizar no debería haberse llamado",
		)
	}
}

func TestSubTareaService_EliminarSubTarea_ErrorObtenerPorId_RetornaError(t *testing.T) {

	errRepositorio := errors.New("error al consultar la subtarea")

	subTareaRepository := &mocks.SubTareaRepositoryMock{
		ErrorObtenerPorId: errRepositorio,
	}

	tareaRepository := &mocks.TareaRepositoryMock{}

	service := serv.NewSubTareaService(
		subTareaRepository,
		tareaRepository,
	)

	err := service.EliminarSubTarea(
		context.Background(),
		10,
		2,
	)

	if err == nil {
		t.Fatal(
			"Se esperaba un error del repositorio",
		)
	}

	if err.Error() != "error al consultar la subtarea" {
		t.Fatalf(
			"Se esperaba el error del repositorio, se obtuvo %q",
			err.Error(),
		)
	}

	if subTareaRepository.ActualizarLlamado {
		t.Fatal(
			"Actualizar no debería haberse llamado",
		)
	}
}

func TestObtenerPorTareaId_Exitoso(t *testing.T) {

	mock := &SubTareaServiceMock{
		SubTareasPorTarea: []entities.SubTarea{
			{
				Id:      1,
				TareaId: 10,
				Titulo:  "Subtarea de prueba",
			},
		},
	}

	handler := NewSubTareaHandler(mock)

	router := gin.Default()

	router.GET(
		"/tareas/:id/subtareas",
		handler.ObtenerPorTareaId,
	)

	request := httptest.NewRequest(
		http.MethodGet,
		"/tareas/10/subtareas",
		nil,
	)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf(
			"Se esperaba status 200, se obtuvo %d",
			response.Code,
		)
	}

	if !mock.ObtenerPorTareaIdLlamado {
		t.Error("Se esperaba que ObtenerPorTareaId fuera llamado")
	}
}

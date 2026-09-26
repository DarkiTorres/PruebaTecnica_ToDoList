package services

import (
	"context"
	"errors"
	"strings"
	"testing"
	mocks "to-do-server/internal/Application/Tests/Mocks"
	validators "to-do-server/internal/Application/Validators"
	entities "to-do-server/internal/Domain/Entities"
)

func TestCompletarTarea_SinSubTareas_CompletaTarea(t *testing.T) {
	tareaRepository := &mocks.TareaRepositoryMock{
		Tarea: &entities.Tarea{
			Id:            1,
			Titulo:        "Test",
			EstaTerminada: false,
		},
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{
		SubTareas: []entities.SubTarea{},
	}

	usuarioRepository := &mocks.UsuarioRepositoryMock{}
	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	err := service.CompletarTarea(
		context.Background(),
		1,
	)

	if err != nil {
		t.Fatalf(
			"Se esperaba que no hubiera error, se obtuvo: %v",
			err,
		)
	}

	if !tareaRepository.Tarea.EstaTerminada {
		t.Fatal("Se esperaba que la tarea estuviera terminada")
	}

	if !tareaRepository.ActualizarLlamado {
		t.Fatal("Se esperaba que Actualizar fuera llamado")
	}
}

func TestCompletarTarea_ConSubTareaPendiente_NoCompleta(t *testing.T) {
	tareaRepository := &mocks.TareaRepositoryMock{
		Tarea: &entities.Tarea{
			Id:            1,
			Titulo:        "Test",
			EstaTerminada: false,
		},
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{
		SubTareas: []entities.SubTarea{
			{
				Id:            1,
				TareaId:       1,
				Titulo:        "Subtarea pendiente",
				EstaTerminada: false,
				EstaEliminada: false,
			},
		},
	}

	usuarioRepository := &mocks.UsuarioRepositoryMock{}
	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	err := service.CompletarTarea(
		context.Background(),
		1,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareaRepository.Tarea.EstaTerminada {
		t.Fatal("La tarea no debería haberse completado")
	}

	if tareaRepository.ActualizarLlamado {
		t.Fatal("Actualizar no debería haberse llamado")
	}
}

func TestCompletarTarea_ConTodasLasSubTareasTerminadas_CompletaTarea(t *testing.T) {
	tareaRepository := &mocks.TareaRepositoryMock{
		Tarea: &entities.Tarea{
			Id:            1,
			Titulo:        "Test",
			EstaTerminada: false,
		},
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{
		SubTareas: []entities.SubTarea{
			{
				Id:            1,
				TareaId:       1,
				EstaTerminada: true,
			},
			{
				Id:            2,
				TareaId:       1,
				EstaTerminada: true,
			},
		},
	}

	usuarioRepository := &mocks.UsuarioRepositoryMock{}
	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	err := service.CompletarTarea(
		context.Background(),
		1,
	)

	if err != nil {
		t.Fatalf("No se esperaba error: %v", err)
	}

	if !tareaRepository.Tarea.EstaTerminada {
		t.Fatal("La tarea debería estar terminada")
	}
}

func TestCompletarTarea_ConSubTareaEliminada_NoBloquea(t *testing.T) {
	tareaRepository := &mocks.TareaRepositoryMock{
		Tarea: &entities.Tarea{
			Id:            1,
			Titulo:        "Test",
			EstaTerminada: false,
		},
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{
		SubTareas: []entities.SubTarea{
			{
				Id:            1,
				TareaId:       1,
				EstaTerminada: false,
				EstaEliminada: true,
			},
		},
	}

	usuarioRepository := &mocks.UsuarioRepositoryMock{}
	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	err := service.CompletarTarea(
		context.Background(),
		1,
	)

	if err != nil {
		t.Fatalf("No se esperaba error: %v", err)
	}

	if !tareaRepository.Tarea.EstaTerminada {
		t.Fatal("La tarea debería estar terminada")
	}
}

func TestCompletarTarea_ErrorAlObtenerTarea_RegresaError(t *testing.T) {

	errorEsperado := errors.New("error de base de datos")

	tareaRepository := &mocks.TareaRepositoryMock{
		ErrorObtenerPorId: errorEsperado,
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{}
	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	err := service.CompletarTarea(
		context.Background(),
		1,
	)

	if !errors.Is(err, errorEsperado) {
		t.Fatalf(
			"Se esperaba el error %v, se obtuvo %v",
			errorEsperado,
			err,
		)
	}
}

func TestCrearTarea_TareaValida_CreaTarea(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{
		Usuarios: []entities.Usuario{
			{
				Id:              1,
				Nombre:          "Lider",
				RolId:           1,
				EstaDesactivado: false,
			},
		},
	}

	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	tarea := &entities.Tarea{
		Titulo:      "Estudiar Go",
		PrioridadId: 1,
		CreadoPor:   1,
	}

	err := service.CrearTarea(
		context.Background(),
		tarea,
		1,
	)

	if err != nil {
		t.Fatalf(
			"No se esperaba error, se obtuvo: %v",
			err,
		)
	}

	if !tareaRepository.CrearLlamado {
		t.Fatal("Se esperaba que Crear fuera llamado")
	}

	if tareaRepository.Tarea != tarea {
		t.Fatal("Se esperaba que se enviara la misma tarea al repository")
	}
}

func TestCrearTarea_TituloVacio_NoCreaTarea(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{}
	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	tarea := &entities.Tarea{
		Titulo:      "",
		PrioridadId: 1,
		CreadoPor:   1,
	}

	err := service.CrearTarea(
		context.Background(),
		tarea,
		1,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareaRepository.CrearLlamado {
		t.Fatal("Crear no debería haberse llamado")
	}
}

func TestCrearTarea_TituloDemasiadoLargo_NoCreaTarea(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{}
	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	titulo := strings.Repeat("A", validators.MaxTituloTarea+1)

	tarea := &entities.Tarea{
		Titulo:      titulo,
		PrioridadId: 1,
		CreadoPor:   1,
	}

	err := service.CrearTarea(
		context.Background(),
		tarea,
		1,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareaRepository.CrearLlamado {
		t.Fatal("Crear no debería haberse llamado")
	}
}

func TestCrearTarea_ErrorAlCrear_RegresaError(t *testing.T) {

	errorEsperado := errors.New("error de base de datos")

	tareaRepository := &mocks.TareaRepositoryMock{
		ErrorCrear: errorEsperado,
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{
		Usuarios: []entities.Usuario{
			{
				Id:              1,
				Nombre:          "Lider",
				RolId:           1,
				EstaDesactivado: false,
			},
		},
	}

	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	tarea := &entities.Tarea{
		Titulo:      "Estudiar Go",
		PrioridadId: 1,
		CreadoPor:   1,
	}

	err := service.CrearTarea(
		context.Background(),
		tarea,
		1,
	)

	if !errors.Is(err, errorEsperado) {
		t.Fatalf(
			"Se esperaba el error %v, se obtuvo %v",
			errorEsperado,
			err,
		)
	}
}

func TestActualizarTarea_TareaValida_ActualizaTarea(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{}
	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	tarea := &entities.Tarea{
		Id:          1,
		Titulo:      "Estudiar Go",
		PrioridadId: 1,
		CreadoPor:   1,
	}

	err := service.ActualizarTarea(
		context.Background(),
		tarea,
	)

	if err != nil {
		t.Fatalf(
			"No se esperaba error, se obtuvo: %v",
			err,
		)
	}

	if !tareaRepository.ActualizarLlamado {
		t.Fatal("Se esperaba que Actualizar fuera llamado")
	}

	if tareaRepository.Tarea != tarea {
		t.Fatal("Se esperaba que se enviara la misma tarea al repository")
	}
}

func TestActualizarTarea_TituloVacio_NoActualiza(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{}
	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	tarea := &entities.Tarea{
		Id:          1,
		Titulo:      "",
		PrioridadId: 1,
		CreadoPor:   1,
	}

	err := service.ActualizarTarea(
		context.Background(),
		tarea,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareaRepository.ActualizarLlamado {
		t.Fatal("Actualizar no debería haberse llamado")
	}
}

func TestActualizarTarea_ErrorAlActualizar_RegresaError(t *testing.T) {

	errorEsperado := errors.New("error de base de datos")

	tareaRepository := &mocks.TareaRepositoryMock{
		ErrorActualizar: errorEsperado,
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{}
	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	tarea := &entities.Tarea{
		Id:          1,
		Titulo:      "Estudiar Go",
		PrioridadId: 1,
		CreadoPor:   1,
	}

	err := service.ActualizarTarea(
		context.Background(),
		tarea,
	)

	if !errors.Is(err, errorEsperado) {
		t.Fatalf(
			"Se esperaba el error %v, se obtuvo %v",
			errorEsperado,
			err,
		)
	}
}

func TestEliminarTarea_IdInvalido_NoElimina(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}
	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{}
	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	err := service.EliminarTarea(
		context.Background(),
		0,
		1,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareaRepository.ActualizarLlamado {
		t.Fatal("Actualizar no debería haberse llamado")
	}
}

func TestEliminarTarea_UsuarioInvalido_NoElimina(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}
	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{}
	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	err := service.EliminarTarea(
		context.Background(),
		1,
		0,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareaRepository.ActualizarLlamado {
		t.Fatal("Actualizar no debería haberse llamado")
	}
}

func TestEliminarTarea_TareaValida_RealizaSoftDelete(t *testing.T) {

	tarea := &entities.Tarea{
		Id:            1,
		Titulo:        "Estudiar Go",
		EstaEliminada: false,
		CreadoPor:     1,
	}

	tareaRepository := &mocks.TareaRepositoryMock{
		Tarea: tarea,
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{}
	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	err := service.EliminarTarea(
		context.Background(),
		1,
		2,
	)

	if err != nil {
		t.Fatalf(
			"No se esperaba error, se obtuvo: %v",
			err,
		)
	}

	if !tareaRepository.ActualizarLlamado {
		t.Fatal("Se esperaba que Actualizar fuera llamado")
	}

	if !tareaRepository.Tarea.EstaEliminada {
		t.Fatal("La tarea debería estar marcada como eliminada")
	}

	if tareaRepository.Tarea.ModificadoEl == nil {
		t.Fatal("ModificadoEl debería tener un valor")
	}

	if tareaRepository.Tarea.ModificadoPor == nil {
		t.Fatal("ModificadoPor debería tener un valor")
	}

	if *tareaRepository.Tarea.ModificadoPor != 2 {
		t.Fatalf(
			"Se esperaba ModificadoPor = 2, se obtuvo %d",
			*tareaRepository.Tarea.ModificadoPor,
		)
	}
}

func TestEliminarTarea_ErrorAlObtenerTarea_RegresaError(t *testing.T) {

	errorEsperado := errors.New("error de base de datos")

	tareaRepository := &mocks.TareaRepositoryMock{
		ErrorObtenerPorId: errorEsperado,
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{}
	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	err := service.EliminarTarea(
		context.Background(),
		1,
		2,
	)

	if !errors.Is(err, errorEsperado) {
		t.Fatalf(
			"Se esperaba el error %v, se obtuvo %v",
			errorEsperado,
			err,
		)
	}

	if tareaRepository.ActualizarLlamado {
		t.Fatal("Actualizar no debería haberse llamado")
	}
}

func TestEliminarTarea_ErrorAlActualizar_RegresaError(t *testing.T) {

	errorEsperado := errors.New("error de base de datos")

	tareaRepository := &mocks.TareaRepositoryMock{
		Tarea: &entities.Tarea{
			Id:            1,
			Titulo:        "Estudiar Go",
			EstaEliminada: false,
			CreadoPor:     1,
		},
		ErrorActualizar: errorEsperado,
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{}
	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	err := service.EliminarTarea(
		context.Background(),
		1,
		2,
	)

	if !errors.Is(err, errorEsperado) {
		t.Fatalf(
			"Se esperaba el error %v, se obtuvo %v",
			errorEsperado,
			err,
		)
	}
}

func TestCrearTarea_LiderAsignaAContribuidor_CreaTareaYAsignacion(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{
		Usuarios: []entities.Usuario{
			{
				Id:              1,
				Nombre:          "Lider",
				RolId:           1,
				EstaDesactivado: false,
			},
			{
				Id:              2,
				Nombre:          "Contribuidor",
				RolId:           2,
				EstaDesactivado: false,
			},
		},
	}

	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	tarea := &entities.Tarea{
		Titulo:      "Tarea del contribuidor",
		PrioridadId: 1,
		CreadoPor:   1,
	}

	err := service.CrearTarea(
		context.Background(),
		tarea,
		2,
	)

	if err != nil {
		t.Fatalf(
			"No se esperaba error, se obtuvo: %v",
			err,
		)
	}

	if !tareaRepository.CrearLlamado {
		t.Fatal("Se esperaba que Crear de TareaRepository fuera llamado")
	}

	if !tareaUsuarioRepository.CrearLlamado {
		t.Fatal("Se esperaba que Crear de TareaUsuarioRepository fuera llamado")
	}
}

func TestCrearTarea_ContribuidorSeAsignaASiMismo_CreaTarea(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{
		Usuarios: []entities.Usuario{
			{
				Id:              2,
				Nombre:          "Contribuidor",
				RolId:           2,
				EstaDesactivado: false,
			},
		},
	}

	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	tarea := &entities.Tarea{
		Titulo:      "Mi tarea",
		PrioridadId: 1,
		CreadoPor:   2,
	}

	err := service.CrearTarea(
		context.Background(),
		tarea,
		2,
	)

	if err != nil {
		t.Fatalf(
			"No se esperaba error, se obtuvo: %v",
			err,
		)
	}

	if !tareaRepository.CrearLlamado {
		t.Fatal("Se esperaba que la tarea fuera creada")
	}

	if !tareaUsuarioRepository.CrearLlamado {
		t.Fatal("Se esperaba que la asignación fuera creada")
	}
}

func TestCrearTarea_ContribuidorAsignaALider_NoCreaTarea(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{
		Usuarios: []entities.Usuario{
			{
				Id:              1,
				Nombre:          "Lider",
				RolId:           1,
				EstaDesactivado: false,
			},
			{
				Id:              2,
				Nombre:          "Contribuidor",
				RolId:           2,
				EstaDesactivado: false,
			},
		},
	}

	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	tarea := &entities.Tarea{
		Titulo:      "Tarea no permitida",
		PrioridadId: 1,
		CreadoPor:   2,
	}

	err := service.CrearTarea(
		context.Background(),
		tarea,
		1,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareaRepository.CrearLlamado {
		t.Fatal("La tarea no debería haberse creado")
	}

	if tareaUsuarioRepository.CrearLlamado {
		t.Fatal("La asignación no debería haberse creado")
	}
}

func TestCrearTarea_CreadorInexistente_NoCreaTarea(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}
	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{
		Usuarios: []entities.Usuario{},
	}

	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	tarea := &entities.Tarea{
		Titulo:      "Tarea de prueba",
		PrioridadId: 1,
		CreadoPor:   99,
	}

	err := service.CrearTarea(
		context.Background(),
		tarea,
		1,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareaRepository.CrearLlamado {
		t.Fatal("La tarea no debería haberse creado")
	}

	if tareaUsuarioRepository.CrearLlamado {
		t.Fatal("La asignación no debería haberse creado")
	}
}

func TestCrearTarea_UsuarioAsignadoInexistente_NoCreaTarea(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}
	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{
		Usuarios: []entities.Usuario{
			{
				Id:              1,
				Nombre:          "Lider",
				RolId:           1,
				EstaDesactivado: false,
			},
		},
	}

	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	tarea := &entities.Tarea{
		Titulo:      "Tarea de prueba",
		PrioridadId: 1,
		CreadoPor:   1,
	}

	err := service.CrearTarea(
		context.Background(),
		tarea,
		99,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareaRepository.CrearLlamado {
		t.Fatal("La tarea no debería haberse creado")
	}

	if tareaUsuarioRepository.CrearLlamado {
		t.Fatal("La asignación no debería haberse creado")
	}
}

func TestCrearTarea_CreadorDesactivado_NoCreaTarea(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}
	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{
		Usuarios: []entities.Usuario{
			{
				Id:              1,
				Nombre:          "Lider",
				RolId:           1,
				EstaDesactivado: true,
			},
		},
	}

	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	tarea := &entities.Tarea{
		Titulo:      "Tarea de prueba",
		PrioridadId: 1,
		CreadoPor:   1,
	}

	err := service.CrearTarea(
		context.Background(),
		tarea,
		1,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareaRepository.CrearLlamado {
		t.Fatal("La tarea no debería haberse creado")
	}

	if tareaUsuarioRepository.CrearLlamado {
		t.Fatal("La asignación no debería haberse creado")
	}
}

func TestCrearTarea_UsuarioAsignadoDesactivado_NoCreaTarea(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}
	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{
		Usuarios: []entities.Usuario{
			{
				Id:              1,
				Nombre:          "Lider",
				RolId:           1,
				EstaDesactivado: false,
			},
			{
				Id:              2,
				Nombre:          "Contribuidor",
				RolId:           2,
				EstaDesactivado: true,
			},
		},
	}

	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	tarea := &entities.Tarea{
		Titulo:      "Tarea de prueba",
		PrioridadId: 1,
		CreadoPor:   1,
	}

	err := service.CrearTarea(
		context.Background(),
		tarea,
		2,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareaRepository.CrearLlamado {
		t.Fatal("La tarea no debería haberse creado")
	}

	if tareaUsuarioRepository.CrearLlamado {
		t.Fatal("La asignación no debería haberse creado")
	}
}

func TestCrearTarea_ErrorAlObtenerUsuario_RegresaError(t *testing.T) {

	errorEsperado := errors.New("error de base de datos")

	tareaRepository := &mocks.TareaRepositoryMock{}
	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{
		ErrorObtenerPorId: errorEsperado,
	}

	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	tarea := &entities.Tarea{
		Titulo:      "Tarea de prueba",
		PrioridadId: 1,
		CreadoPor:   1,
	}

	err := service.CrearTarea(
		context.Background(),
		tarea,
		1,
	)

	if !errors.Is(err, errorEsperado) {
		t.Fatalf(
			"Se esperaba el error %v, se obtuvo %v",
			errorEsperado,
			err,
		)
	}

	if tareaRepository.CrearLlamado {
		t.Fatal("La tarea no debería haberse creado")
	}

	if tareaUsuarioRepository.CrearLlamado {
		t.Fatal("La asignación no debería haberse creado")
	}
}

func TestCrearTarea_ErrorAlCrearTarea_RegresaError(t *testing.T) {

	errorEsperado := errors.New("error al crear tarea")

	tareaRepository := &mocks.TareaRepositoryMock{
		ErrorCrear: errorEsperado,
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{
		Usuarios: []entities.Usuario{
			{
				Id:              1,
				Nombre:          "Lider",
				RolId:           1,
				EstaDesactivado: false,
			},
		},
	}

	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	tarea := &entities.Tarea{
		Titulo:      "Tarea de prueba",
		PrioridadId: 1,
		CreadoPor:   1,
	}

	err := service.CrearTarea(
		context.Background(),
		tarea,
		1,
	)

	if !errors.Is(err, errorEsperado) {
		t.Fatalf(
			"Se esperaba el error %v, se obtuvo %v",
			errorEsperado,
			err,
		)
	}

	if tareaUsuarioRepository.CrearLlamado {
		t.Fatal(
			"La asignación no debería crearse si la tarea no fue creada",
		)
	}
}

func TestCrearTarea_ErrorAlCrearAsignacion_RegresaError(t *testing.T) {

	errorEsperado := errors.New("error al crear asignacion")

	tareaRepository := &mocks.TareaRepositoryMock{}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{
		Usuarios: []entities.Usuario{
			{
				Id:              1,
				Nombre:          "Lider",
				RolId:           1,
				EstaDesactivado: false,
			},
		},
	}

	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{
		ErrorCrear: errorEsperado,
	}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	tarea := &entities.Tarea{
		Titulo:      "Tarea de prueba",
		PrioridadId: 1,
		CreadoPor:   1,
	}

	err := service.CrearTarea(
		context.Background(),
		tarea,
		1,
	)

	if !errors.Is(err, errorEsperado) {
		t.Fatalf(
			"Se esperaba el error %v, se obtuvo %v",
			errorEsperado,
			err,
		)
	}

	if !tareaRepository.CrearLlamado {
		t.Fatal("Se esperaba que la tarea fuera creada")
	}

	if !tareaUsuarioRepository.CrearLlamado {
		t.Fatal("Se esperaba que se intentara crear la asignación")
	}
}

func TestCompletarTarea_ErrorAlObtenerSubTareas_RegresaError(t *testing.T) {

	errorEsperado := errors.New("error de base de datos")

	tareaRepository := &mocks.TareaRepositoryMock{
		Tarea: &entities.Tarea{
			Id:            1,
			Titulo:        "Tarea de prueba",
			EstaTerminada: false,
		},
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{
		ErrorObtenerPorTareaId: errorEsperado,
	}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		&mocks.UsuarioRepositoryMock{},
		&mocks.TareaUsuarioRepositoryMock{},
	)

	err := service.CompletarTarea(
		context.Background(),
		1,
	)

	if !errors.Is(err, errorEsperado) {
		t.Fatalf(
			"Se esperaba el error %v, se obtuvo %v",
			errorEsperado,
			err,
		)
	}

	if tareaRepository.ActualizarLlamado {
		t.Fatal("Actualizar no debería haberse llamado")
	}

	if tareaRepository.Tarea.EstaTerminada {
		t.Fatal("La tarea no debería haberse completado")
	}
}

func TestCompletarTarea_ErrorAlActualizar_RegresaError(t *testing.T) {

	errorEsperado := errors.New("error al actualizar")

	tareaRepository := &mocks.TareaRepositoryMock{
		Tarea: &entities.Tarea{
			Id:            1,
			Titulo:        "Tarea de prueba",
			EstaTerminada: false,
		},
		ErrorActualizar: errorEsperado,
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{
		SubTareas: []entities.SubTarea{},
	}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		&mocks.UsuarioRepositoryMock{},
		&mocks.TareaUsuarioRepositoryMock{},
	)

	err := service.CompletarTarea(
		context.Background(),
		1,
	)

	if !errors.Is(err, errorEsperado) {
		t.Fatalf(
			"Se esperaba el error %v, se obtuvo %v",
			errorEsperado,
			err,
		)
	}

	if !tareaRepository.ActualizarLlamado {
		t.Fatal("Se esperaba que Actualizar fuera llamado")
	}

	if !tareaRepository.Tarea.EstaTerminada {
		t.Fatal("La tarea debería haberse marcado como terminada")
	}
}

func TestCompletarTarea_TareaInexistente_RegresaError(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{
		Tarea: nil,
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		&mocks.UsuarioRepositoryMock{},
		&mocks.TareaUsuarioRepositoryMock{},
	)

	err := service.CompletarTarea(
		context.Background(),
		1,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareaRepository.ActualizarLlamado {
		t.Fatal("Actualizar no debería haberse llamado")
	}
}

func TestCompletarTarea_IdInvalido_RegresaError(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}
	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		&mocks.UsuarioRepositoryMock{},
		&mocks.TareaUsuarioRepositoryMock{},
	)

	err := service.CompletarTarea(
		context.Background(),
		0,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareaRepository.ObtenerPorIdLlamado {
		t.Fatal("No debería haberse consultado la tarea")
	}
}

func TestActualizarTarea_TareaNil_RegresaError(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		&mocks.SubTareaRepositoryMock{},
		&mocks.UsuarioRepositoryMock{},
		&mocks.TareaUsuarioRepositoryMock{},
	)

	err := service.ActualizarTarea(
		context.Background(),
		nil,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareaRepository.ActualizarLlamado {
		t.Fatal("Actualizar no debería haberse llamado")
	}
}

func TestActualizarTarea_IdInvalido_RegresaError(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		&mocks.SubTareaRepositoryMock{},
		&mocks.UsuarioRepositoryMock{},
		&mocks.TareaUsuarioRepositoryMock{},
	)

	tarea := &entities.Tarea{
		Id:          0,
		Titulo:      "Tarea",
		PrioridadId: 1,
		CreadoPor:   1,
	}

	err := service.ActualizarTarea(
		context.Background(),
		tarea,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareaRepository.ActualizarLlamado {
		t.Fatal("Actualizar no debería haberse llamado")
	}
}

func TestActualizarTarea_TituloVacio_RegresaError(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		&mocks.SubTareaRepositoryMock{},
		&mocks.UsuarioRepositoryMock{},
		&mocks.TareaUsuarioRepositoryMock{},
	)

	tarea := &entities.Tarea{
		Id:          1,
		Titulo:      "",
		PrioridadId: 1,
		CreadoPor:   1,
	}

	err := service.ActualizarTarea(
		context.Background(),
		tarea,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareaRepository.ActualizarLlamado {
		t.Fatal("Actualizar no debería haberse llamado")
	}
}
func TestEliminarTarea_TareaValida_MarcaComoEliminada(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{
		Tarea: &entities.Tarea{
			Id:            1,
			Titulo:        "Tarea de prueba",
			EstaEliminada: false,
		},
	}

	service := NewTareaService(
		tareaRepository,
		&mocks.SubTareaRepositoryMock{},
		&mocks.UsuarioRepositoryMock{},
		&mocks.TareaUsuarioRepositoryMock{},
	)

	err := service.EliminarTarea(
		context.Background(),
		1,
		5,
	)

	if err != nil {
		t.Fatalf("No se esperaba error, se obtuvo: %v", err)
	}

	if !tareaRepository.ActualizarLlamado {
		t.Fatal("Se esperaba que Actualizar fuera llamado")
	}

	if !tareaRepository.Tarea.EstaEliminada {
		t.Fatal("La tarea debería estar marcada como eliminada")
	}

	if tareaRepository.Tarea.ModificadoPor == nil {
		t.Fatal("Se esperaba que ModificadoPor tuviera valor")
	}

	if *tareaRepository.Tarea.ModificadoPor != 5 {
		t.Fatalf(
			"Se esperaba ModificadoPor = 5, se obtuvo %d",
			*tareaRepository.Tarea.ModificadoPor,
		)
	}

	if tareaRepository.Tarea.ModificadoEl == nil {
		t.Fatal("Se esperaba que ModificadoEl tuviera valor")
	}
}

func TestEliminarTarea_IdInvalido_RegresaError(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		&mocks.SubTareaRepositoryMock{},
		&mocks.UsuarioRepositoryMock{},
		&mocks.TareaUsuarioRepositoryMock{},
	)

	err := service.EliminarTarea(
		context.Background(),
		0,
		5,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareaRepository.ActualizarLlamado {
		t.Fatal("Actualizar no debería haberse llamado")
	}
}

func TestEliminarTarea_ModificadoPorInvalido_RegresaError(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		&mocks.SubTareaRepositoryMock{},
		&mocks.UsuarioRepositoryMock{},
		&mocks.TareaUsuarioRepositoryMock{},
	)

	err := service.EliminarTarea(
		context.Background(),
		1,
		0,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareaRepository.ActualizarLlamado {
		t.Fatal("Actualizar no debería haberse llamado")
	}
}

func TestEliminarTarea_TareaInexistente_RegresaError(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{
		Tarea: nil,
	}

	service := NewTareaService(
		tareaRepository,
		&mocks.SubTareaRepositoryMock{},
		&mocks.UsuarioRepositoryMock{},
		&mocks.TareaUsuarioRepositoryMock{},
	)

	err := service.EliminarTarea(
		context.Background(),
		1,
		5,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareaRepository.ActualizarLlamado {
		t.Fatal("Actualizar no debería haberse llamado")
	}
}

func TestObtenerTareasPorUsuario_UsuarioValido_DevuelveTareas(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{
		Tareas: []entities.Tarea{
			{
				Id:          1,
				Titulo:      "Estudiar Go",
				PrioridadId: 1,
			},
			{
				Id:          2,
				Titulo:      "Estudiar PostgreSQL",
				PrioridadId: 2,
			},
		},
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{
		Usuario: &entities.Usuario{
			Id:              1,
			Nombre:          "Lider",
			RolId:           1,
			EstaDesactivado: false,
		},
	}

	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	tareas, err := service.ObtenerTareasPorUsuario(
		context.Background(),
		1,
	)

	if err != nil {
		t.Fatalf(
			"No se esperaba error, se obtuvo: %v",
			err,
		)
	}

	if len(tareas) != 2 {
		t.Fatalf(
			"Se esperaban 2 tareas, se obtuvieron %d",
			len(tareas),
		)
	}

	if tareas[0].Id != 1 {
		t.Fatalf("Se esperaba la tarea con Id 1")
	}

	if tareas[1].Id != 2 {
		t.Fatalf("Se esperaba la tarea con Id 2")
	}
}

func TestObtenerTareasPorUsuario_UsuarioInexistente_DevuelveError(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}
	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{
		Usuario: nil,
	}

	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	tareas, err := service.ObtenerTareasPorUsuario(
		context.Background(),
		99,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareas != nil {
		t.Fatal("No se esperaban tareas")
	}
}

func TestObtenerTareasPorUsuario_UsuarioDesactivado_DevuelveError(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{}
	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{
		Usuario: &entities.Usuario{
			Id:              1,
			Nombre:          "Usuario desactivado",
			RolId:           2,
			EstaDesactivado: true,
		},
	}

	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	tareas, err := service.ObtenerTareasPorUsuario(
		context.Background(),
		1,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareas != nil {
		t.Fatal("No se esperaban tareas")
	}
}

func TestObtenerTareasPorUsuario_SinTareas_DevuelveListaVacia(t *testing.T) {

	tareaRepository := &mocks.TareaRepositoryMock{
		Tareas: []entities.Tarea{},
	}

	subTareaRepository := &mocks.SubTareaRepositoryMock{}

	usuarioRepository := &mocks.UsuarioRepositoryMock{
		Usuario: &entities.Usuario{
			Id:              1,
			Nombre:          "Contribuidor",
			RolId:           2,
			EstaDesactivado: false,
		},
	}

	tareaUsuarioRepository := &mocks.TareaUsuarioRepositoryMock{}

	service := NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	tareas, err := service.ObtenerTareasPorUsuario(
		context.Background(),
		1,
	)

	if err != nil {
		t.Fatalf(
			"No se esperaba error, se obtuvo: %v",
			err,
		)
	}

	if tareas == nil {
		t.Fatal("Se esperaba una lista vacía, no nil")
	}

	if len(tareas) != 0 {
		t.Fatalf(
			"Se esperaban 0 tareas, se obtuvieron %d",
			len(tareas),
		)
	}
}

func TestObtenerTareasPorUsuario_IdInvalido_DevuelveError(t *testing.T) {

	service := NewTareaService(
		&mocks.TareaRepositoryMock{},
		&mocks.SubTareaRepositoryMock{},
		&mocks.UsuarioRepositoryMock{},
		&mocks.TareaUsuarioRepositoryMock{},
	)

	tareas, err := service.ObtenerTareasPorUsuario(
		context.Background(),
		0,
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if tareas != nil {
		t.Fatal("No se esperaban tareas")
	}
}

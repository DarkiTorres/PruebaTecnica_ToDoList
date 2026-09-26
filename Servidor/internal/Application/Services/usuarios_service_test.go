package services

import (
	"context"
	"errors"
	"testing"
	mocks "to-do-server/internal/Application/Tests/Mocks"
	entities "to-do-server/internal/Domain/Entities"
)

func TestUsuarioService_ObtenerUsuarioPorId_UsuarioValido_RetornaUsuario(t *testing.T) {

	usuarioRepository := &mocks.UsuarioRepositoryMock{
		Usuario: &entities.Usuario{
			Id:              1,
			Nombre:          "Edi",
			RolId:           1,
			EstaDesactivado: false,
		},
	}

	service := NewUsuarioService(
		usuarioRepository,
	)

	usuario, err := service.ObtenerUsuarioPorId(
		context.Background(),
		1,
	)

	if err != nil {
		t.Fatalf(
			"No se esperaba un error, se obtuvo: %v",
			err,
		)
	}

	if usuario == nil {
		t.Fatal(
			"Se esperaba un usuario",
		)
	}

	if usuario.Id != 1 {
		t.Fatalf(
			"Se esperaba Id 1, se obtuvo %d",
			usuario.Id,
		)
	}

	if usuario.Nombre != "Edi" {
		t.Fatalf(
			"Se esperaba Nombre 'Edi', se obtuvo %s",
			usuario.Nombre,
		)
	}
}

func TestUsuarioService_ObtenerUsuarioPorId_IdInvalido_RetornaError(t *testing.T) {

	usuarioRepository := &mocks.UsuarioRepositoryMock{}

	service := NewUsuarioService(
		usuarioRepository,
	)

	usuario, err := service.ObtenerUsuarioPorId(
		context.Background(),
		0,
	)

	if err == nil {
		t.Fatal(
			"Se esperaba un error",
		)
	}

	if usuario != nil {
		t.Fatal(
			"No se esperaba recibir un usuario",
		)
	}

	if usuarioRepository.ObtenerPorIdLlamado {
		t.Fatal(
			"El repositorio no debería haberse llamado",
		)
	}
}

func TestUsuarioService_ObtenerUsuarios_RetornaUsuarios(t *testing.T) {

	usuarios := []entities.Usuario{
		{
			Id:              1,
			Nombre:          "Edi",
			RolId:           1,
			EstaDesactivado: false,
		},
		{
			Id:              2,
			Nombre:          "Juan",
			RolId:           2,
			EstaDesactivado: false,
		},
	}

	usuarioRepository := &mocks.UsuarioRepositoryMock{
		Usuarios: usuarios,
	}

	service := NewUsuarioService(
		usuarioRepository,
	)

	resultado, err := service.ObtenerUsuarios(
		context.Background(),
	)

	if err != nil {
		t.Fatalf(
			"No se esperaba un error, se obtuvo: %v",
			err,
		)
	}

	if !usuarioRepository.ObtenerTodosLlamado {
		t.Fatal(
			"Se esperaba que ObtenerTodos fuera llamado",
		)
	}

	if len(resultado) != 2 {
		t.Fatalf(
			"Se esperaban 2 usuarios, se obtuvieron %d",
			len(resultado),
		)
	}

	if resultado[0].Nombre != "Edi" {
		t.Fatalf(
			"Se esperaba Edi, se obtuvo %s",
			resultado[0].Nombre,
		)
	}
}

func TestUsuarioService_ObtenerUsuarios_ErrorRepositorio_RetornaError(t *testing.T) {

	esperado := errors.New("error al obtener usuarios")

	usuarioRepository := &mocks.UsuarioRepositoryMock{
		ErrorObtenerTodos: esperado,
	}

	service := NewUsuarioService(
		usuarioRepository,
	)

	usuarios, err := service.ObtenerUsuarios(
		context.Background(),
	)

	if err == nil {
		t.Fatal("Se esperaba un error")
	}

	if !errors.Is(err, esperado) {
		t.Fatalf(
			"Se esperaba el error %v, se obtuvo %v",
			esperado,
			err,
		)
	}

	if usuarios != nil {
		t.Fatal("No se esperaban usuarios")
	}

	if !usuarioRepository.ObtenerTodosLlamado {
		t.Fatal(
			"Se esperaba que ObtenerTodos fuera llamado",
		)
	}
}

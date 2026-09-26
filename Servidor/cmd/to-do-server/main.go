package main

import (
	"context"
	"log"

	handlers "to-do-server/internal/Api/Handlers"
	routes "to-do-server/internal/Api/Routes"
	services "to-do-server/internal/Application/Services"
	initialization "to-do-server/internal/Infrastructure/Database/Initialization"
	postgres "to-do-server/internal/Infrastructure/Database/postgres"
	"to-do-server/internal/Infrastructure/config"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.Connect(
		cfg.GenerateConnectionString(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx := context.Background()

	initializer := initialization.NewDatabaseInitializer(db)

	if err := initializer.Initialize(ctx); err != nil {
		log.Fatal(err)
	}

	// =========================
	// REPOSITORIES
	// =========================

	tareaRepository := postgres.NewTareaRepository(db)

	subTareaRepository := postgres.NewSubTareaRepository(db)

	usuarioRepository := postgres.NewUsuarioRepository(db)

	tareaUsuarioRepository := postgres.NewTareaUsuarioRepository(db)

	prioridadRepository := postgres.NewPrioridadRepository(db)

	// =========================
	// SERVICES
	// =========================

	usuarioService := services.NewUsuarioService(
		usuarioRepository,
	)

	prioridadService := services.NewPrioridadService(
		prioridadRepository,
	)

	subTareaService := services.NewSubTareaService(
		subTareaRepository,
		tareaRepository,
	)

	tareaService := services.NewTareaService(
		tareaRepository,
		subTareaRepository,
		usuarioRepository,
		tareaUsuarioRepository,
	)

	// =========================
	// HANDLERS
	// =========================

	usuarioHandler := handlers.NewUsuarioHandler(
		usuarioService,
	)

	prioridadHandler := handlers.NewPrioridadHandler(
		prioridadService,
	)

	subTareaHandler := handlers.NewSubTareaHandler(
		subTareaService,
	)

	tareaHandler := handlers.NewTareaHandler(
		tareaService,
		subTareaService,
		subTareaRepository,
	)

	// =========================
	// ROUTER
	// =========================

	router := gin.Default()

	routes.ConfigurarRutas(
		router,
		tareaHandler,
		subTareaHandler,
		usuarioHandler,
		prioridadHandler,
		db,
	)

	// =========================
	// SERVER
	// =========================

	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}

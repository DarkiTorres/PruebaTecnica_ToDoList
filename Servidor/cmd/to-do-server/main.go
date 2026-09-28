package main

import (
	"context"
	"log"
	"time"

	handlers "to-do-server/internal/Api/Handlers"
	routes "to-do-server/internal/Api/Routes"
	services "to-do-server/internal/Application/Services"
	database "to-do-server/internal/Infrastructure/Database"
	initialization "to-do-server/internal/Infrastructure/Database/initialization"
	postgres "to-do-server/internal/Infrastructure/Database/postgres"
	"to-do-server/internal/Infrastructure/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	log.Println("[STARTUP] Iniciando servidor...")

	cfg, err := config.Load()

	if err != nil {
		log.Fatalf("[STARTUP] Error cargando configuración: %v", err)
	}

	db, err := postgres.Connect(
		cfg.GenerateConnectionString(),
	)
	if err != nil {
		log.Fatalf("[STARTUP] Error creando conexión con PostgreSQL: %v", err)
	}

	defer db.Close()
	log.Println("[STARTUP] Pool PostgreSQL creado.")

	ctx := context.Background()

	initializer := initialization.NewDatabaseInitializer(db)

	if err := initializer.Initialize(ctx); err != nil {
		log.Fatalf(
			"[STARTUP] Error inicializando la base de datos: %v",
			err,
		)
	}

	log.Println("[STARTUP] Inicialización de base de datos completada.")

	validationCtx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	if err := database.ValidarBaseDeDatos(
		validationCtx,
		db,
	); err != nil {
		log.Fatalf(
			"[STARTUP] Validación de base de datos fallida: %v",
			err,
		)
	}

	log.Println("[STARTUP] Base de datos validada correctamente.")

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

	tareaUsuarioService := services.NewTareaUsuarioService(tareaUsuarioRepository)

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
		tareaUsuarioService,
		usuarioService,
	)

	// =========================
	// ROUTER
	// =========================

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

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

	log.Printf(
		"[STARTUP] Servidor escuchando en :%s",
		cfg.AppPort,
	)

	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf(
			"[STARTUP] Error iniciando servidor HTTP: %v",
			err,
		)
	}
}

package main

import (
	"log"
	"to-do-server/internal/Infrastructure/Database/postgres"
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

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		if err := db.Ping(c.Request.Context()); err != nil {
			c.JSON(503, gin.H{
				"status":   "error",
				"database": "disconnected",
			})
		}
		c.JSON(200, gin.H{
			"status":   "ok",
			"database": "connected",
		})
	})

	router.Run(":" + cfg.AppPort)
}

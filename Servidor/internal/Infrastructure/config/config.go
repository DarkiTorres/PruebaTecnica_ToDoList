package config

import (
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string

	DatabaseHost     string
	DatabasePort     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		AppPort:          os.Getenv("APP_PORT"),
		DatabaseHost:     os.Getenv("DATABASE_HOST"),
		DatabasePort:     os.Getenv("DATABASE_PORT"),
		DatabaseUser:     os.Getenv("DATABASE_USER"),
		DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
		DatabaseName:     os.Getenv("DATABASE_NAME"),
	}, nil
}

func (c *Config) GenerateConnectionString() string {
	connectionUrl := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(c.DatabaseUser, c.DatabasePassword),
		Host:   c.DatabaseHost + ":" + c.DatabasePort,
		Path:   "/" + c.DatabaseName,
	}

	return connectionUrl.String()
}

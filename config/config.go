package config

import (
	"os"
)

type Config struct {
	DB struct {
		Host     string
		Username string
		Password string
		DBName   string
	}

	Routing struct {
		Port string
	}
	SESSION_STORE struct {
		SecretHashKey string
	}
}

func NewConfig() *Config {
	c := new(Config)

	c.DB.Host = os.Getenv("DB_HOST")
	c.DB.Username = os.Getenv("DB_USER")
	c.DB.Password = os.Getenv("DB_PASSWORD")
	c.DB.DBName = os.Getenv("DB_NAME")

	c.SESSION_STORE.SecretHashKey = os.Getenv("SECRET_HASH_KEY")

	c.Routing.Port = "3000"

	return c
}

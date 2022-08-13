package config

import (
	"os"
)

type Config struct {
	DB struct {
		Production struct {
			Host     string
			Username string
			Password string
			DBName   string
		}
	}
	Routing struct {
		Port string
	}
}

func NewConfig() *Config {
	c := new(Config)

	c.DB.Production.Host = os.Getenv("DB_HOST")
	c.DB.Production.Username = os.Getenv("DB_USER")
	c.DB.Production.Password = os.Getenv("DB_PASSWORD")
	c.DB.Production.DBName = os.Getenv("DB_NAME")

	c.Routing.Port = "3000"

	return c
}

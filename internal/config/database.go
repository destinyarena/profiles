package config

import (
	"github.com/destinyarena/profiles/internal/structs"
	"os"
)

func LoadDatabaseConfig() *structs.Database {
	return &structs.Database{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("SECRET_DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		DBName:   os.Getenv("DB_NAME"),
	}
}

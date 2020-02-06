package config

import (
    "github.com/arturoguerra/destinyarena-accounts/internal/structs"
    "os"
)

func LoadConfig() *structs.Config {
    return &structs.Config{
        Host: os.Getenv("HOST"),
        Port: os.Getenv("PORT"),
    }
}

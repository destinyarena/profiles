package config

import (
    "github.com/arturoguerra/destinyarena-accounts/internal/structs"
    "os"
)

func LoadConfig() *structs.Config {
    return &structs.Config{
        GRPCHost: os.Getenv("GRPC_HOST"),
        GRPCPort: os.Getenv("GRPC_PORT"),
    }
}

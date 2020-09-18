package config

import (
	"github.com/destinyarena/profiles/internal/structs"
	"os"
)

func LoadConfig() *structs.Config {
	return &structs.Config{
		GRPCHost: os.Getenv("GRPC_HOST"),
		GRPCPort: os.Getenv("GRPC_PORT"),
	}
}

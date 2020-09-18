package main

/*
This package handles communication between our other microservices and our accounts database
*/

import (
	"fmt"
	"net"

	"github.com/destinyarena/profiles/internal/config"
	"github.com/destinyarena/profiles/internal/logging"
	database "github.com/destinyarena/profiles/pkg/database"
	pb "github.com/destinyarena/profiles/proto"
	"google.golang.org/grpc"
)

var log = logging.New()

type profilesServer struct {
	DB database.Client
	pb.UnimplementedProfilesServer
}

func main() {
	cfg := config.LoadConfig()
	address := fmt.Sprintf("%s:%s", cfg.GRPCHost, cfg.GRPCPort)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Infof("Listening on: %s", address)

	s := grpc.NewServer()
	dbcfg := config.LoadDatabaseConfig()
	db, err := database.New(dbcfg.Username, dbcfg.Password, dbcfg.Host, dbcfg.DBName)
	if err != nil {
		log.Fatal(err)
	}

	as := &profilesServer{
		DB: db,
	}

	pb.RegisterProfilesServer(s, as)
	if err := s.Serve(lis); err != nil {
		db.Close()
		log.Fatalf("Failed to serve: %v", err)
	}

	db.Close()
}

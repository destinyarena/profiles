package main

/*
This package handles communication between our other microservices and our accounts database
*/


import (
    "fmt"
    "net"
    "google.golang.org/grpc"
    pb "github.com/arturoguerra/destinyarena-accounts/proto"
    database "github.com/arturoguerra/destinyarena-accounts/pkg/database"
    "github.com/arturoguerra/destinyarena-accounts/internal/logging"
    "github.com/arturoguerra/destinyarena-accounts/internal/config"

)

var log = logging.New()

type ProfilesServer struct {
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

    as := &ProfilesServer{
        DB: db,
    }

    pb.RegisterProfilesServer(s, as)
    if err := s.Serve(lis); err != nil {
        db.Close()
        log.Fatalf("Failed to serve: %v", err)
    }

    db.Close()
}

package main

import (
    "context"
    pb "github.com/arturoguerra/destinyarena-accounts/proto"
)

func (p *ProfilesServer) GetProfile(ctx context.Context, in *pb.IdRequest) (*pb.ProfileReply, error) {
    err, u := p.DB.GetUser(in.GetId())
    if err != nil {
        log.Error(err)
        return nil, err
    }

    return &pb.ProfileReply{
        Id: u.UUID.String(),
        Discord: u.Discord,
        Bungie: u.Bungie,
        Faceit: u.Faceit,
        Banned: u.Banned,
    }, nil
}

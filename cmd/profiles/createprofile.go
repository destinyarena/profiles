package main

import (
    "context"
    pb "github.com/arturoguerra/destinyarena-accounts/proto"
)

func (p *ProfilesServer) CreateProfile(ctx context.Context, in *pb.ProfileRequest) (*pb.ProfileReply, error) {
    err, u := p.DB.RegisterUser(in.GetDiscord(), in.GetBungie(), in.GetFaceit())
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
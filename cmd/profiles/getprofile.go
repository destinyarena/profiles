package main

import (
	"context"

	pb "github.com/destinyarena/profiles/proto"
)

// GetProfile returns a single profile
func (p *profilesServer) GetProfile(ctx context.Context, in *pb.IdRequest) (*pb.ProfileReply, error) {
	log.Infof("Fetching Profile for: %s", in.GetId())
	u, err := p.DB.GetUser(in.GetId())
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &pb.ProfileReply{
		Id:        u.UUID.String(),
		Discord:   u.Discord,
		Bungie:    u.Bungie,
		Faceit:    u.Faceit,
		Banned:    u.Banned,
		Banreason: u.BanReason,
	}, nil
}

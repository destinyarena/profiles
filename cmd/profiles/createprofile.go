package main

import (
	"context"

	pb "github.com/destinyarena/profiles/proto"
)

// CreateProfile creates new database profile if everything checks out
func (p *profilesServer) CreateProfile(ctx context.Context, in *pb.ProfileRequest) (*pb.ProfileReply, error) {
	log.Infof("Registering user Discord: %s Bungie: %s Faceit: %s IPHash: %s", in.GetDiscord(), in.GetBungie(), in.GetFaceit(), in.GetIphash())
	u, err := p.DB.RegisterUser(in.GetDiscord(), in.GetBungie(), in.GetFaceit(), in.GetIphash())
	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.Infof("Registering Discord: %s Bungie: %s Faceit: %s IPHash: %s", u.Discord, u.Bungie, u.Faceit, u.IPHash)

	return &pb.ProfileReply{
		Id:        u.UUID.String(),
		Discord:   u.Discord,
		Bungie:    u.Bungie,
		Faceit:    u.Faceit,
		Banned:    u.Banned,
		Iphash:    u.IPHash,
		Banreason: u.BanReason,
	}, nil
}

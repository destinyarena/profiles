package main

import (
	"context"

	pb "github.com/destinyarena/profiles/proto"
)

// Ban marks a user a banned in the database
func (p *profilesServer) Ban(ctx context.Context, in *pb.IdRequest) (*pb.Empty, error) {
	if err := p.DB.Ban(in.GetId()); err != nil {
		log.Error(err)
		return nil, err
	}

	return &pb.Empty{}, nil
}

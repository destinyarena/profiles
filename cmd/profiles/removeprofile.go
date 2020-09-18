package main

import (
	"context"

	pb "github.com/destinyarena/profiles/proto"
)

// RemoveProfile deletes a profile from the database
func (p *profilesServer) RemoveProfile(ctx context.Context, in *pb.IdRequest) (*pb.Empty, error) {
	if err := p.DB.DeleteUser(in.GetId()); err != nil {
		log.Error(err)
		return nil, err
	}

	return &pb.Empty{}, nil
}

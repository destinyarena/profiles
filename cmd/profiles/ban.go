package main

import (
    "context"
    pb "github.com/arturoguerra/destinyarena-accounts/proto"
)

func (p *ProfilesServer) Ban(ctx context.Context, in *pb.IdRequest) (*pb.Empty, error) {
    if err := p.DB.Ban(in.GetId()); err != nil {
        log.Error(err)
        return nil, err
    }

    return &pb.Empty{}, nil
}

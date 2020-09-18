package main

import (
	pb "github.com/destinyarena/profiles/proto"
)

// GetProfilesByIP returns a list of users that have this ip
func (f *profilesServer) GetProfilesByIP(in *pb.IPRequest, stream pb.Profiles_GetProfilesByIPServer) error {
	log.Infof("Getting profiles by IP Hash: %s", in.GetIphash())
	profiles, err := f.DB.GetUsersByIP(in.GetIphash())
	if err != nil {
		return err
	}

	for _, profile := range profiles {
		p := &pb.ProfileReply{
			Id:      profile.UUID.String(),
			Discord: profile.Discord,
			Bungie:  profile.Bungie,
			Banned:  profile.Banned,
			Iphash:  profile.IPHash,
		}

		if err := stream.Send(p); err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}

package main

import (
	pb "github.com/destinyarena/profiles/proto"
)

// GetAllProfiles returns a list of every profile
func (f *profilesServer) GetAllProfiles(in *pb.Empty, stream pb.Profiles_GetAllProfilesServer) error {
	log.Infof("Here we go again getting every profile...")
	profiles, err := f.DB.GetAllUsers()
	if err != nil {
		log.Error(err)
		return err
	}

	for _, profile := range profiles {
		p := &pb.ProfileReply{
			Id:        profile.UUID.String(),
			Discord:   profile.Discord,
			Bungie:    profile.Bungie,
			Faceit:    profile.Faceit,
			Banned:    profile.Banned,
			Banreason: profile.BanReason,
		}

		if err := stream.Send(p); err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}

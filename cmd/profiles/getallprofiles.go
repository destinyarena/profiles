package main

import (
    pb "github.com/arturoguerra/destinyarena-accounts/proto"
)

func (f *ProfilesServer) GetAllProfiles(in *pb.Empty, stream pb.Profiles_GetAllProfilesServer) error {
    log.Infof("Here we go again getting every profile...")
    err, profiles := f.DB.GetAllUsers()
    if err != nil {
        log.Error(err)
        return err
    }

    for _, profile := range profiles {
        p := &pb.ProfileReply{
            Id:      profile.UUID.String(),
            Discord: profile.Discord,
            Bungie:  profile.Bungie,
            Faceit:  profile.Faceit,
            Banned:  profile.Banned,
        }

        if err := stream.Send(p); err != nil {
            log.Error(err)
            return err
        }
    }

    return nil
}

package database

/*
This Check if a user is already in the database and adds them if they are not
*/

import (
    "errors"
)

func (c *client) DeleteUser(uid string) (error) {
    if c.Where("discord = ? OR faceit = ? OR bungie = ?", uid, uid, uid).Unscoped().Delete(&User{}).RecordNotFound() {
        return errors.New("User not found")
    }


    return nil
}

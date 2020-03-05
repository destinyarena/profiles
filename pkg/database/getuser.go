package database

/*
This Check if a user is already in the database and adds them if they are not
*/

import (
    "errors"
)

func (c *client) GetUser(uid string) (error, *User) {
    u := User{}
    if c.Where("discord = ? OR faceit = ? OR bungie = ?", uid, uid, uid).First(&u).RecordNotFound() {
        return errors.New("User not found"), nil
    }


    return nil, &u
}

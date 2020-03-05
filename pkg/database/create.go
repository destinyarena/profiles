package database

/*
This Check if a user is already in the database and adds them if they are not
*/

import (
    "errors"
)

func (c *client) RegisterUser(discord, bungie, faceit string) (error, *User) {
    u := User{}
    if c.Where("discord = ? OR faceit = ? OR bungie = ?", discord, faceit, bungie).First(&u).RecordNotFound() {
        newUser := &User{
            Discord: discord,
            Bungie:  bungie,
            Faceit:  faceit,
        }

        c.Create(newUser)

        return nil, newUser
    }

    err := errors.New("Member already exists")

    return err, &u
}

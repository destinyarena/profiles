package database

/*
This Check if a user is already in the database and adds them if they are not
*/

import (
	"errors"
)

// GetUser fetches a user from the database
func (c *client) GetUser(uid string) (*User, error) {
	u := User{}
	if c.Where("discord = ? OR faceit = ? OR bungie = ?", uid, uid, uid).First(&u).RecordNotFound() {
		return nil, errors.New("User not found")
	}

	return &u, nil
}

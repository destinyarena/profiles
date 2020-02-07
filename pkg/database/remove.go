package database

/*
This Check if a user is already in the database and adds them if they are not
*/

import (
    "errors"
)

func (f *DBClient) DeleteUser(uid string) (error) {
    db, err := f.Connect()
    if err != nil {
        return err
    }

    defer db.Close()

    if db.Where("discord = ? OR faceit = ? OR bungie = ?", uid, uid, uid).Unscoped().Delete(&User{}).RecordNotFound() {
        return errors.New("User not found")
    }


    return nil
}

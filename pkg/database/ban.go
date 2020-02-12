package database

/*
Marks user as banned
*/

func (f *DBClient) Ban(uid string) error {
    err, user := f.GetUser(uid)
    if err != nil {
        return err
    }

    db, err := f.Connect()
    if err != nil {
        return err
    }
    defer db.Close()

    user.Banned = true
    db.Save(user)
    return nil
}

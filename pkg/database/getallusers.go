package database

/*
This gets all database entries
*/

func (f *DBClient) GetAllUsers() (error, []User) {
    users := make([]User, 0)
    db, err := f.Connect()
    if err != nil {
        return err, users
    }

    defer db.Close()

    if err = db.Find(&users).Error; err != nil {
        return err, users
    }

    return nil, users
}

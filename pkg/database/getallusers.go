package database

/*
This gets all database entries
*/

func (c *client) GetAllUsers() (error, []User) {
    users := make([]User, 0)

    if err := c.Find(&users).Error; err != nil {
        return err, users
    }

    return nil, users
}

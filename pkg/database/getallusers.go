package database

/*
This gets all database entries
*/

func (c *client) GetAllUsers() ([]*User, error) {
	users := make([]*User, 0)

	if err := c.Find(users).Error; err != nil {
		return users, err
	}

	return users, nil
}

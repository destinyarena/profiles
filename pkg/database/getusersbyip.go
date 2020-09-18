package database

/*
This gets all users that have registered with the same IP
*/

// GetUserByIP returns ([]*User, error)
func (c *client) GetUsersByIP(iphash string) ([]*User, error) {
	users := make([]*User, 0)

	if err := c.Where("ip_hash = ?", iphash).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

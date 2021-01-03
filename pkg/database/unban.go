package database

func (c *client) Unban(uid string) error {
	user, err := c.GetUser(uid)
	if err != nil {
		return err
	}

	user.Banned = false
	user.BanReason = ""
	c.Save(user)
	return nil
}

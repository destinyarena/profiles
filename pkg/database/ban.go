package database

/*
Marks user as banned
*/

func (c *client) Ban(uid, reason string) error {
	user, err := c.GetUser(uid)
	if err != nil {
		return err
	}

	user.Banned = true
	user.BanReason = reason
	c.Save(user)
	return nil
}

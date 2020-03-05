package database

/*
Marks user as banned
*/

func (c client) Ban(uid string) error {
    err, user := c.GetUser(uid)
    if err != nil {
        return err
    }


    user.Banned = true
    c.Save(user)
    return nil
}

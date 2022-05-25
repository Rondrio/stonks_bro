package database

type User struct {
	Id   int
	Name string
}

func (db Database) GetUser(discord_id string) (*User, error) {
	rows, err := db.conn.Query("SELECT * FROM users WHERE id = ?", discord_id)
	if err != nil {
		return nil, err
	}

	user := User{}

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (db Database) Subscribe(discord_id string) error {
	_, err := db.conn.Query("UPDATE users SET subscribed = 1 WHERE id = ?", discord_id)
	return err
}

func (db Database) Unsubscribe(discord_id string) error {
	_, err := db.conn.Query("UPDATE users SET subscribed = 0 WHERE id = ?", discord_id)
	return err
}

func (db Database) CheckSubscriptionStatus(discord_id string) (bool, error) {
	subscribed := false
	rows, err := db.conn.Query("SELECT subscribed FROM users WHERE id = ?", discord_id)
	if err != nil {
		return false, err
	}
	for rows.Next() {
		rows.Scan(&subscribed)
	}

	return subscribed, nil
}

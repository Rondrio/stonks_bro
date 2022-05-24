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

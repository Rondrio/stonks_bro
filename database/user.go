package database

type User struct {
	Id           int
	Name         string
	Stonks_Count int
}

func (db Database) GetUser(discord_id string) (*User, error) {
	rows, err := db.conn.Query("SELECT * FROM users WHERE id = ?", discord_id)
	if err != nil {
		return nil, err
	}

	user := User{}

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Stonks_Count)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (db Database) GetTotalStonks() (int, error) {
	var count = -1
	rows, err := db.conn.Query("SELECT SUM(stonks_count) FROM users WHERE 1")
	if err != nil {
		return count, err
	}

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return count, err
		}
	}
	return count, nil
}

func (db Database) UpdateUser(id string, name string) error {
	user, err := db.GetUser(id)
	if err != nil {
		return err
	}

	_, err = db.conn.Exec("INSERT INTO users (id,name,stonks_count) VALUES (?,?,?) ON DUPLICATE KEY UPDATE name = ?, stonks_count = ? ", id, name, user.Stonks_Count+1, name, user.Stonks_Count+1)
	return err
}

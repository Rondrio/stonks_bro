package database

type Stonks struct {
	ID            int
	User          User
	Channel       Channel
	MessageAuthor User
	Type          string
	TimeStamp     int
}

const (
	STONK_TYPE_MESSAGE  = "Message"
	STONK_TYPE_REACTION = "Reaction"
)

func (db Database) GetTotalStonkCount() (int, error) {
	count := -1
	result, err := db.conn.Query("SELECT COUNT(*) FROM stonks WHERE 1")
	if err != nil {
		return count, err
	}
	defer result.Close()

	for result.Next() {
		err = result.Scan(&count)
		if err != nil {
			return count, err
		}
	}

	return count, nil
}

func (db Database) GetTotalLastMonthStonkCount() (int, error) {
	count := -1
	result, err := db.conn.Query("SELECT COUNT(*) FROM stonks WHERE timestamp BETWEEN CURDATE() - INTERVAL 30 DAY AND CURDATE()")
	if err != nil {
		return count, err
	}
	for result.Next() {
		err = result.Scan(&count)
		if err != nil {
			return count, err
		}
	}

	return count, nil
}

func (db Database) GetTotalStonkCountByUser(id string) (int, error) {
	count := -1
	result, err := db.conn.Query("SELECT COUNT(*) FROM stonks WHERE user_id = ?", id)
	if err != nil {
		return count, err
	}
	defer result.Close()

	for result.Next() {
		err = result.Scan(&count)
		if err != nil {
			return count, err
		}
	}

	return count, nil
}

func (db Database) GetStonkCountByUserLastMonth(id string) (int, error) {
	count := -1
	result, err := db.conn.Query("SELECT COUNT(*) FROM stonks WHERE user_id = ? AND timestamp BETWEEN CURDATE() - INTERVAL 30 DAY AND CURDATE()", id)
	if err != nil {
		return count, err
	}
	for result.Next() {
		err = result.Scan(&count)
		if err != nil {
			return count, err
		}
	}

	return count, nil
}

func (db Database) AddStonks(user_id string, author_id string, channel_id string, stonk_type string) error {

	_, err := db.conn.Exec("INSERT INTO channels (id) VALUES (?) ON DUPLICATE KEY UPDATE name = name", channel_id)
	if err != nil {
		return err
	}

	_, err = db.conn.Exec("INSERT INTO users (id) VALUES (?) ON DUPLICATE KEY UPDATE name = name ", user_id)
	if err != nil {
		return err
	}

	subscribed := 0

	rows, err := db.conn.Query("SELECT subscribed FROM users WHERE id = ?", user_id)
	if err != nil {
		return err
	}
	for rows.Next() {
		rows.Scan(&subscribed)
	}

	if subscribed == 0 {
		return nil
	}

	rows, err = db.conn.Query("SELECT subscribed FROM users WHERE id = ?", author_id)
	for rows.Next() {
		rows.Scan(&subscribed)
	}

	if subscribed == 0 {
		return nil
	}

	_, err = db.conn.Exec("INSERT INTO stonks (user_id, channel_id, message_author_id, type, timestamp) VALUES (?,?,?,?,NOW())", user_id, channel_id, author_id, stonk_type)
	return err
}

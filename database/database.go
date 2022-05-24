package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type IDatabase interface {
	GetUser(string) (*User, error)
	GetTotalStonkCount() (int, error)
	GetTotalLastMonthStonkCount() (int, error)
	GetTotalStonkCountByUser(string) (int, error)
	GetStonkCountByUserLastMonth(string) (int, error)
	AddStonks(user_id string, author_id string, channel_id string, stonk_type string) error
}

type Database struct {
	conn *sql.DB
}

func GetDB() (*Database, error) {
	conn, err := sql.Open("mysql", "stonks:stonks@(alexgz.de)/stonks_counter")
	return &Database{
		conn: conn,
	}, err
}

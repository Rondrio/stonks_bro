package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type IDatabase interface {
	GetUser(string) (*User, error)
	GetTotalStonks() (int, error)
	UpdateUser(string, string) error
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

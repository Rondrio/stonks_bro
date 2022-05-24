package handlers

import (
	"stonks_bot/database"
)

type Handlers struct {
	db database.IDatabase
}

func SetupHandlers(database database.IDatabase) *Handlers {
	return &Handlers{
		db: database,
	}
}

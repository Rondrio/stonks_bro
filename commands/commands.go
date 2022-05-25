package commands

import (
	"stonks_bot/database"

	"github.com/bwmarrin/discordgo"
)

type Command interface {
	Execute(db database.IDatabase, s *discordgo.Session, m *discordgo.MessageCreate) error
	GetKeyword() string
}

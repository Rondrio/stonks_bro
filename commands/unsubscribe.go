package commands

import (
	"fmt"
	"stonks_bot/database"

	"github.com/bwmarrin/discordgo"
)

type Unsubscribe struct {
}

func (uc Unsubscribe) Execute(db database.IDatabase, s *discordgo.Session, m *discordgo.MessageCreate) error {
	err := db.Unsubscribe(m.Author.ID)
	if err != nil {
		return err
	}
	_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("You are no longer subscribed to turtle facts!<:not_stonks:813079160464605264>"))
	return err
}

func (uc Unsubscribe) GetKeyword() string {
	return "!unsubscribe"
}

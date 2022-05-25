package commands

import (
	"fmt"
	"stonks_bot/database"

	"github.com/bwmarrin/discordgo"
)

type Subscribe struct {
}

func (uc Subscribe) Execute(db database.IDatabase, s *discordgo.Session, m *discordgo.MessageCreate) error {
	err := db.Subscribe(m.Author.ID)
	if err != nil {
		return err
	}
	_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("You are now subscribed to turtle facts!<:stonks:813079068928507934>"))
	return err
}

func (uc Subscribe) GetKeyword() string {
	return "!subscribe"
}

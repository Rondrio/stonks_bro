package commands

import (
	"fmt"
	"stonks_bot/database"

	"github.com/bwmarrin/discordgo"
)

type Total struct {
}

func (tc Total) Execute(db database.IDatabase, s *discordgo.Session, m *discordgo.MessageCreate) error {
	count, err := db.GetTotalStonkCount()
	if err != nil {
		return err
	}

	_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Insgesamt wurde %d Mal gestonkt", count))
	return err
}

func (tc Total) GetKeyword() string {
	return "!total"
}

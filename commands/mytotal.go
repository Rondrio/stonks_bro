package commands

import (
	"fmt"
	"stonks_bot/database"

	"github.com/bwmarrin/discordgo"
)

type MyTotal struct {
}

func (mtc MyTotal) Execute(db database.IDatabase, s *discordgo.Session, m *discordgo.MessageCreate) error {
	count, err := db.GetTotalStonkCountByUser(m.Author.ID)
	if err != nil {
		return err
	}

	_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Du hast insgesamt %d Mal gestonkt", count))
	return err
}

func (mtc MyTotal) GetKeyword() string {
	return "!mytotal"
}

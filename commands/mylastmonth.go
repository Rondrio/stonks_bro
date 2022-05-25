package commands

import (
	"fmt"
	"stonks_bot/database"

	"github.com/bwmarrin/discordgo"
)

type MyLastMonth struct {
}

func (mlmc MyLastMonth) Execute(db database.IDatabase, s *discordgo.Session, m *discordgo.MessageCreate) error {
	count, err := db.GetStonkCountByUserLastMonth(m.Author.ID)
	if err != nil {
		return err
	}

	_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Du hast im letzten Monat %d Mal gestonkt", count))
	return err
}

func (mlmc MyLastMonth) GetKeyword() string {
	return "!mylastmonth"
}

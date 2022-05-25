package commands

import (
	"fmt"
	"stonks_bot/database"

	"github.com/bwmarrin/discordgo"
)

type LastMonth struct {
}

func (lmc LastMonth) Execute(db database.IDatabase, s *discordgo.Session, m *discordgo.MessageCreate) error {
	count, err := db.GetTotalLastMonthStonkCount()
	if err != nil {
		return err
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Im letzten Monat wurde %d Mal gestonkt", count))
	return nil
}

func (lmc LastMonth) GetKeyword() string {
	return "!lastmonth"
}

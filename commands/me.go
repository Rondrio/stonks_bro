package commands

import (
	"fmt"
	"math"
	"stonks_bot/database"

	"github.com/bwmarrin/discordgo"
)

type Me struct {
}

func (mc Me) Execute(db database.IDatabase, s *discordgo.Session, m *discordgo.MessageCreate) error {
	count, err := db.GetTotalStonkCountByUser(m.Author.ID)
	if err != nil {
		return err
	}

	_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Bravo! Du hast insgesamt schon %d Mal gestonkt.\nDamit hättest du die Gamestop-Aktien wieder um etwa %.2f%% ansteigen lassen können", count, (float64(count)*math.Pi)/100))

	return err
}

func (mc Me) GetKeyword() string {
	return "!me"
}

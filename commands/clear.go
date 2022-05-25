package commands

import (
	"stonks_bot/database"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Clear struct {
}

func (cc Clear) Execute(db database.IDatabase, s *discordgo.Session, m *discordgo.MessageCreate) error {

	count, err := strconv.Atoi(strings.Fields(m.Content)[1])
	if err != nil {
		count = 10
	}

	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		return err
	}
	msgs, err := s.ChannelMessages(channel.ID, count, "", "", "")
	if err != nil {
		return err
	}

	go func() {
		for _, msg := range msgs {
			if msg.Author.ID == "978338573533212694" {
				s.ChannelMessageDelete(m.ChannelID, msg.ID)
				time.Sleep(250 * time.Millisecond)
			}
		}
	}()
	return nil
}

func (cc Clear) GetKeyword() string {
	return "!clear"
}

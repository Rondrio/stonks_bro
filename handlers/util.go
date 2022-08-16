package handlers

import (
	"fmt"
	"stonks_bot/commands"
	"stonks_bot/database"

	"github.com/bwmarrin/discordgo"
)

type Handlers struct {
	db       database.IDatabase
	commands []commands.Command
}

func SetupHandlers(database database.IDatabase) *Handlers {
	return &Handlers{
		db: database,
		commands: []commands.Command{
			commands.Clear{},
			commands.Help{},
			commands.LastMonth{},
			commands.Me{},
			commands.MyLastMonth{},
			commands.MyTotal{},
			commands.Subscribe{},
			commands.Total{},
			commands.Unsubscribe{},
		},
	}
}

func WriteStonkMessage(s *discordgo.Session, channel *discordgo.Channel, m *discordgo.Message, stonkNr int) error {
	msg, err := s.ChannelMessage(channel.ID, channel.LastMessageID)
	if err != nil {
		return err
	}

	if msg.Author.ID == "978338573533212694" {
		s.ChannelMessageEdit(channel.ID, msg.ID, fmt.Sprintf("Ich habe schon %d Mal gestonkt!", stonkNr+1))
		return nil
	}

	s.ChannelMessageSendReply(m.ChannelID, "<:stonks:813079068928507934>", m.Reference())
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Ich habe schon %d Mal gestonkt!", stonkNr+1))
	return nil
}

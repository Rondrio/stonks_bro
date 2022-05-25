package handlers

import (
	"fmt"
	"log"
	"stonks_bot/database"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (h Handlers) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.Bot {
		return
	}

	if strings.HasPrefix(m.Content, "!") {
		h.handleCommands(s, m)
		return
	}

	subbed, err := h.db.CheckSubscriptionStatus(m.Author.ID)
	if err != nil {
		log.Println(err)
	}
	if !subbed {
		return
	}

	total, err := h.db.GetTotalStonkCount()
	if err != nil {
		log.Println(err)
	}

	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		log.Println(err)
	}

	if m.ChannelID == "822461092122198097" || m.ChannelID == "776441387876614165" {
		fmt.Println("adding reaction")
		s.MessageReactionAdd(m.ChannelID, m.Message.ID, "stonks:813079068928507934")
	}

	if strings.Contains(strings.ToLower(m.Content), "stonks") {
		WriteStonkMessage(s, channel, m.Message, total)
	}

	err = h.db.AddStonks(m.Author.ID, m.Author.ID, m.ChannelID, database.STONK_TYPE_MESSAGE)

	if err != nil {
		log.Println(err)
	}

}

func (h Handlers) handleCommands(s *discordgo.Session, m *discordgo.MessageCreate) {
	for _, cmd := range h.commands {
		if cmd.GetKeyword() == strings.ToLower(strings.Fields(m.Content)[0]) {
			err := cmd.Execute(h.db, s, m)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

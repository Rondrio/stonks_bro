package handlers

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func (h Handlers) ReactionAdded(s *discordgo.Session, m *discordgo.MessageReactionAdd) {

	if m.Member.User.Bot {
		return
	}

	total, err := h.db.GetTotalStonks()
	if err != nil {
		log.Println(err)
		return
	}

	if m.Emoji.ID == "813079068928507934" {
		s.ChannelMessageSend(m.ChannelID, "<:stonks:813079068928507934>")
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Ich habe schon %d Mal gestonkt!", total+1))
	}

	user, err := s.User(m.UserID)
	if err != nil {
		log.Println(err)
	}

	err = h.db.UpdateUser(user.ID, user.Username)
	if err != nil {
		log.Println(err)
	}
}

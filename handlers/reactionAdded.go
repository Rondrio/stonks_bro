package handlers

import (
	"fmt"
	"log"
	"stonks_bot/database"

	"github.com/bwmarrin/discordgo"
)

func (h Handlers) ReactionAdded(s *discordgo.Session, m *discordgo.MessageReactionAdd) {

	if m.Member.User.Bot {
		return
	}

	total, err := h.db.GetTotalStonkCount()
	if err != nil {
		log.Println(err)
		return
	}

	if m.Emoji.ID == "813079068928507934" {
		s.ChannelMessageSend(m.ChannelID, "<:stonks:813079068928507934>")
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Ich habe schon %d Mal gestonkt!", total+1))
	}

	msg, err := s.ChannelMessage(m.ChannelID, m.MessageID)
	if err != nil {
		log.Println(err)
	}

	user, err := s.User(m.UserID)
	if err != nil {
		log.Println(err)
	}

	err = h.db.AddStonks(user.ID, msg.Author.ID, m.ChannelID, database.STONK_TYPE_REACTION)
	if err != nil {
		log.Println(err)
	}

}

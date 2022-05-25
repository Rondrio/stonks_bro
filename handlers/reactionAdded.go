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

	subbed, err := h.db.CheckSubscriptionStatus(m.UserID)
	if err != nil {
		log.Println(err)
	}
	if !subbed {
		return
	}

	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		log.Println(err)
	}

	msg, err := s.ChannelMessage(m.ChannelID, m.MessageID)
	if err != nil {
		log.Println(err)
	}

	subbed, err = h.db.CheckSubscriptionStatus(msg.Author.ID)
	if err != nil {
		log.Println(err)
	}
	if !subbed {
		return
	}

	user, err := s.User(m.UserID)
	if err != nil {
		log.Println(err)
	}

	total, err := h.db.GetTotalStonkCount()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(m.Emoji.ID)

	if m.Emoji.ID == "813079068928507934" {
		WriteStonkMessage(s, channel, msg, total)
	}

	err = h.db.AddStonks(user.ID, msg.Author.ID, m.ChannelID, database.STONK_TYPE_REACTION)
	if err != nil {
		log.Println(err)
	}

}

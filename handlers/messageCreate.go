package handlers

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (h Handlers) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.Bot {
		return
	}

	if strings.HasPrefix(m.Content, "!") {
		h.handleCommand(s, m)
		return
	}

	total, err := h.db.GetTotalStonks()
	if err != nil {
		log.Println(err)
		return
	}

	if m.ChannelID == "822461092122198097" || m.ChannelID == "776441387876614165" {
		fmt.Println("adding reaction")
		s.MessageReactionAdd(m.ChannelID, m.Message.ID, "stonks:813079068928507934")
	}

	if strings.Contains(strings.ToLower(m.Content), "stonks") {
		s.ChannelMessageSendReply(m.ChannelID, "<:stonks:813079068928507934>", m.Reference())
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Ich habe schon %d Mal gestonkt!", total+1))
	}

	err = h.db.UpdateUser(m.Author.ID, m.Author.Username)
	if err != nil {
		log.Println(err)
	}

}

func (h Handlers) handleCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	switch strings.ToLower(strings.Fields(m.Content)[0]) {
	case "!me":
		user, err := h.db.GetUser(m.Author.ID)
		if err != nil {
			log.Println(err)
		}
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Bravo! Du hast insgesamt schon %d Mal gestonkt.\nDamit hättest du die Gamestop-Aktien wieder um etwa %.2f%% ansteigen lassen können", user.Stonks_Count, (float64(user.Stonks_Count)*math.Pi)/100))
	}
}

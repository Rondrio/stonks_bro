package handlers

import (
	"fmt"
	"log"
	"math"
	"stonks_bot/database"
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

	total, err := h.db.GetTotalStonkCount()
	if err != nil {
		log.Println(err)
	}

	if m.ChannelID == "822461092122198097" || m.ChannelID == "776441387876614165" {
		fmt.Println("adding reaction")
		s.MessageReactionAdd(m.ChannelID, m.Message.ID, "stonks:813079068928507934")
	}

	if strings.Contains(strings.ToLower(m.Content), "stonks") {
		s.ChannelMessageSendReply(m.ChannelID, "<:stonks:813079068928507934>", m.Reference())
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Ich habe schon %d Mal gestonkt!", total+1))
	}

	err = h.db.AddStonks(m.Author.ID, m.Author.ID, m.ChannelID, database.STONK_TYPE_MESSAGE)

	if err != nil {
		log.Println(err)
	}

}

func (h Handlers) handleCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	switch strings.ToLower(strings.Fields(m.Content)[0]) {
	case "!me":

		count, err := h.db.GetTotalStonkCountByUser(m.Author.ID)
		if err != nil {
			log.Println(err)
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Bravo! Du hast insgesamt schon %d Mal gestonkt.\nDamit hättest du die Gamestop-Aktien wieder um etwa %.2f%% ansteigen lassen können", count, (float64(count)*math.Pi)/100))

	case "!lastmonth":
		count, err := h.db.GetTotalLastMonthStonkCount()
		if err != nil {
			log.Println(err)
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Im letzten Monat wurde %d Mal gestonkt", count))
	case "!mylastmonth":
		count, err := h.db.GetStonkCountByUserLastMonth(m.Author.ID)
		if err != nil {
			log.Println(err)
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Du hast im letzten Monat %d Mal gestonkt", count))
	case "!mytotal":
		count, err := h.db.GetTotalStonkCountByUser(m.Author.ID)
		if err != nil {
			log.Println(err)
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Du hast insgesamt %d Mal gestonkt", count))

	}

}

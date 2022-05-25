package commands

import (
	"stonks_bot/database"

	"github.com/bwmarrin/discordgo"
)

type Help struct {
}

func (hc Help) Execute(db database.IDatabase, s *discordgo.Session, m *discordgo.MessageCreate) error {
	_, err := s.ChannelMessageSend(m.ChannelID,
		"```!clear [n]          ->          delete this bots messages\n"+
			"!help               ->          show this message\n"+
			"!lastmonth          ->          show total stonks from last month\n"+
			"!me                 ->          show info about you\n"+
			"!mylastmonth        ->          show your stonks from last month\n"+
			"!mytotal            ->          show your total stonks\n"+
			"!subscribe          ->          subscribe to get stonked\n"+
			"!total              ->          show total stonk count\n"+
			"!unsubscribe        ->          unsubscribe from getting stonked\n```")
	return err
}

func (hc Help) GetKeyword() string {
	return "!help"
}

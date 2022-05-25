package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"stonks_bot/config"
	"stonks_bot/database"
	"stonks_bot/handlers"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	cfg           *config.Config
	stonk_counter = 0
)

func main() {
	cfg, err := config.ReadConfig("config.json")
	if err != nil {
		log.Println("failed reading config", err)
		return
	}

	discord, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		log.Panic(err)
	}
	db, err := database.GetDB()
	if err != nil {
		log.Panic(err)
	}

	h := handlers.SetupHandlers(db)

	discord.AddHandler(h.MessageCreate)
	discord.AddHandler(h.ReactionAdded)

	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()

}

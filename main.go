package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"snuzzers.me/Lunara/util/types"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

var (
	// Version is the current version of the application
	Version = "0.0.1"

	BotCredentials = types.BotCredentials{
		BotToken:     os.Getenv("BOT_TOKEN"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		ClientID:     os.Getenv("CLIENT_ID"),
	}
)

func main() {
	discord, err := discordgo.New("Bot " + BotCredentials.BotToken)
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
	}

	discord.AddHandler(messageCreate)

	discord.Identify.Intents = discordgo.IntentsGuildMessages

	err = discord.Open()
	if err != nil {
		log.Fatal("Error opening connection: ", err)
	}
	defer discord.Close()

	log.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	panic("Not implemented")
}

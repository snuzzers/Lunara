package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"snuzzers.me/Lunara/bot"
	"snuzzers.me/Lunara/util/types"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

var (
	// Version is the current version of the application
	Version = "0.0.1"

	BotCredentials = &types.BotCredentials{
		BotToken:     os.Getenv("BOT_TOKEN"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		ClientID:     os.Getenv("CLIENT_ID"),
		GuildID:      "1093003011476963442",
	}
	removeCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")

	s *discordgo.Session
)

func init() {
	flag.Parse()

	var err error
	s, err = discordgo.New("Bot " + BotCredentials.BotToken)
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := bot.CommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v\n", r.User.Username, r.User.Discriminator)
	})
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(bot.Commands))
	for i, v := range bot.Commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, BotCredentials.GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command:%v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Bot is now running. Press CTRL-C to exit.")
	<-stop

	if *removeCommands {
		log.Println("Removing commands...")
		// We need to fetch the commands, since deleting requires the command ID.
		// We are doing this from the returned commands on line 375, because using
		// this will delete all the commands, which might not be desirable, so we
		// are deleting only the commands that we added.
		// registeredCommands, err := s.ApplicationCommands(s.State.User.ID, *GuildID)
		// if err != nil {
		// 	log.Fatalf("Could not fetch registered commands: %v", err)
		// }

		for _, v := range registeredCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, BotCredentials.GuildID, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}

	log.Println("Gracefully shutting down.")
}

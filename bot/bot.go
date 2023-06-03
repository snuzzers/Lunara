package bot

import (
	"github.com/bwmarrin/discordgo"
)

var (
	IntegerOptionMinValue          = 1.0
	DmPermission                   = false
	DefaultMemberPermissions int64 = discordgo.PermissionManageServer

	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Ping the bot",
		},
	}

	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"ping": Ping,
	}
)

package bot

import (
	"github.com/bwmarrin/discordgo"
)

var (
	integerOptionMinValue          = 1.0
	DmPermission                   = false
	DefaultMemberPermissions int64 = discordgo.PermissionManageServer

	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Ping the bot",
		},
		{
			Name:        "command-with-options",
			Description: "Ping the bot with an option",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "string-option",
					Description: "String option",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "integer-option",
					Description: "Integer option",
					MaxValue:    10,
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionNumber,
					Name:        "number-option",
					Description: "Float option",
					MinValue:    &integerOptionMinValue,
					MaxValue:    10.1,
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Name:        "bool-option",
					Description: "Boolean option",
					Required:    true,
				},
			},
		},
	}

	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"ping":                 Ping,
		"command-with-options": CommandWithOptions,
	}
)

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
			Name:        "convert-media",
			Description: "Converts media to a different format.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionAttachment,
					Name:        "attachment",
					Description: "The attachment to convert.",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "format",
					Description: "The format to convert to.",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "mp4",
							Value: "mp4",
						},
						{
							Name:  "webm",
							Value: "webm",
						},
						{
							Name:  "gif",
							Value: "gif",
						},
						{
							Name:  "png",
							Value: "png",
						},
						{
							Name:  "jpg",
							Value: "jpg",
						},
						{
							Name:  "mp3",
							Value: "mp3",
						},
						{
							Name:  "wav",
							Value: "wav",
						},
						{
							Name:  "ogg",
							Value: "ogg",
						},
					},
				},
			},
		},
		// #region ExampleCommands
		// {
		// 	Name:        "ping",
		// 	Description: "Ping the bot",
		// },
		// {
		// 	Name:        "command-with-options",
		// 	Description: "Ping the bot with an option",
		// 	Options: []*discordgo.ApplicationCommandOption{
		// 		{
		// 			Type:        discordgo.ApplicationCommandOptionString,
		// 			Name:        "string-option",
		// 			Description: "String option",
		// 			Required:    true,
		// 		},
		// 		{
		// 			Type:        discordgo.ApplicationCommandOptionInteger,
		// 			Name:        "integer-option",
		// 			Description: "Integer option",
		// 			MaxValue:    10,
		// 			Required:    true,
		// 		},
		// 		{
		// 			Type:        discordgo.ApplicationCommandOptionNumber,
		// 			Name:        "number-option",
		// 			Description: "Float option",
		// 			MinValue:    &integerOptionMinValue,
		// 			MaxValue:    10.1,
		// 			Required:    true,
		// 		},
		// 		{
		// 			Type:        discordgo.ApplicationCommandOptionBoolean,
		// 			Name:        "bool-option",
		// 			Description: "Boolean option",
		// 			Required:    true,
		// 		},
		// 	},
		// },
		// {
		// 	Name:        "command-with-subcommands",
		// 	Description: "Subcommands example",
		// 	Options: []*discordgo.ApplicationCommandOption{
		// 		// When a command has subcommands/subcommand groups
		// 		// It must not have top-level options, they aren't accesible in the UI
		// 		// in this case (at least not yet), so if a command has
		// 		// subcommands/subcommand any groups registering top-level options
		// 		// will cause the registration of the command to fail
		// 		{
		// 			Name:        "subcommand-group",
		// 			Description: "Subcommands group",
		// 			Options: []*discordgo.ApplicationCommandOption{
		// 				// Also, subcommand groups aren't capable of
		// 				// containing options, by the name of them, you can see
		// 				// they can only contain subcommands
		// 				{
		// 					Name:        "nested-subcommand",
		// 					Description: "Nested subcommand",
		// 					Type:        discordgo.ApplicationCommandOptionSubCommand,
		// 				},
		// 			},
		// 			Type: discordgo.ApplicationCommandOptionSubCommandGroup,
		// 		},
		// 		// Also, you can create both subcommand groups and subcommands
		// 		// in the command at the same time. But, there's some limits to
		// 		// nesting, count of subcommands (top level and nested) and options.
		// 		// Read the intro of slash-commands docs on Discord dev portal
		// 		// to get more information
		// 		{
		// 			Name:        "subcommand",
		// 			Description: "Top-level subcommand",
		// 			Type:        discordgo.ApplicationCommandOptionSubCommand,
		// 		},
		// 	},
		// },
		// #endregion ExampleCommands
	}

	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		// #region ExampleCommands
		// "ping":                     Ping,
		// "command-with-options":     CommandWithOptions,
		// "command-with-subcommands": CommandWithSubcommands,
		// #endregion ExampleCommands

		"convert-media": ConvertMedia,
	}
)

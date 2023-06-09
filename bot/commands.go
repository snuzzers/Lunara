package bot

import (
	"github.com/bwmarrin/discordgo"
)

func ConvertMedia(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// TODO: Implement
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "This command is not yet implemented",
		},
	})
}

// #region ExampleCommands
// func Ping(s *discordgo.Session, i *discordgo.InteractionCreate) {
// 	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 		Type: discordgo.InteractionResponseChannelMessageWithSource,
// 		Data: &discordgo.InteractionResponseData{
// 			Content: "Pong!",
// 		},
// 	})
// }

// func CommandWithOptions(s *discordgo.Session, i *discordgo.InteractionCreate) {
// 	// Access options in the order provided by the user.
// 	options := i.ApplicationCommandData().Options

// 	// Or convert the slice into a map
// 	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
// 	for _, opt := range options {
// 		optionMap[opt.Name] = opt
// 	}

// 	// This example stores the provided arguments in an []interface{}
// 	// which will be used to format the bot's response
// 	margs := make([]interface{}, 0, len(options))
// 	msgformat := "You provided the following arguments:\n"

// 	// Get the value from the option map.
// 	// When the option exists, ok = true
// 	if option, ok := optionMap["string-option"]; ok {
// 		// Option values must be type asserted from interface{}.
// 		// Discordgo provides utility functions to make this simple.
// 		margs = append(margs, option.StringValue())
// 		msgformat += "> string-option: %s\n"
// 	}

// 	if opt, ok := optionMap["integer-option"]; ok {
// 		margs = append(margs, opt.IntValue())
// 		msgformat += "> integer-option: %d\n"
// 	}

// 	if opt, ok := optionMap["number-option"]; ok {
// 		margs = append(margs, opt.FloatValue())
// 		msgformat += "> number-option: %f\n"
// 	}

// 	if opt, ok := optionMap["bool-option"]; ok {
// 		margs = append(margs, opt.BoolValue())
// 		msgformat += "> bool-option: %v\n"
// 	}

// 	if opt, ok := optionMap["channel-option"]; ok {
// 		margs = append(margs, opt.ChannelValue(nil).ID)
// 		msgformat += "> channel-option: <#%s>\n"
// 	}

// 	if opt, ok := optionMap["user-option"]; ok {
// 		margs = append(margs, opt.UserValue(nil).ID)
// 		msgformat += "> user-option: <@%s>\n"
// 	}

// 	if opt, ok := optionMap["role-option"]; ok {
// 		margs = append(margs, opt.RoleValue(nil, "").ID)
// 		msgformat += "> role-option: <@&%s>\n"
// 	}

// 	// Send the response
// 	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 		Type: discordgo.InteractionResponseChannelMessageWithSource,
// 		Data: &discordgo.InteractionResponseData{
// 			Content: fmt.Sprintf(
// 				msgformat,
// 				margs...,
// 			),
// 		},
// 	})
// }

// func CommandWithSubcommands(s *discordgo.Session, i *discordgo.InteractionCreate) {
// 	options := i.ApplicationCommandData().Options
// 	content := ""

// 	// As you can see, names of subcommands (nested, top-level)
// 	// and subcommand groups are provided through the arguments.
// 	switch options[0].Name {
// 	case "subcommand":
// 		content = "The top-level subcommand is executed. Now try to execute the nested one."
// 	case "subcommand-group":
// 		options = options[0].Options
// 		switch options[0].Name {
// 		case "nested-subcommand":
// 			content = "Nice, now you know how to execute nested commands too"
// 		default:
// 			content = "Oops, something went wrong.\n" +
// 				"Hol' up, you aren't supposed to see this message."
// 		}
// 	}

// 	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 		Type: discordgo.InteractionResponseChannelMessageWithSource,
// 		Data: &discordgo.InteractionResponseData{
// 			Content: content,
// 		},
// 	})
// }

// #endregion ExampleCommands

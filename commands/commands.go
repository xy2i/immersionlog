package commands

import "github.com/bwmarrin/discordgo"

var (
	//Commands are represented with a struct with the fields below
	//https://pkg.go.dev/github.com/bwmarrin/discordgo#ApplicationCommand
	Commands = []*discordgo.ApplicationCommand{
		{
			Name: "basic-command",
			// All commands and options must have a description
			// Commands/options without description will fail the registration
			// of the command.
			Description: "Basic command",
		},
		{
			Name:        "log",
			Description: "You can use vndb Ids and titles for VN",
			Options: []*discordgo.ApplicationCommandOption{
				{
					// make this a choice, user can only pick one of the given ones.
					Name:         "mediatype",
					Description:  "mediatype",
					Type:         discordgo.ApplicationCommandOptionString,
					Required:     true,
					Autocomplete: true,
				},
				{
					Name:         "amount",
					Description:  "how much time you have spent on this mediatype in HH:MM format or MM",
					Type:         discordgo.ApplicationCommandOptionString,
					Autocomplete: false,
					Required:     true,
				},
				{
					Name:         "name",
					Description:  "You can use vndb ids and titles for VNs",
					Type:         discordgo.ApplicationCommandOptionString,
					Required:     true,
					Autocomplete: true,
				},
			},
		},
	}

	// a map with the name of the command as the key and the
	// function that handles that command as the value.
	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"basic-command": handleBasicCommand,
		"log":           handleLogCommand,
	}
)

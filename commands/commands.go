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
					Name:         "mediatype",
					Description:  "mediatype",
					Type:         discordgo.ApplicationCommandOptionString,
					Required:     true,
					Autocomplete: true,
				},
				{
					Name:         "amount",
					Description:  "time in minutes",
					Type:         discordgo.ApplicationCommandOptionNumber,
					Required:     true,
					Autocomplete: false,
				},
				{
					Name:         "name",
					Description:  "For VNs you can use the title or the vndb ID, for books BLAH BLAH BLAH",
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

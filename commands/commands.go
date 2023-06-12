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
	}

	// a map with the name of the command as the key and the
	// function that handles that command as the value.
	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"basic-command": handleBasicCommand,
		"false-command": handleBasicCommand,
	}
)

package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func handleBasicCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "itsover",
		},
	})
}

func handleLogCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		data := i.ApplicationCommandData()
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf(
					"Option 1: %s",
					data.Options[0].StringValue(),
				),
			},
		})
		if err != nil {
			panic(err)
		}
	case discordgo.InteractionApplicationCommandAutocomplete:
		data := i.ApplicationCommandData()
		var choices []*discordgo.ApplicationCommandOptionChoice
		switch {
		// In this case there are multiple autocomplete options. The Focused field shows which option user is focused on.
		case data.Options[0].Focused:
			choices = []*discordgo.ApplicationCommandOptionChoice{
				{
					Name:  "Visual Novel",
					Value: "VN",
				},
				{
					Name:  "Manga",
					Value: "Manga",
				},
				{
					Name:  "Anime",
					Value: "Anime",
				},
				{
					Name:  "Book",
					Value: "Book",
				},
				{
					Name:  "Readtime",
					Value: "Readtime",
				},
				{
					Name:  "Listening",
					Value: "Listening",
				},
				{
					Name:  "Reading",
					Value: "Reading",
				},
			}
			if data.Options[0].StringValue() != "" {
				choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
					Name:  data.Options[0].StringValue(),
					Value: "choice_custom",
				})
			}
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionApplicationCommandAutocompleteResult,
				Data: &discordgo.InteractionResponseData{
					Choices: choices,
				},
			})
			if err != nil {
				panic(err)
			}
		}
	}
}

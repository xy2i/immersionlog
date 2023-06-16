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
					"Option 1: %s for %s",
					data.Options[0].StringValue(),
					data.Options[1].StringValue(),
				),
			},
		})
		if err != nil {
			panic(err)
		}
	case discordgo.InteractionApplicationCommandAutocomplete:
		data := i.ApplicationCommandData()
		options := data.Options
		for k, v := range options {
			fmt.Printf("%s: %s\n", k, v.Name)
		}
		var choices []*discordgo.ApplicationCommandOptionChoice
		switch {
		// In this case there are multiple autocomplete options. The Focused field shows which option user is focused on.
		// replace this by working with a map
		//instead of array indexes.
		case data.Options[0].Focused:
			choices = []*discordgo.ApplicationCommandOptionChoice{
				{
					Name:  "Visual Novel",
					Value: "Visual Novel",
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
		case data.Options[2].Focused:
			if data.Options[0].StringValue() == "Visual Novel" {
				choices = []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Visual Novel",
						Value: "VN",
					},
				}
			} else {
				choices = []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Not a visual novel",
						Value: "not",
					},
				}
			}

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

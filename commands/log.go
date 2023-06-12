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
					"Option 1: %s\nOption 2: %s",
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
		var choices []*discordgo.ApplicationCommandOptionChoice
		switch {
		// In this case there are multiple autocomplete options. The Focused field shows which option user is focused on.
		case data.Options[0].Focused:
			choices = []*discordgo.ApplicationCommandOptionChoice{
				{
					Name:  "Visual Novel",
					Value: "visualnovel",
				},
				{
					Name:  "Novel",
					Value: "novel",
				},
				{
					Name:  "Anime",
					Value: "anime",
				},
				{
					Name:  "Choice 5",
					Value: "choice5",
				},
			}
			if data.Options[0].StringValue() != "" {
				choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
					Name:  data.Options[0].StringValue(),
					Value: "choice_custom",
				})
			}

		case data.Options[1].Focused:
			choices = []*discordgo.ApplicationCommandOptionChoice{
				{
					Name:  "Autocomplete 4 second option",
					Value: "autocomplete_1_default",
				},
				{
					Name:  "Choice 3.1",
					Value: "choice3_1",
				},
				{
					Name:  "Choice 4.1",
					Value: "choice4_1",
				},
				{
					Name:  "Choice 5.1",
					Value: "choice5_1",
				},
			}
			if data.Options[1].StringValue() != "" {
				choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
					Name:  data.Options[1].StringValue(),
					Value: "choice_custom_2",
				})
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

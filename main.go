package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var botToken string

//const guildId string

func main() {
	dg, err := discordgo.New("Bot " + botToken)
	if err != nil {
		fmt.Println("error starting Discord session", err)
		return
	}
	dg.AddHandler(messageCreate)
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
	}
	defer dg.Close()

	fmt.Println("Bot running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore bot messages
	if m.Author.ID == s.State.User.ID {
		return
	}
	// BOT can only view messages that mention it, or a DM.
	switch {
	case strings.Contains(m.Content, "hello"):
		s.ChannelMessageSend(m.ChannelID, "hello!!")
	case strings.Contains(m.Content, "bye"):
		s.ChannelMessageSend(m.ChannelID, "bye!")
	}
}

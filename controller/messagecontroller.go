package controller

import "github.com/bwmarrin/discordgo"

func MessageController(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
	
	if m.Content == "Hello" {
		s.ChannelMessageSend(m.ChannelID, "Henlo!")
	}

	if m.Content == "Hi" {
		s.ChannelMessageSend(m.ChannelID, "Henlo!")
	}
}
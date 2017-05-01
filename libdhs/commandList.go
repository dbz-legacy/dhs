package dhs

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func showHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	logger.Debug("Called help handler")

	var err error

	msg := &discordgo.MessageEmbed{}

	msg.Author = &discordgo.MessageEmbedAuthor{Name: "Discord Hetzner Service bot"}
	msg.Description = "Available commands for !dhs"
	msg.Fields = []*discordgo.MessageEmbedField {
		&discordgo.MessageEmbedField{
			Name: "help",
			Value: "show this message",
			Inline: false,
		},
		&discordgo.MessageEmbedField{
			Name: "status",
			Value: "show summary about servers",
			Inline: false,
		},
	}
	_, err = s.ChannelMessageSendEmbed(m.ChannelID, msg)
	if err != nil {
		logger.Error("Failed to send message", zap.Error(err))
	}
}
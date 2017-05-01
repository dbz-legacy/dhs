package dhs

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func commandRouter(s *discordgo.Session, m *discordgo.MessageCreate) {
	var tokens []string
	var command string

	if m.Author.ID == app.User.ID { return }

	tokens = strings.Split(m.Content, " ")
	if len(tokens) < 2 { return }
	if tokens[0] != commandPrefix {	return }

	command = tokens[1]
	switch command {
		case "status": getServerStatus(s, m)
		default: showHelp(s, m)
	}
}
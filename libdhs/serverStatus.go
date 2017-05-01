package dhs

import (
	"github.com/bwmarrin/discordgo"
	"github.com/appscode/go-hetzner"
	"go.uber.org/zap"
	"fmt"
	"github.com/ryanuber/columnize"
)

func getServerStatus(s *discordgo.Session, m *discordgo.MessageCreate) {
	logger.Debug("Called getServerStatus handler")

	var err error
	var servers []*hetzner.ServerSummary

	servers, _, err  = app.Hetzner.Server.ListServers()
	if err != nil {
		logger.Error("Can't get servers summary", zap.Error(err))
	}

	var message []string

	message = append(message, "#|Server IP|Server Number|Server Name|DC|Status")
	for idx, server := range servers {
		message = append(message,
			fmt.Sprintf("%d|%s|%d|%s|%s|%s",
				idx,
				server.ServerIP,
				server.ServerNumber,
				server.ServerName,
				server.Dc,
				server.Status))
	}

	result := columnize.SimpleFormat(message)
	result = "```md\n" + result + "```"

	s.ChannelMessageSend(m.ChannelID, result)
}

package dhs

import (
	"os"
	"go.uber.org/zap"
	"github.com/bwmarrin/discordgo"
	"github.com/appscode/go-hetzner"
	"os/signal"
	"syscall"
	"fmt"
)

const (
	version = "0.1.0"

	commandPrefix = "!dhs" //don't forget about " "

)

var logger *zap.Logger

type AppData struct {
	Mode            string

	DiscordToken    string
	Discord         *discordgo.Session
	User            *discordgo.User

	HetznerLogin    string
	HetznerPassword string
	Hetzner			*hetzner.Client
}

var app AppData

func init() {
	var err error

	app.Mode = os.Getenv("DHS_MODE")
	app.DiscordToken = os.Getenv("DHS_DISCORD_TOKEN")
	app.HetznerLogin= os.Getenv("DHS_HETZNER_LOGIN")
	app.HetznerPassword = os.Getenv("DHS_HETZNER_PASSWORD")

	if app.Mode == "release" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		panic("Can't initialize logger")
	}

	app.Hetzner = hetzner.NewClient(app.HetznerLogin, app.HetznerPassword)

	app.Discord, err = discordgo.New("Bot " + app.DiscordToken)
	if err != nil {
		logger.Error("Can't create bot", zap.Error(err))
	}

	// add handlers here
	app.Discord.AddHandler(commandRouter)

}

func Run() {
	var err error

	app.User, err = app.Discord.User("@me")

	if err != nil {
		logger.Error("Can't get self info", zap.Error(err))
	}

	err = app.Discord.Open()
	if err != nil {
		logger.Error("Can't open connection", zap.Error(err))
	}
	defer app.Discord.Close()

	addLink := fmt.Sprintf("Use this link: https://discordapp.com/oauth2/authorize?client_id=%s&scope=bot&permissions=0 to add bot on server.", app.User.ID)
	logger.Info(addLink)
	logger.Info("DHS bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

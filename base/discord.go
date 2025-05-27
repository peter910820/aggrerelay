package base

import (
	"context"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"

	"aggrerelay/model"
	"aggrerelay/relay"
)

func DiscordStart(ctx context.Context) {
	token := os.Getenv("DISCORD_TOKEN")
	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		logrus.Fatal(err)
	}

	// add handler
	bot.AddHandler(relay.DiscordRelay)

	err = bot.Open() // websocket connect
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("discord bot is now running...")

	<-ctx.Done()
	bot.Close() // websocket disconnect
	logrus.Info("discord bot is closing...")
	model.DiscordShutdownChan <- struct{}{}
}

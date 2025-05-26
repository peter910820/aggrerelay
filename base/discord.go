package base

import (
	"context"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func DiscordStart(ctx context.Context) {
	token := os.Getenv("DISCORD_TOKEN")
	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		logrus.Fatal(err)
	}

	// add handler
	// bot.AddHandler()

	err = bot.Open() // websocket connect
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("discord bot is now running...")

	<-ctx.Done()
	bot.Close() // websocket disconnect
	logrus.Info("discord bot is closing...")
	DiscordShutdownChan <- struct{}{}
}

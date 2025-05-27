package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"aggrerelay/base"
	"aggrerelay/model"
)

func init() {
	// init logrus settings
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.DebugLevel)
	// init env file
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf(".env file load error: %v", err)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// 定義一個全局變數存放所有的bot instance
	// open bot
	go base.LineStart(ctx)
	go base.DiscordStart(ctx)

	go func(ctx context.Context) {
		for {
			select {
			case msg := <-model.msgChan:
				// 在這邊處理轉送的邏輯
				go Translate(msg)
			case <-ctx.Done():
				return
			}
		}
	}(ctx)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	cancel()
	<-model.DiscordShutdownChan
	<-model.LineShutdownChan
	logrus.Println("program terminated...")
}

func Translate(msg) {
	// 轉送到非自己BotID的平台
}

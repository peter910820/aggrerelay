package base

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"

	"aggrerelay/relay"
)

func LineStart(ctx context.Context) {
	channelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	_, err := messaging_api.NewMessagingApiAPI(
		os.Getenv("LINE_CHANNEL_TOKEN"),
	)
	if err != nil {
		logrus.Fatal(err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	app.Post("/callback", func(c *fiber.Ctx) error {
		// convert *fiber.Ctx to *http.Request
		req, err := adaptor.ConvertRequest(c, false)
		if err != nil {
			logrus.Errorf("translate failed: %s", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		events, err := webhook.ParseRequest(channelSecret, req)
		if err != nil {
			if err == webhook.ErrInvalidSignature {
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		// main event handler
		for _, event := range events.Events {
			// event type
			switch e := event.(type) {
			case webhook.MessageEvent:
				switch e.Message.(type) {
				case webhook.TextMessageContent:
					go relay.LineRelay()
					// add your bot feature
				}
			}
		}

		return c.SendStatus(fiber.StatusOK)
	})

	go func() {
		if err := app.Listen(fmt.Sprintf("127.0.0.1:%s", os.Getenv("LINE_FIBER_PORT"))); err != nil {
			logrus.Errorf("Fiber server stopped: %v", err)
		}
	}()

	<-ctx.Done()
	if err := app.Shutdown(); err != nil {
		logrus.Errorf("Fiber shutdown error: %v", err)
	}
	logrus.Info("Fiber server is shutting down...")
	LineShutdownChan <- struct{}{}
}

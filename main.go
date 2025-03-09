package main

import (
	"github.com/famesensor/playground-go-telegram-bot/adapter/telegram"
	"github.com/famesensor/playground-go-telegram-bot/config"
	"github.com/famesensor/playground-go-telegram-bot/handler"
	"github.com/famesensor/playground-go-telegram-bot/helper/resty"
	"github.com/famesensor/playground-go-telegram-bot/infra"
	"github.com/famesensor/playground-go-telegram-bot/service/notify"
)

func main() {
	cfg := config.NewConfig()

	bot := infra.NewTelegramBot(cfg.Telegram.Token)

	// u := tgbotapi.NewUpdate(0)
	// u.Timeout = 60
	// updates := bot.GetUpdatesChan(u)

	// for update := range updates {
	// 	if update.Message != nil { // If we got a message
	// 		log.Printf("[%s] text: %s", update.Message.From.UserName, update.Message.Text)
	// 		log.Printf("[%s] chat id: %d", update.Message.From.UserName, update.Message.Chat.ID)
	// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	// 		msg.ReplyToMessageID = update.Message.MessageID

	// 		bot.Send(msg)
	// 	}
	// }

	clientHTTP := resty.NewResty()

	notifyAdapter := telegram.NewAdapter(clientHTTP, bot, cfg.Telegram.Token)

	notifyService := notify.NewService(notifyAdapter)

	notifyHandler := handler.NewNotifyHandler(notifyService)

	httpServer := infra.NewHTTPServer(cfg)

	httpServer.App.Post("/v1/notify", notifyHandler.Notify)

	httpServer.Start()
}

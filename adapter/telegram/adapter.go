package telegram

import (
	"fmt"

	"github.com/famesensor/playground-go-telegram-bot/port"
	"github.com/go-resty/resty/v2"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Adapter struct {
	client *resty.Client
	bot    *tgbotapi.BotAPI
	token  string
}

func NewAdapter(client *resty.Client, bot *tgbotapi.BotAPI, token string) port.NotifyPort {
	return &Adapter{client: client, token: token, bot: bot}
}

// Send implements port.NotifyPort.
func (a *Adapter) Send(chatID int, msg string) error {
	_, err := a.client.R().
		SetBody(map[string]interface{}{
			"chat_id": chatID,
			"text":    msg,
		}).
		Post(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", a.token))
	if err != nil {
		return err
	}

	return nil
}

func (a *Adapter) SendBot(chatID int, msg string) error {
	ms := tgbotapi.NewMessage(int64(chatID), fmt.Sprintf("bot: %s", msg))
	_, err := a.bot.Send(ms)
	if err != nil {
		return err
	}

	return nil
}

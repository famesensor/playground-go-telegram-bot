package notify

import (
	"github.com/famesensor/playground-go-telegram-bot/port"
)

type Port interface {
	Send(message string) error
}

type service struct {
	telegramAdapter port.NotifyPort
}

func NewService(telegramAdapter port.NotifyPort) Port {
	return &service{
		telegramAdapter: telegramAdapter,
	}
}

func (s *service) Send(message string) error {
	_ = s.telegramAdapter.Send(5429734773, message)
	_ = s.telegramAdapter.SendBot(5429734773, message)
	return nil
}

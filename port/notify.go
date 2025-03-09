package port

type NotifyPort interface {
	Send(chatID int, msg string) error
	SendBot(chatID int, msg string) error
}

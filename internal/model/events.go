package model

type Type int

const (
	Unknown Type = iota
	Message
	CallbackQuery
)

type Event struct {
	Type     Type
	Text     string
	ChatID   int
	UserID   int
	Username string
	Meta     interface{}
}

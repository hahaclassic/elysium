package logger

import (
	"context"
	"fmt"

	"github.com/hahaclassic/elysium/config"
	tgclient "github.com/hahaclassic/elysium/internal/client/telegram"
)

const (
	InfoLevel     = "info"
	ErrorLevel    = "error"
	FeedbackLevel = "feedback"
)

type LoggerBot struct {
	logLevel    string
	adminChatID int64
	client      *tgclient.Client
}

func New(conf config.LoggerConfig, client *tgclient.Client) *LoggerBot {
	return &LoggerBot{
		client:      client,
		adminChatID: conf.AdminID,
		logLevel:    conf.Level,
	}
}

func (l *LoggerBot) Error(ctx context.Context, msg string) error {
	if l.logLevel == ErrorLevel {
		return l.SendMessage(ctx, l.wrap(ErrorLevel, msg))
	}

	return nil
}

func (l *LoggerBot) Info(ctx context.Context, msg string) error {
	if l.logLevel == InfoLevel {
		return l.SendMessage(ctx, l.wrap(InfoLevel, msg))
	}

	return nil
}

func (l *LoggerBot) Feedback(ctx context.Context, msg string) error {
	if l.logLevel == FeedbackLevel {
		return l.SendMessage(ctx, l.wrap(FeedbackLevel, msg))
	}

	return nil
}

func (l *LoggerBot) SendMessage(ctx context.Context, msg string) error {
	outMsg := &tgclient.OutputMessage{
		ChatID: int(l.adminChatID),
		Text:   msg,
	}

	_, err := l.client.SendMessage(ctx, outMsg)

	return err
}

func (l *LoggerBot) wrap(tag string, msg string) string {
	return fmt.Sprintf("#%s\n\n%s", tag, msg)
}

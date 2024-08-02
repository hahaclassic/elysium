package service

import (
	"context"

	tgclient "github.com/hahaclassic/elysium/internal/client/telegram"
	"github.com/hahaclassic/elysium/internal/storage"
	"github.com/hahaclassic/elysium/pkg/syncmap"
)

type Logger interface {
	Error(ctx context.Context, msg string)
	Info(ctx context.Context, msg string)
	Feedback(ctx context.Context, msg string)
}

type Processor struct {
	tg       *tgclient.Client
	logger   Logger
	offset   int
	storage  storage.Storage
	sessions *syncmap.Map
}

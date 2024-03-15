package storage

import (
	"context"

	"github.com/Shemistan/chat_server/internal/model"
)

// Chat - чат
type Chat interface {
	CreateChat(ctx context.Context, req model.Chat) (int64, error)
	AddMessage(ctx context.Context, req model.Message) error
	DeactivateChat(ctx context.Context, chatID int64) error
}

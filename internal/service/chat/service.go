package chat

import (
	"context"

	"github.com/Shemistan/chat_server/internal/model"
	def "github.com/Shemistan/chat_server/internal/service"
	"github.com/Shemistan/chat_server/internal/storage"
)

type service struct {
	storage storage.Chat
}

// NewService - новый сервис
func NewService(storage storage.Chat) def.Chat {
	return &service{storage: storage}
}

// CreateChat - создать чат
func (s *service) CreateChat(ctx context.Context, req model.Chat) (int64, error) {
	return s.storage.CreateChat(ctx, req)
}

// AddMessage - добавить сообщение в чат
func (s *service) AddMessage(ctx context.Context, req model.Message) error {
	return s.storage.AddMessage(ctx, req)
}

// DeactivateChat - деактивировать чат
func (s *service) DeactivateChat(ctx context.Context, chatID int64) error {
	return s.storage.DeactivateChat(ctx, chatID)
}

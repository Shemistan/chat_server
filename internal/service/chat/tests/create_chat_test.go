package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Shemistan/chat_server/internal/model"
	"github.com/Shemistan/chat_server/internal/service/chat"
	"github.com/Shemistan/chat_server/internal/service/mocks"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()
	storage := new(mocks.Chat)
	service := chat.NewService(storage)
	errorTest := errors.New("test error")

	t.Run("storage error", func(t *testing.T) {
		storage.On("CreateChat", ctx, model.Chat{}).
			Return(int64(0), errorTest).Once()

		actual, err := service.CreateChat(ctx, model.Chat{})
		assert.ErrorIs(t, err, errorTest)
		assert.Equal(t, int64(0), actual)
	})

	t.Run("storage valid", func(t *testing.T) {
		storage.On("CreateChat", ctx, model.Chat{}).
			Return(int64(1), nil).Once()

		actual, err := service.CreateChat(ctx, model.Chat{})
		assert.Nil(t, err)
		assert.Equal(t, int64(1), actual)
	})
}

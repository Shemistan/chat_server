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

func TestAddMessage(t *testing.T) {
	ctx := context.Background()
	storage := new(mocks.Chat)
	service := chat.NewService(storage)
	errorTest := errors.New("test error")

	t.Run("storage error", func(t *testing.T) {
		storage.On("AddMessage", ctx, model.Message{}).
			Return(errorTest).Once()

		err := service.AddMessage(ctx, model.Message{})
		assert.ErrorIs(t, err, errorTest)
	})

	t.Run("storage valid", func(t *testing.T) {
		storage.On("AddMessage", ctx, model.Message{}).
			Return(nil).Once()

		err := service.AddMessage(ctx, model.Message{})
		assert.Nil(t, err)
	})
}

package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	chatAPI "github.com/Shemistan/chat_server/internal/api/chat_v1"
	"github.com/Shemistan/chat_server/internal/converter"
	"github.com/Shemistan/chat_server/internal/service/mocks"
	pb "github.com/Shemistan/chat_server/pkg/chat_api_v1"
)

func TestCreate(t *testing.T) {
	service := new(mocks.Chat)
	ctx := context.Background()
	api := chatAPI.New(service)
	errorTest := errors.New("test error")

	reqAPI := &pb.CreateRequest{
		ChatName:   "ChatName",
		UserLogins: []string{"test1", "test2"},
	}

	t.Run("service error", func(t *testing.T) {
		service.On("CreateChat", ctx, converter.RPCCreateChatToModelChat(reqAPI)).
			Return(int64(0), errorTest).Once()

		actual, err := api.Create(ctx, reqAPI)
		assert.ErrorIs(t, err, errorTest)
		assert.Nil(t, actual)
	})

	t.Run("service valid", func(t *testing.T) {
		service.On("CreateChat", ctx, converter.RPCCreateChatToModelChat(reqAPI)).
			Return(int64(1), nil).Once()

		actual, err := api.Create(ctx, reqAPI)
		assert.Nil(t, err)
		assert.Equal(t, int64(1), actual.Id)
	})
}

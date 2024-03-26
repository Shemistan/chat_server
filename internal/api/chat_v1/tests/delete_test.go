package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"

	chatAPI "github.com/Shemistan/chat_server/internal/api/chat_v1"
	"github.com/Shemistan/chat_server/internal/service/mocks"
	pb "github.com/Shemistan/chat_server/pkg/chat_api_v1"
)

func TestDelete(t *testing.T) {
	service := new(mocks.Chat)
	ctx := context.Background()
	api := chatAPI.New(service)
	errorTest := errors.New("test error")
	actual := &emptypb.Empty{}

	t.Run("service error", func(t *testing.T) {
		service.On("DeactivateChat", ctx, int64(1)).Return(errorTest).Once()
		expected, err := api.DeleteChat(ctx, &pb.DeleteChatRequest{ChatId: int64(1)})

		assert.ErrorIs(t, err, errorTest)
		assert.Nil(t, expected)
	})

	t.Run("service valid", func(t *testing.T) {
		service.On("DeactivateChat", ctx, int64(1)).Return(nil).Once()
		expected, err := api.DeleteChat(ctx, &pb.DeleteChatRequest{ChatId: int64(1)})

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
}

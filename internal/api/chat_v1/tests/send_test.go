package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"

	chatAPI "github.com/Shemistan/chat_server/internal/api/chat_v1"
	"github.com/Shemistan/chat_server/internal/converter"
	"github.com/Shemistan/chat_server/internal/service/mocks"
	pb "github.com/Shemistan/chat_server/pkg/chat_api_v1"
)

func TestSend(t *testing.T) {
	service := new(mocks.Chat)
	ctx := context.Background()
	api := chatAPI.New(service)

	reqAPI := &pb.SendMessageRequest{
		ChatName:  "ChatName",
		Message:   "Message",
		UserLogin: "UserLogin",
	}

	t.Run("service error", func(t *testing.T) {
		errorTest := errors.New("test error")

		service.On("AddMessage", ctx, converter.RPCSendMessageToModelMessage(reqAPI)).
			Return(errorTest).Once()

		expected, err := api.SendMessage(ctx, reqAPI)
		assert.ErrorIs(t, err, errorTest)
		assert.Nil(t, expected)
	})

	t.Run("service valid", func(t *testing.T) {
		actual := &emptypb.Empty{}

		service.On("AddMessage", ctx, converter.RPCSendMessageToModelMessage(reqAPI)).
			Return(nil).Once()

		expected, err := api.SendMessage(ctx, reqAPI)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)

	})
}

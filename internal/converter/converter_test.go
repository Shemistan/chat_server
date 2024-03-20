package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Shemistan/chat_server/internal/model"
	pb "github.com/Shemistan/chat_server/pkg/chat_api_v1"
)

func TestRPCCreateChatToModelChat(t *testing.T) {
	t.Run("req equal nil", func(t *testing.T) {
		expected := model.Chat{}
		actual := RPCCreateChatToModelChat(nil)
		assert.Equal(t, expected, actual)
	})

	t.Run("not users in req", func(t *testing.T) {
		expected := model.Chat{
			ID:    0,
			Name:  "test",
			Users: nil,
		}
		actual := RPCCreateChatToModelChat(&pb.CreateRequest{
			ChatName:   "test",
			UserLogins: nil,
		})
		assert.Equal(t, expected, actual)
	})

	t.Run("req is valid", func(t *testing.T) {
		expected := model.Chat{
			ID:    0,
			Name:  "test",
			Users: []string{"test1", "test2"},
		}
		actual := RPCCreateChatToModelChat(&pb.CreateRequest{
			ChatName:   "test",
			UserLogins: []string{"test1", "test2"},
		})
		assert.Equal(t, expected, actual)
	})
}

func TestRPCSendMessageToModelMessage(t *testing.T) {
	t.Run("req equal nil", func(t *testing.T) {
		expected := model.Message{}
		actual := RPCSendMessageToModelMessage(nil)
		assert.Equal(t, expected, actual)
	})

	t.Run("not message in req", func(t *testing.T) {
		expected := model.Message{
			ChatName:  "ChatName",
			UserLogin: "UserLogin",
			Message:   "",
		}
		actual := RPCSendMessageToModelMessage(&pb.SendMessageRequest{
			ChatName:  "ChatName",
			Message:   "",
			UserLogin: "UserLogin",
		})
		assert.Equal(t, expected, actual)
	})

	t.Run("req is valid", func(t *testing.T) {
		expected := model.Message{
			ChatName:  "ChatName",
			UserLogin: "UserLogin",
			Message:   "Message",
		}
		actual := RPCSendMessageToModelMessage(&pb.SendMessageRequest{
			ChatName:  "ChatName",
			Message:   "Message",
			UserLogin: "UserLogin",
		})
		assert.Equal(t, expected, actual)
	})
}

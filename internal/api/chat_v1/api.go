package chat_v1

import (
	"github.com/Shemistan/chat_server/internal/service"
	pb "github.com/Shemistan/chat_server/pkg/chat_api_v1"
)

// Chat - структура реализующая методы АПИ
type Chat struct {
	pb.UnimplementedChatV1Server

	chatService service.Chat
}

// New - новый экземпляр АПИ
func New(chatService service.Chat) *Chat {
	return &Chat{
		chatService: chatService,
	}
}

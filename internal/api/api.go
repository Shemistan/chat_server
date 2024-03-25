package api

import (
	"github.com/Shemistan/chat_server/internal/service"
	pb "github.com/Shemistan/chat_server/pkg/chat_api_v1"
)

// Chat - структура реализующая методы АПИ
type Chat struct {
	pb.UnimplementedChatV1Server

	Service service.Chat
}

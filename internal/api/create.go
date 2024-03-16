package api

import (
	"context"

	"github.com/Shemistan/chat_server/internal/converter"
	pb "github.com/Shemistan/chat_server/pkg/chat_api_v1"
)

// Create - создать новый чат
func (c *Chat) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	chatID, err := c.chatService.CreateChat(ctx, converter.RPCCreateChatToModelChat(req))
	if err != nil {
		return nil, err
	}
	return &pb.CreateResponse{Id: chatID}, nil
}

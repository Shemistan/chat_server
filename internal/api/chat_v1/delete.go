package chat_v1

import (
	"context"

	pb "github.com/Shemistan/chat_server/pkg/chat_api_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// DeleteChat - удаление чата из системы по его идентификатору
func (c *Chat) DeleteChat(ctx context.Context, req *pb.DeleteChatRequest) (*emptypb.Empty, error) {
	err := c.chatService.DeactivateChat(ctx, req.GetChatId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

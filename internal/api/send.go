package api

import (
	"context"

	"github.com/Shemistan/chat_server/internal/converter"
	pb "github.com/Shemistan/chat_server/pkg/chat_api_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// SendMessage - отправить сообщение на сервер
func (c *Chat) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*emptypb.Empty, error) {
	err := c.Service.AddMessage(ctx, converter.RPCSendMessageToModelMessage(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

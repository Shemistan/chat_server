package api

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Shemistan/chat_server/pkg/chat_api_v1"
)

// SendMessage - отправить сообщение на сервер
func (u *Chat) SendMessage(_ context.Context, req *pb.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("request: %+v", req)
	return &emptypb.Empty{}, nil
}

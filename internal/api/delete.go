package api

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Shemistan/chat_server/pkg/chat_api_v1"
)

// Delete - удаление чата из системы по его идентификатору
func (c *Chat) Delete(_ context.Context, req *pb.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("request: %+v", req)
	return &emptypb.Empty{}, nil
}

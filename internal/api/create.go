package api

import (
	"context"
	"log"

	pb "github.com/Shemistan/chat_server/pkg/chat_api_v1"
)

// Create - создать новый чат
func (c *Chat) Create(_ context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	log.Printf("request: %+v", req)
	return &pb.CreateResponse{Id: 1}, nil
}

package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Shemistan/chat_server/internal/api"
	pb "github.com/Shemistan/chat_server/pkg/chat_api_v1"
)

const grpcPort = 50052

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalln("failed to listen: ", err.Error())
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterChatV1Server(s, &api.Chat{})

	log.Println("server listening at:", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalln("failed to server:", err.Error())
	}
}

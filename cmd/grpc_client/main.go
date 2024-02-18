package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Shemistan/chat_server/pkg/chat_api_v1"
	"github.com/golang/protobuf/ptypes/timestamp"
)

const (
	serviceAddress = "localhost:50052"
	userID         = 1
)

func main() {
	conn, err := grpc.Dial(serviceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("failed to connect to service:", err.Error())
	}
	defer conn.Close() //nolint:errcheck

	c := pb.NewChatV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	runClientMethods(ctx, c)
}

func runClientMethods(ctx context.Context, c pb.ChatV1Client) {
	_, err := c.Delete(ctx, &pb.DeleteRequest{Id: userID})
	if err != nil {
		log.Println("failed to delete user: ", err.Error())
	}

	resCreate, err := c.Create(ctx, &pb.CreateRequest{
		UserNames: []string{"name_1", "name2"},
	})
	if err != nil {
		log.Println("failed to Create user: ", err.Error())
	}
	log.Printf("create response: %+v\n", resCreate)

	_, err = c.SendMessage(ctx, &pb.SendMessageRequest{
		From:      "test_from",
		Text:      "test_text",
		Timestamp: &timestamp.Timestamp{},
	})
	if err != nil {
		log.Println("failed to SendMessage: ", err.Error())
	}
}

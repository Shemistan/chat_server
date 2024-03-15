package main

import (
	"context"
	"flag"

	"log"
	"net"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Shemistan/chat_server/internal/api"
	"github.com/Shemistan/chat_server/internal/config"
	"github.com/Shemistan/chat_server/internal/config/env"
	"github.com/Shemistan/chat_server/internal/service/chat"
	user "github.com/Shemistan/chat_server/internal/storage/chat"
	pb "github.com/Shemistan/chat_server/pkg/chat_api_v1"
)

var configPath string

func init() {
	//configPath = ".env"
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	flag.Parse()
	ctx := context.Background()

	//Считываем переменные окружения
	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalln("failed to listen: ", err.Error())
	}

	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	if errPing := pool.Ping(ctx); errPing != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	storage := user.NewStorage(pool)
	service := chat.NewService(storage)

	pb.RegisterChatV1Server(s, &api.Chat{
		Service: service,
	})

	log.Println("server listening at:", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalln("failed to server:", err.Error())
	}
}

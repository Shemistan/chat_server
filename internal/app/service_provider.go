package app

import (
	"context"
	"log"

	"github.com/Shemistan/chat_server/internal/api"
	"github.com/Shemistan/chat_server/internal/client/db"
	"github.com/Shemistan/chat_server/internal/client/db/pg"
	"github.com/Shemistan/chat_server/internal/client/db/transaction"
	"github.com/Shemistan/chat_server/internal/closer"
	"github.com/Shemistan/chat_server/internal/config"
	"github.com/Shemistan/chat_server/internal/config/env"
	"github.com/Shemistan/chat_server/internal/service"
	chatService "github.com/Shemistan/chat_server/internal/service/chat"
	"github.com/Shemistan/chat_server/internal/storage"
	chatStorage "github.com/Shemistan/chat_server/internal/storage/chat"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient  db.Client
	txManager db.TxManager

	chatStorage storage.Chat
	chatService service.Chat
	chatAPI     *api.Chat
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) ChatStorage(ctx context.Context) storage.Chat {
	if s.chatStorage == nil {
		s.chatStorage = chatStorage.NewStorage(s.DBClient(ctx), s.TxManager(ctx))
	}

	return s.chatStorage
}

func (s *serviceProvider) ChatService(ctx context.Context) service.Chat {
	if s.chatService == nil {
		s.chatService = chatService.NewService(
			s.ChatStorage(ctx),
		)
	}

	return s.chatService
}

func (s *serviceProvider) ChatAPI(ctx context.Context) *api.Chat {
	if s.chatAPI == nil {
		s.chatAPI = api.New(s.ChatService(ctx))
	}

	return s.chatAPI
}

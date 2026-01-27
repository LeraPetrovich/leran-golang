package service

import (
	"example.com/learn/pkg/learn/core"
	appStorage "example.com/learn/pkg/learn/storage"
	"go.uber.org/zap"
)

type Service struct {
	logger            *zap.Logger
	privateAppStorage core.AppStorage
}

type Config struct {
	AppPostgresURI string
}

func New(logger *zap.Logger, conf Config) (*Service, error) {
	appStorage, err := appStorage.New(logger, conf.AppPostgresURI, appStorage.Config{PostgresConnections: 1})

	if err != nil {
		return nil, err
	}
	return &Service{
		logger:            logger,
		privateAppStorage: appStorage,
	}, nil
}

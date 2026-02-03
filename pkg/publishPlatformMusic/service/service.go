package service

import (
	"example.com/learn/pkg/learn/core"
	appStorage "example.com/learn/pkg/learn/storage"
	"go.uber.org/zap"
)

type service struct {
	logger            *zap.Logger
	privateAppStorage core.AppStorage
}

type ServiceConfig struct {
	AppPostgresURI string
}

func NewService(logger *zap.Logger, conf *ServiceConfig) (*service, error) {
	storage, err := appStorage.New(logger, conf.AppPostgresURI, appStorage.Config{PostgresConnections: 1})
	if err != nil {
		return nil, err
	}
	return &service{
		logger:            logger,
		privateAppStorage: storage,
	}, nil
}

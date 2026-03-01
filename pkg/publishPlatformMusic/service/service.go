package service

import (
	"example.com/learn/pkg/publishPlatformMusic/core"
	appStorage "example.com/learn/pkg/publishPlatformMusic/storage"
	"go.uber.org/zap"
)

type Service struct {
	logger            *zap.Logger
	privateAppStorage core.AppStorage
}

type ServiceConfig struct {
	AppPostgresURI string
}

func NewService(logger *zap.Logger, conf *ServiceConfig) (*Service, error) {
	storage, err := appStorage.New(logger, conf.AppPostgresURI, appStorage.StorageConf{PostgresConnections: 1})
	if err != nil {
		return nil, err
	}
	return &Service{
		logger:            logger,
		privateAppStorage: storage,
	}, nil
}

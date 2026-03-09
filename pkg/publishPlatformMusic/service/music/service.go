package music

import (
	"example.com/learn/pkg/publishPlatformMusic/core"
	"example.com/learn/pkg/publishPlatformMusic/service"
	"go.uber.org/zap"
)

var _ service.ServiceInterface = (*Service)(nil)

type Service struct {
	logger            *zap.Logger
	privateAppStorage core.AppStorage
}

func NewService(logger *zap.Logger, storage core.AppStorage) *Service {
	return &Service{
		logger:            logger,
		privateAppStorage: storage,
	}
}

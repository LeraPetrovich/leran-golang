package api

import (
	appService "example.com/learn/pkg/learn/service"
	"go.uber.org/zap"
)

type Handler struct {
	logger            *zap.Logger
	privateAppService *appService.Service
	// jwtSecret         []byte
}

type HandlerConfig struct {
	AppPostgresURI string
	// JwtSecret      []byte
}

func NewHandler(logger *zap.Logger, config HandlerConfig) (*Handler, error) {
	service, err := appService.New(logger, appService.Config{AppPostgresURI: config.AppPostgresURI})

	if err != nil {
		return nil, err
	}

	return &Handler{
		logger:            logger,
		privateAppService: service,
	}, nil
}

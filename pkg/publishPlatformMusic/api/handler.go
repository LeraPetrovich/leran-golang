package api

import (
	appService "example.com/learn/pkg/publishPlatformMusic/service"
	"go.uber.org/zap"
)

type Handler struct {
	logger            *zap.Logger
	privateAppService *appService.Service
	jwtSecret         []byte
}

type HandlerConfig struct {
	AppPostgresURI string
	JwtSecret      []byte
}

func NewHandler(logger *zap.Logger, config HandlerConfig) (*Handler, error) {
	service, err := appService.NewService(logger, &appService.ServiceConfig{
		AppPostgresURI: config.AppPostgresURI,
	})

	if err != nil {
		return nil, err
	}

	return &Handler{
		logger:            logger,
		privateAppService: service,
		jwtSecret:         config.JwtSecret,
	}, nil
}

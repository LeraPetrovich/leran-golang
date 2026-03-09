package api

import (
	appService "example.com/learn/pkg/publishPlatformMusic/service"
	"go.uber.org/zap"
)

type Handler struct {
	logger            *zap.Logger
	privateAppService appService.ServiceInterface
	jwtSecret         []byte
}

type HandlerConfig struct {
	JwtSecret []byte
}

func NewHandler(logger *zap.Logger, service appService.ServiceInterface, config HandlerConfig) *Handler {

	return &Handler{
		logger:            logger,
		privateAppService: service,
		jwtSecret:         config.JwtSecret,
	}
}

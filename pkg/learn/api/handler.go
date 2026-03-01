package api

import (
	"context"
	"net/http"

	"example.com/learn/pkg/learn/api/oas"
	"example.com/learn/pkg/learn/service"
	"go.uber.org/zap"
)

type Handler struct {
	logger            *zap.Logger
	appPrivateService *service.Service
	jwtSecret         []byte
}

type HandlerConfig struct {
	PostgresURI string
	JwtSecret   []byte
}

func NewHandler(logger *zap.Logger, cnf HandlerConfig) (*Handler, error) {
	service, err := service.New(logger, service.Config{AppPostgresURI: cnf.PostgresURI})

	if err != nil {
		return nil, err
	}

	return &Handler{
		logger:            logger,
		appPrivateService: service,
		jwtSecret:         []byte(cnf.PostgresURI),
	}, nil
}

func (h *Handler) NewError(ctx context.Context, err error) *oas.ErrorStatusCode {
	return &oas.ErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: oas.Error{
			Error: err.Error(),
		},
	}
}

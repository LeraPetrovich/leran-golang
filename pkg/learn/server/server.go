package server

import (
	"context"
	"errors"
	"net/http"

	"example.com/learn/pkg/learn/api"
	"example.com/learn/pkg/learn/api/oas"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

type Server struct {
	logger     *zap.Logger
	httpServer *http.Server
	mux        *http.ServeMux
}

type ServerConfig struct {
	Handler   *api.Handler
	JwtSecret []byte
	Address   string
}

func New(logger *zap.Logger, config ServerConfig) (*Server, error) {
	securityHandler := api.NewSecurityHandler(config.JwtSecret)
	ogenServer, err := oas.NewServer(config.Handler, securityHandler)
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	mux.Handle("/", cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	}).Handler(ogenServer))

	return &Server{
		logger: logger,
		mux:    mux,
		httpServer: &http.Server{
			Addr:    config.Address,
			Handler: mux,
		},
	}, nil
}

func (s *Server) Run() {
	s.logger.Info("starting server", zap.String("addr", s.httpServer.Addr))
	err := s.httpServer.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		s.logger.Info("homepage server quit")
		return
	}
	s.logger.Fatal("ListenAndServe() failed", zap.Error(err))
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

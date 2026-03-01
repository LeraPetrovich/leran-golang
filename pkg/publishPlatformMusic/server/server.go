package server

import (
	"errors"
	"log"

	"example.com/learn/pkg/publishPlatformMusic/api"
	oas "example.com/learn/pkg/publishPlatformMusic/api/oas"
	"github.com/rs/cors"

	"net/http"
)

type Server struct {
	httpServer *http.Server
	mux        *http.ServeMux
}

type ServerConfig struct {
	JwtSecret []byte
	Handler   *api.Handler
	Address   string
}

func NewServer(conf *ServerConfig) (*Server, error) {
	securityHandler := api.NewSecurityHandler(conf.JwtSecret)
	ogenServer, err := oas.NewServer(conf.Handler, securityHandler)
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
		httpServer: &http.Server{
			Addr:    conf.Address,
			Handler: mux,
		},
		mux: mux,
	}, nil
}

func (s *Server) Start() {
	err := s.httpServer.ListenAndServe()
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			log.Printf("server quit: %v", err)
			return
		}
		log.Fatalf("ListenAndServe() failed %v", err)
	}
}

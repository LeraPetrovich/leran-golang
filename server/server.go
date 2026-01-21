package server

import (
	"context"
	"net/http"
	"time"
)

// создаем сущность сервера (ее можно переиспользовать)
type Server struct {
	httpServer *http.Server
}

// функции работы с этой сущностью
// hendler это структура в которой хранится рунтинг с функциями обработки запросов с клиента
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, //1mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

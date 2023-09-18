package webserver

import (
	"context"
	"net/http"

	"main/structuredlogger"
)

type Server struct {
	httpServer *http.Server
	mux        *http.ServeMux
	ErrChan    chan error
	logger     *structuredlogger.CustomLogger
}

type ServerOption func(*Server)

func NewServer(addr string) *Server {
	mux := http.NewServeMux()
	server := &Server{
		httpServer: &http.Server{
			Addr:    addr,
			Handler: mux,
		},
		mux:     mux,
		ErrChan: make(chan error),
	}

	return server
}

func (s *Server) AddLogger(l *structuredlogger.CustomLogger) *Server {
	s.logger = l
	return s
}

func (s *Server) AddRoute(path string, handler http.HandlerFunc) {
	s.mux.HandleFunc(path, handler)
}

func (s *Server) Start() {
	go func() {
		s.ErrChan <- s.httpServer.ListenAndServe()
	}()
}

func (s *Server) GracefulShutdown(ctx context.Context) {
	if s.logger != nil {
		s.logger.Logger.Info("Shutting down HTTP Server")
	}

	s.httpServer.Shutdown(ctx)
}

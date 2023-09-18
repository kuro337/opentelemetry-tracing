package webserver

import (
	"context"
	"net/http"

	"main/constants"
	"main/structuredlogger"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
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
	t := otel.GetTracerProvider().Tracer(constants.NAME)

	cleanupCtx, span := t.Start(ctx, "SHUTDOWN")

	span.SetAttributes(attribute.String("SERVER_SHUTTING_DOWN", "SHUTDOWN"))

	span.End()

	if err := s.httpServer.Shutdown(cleanupCtx); err != nil {
		s.logger.Logger.Error("Error during server shutdown:", err)
	} else {
		s.logger.Logger.Info("Shutting down HTTP Server from STOPHANDLER.")
	}
}

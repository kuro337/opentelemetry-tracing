package app

import (
	"main/structuredlogger"
	"main/webserver"

	"go.opentelemetry.io/otel/trace"
)

type FibonacciApp struct {
	server *webserver.Server
	tracer trace.Tracer
	logger *structuredlogger.CustomLogger
	deps   *HandlerDependencies
}

type HandlerDependencies struct {
	tracer trace.Tracer
	logger *structuredlogger.CustomLogger
}

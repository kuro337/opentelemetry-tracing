package app

import (
	"context"
	"net/http"

	"main/structuredlogger"
	"main/webserver"

	"go.opentelemetry.io/otel/trace"
)

func initializeServer(tracer trace.Tracer, l *structuredlogger.CustomLogger) (context.Context, *webserver.Server, *HandlerDependencies) {
	initCtx, span := tracer.Start(context.Background(), "Server Initialization")
	defer span.End()

	server := webserver.NewServer(":8080").AddLogger(l)

	deps := &HandlerDependencies{tracer: tracer, logger: l}

	l.Logger.Debug("Created Web Server")
	span.AddEvent("Success Initializing Web Server")

	return initCtx, server, deps
}

func InitFibonacciServer(ctx context.Context, tracer trace.Tracer, l *structuredlogger.CustomLogger) {
	_, server, deps := initializeServer(tracer, l)

	app := &FibonacciApp{server: server, tracer: tracer, logger: l, deps: deps}

	activeCtx, span := tracer.Start(context.Background(), "HTTPServerActive")

	app.setupRoutes(activeCtx, tracer)
	span.End()

	server.Start()

	l.Logger.Debug("Started HTTP Server")

	// Listen for the server to stop
	err := <-server.ErrChan
	l.Logger.Debug("HTTP Server Stopped and Exit Signal Received")
	if err != nil && err != http.ErrServerClosed {
		l.Logger.Error("HTTP server stopped with error:", err)
	}
}

func (app *FibonacciApp) setupRoutes(ctx context.Context, tracer trace.Tracer) {
	_, span := tracer.Start(ctx, "HTTPServerActive")
	span.AddEvent("Success Adding Routes")
	defer span.End()

	app.server.AddRoute("/fibonacci/", TraceRequest(fibonacciHandler(app.deps), tracer))
	app.server.AddRoute("/stop", TraceRequest(app.stopHandler, tracer))
	app.logger.Logger.Debug("Added Routes")
}

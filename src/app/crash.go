package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"main/structuredlogger"
	"main/webserver"

	"go.opentelemetry.io/otel/trace"
)

func listenForError(l *structuredlogger.CustomLogger, tracer trace.Tracer, ctx context.Context, server *webserver.Server) {
	l.Logger.DebugContext(ctx, "HTTP Server Active - listening for Shutdown or Error for Graceful Termination.")

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	select {
	case <-sigCh:
		l.Print("\ngoodbye")

		_, signalSpan := tracer.Start(ctx, "OSInterruptReceived")
		defer signalSpan.End()

		shutdownAfterSignal(tracer, server)

	case err := <-server.ErrChan:
		if err != nil && err != http.ErrServerClosed {
			l.Fatal(err)
		}
	}
}

func shutdownAfterSignal(tracer trace.Tracer, server *webserver.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	shutdownCtx, shutdownSpan := tracer.Start(ctx, "SERVER_CRASH")
	defer shutdownSpan.End()

	server.GracefulShutdown(shutdownCtx)
}

package main

import (
	"context"
	"log/slog"
	"os"

	"main/app"
	"main/constants"
	"main/opentelemetry"
	"main/structuredlogger"

	"go.opentelemetry.io/otel"
)

func main() {
	debugLevel := slog.LevelDebug

	l := structuredlogger.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: debugLevel}))
	l.Logger.Info("App Init", "user", "default")

	cleanup, err := opentelemetry.SetupTelemetry(l)
	if err != nil {
		l.Fatal(err)
	}
	defer cleanup()

	l.Logger.Info("gRPC Tracing Exporter enabled")

	tracer := otel.Tracer(constants.NAME)

	mainCtx, mainSpan := tracer.Start(context.Background(), "BINARY_EXECUTED")
	mainSpan.End()

	app.InitFibonacciServer(mainCtx, tracer, l)

	l.Logger.Info("App Exited")
}

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
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func main() {
	debugLevel := slog.LevelDebug

	l := structuredlogger.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: debugLevel}))
	l.Logger.Info("App Init", "user", "default")

	cleanup, err := opentelemetry.SetupTelemetry(l)
	if err != nil {
		l.Fatal(err)
	}
	ctx := context.Background()
	defer cleanup(ctx)

	l.Logger.Info("gRPC Tracing Exporter enabled")

	tracer := otel.GetTracerProvider().Tracer(constants.NAME)

	ctx, span := tracer.Start(ctx, "BINARY_EXECUTED")
	defer func() {
		span.AddEvent("EXIT", trace.WithAttributes(attribute.String("BINARY_EXIT", "EXIT")))
		span.End()
	}()

	app.InitFibonacciServer(ctx, tracer, l)

	l.Logger.Info("App Exited")
}

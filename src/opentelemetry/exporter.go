package opentelemetry

import (
	"context"

	"main/structuredlogger"
)

func SetupTelemetry(l *structuredlogger.CustomLogger) (cleanup func(context.Context), err error) {
	l.Logger.Info("Creating gRPC OTLP Trace Exporter")

	shutdown, err := CreateOTLPTraceExporterGRPC(l) // Note: Assuming you have moved the function to the opentelemetry package. Otherwise, use utils.
	if err != nil {
		l.Fatal(err)
		return nil, err
	}

	// Return this cleanup function for caller to defer

	return func(ctx context.Context) {
		shutdown(ctx) // Execute the shutdown without a context
		l.Logger.Info("Tracer Shutdown complete")
	}, nil
}

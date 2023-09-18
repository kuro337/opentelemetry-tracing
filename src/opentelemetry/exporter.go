package opentelemetry

import (
	"context"
	"time"

	"main/structuredlogger"
)

func SetupTelemetry(l *structuredlogger.CustomLogger) (cleanup func(), err error) {
	l.Logger.Info("Creating gRPC OTLP Trace Exporter")

	shutdown, err := CreateOTLPTraceExporterGRPC(l) // Note: Assuming you have moved the function to the opentelemetry package. Otherwise, use utils.
	if err != nil {
		l.Fatal(err)
		return nil, err
	}

	// Return this cleanup function for caller to defer
	//	return shutdown, nil

	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		shutdown() // Execute the shutdown without a context

		if ctx.Err() == context.DeadlineExceeded {
			l.Logger.Warn("Telemetry cleanup was interrupted by timeout.")
		} else {
			l.Logger.Info("Telemetry cleanup completed gracefully.")
		}
	}, nil
}

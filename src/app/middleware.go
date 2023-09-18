package app

import (
	"net"
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func TraceRequest(next http.HandlerFunc, tracer trace.Tracer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqCtx, span := tracer.Start(r.Context(), "REQUEST")

		addRequestMetadataToSpan(span, r)
		r = r.WithContext(reqCtx)

		defer span.End()

		next(w, r)
	}
}

func addRequestMetadataToSpan(span trace.Span, r *http.Request) {
	span.SetAttributes(
		attribute.String("handler", "TraceRequest"),
		attribute.String("request.type", r.Method),
		attribute.String("request.url", r.URL.String()),
		attribute.String("request.user_agent", r.UserAgent()),
		attribute.String("request.referer", r.Referer()),
		attribute.String("request.protocol", r.Proto),
	)

	// Extracting IP address
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		span.SetAttributes(attribute.String("request.ip", ip))
	}

	if accept := r.Header.Get("Accept"); accept != "" {
		span.SetAttributes(attribute.String("request.header.accept", accept))
	}
}

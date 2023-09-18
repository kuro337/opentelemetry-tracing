package app

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func Fibonacci(ctx context.Context, tracer trace.Tracer, n uint) (uint64, error) {
	_, span := tracer.Start(ctx, "COMPUTE")
	defer span.End()

	if n <= 1 {
		span.SetAttributes(attribute.Int64("response.fib", int64(n)))
		return uint64(n), nil
	}

	var n2, n1 uint64 = 0, 1
	for i := uint(2); i < n; i++ {
		n2, n1 = n1, n1+n2
	}
	span.SetAttributes(attribute.Int64("response.fib", int64(n)))
	return n2 + n1, nil
}

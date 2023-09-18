package utils

import (
	"context"

	"go.opentelemetry.io/otel/baggage"
)

func AddBaggageToCtx(ctx context.Context) context.Context {
	m0, _ := baggage.NewMember(string("foo"), "foo1")
	m1, _ := baggage.NewMember(string("bar"), "bar1")
	b, _ := baggage.New(m0, m1)
	ctx = baggage.ContextWithBaggage(ctx, b)
	return ctx
}

/*

l.Logger.InfoContext(ctx, "Baggage")


*/

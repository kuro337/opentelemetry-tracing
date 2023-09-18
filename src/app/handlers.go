package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

func (app *FibonacciApp) fibonacciHandler(w http.ResponseWriter, r *http.Request) {
	_, span := app.deps.tracer.Start(r.Context(), "FIB_HANDLER")
	span.SetAttributes(attribute.String("handler", "fibonacciHandler"))
	defer span.End()

	if r.Method == http.MethodGet {
		span.SetAttributes(attribute.String("request.type", "GET"))

		// Retrieve Number from Query String Param
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 3 {
			span.RecordError(errors.New("Invalid Req"))
			span.SetStatus(codes.Error, "Invalid Req")
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// Convert to Integer
		num, err := strconv.Atoi(parts[2])
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		_, span2 := app.deps.tracer.Start(r.Context(), "ReqGetFib")
		defer span2.End()

		// Call the Fibonacci function with the input number.
		result, err := ComputeFibonacci(r.Context(), app.deps.tracer, uint(num))
		if err != nil {
			span2.RecordError(err)
			span2.SetStatus(codes.Error, err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Return the result as an HTTP response.
		fmt.Fprintf(w, "Fibonacci(%d) = %d\n", num, result)
		span.SetAttributes(attribute.String("success", "200"), attribute.Int64("response.fib", int64(result)))

	} else {
		span.RecordError(errors.New("Method Not Allowed"))
		span.SetStatus(codes.Error, "Method Not Allowed")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (app *FibonacciApp) stopHandler(w http.ResponseWriter, r *http.Request) {
	app.logger.Logger.Info("Req sent to Stop Route")
	w.Write([]byte("Shutting down..."))

	app.server.GracefulShutdown(context.Background())
}

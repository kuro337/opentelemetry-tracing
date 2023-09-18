package structuredlogger

import (
	"context"
	"log/slog"
	"os"
)

func TestSlogger() {
	slog.Info("Starting Fibonacci server...")

	// create logger explicitly
	//	logger := slog.Default()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Info("hello, world", slog.String("user", "default"))
	logger.Info("hello, world", slog.Int("user", 24))

	// {"time":"2023-09-16T13:49:39.1894232-04:00","level":"INFO","msg":"hello, world","user":"default"}

	// Adding an additional attribute to log in the output

	logger2 := logger.With("level2", "nested")
	logger2.Info("hello, world", "user", "default")
	logger2.InfoContext(context.Background(), "hello, world", "user", "default")
}

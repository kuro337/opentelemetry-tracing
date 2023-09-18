package structuredlogger

import (
	"fmt"
	"log/slog"
	"os"
)

type CustomLogger struct {
	Logger *slog.Logger
}

func New(h slog.Handler) *CustomLogger {
	return &CustomLogger{
		Logger: slog.New(h),
	}
}

func (l *CustomLogger) Print(v ...interface{}) {
	fmt.Println(v...)
}

func (l *CustomLogger) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

// Logs a fatal error and exits the program
func (l *CustomLogger) Fatal(err error) {
	l.Logger.Error(err.Error()) // Assuming slog.Logger has an Error method
	os.Exit(1)
}

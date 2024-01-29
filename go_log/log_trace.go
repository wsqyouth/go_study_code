package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Logger is a simple logger interface that we could plug any loggers like zap, logrus etc.
type Logger interface {
	InfoContextf(ctx context.Context, format string, args ...any)
	DebugContextf(ctx context.Context, format string, args ...any)
	WarnContextf(ctx context.Context, format string, args ...any)
	ErrorContextf(ctx context.Context, format string, args ...any)
}

// DefaultLogger implements Logger using builtin log package. It should be replaced with your own implementation.
type DefaultLogger struct{}

func (l *DefaultLogger) InfoContextf(ctx context.Context, format string, args ...any) {
	traceID := getTraceIDFromContext(ctx)
	msg := fmt.Sprintf("%s INFO [%s] - %s", time.Now().Format(time.RFC3339), traceID, fmt.Sprintf(format, args...))
	_ = l.writeToFile(msg)
}

func (l *DefaultLogger) DebugContextf(ctx context.Context, format string, args ...any) {
	traceID := getTraceIDFromContext(ctx)
	msg := fmt.Sprintf("%s DEBUG [%s] - %s", time.Now().Format(time.RFC3339), traceID, fmt.Sprintf(format, args...))
	_ = l.writeToFile(msg)
}

func (l *DefaultLogger) WarnContextf(ctx context.Context, format string, args ...any) {
	traceID := getTraceIDFromContext(ctx)
	msg := fmt.Sprintf("%s WARN [%s] - %s", time.Now().Format(time.RFC3339), traceID, fmt.Sprintf(format, args...))
	_ = l.writeToFile(msg)
}

func (l *DefaultLogger) ErrorContextf(ctx context.Context, format string, args ...any) {
	traceID := getTraceIDFromContext(ctx)
	msg := fmt.Sprintf("%s ERROR [%s] - %s", time.Now().Format(time.RFC3339), traceID, fmt.Sprintf(format, args...))
	_ = l.writeToFile(msg)
}

func (l *DefaultLogger) writeToFile(message string) error {
	fileName := "/tmp/" + strings.ReplaceAll(time.Now().Format(time.RFC3339), ":", "_") + ".txt"
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return errors.New("failed to open file")
	}
	defer f.Close()

	if _, err := f.WriteString(message + "\n"); err != nil {
		log.Println(err)
		return errors.New("failed to write to file")
	}
	return nil
}

// getTraceIDFromContext extracts the trace_id from the context.
func getTraceIDFromContext(ctx context.Context) string {
	traceID, ok := ctx.Value("trace_id").(string)
	if !ok {
		return "unknown"
	}
	return traceID
}

func main() {
	logger := &DefaultLogger{}
	ctx := context.WithValue(context.Background(), "trace_id", "1234567890")

	logger.InfoContextf(ctx, "This is an info message with value: %d", 42)
	logger.WarnContextf(ctx, "This is a warning message with value: %s", "warning")
	logger.DebugContextf(ctx, "This is a debug message with value: %f", 3.14)
	logger.ErrorContextf(ctx, "This is an error message with value: %v", errors.New("example error"))
}

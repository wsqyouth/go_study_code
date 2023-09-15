package main

import (
	"context"
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	loggerKey string = "logger"
)

type Logger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
}

type ZapLogger struct {
	*zap.Logger
}

func (l *ZapLogger) Debug(args ...interface{}) {
	l.Sugar().Debug(args...)
}

func (l *ZapLogger) Debugf(template string, args ...interface{}) {
	l.Sugar().Debugf(template, args...)
}

func DebugContext(ctx context.Context, args ...interface{}) {
	logger, ok := ctx.Value(loggerKey).(Logger)
	if !ok {
		fmt.Println("not found logger")
		return
	}
	// msg := fmt.Sprintf()
	logger.Debug(args)
}

func DebugContextf(ctx context.Context, template string, args ...interface{}) {
	logger, ok := ctx.Value(loggerKey).(Logger)
	if !ok {
		fmt.Println("not found logger")
		return
	}
	// msg := fmt.Sprintf()
	logger.Debugf(template, args)
}

func main() {
	// 创建一个Logger
	encoderConfig := zap.NewProductionEncoderConfig()
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.Lock(os.Stdout), zapcore.DebugLevel)
	zapLogger := zap.New(core)
	defer zapLogger.Sync() // flushes buffer, if any
	// 设置日志logger到context里
	ctx := context.WithValue(context.Background(), loggerKey, &ZapLogger{zapLogger})
	DebugContext(ctx, "wsq")
	DebugContextf(ctx, "wsq: %v", 233)
}

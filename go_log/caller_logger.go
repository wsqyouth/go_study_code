package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Info struct {
	Message string
}

type key int

const (
	callerLoggerKey key = iota
	calledLoggerKey
)

type Logger interface {
	Info(args ...interface{})
}

type OsLogger struct {
	*log.Logger
}

func (l *OsLogger) Info(args ...interface{}) {
	l.Println(args...)
}

type ZapLogger struct {
	*zap.Logger
}

func (l *ZapLogger) Info(args ...interface{}) {
	l.Sugar().Infow("", "msg:", fmt.Sprint(args...))
}

// WithLogger 通过键区分调用者和被调用者的日志记录器
func WithLogger(ctx context.Context, key key, logger Logger) context.Context {
	return context.WithValue(ctx, key, logger)
}

func logPrint(ctx context.Context, info Info) {
	logger, ok := ctx.Value(callerLoggerKey).(Logger)
	if !ok {
		fmt.Println("not found logger")
		return
	}

	// Convert info to log fields
	fields := info.Message

	logger.Info(fields)
}

// 将包含两条日志记录，一条来自标准的log包，另一条来自zap包。
func main() {
	logger := &OsLogger{log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)}
	// 主调日志通过主键key设置，使用系统日志记录
	ctx := WithLogger(context.Background(), callerLoggerKey, logger)
	info := Info{Message: "This is a test message."}
	logPrint(ctx, info)

	// 主调日志通过主键key设置，使用zap日志记录
	encoderConfig := zap.NewProductionEncoderConfig()
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.Lock(os.Stdout), zapcore.InfoLevel)
	zapLogger := zap.New(core)
	defer zapLogger.Sync() // flushes buffer, if any
	zapLoggerCtx := WithLogger(context.Background(), callerLoggerKey, &ZapLogger{zapLogger})
	logPrint(zapLoggerCtx, info)
}

/*
总结：
这种将要记录的场景(主调,被调等)作为key放到ctx里，之后注入想要对应的日志记录器类型(系统日志,zap日志),非常灵活
从而实现了可插拔、可配置

todo: 在系统启动时通过加载conf配置，实现日志类型的注入
*/

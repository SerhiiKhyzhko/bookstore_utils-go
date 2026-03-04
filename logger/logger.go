package logger

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)
type Logger struct {
	logger *zap.Logger
}

type Config struct {
    Level       string   
    OutputPaths []string // e.g. []string{"stdout"}
}

func getLogLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "panic":
		return zap.PanicLevel
	default:
		return zap.InfoLevel
	}
}

func NewLogger(cfg Config) (*Logger, error) {
	var logger *zap.Logger
	var err error
	
	logConfig := zap.Config{
		OutputPaths: cfg.OutputPaths,
		Level: zap.NewAtomicLevelAt(getLogLevel(cfg.Level)),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey: "level",
			TimeKey: "time",
			MessageKey: "msg",
			EncodeTime: zapcore.ISO8601TimeEncoder,
			EncodeLevel: zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	if logger, err = logConfig.Build(); err != nil {
		return nil, err
	}
	return &Logger{logger: logger}, nil
}

func (l *Logger) Info(msg string, tags ...zap.Field) {
	l.logger.Info(msg, tags...)
	l.logger.Sync()
}

func (l *Logger) Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))

	l.logger.Error(msg, tags...)
	l.logger.Sync()
}

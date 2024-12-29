package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// Logger 日志接口
type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	WithField(key string, value interface{}) Logger
	WithFields(fields map[string]interface{}) Logger
}

// LogrusLogger logrus实现
type LogrusLogger struct {
	logger *logrus.Logger
	fields logrus.Fields
}

// NewLogger 创建新的日志记录器
func NewLogger() Logger {
	logger := logrus.New()

	// 设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	// 设置输出
	logger.SetOutput(os.Stdout)

	// 设置日志级别
	level := os.Getenv("LOG_LEVEL")
	if level == "" {
		level = "info"
	}
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	logger.SetLevel(logLevel)

	return &LogrusLogger{
		logger: logger,
		fields: make(logrus.Fields),
	}
}

func (l *LogrusLogger) Debug(args ...interface{}) {
	l.logger.WithFields(l.fields).Debug(args...)
}

func (l *LogrusLogger) Debugf(format string, args ...interface{}) {
	l.logger.WithFields(l.fields).Debugf(format, args...)
}

func (l *LogrusLogger) Info(args ...interface{}) {
	l.logger.WithFields(l.fields).Info(args...)
}

func (l *LogrusLogger) Infof(format string, args ...interface{}) {
	l.logger.WithFields(l.fields).Infof(format, args...)
}

func (l *LogrusLogger) Warn(args ...interface{}) {
	l.logger.WithFields(l.fields).Warn(args...)
}

func (l *LogrusLogger) Warnf(format string, args ...interface{}) {
	l.logger.WithFields(l.fields).Warnf(format, args...)
}

func (l *LogrusLogger) Error(args ...interface{}) {
	l.logger.WithFields(l.fields).Error(args...)
}

func (l *LogrusLogger) Errorf(format string, args ...interface{}) {
	l.logger.WithFields(l.fields).Errorf(format, args...)
}

func (l *LogrusLogger) Fatal(args ...interface{}) {
	l.logger.WithFields(l.fields).Fatal(args...)
}

func (l *LogrusLogger) Fatalf(format string, args ...interface{}) {
	l.logger.WithFields(l.fields).Fatalf(format, args...)
}

func (l *LogrusLogger) WithField(key string, value interface{}) Logger {
	newFields := make(logrus.Fields)
	for k, v := range l.fields {
		newFields[k] = v
	}
	newFields[key] = value
	return &LogrusLogger{
		logger: l.logger,
		fields: newFields,
	}
}

func (l *LogrusLogger) WithFields(fields map[string]interface{}) Logger {
	newFields := make(logrus.Fields)
	for k, v := range l.fields {
		newFields[k] = v
	}
	for k, v := range fields {
		newFields[k] = v
	}
	return &LogrusLogger{
		logger: l.logger,
		fields: newFields,
	}
}

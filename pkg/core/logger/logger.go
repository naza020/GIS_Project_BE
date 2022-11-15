package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})

	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	GetLogger() *zap.Logger
}

// logger @impl logger.Logger
type logger struct {
	ZapLogger *zap.Logger
}

// NewLogger ...
// https://nyogjtrc.github.io/posts/2019/09/log-rotate-with-zap-logger/
func NewLogger(filePath string) Logger {
	var ioWriter = &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    10, // MB
		MaxBackups: 3,  // number of backups
		MaxAge:     30, //days
		LocalTime:  true,
		Compress:   false, // disabled by default
	}

	writeFile := zapcore.AddSync(ioWriter)
	writeStdout := zapcore.AddSync(os.Stdout)

	//encoderConfig := zap.NewDevelopmentEncoderConfig()
	//encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writeFile, writeStdout),
		zap.DebugLevel,
	)
	return &logger{ZapLogger: zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))}
}

func (l *logger) Debug(args ...interface{}) {
	l.ZapLogger.Sugar().Debug(args)
}

func (l *logger) Info(args ...interface{}) {
	l.ZapLogger.Sugar().Info(args)
}

func (l *logger) Warn(args ...interface{}) {
	l.ZapLogger.Sugar().Warn(args)
}

func (l *logger) Error(args ...interface{}) {
	l.ZapLogger.Sugar().Error(args)
}

func (l *logger) GetLogger() *zap.Logger {
	return l.ZapLogger
}

func (l *logger) Debugf(template string, args ...interface{}) {
	l.ZapLogger.Sugar().Debugf(template, args)
}

func (l *logger) Infof(template string, args ...interface{}) {
	l.ZapLogger.Sugar().Infof(template, args)
}

func (l *logger) Warnf(template string, args ...interface{}) {
	l.ZapLogger.Sugar().Warnf(template, args)
}

func (l *logger) Errorf(template string, args ...interface{}) {
	l.ZapLogger.Sugar().Errorf(template, args)
}

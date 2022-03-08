package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type Logger struct {
	level     Level
	zapLogger *zap.Logger
}

var logger = NewLogger(DefaultOption())

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Debugf(msg string, args ...interface{}) {
	logger.Debugf(msg, args...)
}

func Info(msg string, fields ...Field) {
	logger.Info(msg, fields...)
}

func Infof(msg string, args ...interface{}) {
	logger.Infof(msg, args...)
}

func Warn(msg string, fields ...Field) {
	logger.Warn(msg, fields...)
}

func Warnf(msg string, args ...interface{}) {
	logger.Warnf(msg, args...)
}

func Error(msg string, fields ...Field) {
	logger.Error(msg, fields...)
}

func Errorf(msg string, args ...interface{}) {
	logger.Errorf(msg, args...)
}

func Errorw(msg string, args ...interface{}) {
	logger.Errorw(msg, args...)
}

func Fatal(msg string, fields ...Field) {
	logger.Fatal(msg, fields...)
}

func Fatalf(msg string, args ...interface{}) {
	logger.Fatalf(msg, args...)
}

func Fatalw(msg string, args ...interface{}) {
	logger.Fatalw(msg, args...)
}

func Panic(msg string, fields ...Field) {
	logger.Panic(msg, fields...)
}

func Panicf(msg string, args ...interface{}) {
	logger.Panicf(msg, args...)
}

func Panicw(msg string, args ...interface{}) {
	logger.Panicw(msg, args...)
}

func Flush() {
	logger.Flush()
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func milliSecondsDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendFloat64(float64(d) / float64(time.Millisecond))
}

func Init(opt *Option) {
	logger = NewLogger(opt)
}

func NewLogger(opt *Option) *Logger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "zapLogger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: milliSecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	// when output to local path, with color is forbidden
	if opt.EnableColor {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opt.Level)); err != nil {
		zapLevel = InfoLevel
	}

	loggerConfig := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Development:       false,
		DisableCaller:     !opt.EnableCaller,
		DisableStacktrace: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         opt.Format,
		EncoderConfig:    encoderConfig,
		OutputPaths:      opt.OutputPaths,
		ErrorOutputPaths: opt.ErrorOutputPaths,
	}

	var err error
	zl, err := loggerConfig.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	logger := &Logger{
		level:     zapLevel,
		zapLogger: zl,
	}
	zap.RedirectStdLog(zl)
	return logger
}

func (l *Logger) Flush() {
	_ = l.zapLogger.Sync()
}

// Debug logs a message at DebugLevel. The message includes any fields passed
func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.zapLogger.Debug(msg, fields...)
}

// Debugf logs a message at DebugLevel.
func (l *Logger) Debugf(msg string, args ...interface{}) {
	l.zapLogger.Sugar().Debugf(msg, args...)
}

// Info logs a message at level Info on the standard logger.
func (l *Logger) Info(msg string, fields ...Field) {
	fmt.Println(l)
	l.zapLogger.Info(msg, fields...)
}

// Infof logs a message at level Info on the standard logger.
func (l *Logger) Infof(msg string, args ...interface{}) {
	l.zapLogger.Sugar().Infof(msg, args...)
}

// Warn logs a message at level Warn on the standard logger.
func (l *Logger) Warn(msg string, fields ...Field) {
	l.zapLogger.Warn(msg, fields...)
}

// Warnf logs a message at level Warn on the standard logger.
func (l *Logger) Warnf(msg string, args ...interface{}) {
	l.zapLogger.Sugar().Warn(msg, args)
}

// Error logs a message at level Error on the standard logger.
func (l *Logger) Error(msg string, fields ...Field) {
	l.zapLogger.Error(msg, fields...)
}

// Errorf logs a message at level Error on the standard logger.
func (l *Logger) Errorf(msg string, args ...interface{}) {
	l.zapLogger.Sugar().Errorf(msg, args)
}

// Errorw method output error level log.
func (l *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Errorw(msg, keysAndValues...)
}

// Panic method output panic level log and shutdown application.
func (l *Logger) Panic(msg string, fields ...Field) {
	l.zapLogger.Panic(msg, fields...)
}

// Panicf method output panic level log and shutdown application.
func (l *Logger) Panicf(format string, v ...interface{}) {
	l.zapLogger.Sugar().Panicf(format, v...)
}

// Panicw method output panic level log.
func (l *Logger) Panicw(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Panicw(msg, keysAndValues...)
}

// Fatal method output fatal level log.
func (l *Logger) Fatal(msg string, fields ...Field) {
	l.zapLogger.Fatal(msg, fields...)
}

// Fatalf method output fatal level log.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.zapLogger.Sugar().Fatalf(format, v...)
}

// Fatalw method output Fatalw level log.
func (l *Logger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Fatalw(msg, keysAndValues...)
}

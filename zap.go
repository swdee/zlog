package zlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Field wraps zapcore Field
type Field = zapcore.Field

// LogMode defines the logger operation mode, eg: development, production, disabled
type LogMode string

// LogEncoding defines the logger output encoding of json or console/text
type LogEncoding string

// DebugVerbosity allows verbosity levels to be set to Debug() calls where 0 is
// the lowest (quietest) level and 9 is the nosiest level
type DebugVerbosity int

const (
	// define the logger operation mode which effects logging output
	ModeDevel      LogMode = "development"
	ModeProduction LogMode = "production"
	ModeDisabled   LogMode = "disabled"

	// define the logger encoding output method
	EncodingJson    LogEncoding = "json"
	EncodingConsole LogEncoding = "console"
)

// Logger wraps the zap logger
type Logger struct {
	verbosity DebugVerbosity
	zap       *zap.Logger
}

// initialise a disabled singleton logger
var logger = New(ModeDisabled, EncodingConsole, 0)

// SetMode sets the operation mode of the singleton logger
func SetMode(mode LogMode, enc LogEncoding, v DebugVerbosity) {
	logger = New(mode, enc, v)
}

// New creates a new logger instance of a given operation mode
func New(mode LogMode, enc LogEncoding, v DebugVerbosity) *Logger {

	var instance *zap.Logger

	// set JSON encoding
	encConf := zap.NewProductionEncoderConfig()

	if enc == EncodingConsole {
		encConf = zap.NewDevelopmentEncoderConfig()
	}

	// create new zap Logger instance based on log mode
	switch mode {

	case ModeDevel:
		// AddCallerSkip(1) is added since we have wrapped the zap logger with
		// this package so the caller trace shows the correct code entry point
		// the logging was called at.

		// Copies zap.NewDevelopment() functionality but allows us to modify
		// the Config.Encoding parameter
		zcfg := zap.Config{
			Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
			Development:      true,
			Encoding:         string(enc),
			EncoderConfig:    encConf,
			OutputPaths:      []string{"stderr"},
			ErrorOutputPaths: []string{"stderr"},
		}

		instance, _ = zcfg.Build(zap.AddCallerSkip(1))

	case ModeProduction:
		// Copies zap.NewProduction()
		zcfg := zap.Config{
			Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
			Development: false,
			Sampling: &zap.SamplingConfig{
				Initial:    100,
				Thereafter: 100,
			},
			Encoding:         string(enc),
			EncoderConfig:    encConf,
			OutputPaths:      []string{"stderr"},
			ErrorOutputPaths: []string{"stderr"},
		}

		instance, _ = zcfg.Build(zap.AddCallerSkip(1))

	case ModeDisabled:
		instance = zap.NewNop()
	}

	defer instance.Sync()

	return &Logger{
		verbosity: v,
		zap:       instance,
	}
}

// NewNoop returns a logger that does nothing
func NewNoop() *Logger {
	return New(ModeDisabled, EncodingConsole, 0)
}

// Wrappers for singleton

// Info logs Info level entries to the singleton instance of the logger
func Info(msg string, f ...Field) {
	logger.zap.Info(msg, f...)
}

func Debug(v DebugVerbosity, msg string, f ...Field) {
	if logger.verbosity >= v {
		logger.zap.Debug(msg, f...)
	}
}

func Warn(msg string, f ...Field) {
	logger.zap.Warn(msg, f...)
}

func Error(msg string, f ...Field) {
	logger.zap.Error(msg, f...)
}

func Fatal(msg string, f ...Field) {
	logger.zap.Fatal(msg, f...)
}

// Wrappers for Logger object instances

// Info logs Info level entries to the given logger instance
func (l *Logger) Info(msg string, f ...Field) {
	l.zap.Info(msg, f...)
}

func (l *Logger) Debug(v DebugVerbosity, msg string, f ...Field) {
	if l.verbosity >= v {
		l.zap.Debug(msg, f...)
	}
}

func (l *Logger) Warn(msg string, f ...Field) {
	l.zap.Warn(msg, f...)
}

func (l *Logger) Error(msg string, f ...Field) {
	l.zap.Error(msg, f...)
}

func (l *Logger) Fatal(msg string, f ...Field) {
	l.zap.Fatal(msg, f...)
}

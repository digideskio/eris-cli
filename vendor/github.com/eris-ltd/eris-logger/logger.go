package logger

import (
	"io"
	"os"
	"sync"
)

type Logger struct {
	// The logs are `io.Copy`'d to this in a mutex. It's common to set this to a
	// file, or leave it default which is `os.Stderr`. You can also set this to
	// something more adventorous, such as logging to Kafka.
	Out io.Writer
	// Hooks for the logger instance. These allow firing events based on logging
	// levels and log entries. For example, to send errors to an error tracking
	// service, log to StatsD or dump the core on fatal errors.
	Hooks LevelHooks
	// All log entries pass through the formatter before logged to Out.
	// The default formmatter is the ErisFormatter. You can easily implement your
	// own that implements the `Formatter` interface, see the `README` or included
	// formatters for examples.
	Formatter Formatter
	// The logging level the logger should log at. This is typically (and defaults
	// to) `logger.Info`, which allows Info(), Warn(), Error() and Fatal() to be
	// logged. `logger.Debug` is useful in
	Level Level
	// Used to sync writing to the log.
	mu sync.Mutex
}

// Creates a new logger. Configuration should be set by changing `Formatter`,
// `Out` and `Hooks` directly on the default logger instance. You can also just
// instantiate your own:
//
//    var log = &Logger{
//      Out: os.Stderr,
//      Formatter: new(JSONFormatter),
//      Hooks: make(LevelHooks),
//      Level: logger.DebugLevel,
//    }
//
// It's recommended to make this a global instance called `log`.
func New() *Logger {
	return &Logger{
		Out:       os.Stderr,
		Formatter: &ErisFormatter{Color: true},
		Hooks:     make(LevelHooks),
		Level:     InfoLevel,
	}
}

// Adds a field to the log entry, note that you it doesn't log until you call
// Debug, Print, Info, Warn, Fatal or Panic. It only creates a log entry.
// If you want multiple fields, use `WithFields`.
func (logger *Logger) WithField(key string, value interface{}) *Entry {
	return NewEntry(logger).WithField(key, value)
}

// Adds a struct of fields to the log entry. All it does is call `WithField` for
// each `Field`.
func (logger *Logger) WithFields(fields Fields) *Entry {
	return NewEntry(logger).WithFields(fields)
}

// Add an error as single field to the log entry.  All it does is call
// `WithError` for the given `error`.
func (logger *Logger) WithError(err error) *Entry {
	return NewEntry(logger).WithError(err)
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	NewEntry(logger).Debugf(format, args...)
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	NewEntry(logger).Infof(format, args...)
}

func (logger *Logger) Printf(format string, args ...interface{}) {
	NewEntry(logger).Printf(format, args...)
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	NewEntry(logger).Warnf(format, args...)
}

func (logger *Logger) Warningf(format string, args ...interface{}) {
	NewEntry(logger).Warnf(format, args...)
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	NewEntry(logger).Errorf(format, args...)
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	NewEntry(logger).Fatalf(format, args...)
	os.Exit(1)
}

func (logger *Logger) Panicf(format string, args ...interface{}) {
	NewEntry(logger).Panicf(format, args...)
}

func (logger *Logger) Debug(args ...interface{}) {
	NewEntry(logger).Debug(args...)
}

func (logger *Logger) Info(args ...interface{}) {
	NewEntry(logger).Info(args...)
}

func (logger *Logger) Print(args ...interface{}) {
	NewEntry(logger).Info(args...)
}

func (logger *Logger) Warn(args ...interface{}) {
	NewEntry(logger).Warn(args...)
}

func (logger *Logger) Warning(args ...interface{}) {
	NewEntry(logger).Warn(args...)
}

func (logger *Logger) Error(args ...interface{}) {
	NewEntry(logger).Error(args...)
}

func (logger *Logger) Fatal(args ...interface{}) {
	NewEntry(logger).Fatal(args...)
	os.Exit(1)
}

func (logger *Logger) Panic(args ...interface{}) {
	NewEntry(logger).Panic(args...)
}

func (logger *Logger) Debugln(args ...interface{}) {
	NewEntry(logger).Debugln(args...)
}

func (logger *Logger) Infoln(args ...interface{}) {
	NewEntry(logger).Infoln(args...)
}

func (logger *Logger) Println(args ...interface{}) {
	NewEntry(logger).Println(args...)
}

func (logger *Logger) Warnln(args ...interface{}) {
	NewEntry(logger).Warnln(args...)
}

func (logger *Logger) Warningln(args ...interface{}) {
	NewEntry(logger).Warnln(args...)
}

func (logger *Logger) Errorln(args ...interface{}) {
	NewEntry(logger).Errorln(args...)
}

func (logger *Logger) Fatalln(args ...interface{}) {
	NewEntry(logger).Fatalln(args...)
	os.Exit(1)
}

func (logger *Logger) Panicln(args ...interface{}) {
	NewEntry(logger).Panicln(args...)
}

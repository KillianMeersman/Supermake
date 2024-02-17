package log

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Logger struct {
	name      string
	fields    map[string]string
	stdout    io.Writer
	stderr    io.Writer
	Level     LogLevel
	Formatter LogFormatter
}

func GetLogLevel() LogLevel {
	level, exists := os.LookupEnv("LOG_LEVEL")
	if !exists {
		return INFO
	}

	switch strings.ToLower(level) {
	case "trace":
		return TRACE
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "warn", "warning":
		return WARNING
	case "error":
		return ERROR
	case "fatal":
		return FATAL
	default:
		DefaultLogger().Error(fmt.Sprintf("invalid LOG_LEVEL environment variable: '%s'. Defaulting to INFO.", level))
		return INFO
	}
}

var defaultLogger = NewLogger("", INFO, ShellColoredLevels, os.Stdout, os.Stderr)

func init() {
	defaultLogger.Level = GetLogLevel()
}

func DefaultLogger() *Logger {
	return defaultLogger
}

func NamedLogger(name string) *Logger {
	return &Logger{
		name:      name,
		fields:    make(map[string]string),
		Level:     GetLogLevel(),
		Formatter: ShellColoredLevels,
		stdout:    os.Stdout,
		stderr:    os.Stderr,
	}
}

func NewLogger(name string, level LogLevel, formatter LogFormatter, stdout, stderr io.Writer) *Logger {
	return &Logger{
		name:      name,
		fields:    make(map[string]string),
		Level:     level,
		Formatter: formatter,
		stdout:    stdout,
		stderr:    stderr,
	}
}

func (l *Logger) log(level LogLevel, msg string, out io.Writer, fields ...string) {
	if len(fields)%2 != 0 {
		panic("odd number of fields provided")
	}

	if msg == "" {
		return
	}

	if level < l.Level {
		return
	}

	combinedFields := make(map[string]string)

	for k, v := range l.fields {
		combinedFields[k] = v
	}

	for i := 1; i < len(fields); i += 2 {
		combinedFields[fields[i-1]] = fields[i]
	}

	formattedMessage := l.Formatter(level, msg, combinedFields)
	out.Write(formattedMessage)
}

func (l *Logger) Trace(msg string, fields ...string) {
	l.log(TRACE, msg, l.stdout, fields...)
}

func (l *Logger) Debug(msg string, fields ...string) {
	l.log(DEBUG, msg, l.stdout, fields...)
}

func (l *Logger) Info(msg string, fields ...string) {
	l.log(INFO, msg, l.stdout, fields...)
}

func (l *Logger) Warning(msg string, fields ...string) {
	l.log(WARNING, msg, l.stdout, fields...)
}

func (l *Logger) Error(msg string, fields ...string) {
	l.log(ERROR, msg, l.stdout, fields...)
}

func (l *Logger) Fatal(msg string, fields ...string) {
	l.log(FATAL, msg, l.stdout, fields...)
	os.Exit(1)
}

func (l *Logger) Panic(msg string, fields ...string) {
	l.log(FATAL, msg, l.stdout, fields...)
	panic(msg)
}

func (l *Logger) With(fields ...string) *Logger {
	if len(fields)%2 != 0 {
		panic("odd number of fields provided")
	}

	combinedFields := make(map[string]string)

	for k, v := range l.fields {
		combinedFields[k] = v
	}

	for i := 1; i < len(fields); i += 2 {
		combinedFields[fields[i-1]] = fields[i]
	}

	return &Logger{
		fields:    combinedFields,
		Level:     l.Level,
		Formatter: l.Formatter,
		stdout:    l.stdout,
		stderr:    l.stderr,
	}
}

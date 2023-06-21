package log

import (
	"fmt"
	"io"
	"strings"
)

func ShellColoredLevels(level LogLevel, message string, fields map[string]string) []byte {
	fieldString := strings.Builder{}

	for k, v := range fields {
		fieldString.WriteString(fmt.Sprintf("%s=%s", k, v))
	}

	levelStr := "???"
	levelColor := Reset
	switch level {
	case DEBUG:
		levelStr = "DEBUG"
		levelColor = White
	case INFO:
		levelStr = "INFO"
		levelColor = Blue
	case WARNING:
		levelStr = "WARNING"
		levelColor = Yellow
	case ERROR:
		levelStr = "ERROR"
		levelColor = Red
	case FATAL:
		levelStr = "FATAL"
		levelColor = Purple
	}

	spaces := strings.Repeat(" ", 7-len(levelStr))
	return []byte(fmt.Sprintf("%s%s%s|%s %s %s\n", levelColor, levelStr, spaces, Reset, message, fieldString.String()))
}

type Logger struct {
	fields    map[string]string
	output    io.Writer
	Level     LogLevel
	Formatter LogFormatter
}

func NewLogger(level LogLevel, formatter LogFormatter, output io.Writer) *Logger {
	return &Logger{
		fields:    make(map[string]string),
		Level:     level,
		Formatter: formatter,
		output:    output,
	}
}

func (l *Logger) log(level LogLevel, msg string, fields ...string) {
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
	l.output.Write(formattedMessage)
}

func (l *Logger) Debug(msg string, fields ...string) {
	l.log(DEBUG, msg, fields...)
}

func (l *Logger) Info(msg string, fields ...string) {
	l.log(INFO, msg, fields...)
}

func (l *Logger) Warning(msg string, fields ...string) {
	l.log(WARNING, msg, fields...)
}

func (l *Logger) Error(msg string, fields ...string) {
	l.log(ERROR, msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...string) {
	l.log(FATAL, msg, fields...)
	panic(msg)
}

func (l *Logger) With(fields ...string) *Logger {
	newFields := make(map[string]string)

	for k, v := range l.fields {
		newFields[k] = v
	}

	return &Logger{
		fields:    newFields,
		Level:     l.Level,
		Formatter: l.Formatter,
		output:    l.output,
	}
}

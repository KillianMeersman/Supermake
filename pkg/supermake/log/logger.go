package log

import (
	"fmt"
	"io"
	"strings"
)

func ShellColoredLevels(level LogLevel, message string, fields map[string]string) []byte {
	target := "root"

	fieldStrings := make([]string, 0, len(fields))
	for k, v := range fields {
		if k == "target" {
			target = v
		} else {
			fieldStrings = append(fieldStrings, fmt.Sprintf("%s=%s", k, v))
		}
	}

	fieldString := strings.Join(fieldStrings, ", ")
	if len(fieldString) > 0 {
		fieldString = "[" + fieldString + "]"
	}

	levelStr := "???"
	levelFormatting := []ShellFormatting{Reset}
	switch level {
	case DEBUG:
		levelStr = "DEBUG"
		levelFormatting = []ShellFormatting{FgWhite, Faint}
	case INFO:
		levelStr = "INFO"
		levelFormatting = []ShellFormatting{FgBlue}
	case WARNING:
		levelStr = "WARNING"
		levelFormatting = []ShellFormatting{FgYellow, Bold}
	case ERROR:
		levelStr = "ERROR"
		levelFormatting = []ShellFormatting{FgHiRed, Bold}
	case FATAL:
		levelStr = "FATAL"
		levelFormatting = []ShellFormatting{FgHiMagenta, Bold}
	}

	return []byte(TabulateRow().Field(levelStr, 9, levelFormatting...).Field(target, 15).Field("|", 1).Field(message, -1).Field(fieldString, -1).String())
}

type Logger struct {
	fields    map[string]string
	stdout    io.Writer
	stderr    io.Writer
	Level     LogLevel
	Formatter LogFormatter
}

func NewLogger(level LogLevel, formatter LogFormatter, stdout, stderr io.Writer) *Logger {
	return &Logger{
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
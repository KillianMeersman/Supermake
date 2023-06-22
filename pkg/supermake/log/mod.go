package log

type LogLevel byte

const (
	TRACE   LogLevel = 1
	DEBUG   LogLevel = 2
	INFO    LogLevel = 3
	WARNING LogLevel = 4
	ERROR   LogLevel = 5
	FATAL   LogLevel = 6
)

type LogFormatter = func(level LogLevel, message string, fields map[string]string) []byte

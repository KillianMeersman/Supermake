package log

type LogLevel byte

const (
	DEBUG   LogLevel = 1
	INFO    LogLevel = 2
	WARNING LogLevel = 3
	ERROR   LogLevel = 4
	FATAL   LogLevel = 5
)

type LogFormatter = func(level LogLevel, message string, fields map[string]string) []byte

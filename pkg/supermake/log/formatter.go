package log

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ShellColoredLevels(level LogLevel, message string, fields map[string]string) []byte {
	target := "root"
	container := "@local"

	fieldStrings := make([]string, 0, len(fields))
	for k, v := range fields {
		switch k {
		case "target":
			target = v
		case "container":
			container = fmt.Sprintf("@%s", v)
		default:
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
	case TRACE:
		levelStr = "TRACE"
		levelFormatting = []ShellFormatting{FgWhite, Faint}
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

	return []byte(TabulateRow().Field(levelStr, 5, levelFormatting...).Field(target, 20).Field(container, len(container)).Field("|", 1).Field(message, -1).Field(fieldString, -1).String())
}

type jsonMessage struct {
	Message string            `json:"message"`
	Fields  map[string]string `json:"fields"`
}

func JSON(level LogLevel, message string, fields map[string]string) []byte {
	messageStruct := jsonMessage{
		message,
		fields,
	}
	data, err := json.Marshal(messageStruct)
	if err != nil {
		panic(err)
	}

	return data
}

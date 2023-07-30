package log

import (
	"bufio"
	"io"
)

// Stream every newline as a new INFO-level entry.
func StreamReaderNewLines(logger *Logger, reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		m := scanner.Text()
		logger.Info(m)
	}
}

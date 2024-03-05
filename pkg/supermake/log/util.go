package log

import (
	"bufio"
	"io"
)

// Stream every newline as a new INFO-level entry.
func StreamReaderNewLines(logger func(message string, fields ...string), reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		m := scanner.Text()
		logger(m)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

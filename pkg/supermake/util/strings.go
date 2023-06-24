package util

import (
	"strings"

	"github.com/KillianMeersman/Supermake/pkg/supermake/datastructures"
)

func IndexAll(s, substr string) []int {
	indices := make([]int, 0)

	for i := 0; i <= len(s)-len(substr); i++ {
		subStr := s[i : i+len(substr)]
		if subStr == substr {
			indices = append(indices, i)
		}
	}

	return indices
}

// Split words, using space as default delimiter.
// Treats quoted text as one word and keeps the quotation marks in.
func SplitWords(data string) []string {
	currentWord := &strings.Builder{}
	words := make([]string, 0)
	delimiterStack := datastructures.NewStack[rune]()

	for _, char := range data {
		delimiter := delimiterStack.Peek(' ')
		if char == delimiter {
			if delimiter != ' ' {
				currentWord.WriteRune(char)
			}
			words = append(words, currentWord.String())
			currentWord.Reset()
			delimiterStack.Pop(' ')
		} else if char == '"' || char == '\'' {
			currentWord.WriteRune(char)
			delimiterStack.Push(char)
		} else {
			currentWord.WriteRune(char)
		}
	}

	if currentWord.Len() > 0 {
		words = append(words, currentWord.String())
	}

	return words
}

// Split a string into lines.
// Respects escaped newlines, combining both lines into one.
func SplitLines(data string) []string {
	lines := strings.Split(data, "\n")
	escapedLines := make([]string, 0, len(lines))
	escapedLineBuilder := strings.Builder{}

	for i := 0; i < len(lines); i++ {
		line := strings.TrimRight(lines[i], " ")
		escapedLineBuilder.WriteString(line)

		if !strings.HasSuffix(line, `\`) {
			escapedLines = append(escapedLines, escapedLineBuilder.String())
			escapedLineBuilder.Reset()
		}
	}
	return escapedLines
}

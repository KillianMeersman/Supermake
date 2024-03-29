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
func SplitEscapedLines(data string) []string {
	lines := make([]string, 0)
	escaped := false
	line := strings.Builder{}

	for _, char := range data {
		switch char {
		case '\\':
			escaped = true
			continue
		case ' ':
			line.WriteRune(char)
			continue
		case '\n':
			if !escaped {
				lines = append(lines, line.String())
				line.Reset()
			}
		default:
			line.WriteRune(char)
		}
		escaped = false
	}

	return lines
}

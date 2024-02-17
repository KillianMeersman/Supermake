package parse

import (
	"strings"

	"github.com/KillianMeersman/Supermake/pkg/supermake/datastructures"
)

const VARIABLE_START = "${{"
const VARIABLE_END = "}}"

// Replace every variable with the return value of f.
// Supports nested variables, such that variable references may be constructed out of variables themselves.
// E.g. if X=A, Y=B, A-B=test, $($(X)-$(Y)) becomes 'test'.
// Errors returned by f are propagated to the function return value.
func ReplaceVariables(s string, f func(v string) (string, error)) (string, error) {
	rootString := &strings.Builder{}
	variableDataStack := datastructures.NewStack[*strings.Builder]()

	// This code assumes the start character's length >= end character's length.
	for i := 0; i < len(s); i++ {
		if i+len(VARIABLE_START) <= len(s) && s[i:i+len(VARIABLE_START)] == VARIABLE_START {
			// variable start detected, push to stack
			variableDataStack.Push(&strings.Builder{})
			i += len(VARIABLE_START) - 1
		} else if i+len(VARIABLE_END) <= len(s) && s[i:i+len(VARIABLE_END)] == VARIABLE_END && variableDataStack.Count() > 0 {
			// if the end character is detected and there's variables defined, terminate the topmost variable
			key := variableDataStack.Pop(nil).String()
			key = strings.TrimSpace(key)
			value, err := f(key)
			if err != nil {
				return "", err
			}
			variableDataStack.Peek(rootString).WriteString(value)
			i += len(VARIABLE_END) - 1
		} else {
			// no variable start or end detected, just copy the character as-is
			variableDataStack.Peek(rootString).WriteByte(s[i])
		}
	}

	return rootString.String(), nil
}

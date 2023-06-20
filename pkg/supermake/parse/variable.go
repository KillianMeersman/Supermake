package parse

import (
	"strings"

	"github.com/KillianMeersman/Supermake/pkg/supermake/datastructures"
)

// Replace every variable with the return value of f.
// Supports nested variables, such that if X=A, Y=B, A-B=test, $(X-$(Y)) becomes 'test'.
// Errors returned by f are progatated to the function return value.
func ReplaceVariables(s string, f func(v string) (string, error)) (string, error) {
	rootString := &strings.Builder{}
	variableDataStack := datastructures.NewStack[*strings.Builder]()

	for i := 0; i < len(s); i++ {
		// prevent overflows
		substrEnd := i + 2
		if substrEnd > len(s) {
			substrEnd = len(s)
		}

		if s[i:substrEnd] == "$(" {
			// variable start detected, push to stack
			variableDataStack.Push(&strings.Builder{})
			i++ // skip next character as this will be '('
		} else if s[i] == ')' && variableDataStack.Count() > 0 {
			// if the end character is detected and there's variables defined, terminate the topmost variable
			key := variableDataStack.Pop(nil).String()
			value, err := f(key)
			if err != nil {
				return "", err
			}
			variableDataStack.Peek(rootString).WriteString(value)
		} else {
			// no variable start or end detected, just copy the character as-is
			variableDataStack.Peek(rootString).WriteByte(s[i])
		}
	}

	return rootString.String(), nil
}

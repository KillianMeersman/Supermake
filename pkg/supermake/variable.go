package supermake

import (
	"fmt"
	"os"
)

type Variables map[string]*Variable

func (v Variables) EnvStrings() []string {
	strings := make([]string, 0, len(v))
	for _, v := range v {
		strings = append(strings, v.EnvString())
	}

	return strings
}

type VariableEvaluationType int

const (
	RECURSIVE VariableEvaluationType = iota
	FALLBACK
)

type Variable struct {
	Name           string
	EvaluationType VariableEvaluationType
	Export         bool
	value          string
}

func NewVariable(name string, evaluationType VariableEvaluationType, export bool, value string) *Variable {
	return &Variable{
		Name:           name,
		EvaluationType: evaluationType,
		Export:         export,
		value:          value,
	}
}

func (v *Variable) Value() string {
	if v.EvaluationType == FALLBACK {
		if v, ok := os.LookupEnv(v.Name); ok {
			return v
		}
	}
	return v.value
}

func (v *Variable) EnvString() string {
	return fmt.Sprintf("%s=%s", v.Name, v.Value())
}

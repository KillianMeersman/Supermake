package supermake

import (
	"context"
	"fmt"
	"os"
)

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

type SupermakeFile struct {
	Targets   map[string]*Target
	Variables map[string]*Variable
}

func (s *SupermakeFile) Run(target string) error {
	t, ok := s.Targets[target]
	if !ok {
		return fmt.Errorf("no such target: %s", target)
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	return t.Run(context.TODO(), s, cwd)
}

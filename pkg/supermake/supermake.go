package supermake

import (
	"context"
	"fmt"
	"os"

	"github.com/KillianMeersman/Supermake/pkg/supermake/executors"
	"github.com/KillianMeersman/Supermake/pkg/supermake/log"
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
	Targets   map[string]executors.Runable
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

	logger := log.NewLogger(log.DEBUG, log.ShellColoredLevels, os.Stdout, os.Stderr)
	err = t.Run(context.TODO(), executors.ExecutorContext{
		map[string]string{},
		s.Targets,
		make(map[string]executors.Runable),
		cwd,
		logger,
	})

	if err != nil {
		logger.Fatal(err.Error())
	}

	return nil
}

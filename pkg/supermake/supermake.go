package supermake

import (
	"context"
	"fmt"
	"os"

	"github.com/KillianMeersman/Supermake/pkg/supermake/log"
)

type SupermakeFile struct {
	Targets   map[string]Runable
	Variables Variables
}

func (s *SupermakeFile) Run(ctx context.Context, target string) error {
	t, ok := s.Targets[target]
	if !ok {
		return fmt.Errorf("no such target: %s", target)
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	logger := log.NewLogger(log.TRACE, log.ShellColoredLevels, os.Stdout, os.Stderr)
	err = t.Run(ctx, ExecutorContext{
		EnvVars:       s.Variables,
		Targets:       s.Targets,
		ParentTargets: make(map[string]Runable),
		WorkingDir:    cwd,
		Logger:        logger,
	})

	if err != nil {
		logger.Fatal(err.Error())
	}

	return err
}

package supermake

import (
	"context"
	"fmt"
	"os"

	"github.com/KillianMeersman/Supermake/pkg/supermake/log"
)

type SupermakeFile struct {
	Targets   Targets
	Variables Variables
}

func (s *SupermakeFile) Reset() {
	for _, target := range s.Targets {
		target.Reset()
	}
}

func (s *SupermakeFile) Run(ctx context.Context, target, cwd string) error {
	t, ok := s.Targets[target]
	if !ok {
		return fmt.Errorf("no such target: %s", target)
	}

	logger := log.NewLogger(log.INFO, log.ShellColoredLevels, os.Stdout, os.Stderr)

	logger.Debug("VARIABLES:")
	for k, v := range t.Variables {
		logger.Debug(fmt.Sprintf("%s = %s", k, v.Value()))
	}

	scheduler := &LocalScheduler{}
	execCtx := ExecutorContext{
		EnvVars:       s.Variables,
		Targets:       s.Targets,
		ParentTargets: make(map[string]Runable),
		WorkingDir:    cwd,
		Logger:        logger,
		Scheduler:     scheduler,
	}
	err := scheduler.ScheduleTarget(ctx, execCtx, t)
	if err != nil {
		logger.Fatal(err.Error())
	}

	return err
}

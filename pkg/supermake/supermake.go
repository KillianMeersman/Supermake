package supermake

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/KillianMeersman/Supermake/pkg/supermake/datastructures"
	"github.com/KillianMeersman/Supermake/pkg/supermake/log"
	"github.com/KillianMeersman/Supermake/pkg/supermake/util"
)

type SupermakeFile struct {
	Targets   Targets
	Variables Variables
}

func (s *SupermakeFile) Run(ctx context.Context, scheduler Scheduler, cwd string, targets ...string) error {
	logger := log.NewLogger(log.INFO, log.ShellColoredLevels, os.Stdout, os.Stderr)
	execCtx := ExecutorContext{
		EnvVars:       s.Variables,
		Targets:       s.Targets,
		ParentTargets: make(Targets),
		WorkingDir:    cwd,
		Logger:        logger,
		Scheduler:     scheduler,
	}

	wg := sync.WaitGroup{}
	errs := make(chan error, len(targets))

	for _, target := range targets {
		t, ok := s.Targets[target]
		if !ok {
			return fmt.Errorf("no such target: %s", target)
		}

		logger.Debug("VARIABLES:")
		for k, v := range t.Variables {
			logger.Debug(fmt.Sprintf("%s = %s", k, v.Value()))
		}

		wg.Add(1)
		go func() {
			err := scheduler.ScheduleTarget(ctx, execCtx, t)
			if err != nil {
				errs <- err
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		errs <- nil
	}()

	return <-errs
}

func (s *SupermakeFile) Help() string {
	text := strings.Builder{}

	text.WriteString("Targets:\n")
	text.WriteString("========\n")

	for _, target := range s.Targets {
		text.WriteString(fmt.Sprintf("%s - No description\n", target.FQN()))
	}

	return text.String()
}

func (s *SupermakeFile) GetDoneFileTargets() []*Target {
	phonies := datastructures.NewUnorderedSet[string]()
	doneFileTargets := make([]*Target, 0)

	phony, ok := s.Targets[".PHONY"]
	if ok {
		for _, dep := range phony.Dependencies {
			phonies.Add(dep)
		}
	}

	for _, target := range s.Targets {
		if !phonies.Contains(target.Name()) && util.IsFileOrFolder(target.Name()) {
			doneFileTargets = append(doneFileTargets, target)
		}
	}

	return doneFileTargets
}

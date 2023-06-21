package supermake

import (
	"context"
	"fmt"
	"sync"

	"github.com/KillianMeersman/Supermake/pkg/supermake/executors"
)

type Runable interface {
	Run(ctx context.Context, execCtx executors.ExecutorContext, file *SupermakeFile) error
}

type Target struct {
	Name         string
	Node         string
	Dependencies []string
	Steps        []Runable
	Variables    map[string]*Variable
	SubTargets   map[string]*Target
	Parent       *Target
}

func (t *Target) runDependencies(ctx context.Context, execCtx executors.ExecutorContext, file *SupermakeFile) error {
	errChan := make(chan error)
	doneChan := make(chan struct{})
	wg := new(sync.WaitGroup)

	subTargetCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, subTargetName := range t.Dependencies {
		subTarget, ok := file.Targets[subTargetName]
		if !ok {
			return fmt.Errorf("unknown dependency target '%s'", subTargetName)
		}
		wg.Add(1)
		go func(t *Target) {
			defer wg.Done()
			err := t.Run(subTargetCtx, execCtx, file)
			if err != nil {
				errChan <- err
			}
		}(subTarget)
	}

	go func() {
		wg.Wait()
		close(doneChan)
	}()

	select {
	case err := <-errChan:
		return err
	case <-doneChan:
		return nil
	}
}

func (t *Target) Run(ctx context.Context, execCtx executors.ExecutorContext, file *SupermakeFile) error {
	err := t.runDependencies(ctx, execCtx, file)
	if err != nil {
		return err
	}

	logger := execCtx.Logger.With("target", t.Name)
	execCtx.Logger = logger

	for _, step := range t.Steps {
		err := step.Run(ctx, execCtx, file)
		if err != nil {
			return err
		}
	}

	return nil
}

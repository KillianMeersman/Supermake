package supermake

import (
	"context"
	"fmt"
	"sync"
)

type Runable interface {
	Run(ctx context.Context, file *SupermakeFile, workDir string) error
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

func (t *Target) runDependencies(ctx context.Context, file *SupermakeFile, workDir string) error {
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
			err := t.Run(subTargetCtx, file, workDir)
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

func (t *Target) Run(ctx context.Context, file *SupermakeFile, workDir string) error {
	err := t.runDependencies(ctx, file, workDir)
	if err != nil {
		return err
	}

	for _, step := range t.Steps {
		err := step.Run(ctx, file, workDir)
		if err != nil {
			return err
		}
	}

	return nil
}

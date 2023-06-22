package supermake

import (
	"context"
	"fmt"
	"sync"

	"github.com/KillianMeersman/Supermake/pkg/supermake/executors"
)

type Runable interface {
	Run(ctx context.Context, execCtx executors.ExecutorContext, targets map[string]*Target) error
}

type Target struct {
	Name         string
	Node         string
	Dependencies []string
	Steps        []Runable
	Variables    map[string]*Variable
	SubTargets   map[string]*Target
	Parent       *Target
	lock         *sync.Mutex
	done         bool
}

func NewTarget(name, node string, dependencies []string, steps []Runable, Variables map[string]*Variable, subTargets map[string]*Target, Parent *Target) *Target {
	return &Target{
		Name:         name,
		Dependencies: dependencies,
		Node:         node,
		Steps:        steps,
		SubTargets:   subTargets,
		Parent:       nil,
		lock:         &sync.Mutex{},
		done:         false,
	}
}

// Run the target's dependencies
func (t *Target) runDependencies(ctx context.Context, execCtx executors.ExecutorContext, targets map[string]*Target) error {
	errChan := make(chan error)
	doneChan := make(chan struct{})
	wg := new(sync.WaitGroup)

	subTargetCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, subTargetName := range t.Dependencies {
		dependencyTarget, ok := targets[subTargetName]
		if !ok {
			return fmt.Errorf("unknown dependency target '%s'", subTargetName)
		}

		wg.Add(1)

		execCtx.Logger.Debug("running dependency target", "dependency", dependencyTarget.Name)
		go func() {
			defer wg.Done()
			err := dependencyTarget.Run(subTargetCtx, execCtx, targets)
			if err != nil {
				errChan <- err
			}
		}()
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

func (t *Target) Reset() {
	t.done = false
}

func (t *Target) Run(ctx context.Context, execCtx executors.ExecutorContext, targets map[string]*Target) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	if t.done {
		return nil
	}
	t.done = true

	logger := execCtx.Logger.With("target", t.Name)
	execCtx.Logger = logger

	logger.Debug("running target")

	targetsWithSubtargets := make(map[string]*Target)
	for k, v := range targets {
		targetsWithSubtargets[k] = v
	}
	for k, v := range t.SubTargets {
		targetsWithSubtargets[k] = v
	}

	err := t.runDependencies(ctx, execCtx, targetsWithSubtargets)
	if err != nil {
		return err
	}

	for _, step := range t.Steps {
		err := step.Run(ctx, execCtx, targetsWithSubtargets)
		if err != nil {
			return err
		}
	}

	return nil
}

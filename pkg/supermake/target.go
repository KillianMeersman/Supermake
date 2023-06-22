package supermake

import (
	"context"
	"fmt"
	"sync"

	"github.com/KillianMeersman/Supermake/pkg/supermake/executors"
)

type Target struct {
	name         string
	Node         string
	Dependencies []string
	Steps        []executors.Runable
	Variables    map[string]*Variable
	SubTargets   map[string]*Target
	Parent       *Target
	lock         *sync.Mutex
	done         bool
}

func NewTarget(name, node string, dependencies []string, steps []executors.Runable, Variables map[string]*Variable, subTargets map[string]*Target, Parent *Target) *Target {
	return &Target{
		name:         name,
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
func (t *Target) runDependencies(ctx context.Context, execCtx executors.ExecutorContext, targets map[string]executors.Runable) error {
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

		for name := range execCtx.ParentTargets {
			if name == dependencyTarget.Name() {
				return fmt.Errorf("target dependency loop %s -> %s", t.name, subTargetName)
			}
		}

		wg.Add(1)

		execCtx.Logger.Debug("running dependency target", "dependency", dependencyTarget.Name())
		go func() {
			defer wg.Done()
			err := dependencyTarget.Run(subTargetCtx, execCtx)
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

func (t *Target) Name() string {
	return t.name
}

func (t *Target) Run(ctx context.Context, execCtx executors.ExecutorContext) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	if t.done {
		return nil
	}
	t.done = true

	logger := execCtx.Logger.With("target", t.Name())
	execCtx.Logger = logger

	logger.Debug("running target")

	targetsWithSubtargets := make(map[string]executors.Runable)
	for k, v := range execCtx.Targets {
		targetsWithSubtargets[k] = v
	}
	for k, v := range t.SubTargets {
		targetsWithSubtargets[k] = v
	}

	execCtx.Targets = targetsWithSubtargets
	execCtx.ParentTargets[t.name] = t

	err := t.runDependencies(ctx, execCtx, targetsWithSubtargets)
	if err != nil {
		return err
	}

	for _, step := range t.Steps {
		err := step.Run(ctx, execCtx)
		if err != nil {
			return err
		}
	}

	return nil
}

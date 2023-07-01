package supermake

import (
	"context"
	"fmt"
	"strings"
	"sync"
)

type Targets map[string]*Target

type Target struct {
	name         string
	Node         string
	Dependencies []string
	Variables    Variables
	Executor     CommandExecutor
	Steps        []Command
	lock         *sync.Mutex
	done         bool
}

func NewTarget(name, node string, dependencies []string, executor CommandExecutor, steps []Command, variables Variables) *Target {
	return &Target{
		name:         name,
		Dependencies: dependencies,
		Node:         node,
		Executor:     executor,
		Steps:        steps,
		Variables:    variables,
		lock:         &sync.Mutex{},
		done:         false,
	}
}

// Run the target's dependencies in parallel.
func (t *Target) runDependencies(ctx context.Context, execCtx ExecutorContext) error {
	errChan := make(chan error)
	doneChan := make(chan struct{})
	wg := new(sync.WaitGroup)
	wg.Add(len(t.Dependencies))

	subTargetCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, subTargetName := range t.Dependencies {
		dependencyTarget, ok := execCtx.Targets[subTargetName]
		if !ok {
			return fmt.Errorf("unknown dependency target '%s'", subTargetName)
		}

		for name := range execCtx.ParentTargets {
			if name == dependencyTarget.Name() {
				return fmt.Errorf("target dependency loop %s -> %s", t.name, subTargetName)
			}
		}

		// execCtx.Logger.Debug("running dependency target", "dependency", dependencyTarget.FQN())
		go func() {
			defer wg.Done()
			err := execCtx.Scheduler.ScheduleTarget(subTargetCtx, execCtx, dependencyTarget)
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

func (t *Target) Name() string {
	parts := strings.Split(t.name, "::")
	return parts[len(parts)-1]

}

func (t *Target) FQN() string {
	return t.name
}

func (t *Target) Done() {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.done = true
}

func (t *Target) Reset() {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.done = false
}

func (t *Target) Run(ctx context.Context, execCtx ExecutorContext) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	if t.done {
		return nil
	}

	logger := execCtx.Logger.With("target", t.FQN())
	execCtx.Logger = logger

	err := t.runDependencies(ctx, execCtx)
	if err != nil {
		return err
	}

	logger.Debug("running target")

	for _, step := range t.Steps {
		err := t.Executor.Execute(ctx, execCtx, step)
		if err != nil {
			logger.Fatal(err.Error())
			return err
		}
	}

	t.done = true
	return nil
}

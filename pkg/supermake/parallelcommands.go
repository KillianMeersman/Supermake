package supermake

import (
	"context"
	"sync"
)

type ParallelCommands struct {
	Environment CommandExecutor
	Steps       []Command
}

func (c *ParallelCommands) GetShellCommands() []string {
	commands := make([]string, 0)
	for _, step := range c.Steps {
		commands = append(commands, step.GetShellCommands()...)
	}

	return commands
}

func (c *ParallelCommands) Name() string {
	return ""
}

func (c *ParallelCommands) FQN() string {
	return ""
}

func (c *ParallelCommands) Run(ctx context.Context, params map[string]string, execCtx ExecutorContext) error {
	wg := new(sync.WaitGroup)
	errChan := make(chan error)

	wg.Add(len(c.Steps))

	for _, step := range c.Steps {
		go func(s Command) {
			err := c.Environment.Execute(ctx, execCtx, s)
			if err != nil {
				errChan <- err
			}
			wg.Done()
		}(step)
	}

	go func() {
		wg.Wait()
		errChan <- nil
	}()

	return <-errChan
}

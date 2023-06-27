package supermake

import (
	"context"
)

type SerialCommands struct {
	Environment CommandExecutor
	Steps       []Command
}

func (c *SerialCommands) GetShellCommands() []string {
	commands := make([]string, 0)
	for _, step := range c.Steps {
		commands = append(commands, step.GetShellCommands()...)
	}

	return commands
}

func (c *SerialCommands) Name() string {
	return ""
}

func (c *SerialCommands) FQN() string {
	return ""
}

func (c *SerialCommands) Run(ctx context.Context, execCtx ExecutorContext) error {
	return c.Environment.Execute(ctx, execCtx, c)
}

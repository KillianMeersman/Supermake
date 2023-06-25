package supermake

import (
	"context"
)

type CommandGroup struct {
	Environment CommandExecutor
	Steps       []Command
}

func (c *CommandGroup) GetShellCommands() []string {
	commands := make([]string, 0)
	for _, step := range c.Steps {
		commands = append(commands, step.GetShellCommands()...)
	}

	return commands
}

func (c *CommandGroup) Name() string {
	return ""
}

func (c *CommandGroup) Run(ctx context.Context, execCtx ExecutorContext) error {
	return c.Environment.Execute(ctx, execCtx, c)
}

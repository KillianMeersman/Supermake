package supermake

import (
	"context"

	"github.com/KillianMeersman/Supermake/pkg/supermake/executables"
	"github.com/KillianMeersman/Supermake/pkg/supermake/executors"
)

type CommandGroup struct {
	Environment executors.CommandExecutor
	Steps       []executables.Command
}

func (c *CommandGroup) GetShellCommands() []string {
	commands := make([]string, 0)
	for _, step := range c.Steps {
		commands = append(commands, step.GetShellCommands()...)
	}

	return commands
}

func (c *CommandGroup) Run(ctx context.Context, execCtx executors.ExecutorContext, targets map[string]*Target) error {
	return c.Environment.Execute(ctx, execCtx, c)
}

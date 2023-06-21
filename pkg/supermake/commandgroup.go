package supermake

import (
	"context"
	"strings"

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

func (c *CommandGroup) Run(ctx context.Context, execCtx executors.ExecutorContext, file *SupermakeFile) error {
	command := &executables.ShellCommand{Command: strings.Join(c.GetShellCommands(), "\n")}
	return c.Environment.Execute(ctx, execCtx, command)
}

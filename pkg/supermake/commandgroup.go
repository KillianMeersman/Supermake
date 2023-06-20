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

func (c *CommandGroup) Run(ctx context.Context, file *SupermakeFile, workDir string) error {
	for _, step := range c.Steps {
		err := c.Environment.Execute(ctx, workDir, step)
		if err != nil {
			return err
		}
	}

	return nil
}

package supermake

import "context"

type Command interface {
	GetShellCommands() []string
}

type Runable interface {
	Run(ctx context.Context, workDir string) error
}

type ShellCommand struct {
	Command string
}

func (c *ShellCommand) GetShellCommands() []string {
	return []string{c.Command}
}

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

func (c *CommandGroup) Run(ctx context.Context, workDir string) error {
	for _, step := range c.Steps {
		err := c.Environment.Execute(ctx, workDir, step)
		if err != nil {
			return err
		}
	}

	return nil
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

func (t *Target) Run(ctx context.Context, workDir string) error {
	for _, step := range t.Steps {
		err := step.Run(ctx, workDir)
		if err != nil {
			return err
		}
	}

	return nil
}

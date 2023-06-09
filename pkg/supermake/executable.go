package supermake

import "context"

type Executable interface {
	GetCommands() []string
}

type Runable interface {
	Run(ctx context.Context, workDir string) error
}

type ShellCommand struct {
	Command string
}

func (c *ShellCommand) GetCommands() []string {
	return []string{c.Command}
}

type CommandGroup struct {
	Environment Executor
	Steps       []Executable
}

func (c *CommandGroup) GetCommands() []string {
	commands := make([]string, 0)
	for _, step := range c.Steps {
		commands = append(commands, step.GetCommands()...)
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

package supermake

type Executable interface {
	Execute() error
}

type ShellCommand struct {
	Command string
}

func (c *ShellCommand) Execute() error {
	return nil
}

type CommandGroup struct {
	Environment Executor
	Steps       []Executable
}

func (c *CommandGroup) Execute() error {
	return nil
}

type Target struct {
	Name         string
	Dependencies []string
	Steps        []Executable
}

func (t *Target) Execute() error {
	return nil
}

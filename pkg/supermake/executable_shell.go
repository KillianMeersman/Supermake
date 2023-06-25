package supermake

type ShellCommand struct {
	Command string
}

func (c *ShellCommand) GetShellCommands() []string {
	return []string{c.Command}
}

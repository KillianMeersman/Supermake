package supermake

type CommandGroup []string

func (c CommandGroup) GetShellCommands() []string {
	return c
}

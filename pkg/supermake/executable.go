package supermake

type Command interface {
	GetShellCommands() []string
}

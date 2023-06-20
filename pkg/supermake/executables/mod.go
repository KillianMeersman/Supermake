package executables

type Command interface {
	GetShellCommands() []string
}

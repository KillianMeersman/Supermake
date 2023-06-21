package executors

import (
	"context"
	"os"
	"os/exec"

	"github.com/KillianMeersman/Supermake/pkg/supermake/executables"
	"github.com/KillianMeersman/Supermake/pkg/supermake/util"
)

type LocalEnvironment struct {
	Entrypoint string
}

func (l *LocalEnvironment) Execute(ctx context.Context, execCtx ExecutorContext, command executables.Command) error {

	for _, command := range command.GetShellCommands() {
		// Split the command and arguments
		parts := util.SplitWords(command)
		command := parts[0]
		args := parts[1:]
		if l.Entrypoint != "" {
			command = l.Entrypoint
			args = parts
		}
		cmd := exec.Command(command, args...)
		// Couple host streams to command, so that the user can see the output
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Run it
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}

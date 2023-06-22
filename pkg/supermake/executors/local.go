package executors

import (
	"bufio"
	"bytes"
	"context"
	"io"
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

		stdout := new(bytes.Buffer)

		// Couple host streams to command, so that the user can see the output
		cmd.Stdout = stdout
		cmd.Stderr = stdout

		// Run it
		err := cmd.Run()
		if err != nil {
			return err
		}

		go func() {
			reader := bufio.NewReader(stdout)

			for {
				line, err := reader.ReadBytes('\n')
				if err != nil {
					if err != io.EOF {
						execCtx.Logger.Fatal(err.Error())
					}
					return
				}

				line = bytes.TrimRight(line, "\n\r")
				execCtx.Logger.Info(string(line))
			}
		}()
	}
	return nil
}

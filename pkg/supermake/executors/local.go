package executors

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"os/exec"
	"strings"

	"github.com/KillianMeersman/Supermake/pkg/supermake/executables"
	"github.com/KillianMeersman/Supermake/pkg/supermake/util"
)

type LocalEnvironment struct {
	Entrypoint []string
}

func NewLocalEnvironment() *LocalEnvironment {
	return &LocalEnvironment{
		Entrypoint: []string{"bash", "-c"},
	}
}

func (l *LocalEnvironment) Execute(ctx context.Context, execCtx ExecutorContext, command executables.Command) error {

	for _, command := range command.GetShellCommands() {
		// Split the command and arguments
		parts := util.SplitWords(command)
		program := parts[0]
		args := parts[1:]

		if len(l.Entrypoint) > 0 {
			program = l.Entrypoint[0]
			args = l.Entrypoint[1:]
			args = append(args, strings.Join(parts, " "))
		}

		execCtx.Logger.Trace("running command", "command", command)
		cmd := exec.Command(program, args...)

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

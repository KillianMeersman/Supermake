package executors

import (
	"bytes"
	"context"
	"os/exec"
	"strings"

	"github.com/KillianMeersman/Supermake/pkg/supermake/executables"
	"github.com/KillianMeersman/Supermake/pkg/supermake/log"
)

type LocalEnvironment struct {
	Entrypoint []string
}

func NewLocalEnvironment() *LocalEnvironment {
	return &LocalEnvironment{
		Entrypoint: []string{"bash", "-c"},
	}
}

// Start an interpreter and feed commands into stdin, logs stdout and stderr to the logger as INFO level logs.
func startAndStreamOutput(ctx context.Context, command string, args []string, logger *log.Logger) error {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	output = bytes.TrimRight(output, "\n\r")
	logger.Info(string(output))
	return err
}

func (l *LocalEnvironment) Execute(ctx context.Context, execCtx ExecutorContext, command executables.Command) error {
	args := []string{strings.Join(l.Entrypoint[1:], " ")}
	args = append(args, strings.Join(command.GetShellCommands(), "\n"))
	return startAndStreamOutput(ctx, l.Entrypoint[0], args, execCtx.Logger)
}

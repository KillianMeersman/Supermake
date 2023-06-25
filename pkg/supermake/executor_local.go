package supermake

import (
	"bytes"
	"context"
	"os/exec"

	"github.com/KillianMeersman/Supermake/pkg/supermake/log"
)

type LocalEnvironment struct {
	Entrypoint []string
}

func NewLocalEnvironment() *LocalEnvironment {
	return &LocalEnvironment{
		Entrypoint: []string{"sh", "-ce"},
	}
}

// Start an interpreter and feed commands into stdin, logs stdout and stderr to the logger as INFO level logs.
func startAndStreamOutput(ctx context.Context, command string, args []string, vars Variables, logger *log.Logger) error {
	cmd := exec.Command(command, args...)

	cmd.Env = append(cmd.Env, vars.EnvStrings()...)

	output, err := cmd.CombinedOutput()
	output = bytes.TrimRight(output, "\n\r")
	logger.Info(string(output))
	return err
}

func (l *LocalEnvironment) Execute(ctx context.Context, execCtx ExecutorContext, command Command) error {
	entrypoint, cmd := ParseInterpreterCommand(l.Entrypoint, command.GetShellCommands())
	return startAndStreamOutput(ctx, entrypoint, cmd, execCtx.EnvVars, execCtx.Logger)
}

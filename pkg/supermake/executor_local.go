package supermake

import (
	"context"
	"os/exec"
	"sync"

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
	cmd := exec.CommandContext(ctx, command, args...)

	// Set env
	cmd.Env = append(cmd.Env, vars.EnvStringsInherited()...)

	// Stream output
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		log.StreamReaderNewLines(logger, stdout)
		wg.Done()
	}()

	go func() {
		log.StreamReaderNewLines(logger, stderr)
		wg.Done()
	}()

	wg.Wait()
	return cmd.Wait()
}

func (l *LocalEnvironment) Execute(ctx context.Context, execCtx ExecutorContext, command Command) error {
	entrypoint, cmd := ParseInterpreterCommand(l.Entrypoint, command.GetShellCommands())
	return startAndStreamOutput(ctx, entrypoint, cmd, execCtx.EnvVars, execCtx.Logger)
}

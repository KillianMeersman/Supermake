package supermake

import (
	"context"
	"strings"

	"github.com/KillianMeersman/Supermake/pkg/supermake/log"
)

func ParseInterpreterCommand(entrypoint, commands []string) (string, []string) {
	args := entrypoint[1:]

	cmd := strings.Join(commands, "\n")
	// cmd = strings.ReplaceAll(cmd, `"`, `\"`)
	// cmd = strings.ReplaceAll(cmd, `!`, `\!`)
	args = append(args, cmd)

	return entrypoint[0], args
}

type CommandExecutor interface {
	Execute(ctx context.Context, execCtx ExecutorContext, command Command) error
}

type Runable interface {
	Run(ctx context.Context, execCtx ExecutorContext) error
	Name() string
}

type ExecutorContext struct {
	EnvVars       Variables
	Targets       map[string]Runable
	ParentTargets map[string]Runable
	WorkingDir    string
	Logger        *log.Logger
}

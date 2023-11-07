package supermake

import (
	"context"
	"strings"

	"github.com/KillianMeersman/Supermake/pkg/supermake/log"
)

func parseDefaultEntrypoint(variables Variables) []string {
	ep, ok := variables[".SHELL"]
	if !ok {
		return []string{"sh", "-cex"}
	}

	return strings.Split(ep.Value(), " ")
}

func ParseInterpreterCommand(variables Variables, entrypoint, commands []string) (string, []string) {
	if len(entrypoint) == 0 {
		entrypoint = parseDefaultEntrypoint(variables)
	}
	args := entrypoint[1:]

	cmd := strings.Join(commands, "\n")
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
	Targets       Targets
	ParentTargets Targets
	WorkingDir    string
	Logger        *log.Logger
	Scheduler     Scheduler
}

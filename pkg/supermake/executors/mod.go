package executors

import (
	"context"

	"github.com/KillianMeersman/Supermake/pkg/supermake/executables"
	"github.com/KillianMeersman/Supermake/pkg/supermake/log"
)

type CommandExecutor interface {
	Execute(ctx context.Context, execCtx ExecutorContext, command executables.Command) error
}

type Runable interface {
	Run(ctx context.Context, execCtx ExecutorContext) error
	Name() string
}

type ExecutorContext struct {
	EnvVars       map[string]string
	Targets       map[string]Runable
	ParentTargets map[string]Runable
	WorkingDir    string
	Logger        *log.Logger
}

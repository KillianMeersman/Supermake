package executors

import (
	"context"

	"github.com/KillianMeersman/Supermake/pkg/supermake/executables"
	"github.com/KillianMeersman/Supermake/pkg/supermake/log"
)

type CommandExecutor interface {
	Execute(ctx context.Context, execCtx ExecutorContext, command executables.Command) error
}

type ExecutorContext struct {
	EnvVars    map[string]string
	WorkingDir string
	Logger     *log.Logger
}

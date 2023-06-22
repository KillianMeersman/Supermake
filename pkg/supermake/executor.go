package supermake

import (
	"context"
	"log"

	"github.com/KillianMeersman/Supermake/pkg/supermake/executables"
)

type CommandExecutor interface {
	Execute(ctx context.Context, execCtx ExecutorContext, command executables.Command) error
}

type ExecutorContext struct {
	EnvVars    map[string]string
	WorkingDir string
	Logger     *log.Logger
}

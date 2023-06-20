package executors

import (
	"context"

	"github.com/KillianMeersman/Supermake/pkg/supermake/executables"
)

type CommandExecutor interface {
	Execute(ctx context.Context, workingDir string, command executables.Command) error
}

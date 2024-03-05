package supermake_test

import (
	"context"
	"os"
	"testing"

	"github.com/KillianMeersman/Supermake/pkg/supermake"
	"github.com/KillianMeersman/Supermake/pkg/supermake/parse"
)

func TestContainerLogs(t *testing.T) {
	smakeFileData := `
	test:
		error:
			@alpine:3
			echo 'Test error start'
			echo 'Test error stop'
	`

	ctx := context.Background()
	cwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	smakeFile, err := parse.ParseSupermakeString(cwd, smakeFileData)
	if err != nil {
		t.Error(err)
	}

	scheduler := supermake.NewLocalScheduler()
	smakeFile.Run(ctx, scheduler, cwd, "test")
}

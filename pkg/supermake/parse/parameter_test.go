package parse_test

import (
	"context"
	"testing"

	"github.com/KillianMeersman/Supermake/pkg/supermake"
	"github.com/KillianMeersman/Supermake/pkg/supermake/parse"
)

func TestParameters(t *testing.T) {
	definition := `
	echo(a, b):
		echo '${{ a }} ${{ b }}'

	test: echo(a=test1, b=test2)
	`

	file, err := parse.ParseSupermakeString(".", definition)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	scheduler := supermake.NewLocalScheduler()

	file.Run(ctx, scheduler, ".", "test")
}

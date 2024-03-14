package parse_test

import (
	"testing"

	"github.com/KillianMeersman/Supermake/pkg/supermake/parse"
)

func TestVariableSubstitution(t *testing.T) {
	test := "This is a ${{ OPERATION }} for ${{WORKLOAD}} ${{VERSION}} ${{WORKLOAD-${{VERSION}}}}"
	vars := map[string]string{
		"OPERATION":   "test",
		"WORKLOAD":    "replacevars",
		"VERSION":     "v1",
		"WORKLOAD-v1": "YAY",
	}

	value, err := parse.ReplaceVariables(test, func(v string) (string, error) {
		return vars[v], nil
	})

	if err != nil {
		t.Error(err)
	}

	if value != "This is a test for replacevars v1 YAY" {
		t.Errorf("variable should expand to '%s', but got '%s'", "This is a test for replacevars v1 YAY", value)
	}
}

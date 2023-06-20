package parse_test

import (
	"testing"

	"github.com/KillianMeersman/Supermake/pkg/supermake/parse"
)

func TestVariableSubstitution(t *testing.T) {
	test := "This is a $(OPERATION) for $(WORKLOAD) v$(VERSION) $(WORKLOAD-$(VERSION))"
	vars := map[string]string{
		"OPERATION":  "test",
		"WORKLOAD":   "replacevars",
		"VERSION":    "1",
		"WORKLOAD-1": "YAY",
	}

	value, err := parse.ReplaceVariables(test, func(v string) string {
		return vars[v]
	})

	if err != nil {
		t.Fail()
	}

	if value != "This is a test for replacevars v1" {
		t.Fail()
	}
}

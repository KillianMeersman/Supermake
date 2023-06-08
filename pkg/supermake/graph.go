package supermake

import "fmt"

type TargetDAG struct {
	levels [][]string
}

func ComputeTargetDAG(targets map[string]Target, start *Target) error {
	seenTargets := make(map[string]struct{})
	levels := make([][]string, 0)

	currentTarget := start
	currentLevel := 0

	for { // level
		for _, dep := range currentTarget.Dependencies {
			// Check for circular references
			if _, ok := seenTargets[dep]; ok {
				return fmt.Errorf("circular reference to '%s'", dep)
			}

			level := make([]string, 0)
			levels = append(levels, level)

			// Mark targets in this level as seen
			for _, dep := range level {
				seenTargets[dep] = struct{}{}
			}

			levels[currentLevel] = append(levels[currentLevel], currentTarget.Dependencies...)
		}
	}

	return nil
}

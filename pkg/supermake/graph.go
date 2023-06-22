package supermake

import (
	"errors"
	"fmt"
)

type TargetDAG struct {
	levels [][]string
}

// Get the target's dependencies, including the
func getTargetDependencies(target *Target, deps map[string]struct{}) error {
	for _, dep := range target.Dependencies {
		if dep == target.Name() {
			return errors.New("nested targets may not list their parent as a dependency")
		}
		deps[dep] = struct{}{}
	}

	for _, subTarget := range target.SubTargets {
		getTargetDependencies(subTarget, deps)
	}

	if target.Parent != nil {
		deps[target.Parent.Name()] = struct{}{}
	}

	return nil
}

func ComputeTargetDAG(targets map[string]*Target, start string) error {
	seenTargets := make(map[string]struct{})
	levels := make([][]string, 0)

	currentTarget, ok := targets[start]
	if !ok {
		return fmt.Errorf("unknown target: %s", start)
	}
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

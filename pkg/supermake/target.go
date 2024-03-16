package supermake

import (
	"context"
	"fmt"
	"regexp"
	"slices"
	"strings"
	"sync"
)

// Parse a target call, including parameters. e.g.
// `build_docker(stage=dev)`
func ParseTargetCall(call string) (string, map[string]string, error) {
	nameRe := regexp.MustCompile(`([\w-\.]+)\(`)
	paramRe := regexp.MustCompile(`([\w-\.]+)=([\w-\.]+)`)

	name := nameRe.FindString(call)

	paramMatches := paramRe.FindAllStringSubmatch(call, 50)
	params := make(map[string]string)
	for _, match := range paramMatches {
		params[strings.TrimSpace(match[1])] = strings.TrimSpace(match[2])
	}

	return name, params, nil
}

type Targets map[string]*Target

type Target struct {
	name         string
	Parameters   []string
	Node         string
	Dependencies []string
	Variables    Variables
	Executor     CommandExecutor
	Steps        []Command
}

func NewTarget(name, node string, parameters, dependencies []string, executor CommandExecutor, steps []Command, variables Variables) *Target {
	return &Target{
		name:         name,
		Parameters:   parameters,
		Dependencies: dependencies,
		Node:         node,
		Executor:     executor,
		Steps:        steps,
		Variables:    variables,
	}
}

// Run the target's dependencies in parallel.
func (t *Target) runDependencies(ctx context.Context, execCtx ExecutorContext) error {
	errChan := make(chan error)
	doneChan := make(chan struct{})
	wg := new(sync.WaitGroup)
	wg.Add(len(t.Dependencies))

	subTargetCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, subTargetName := range t.Dependencies {
		name, parameters, err := ParseTargetCall(subTargetName)
		if err != nil {
			return err
		}

		dependencyTarget, ok := execCtx.Targets[name]
		if !ok {
			return fmt.Errorf("unknown dependency target '%s'", subTargetName)
		}

		// Detect loops
		for name := range execCtx.ParentTargets {
			if name == dependencyTarget.Name() {
				return fmt.Errorf("target dependency loop %s -> %s", t.name, subTargetName)
			}
		}

		execCtx.Logger.Debug("running dependency target", "dependency", dependencyTarget.FQN())
		go func() {
			defer wg.Done()
			err := execCtx.Scheduler.ScheduleTarget(subTargetCtx, parameters, execCtx, dependencyTarget)
			if err != nil {
				errChan <- err
			}
		}()
	}

	go func() {
		wg.Wait()
		close(doneChan)
	}()

	select {
	case err := <-errChan:
		return err
	case <-doneChan:
		return nil
	}
}

func (t *Target) Name() string {
	parts := strings.Split(t.name, "::")
	return parts[len(parts)-1]
}

func (t *Target) FQN() string {
	return t.name
}

func (t *Target) FQNWithParameters(parameters map[string]string) string {
	sortedParameters := make([]string, 0, len(parameters))
	for key := range parameters {
		sortedParameters = append(sortedParameters, key)
	}
	slices.Sort(sortedParameters)

	paramString := strings.Builder{}
	for _, key := range sortedParameters {
		paramString.WriteString(fmt.Sprintf("%s=%s", key, parameters[key]))
	}

	return fmt.Sprintf("%s(%s)", t.FQN(), paramString.String())
}

func (t *Target) Run(ctx context.Context, params map[string]string, execCtx ExecutorContext) error {
	logger := execCtx.Logger.With("target", t.FQN())
	execCtx.Logger = logger

	for key, val := range params {
		execCtx.EnvVars[key] = NewVariable(key, RECURSIVE, true, val)
	}

	logger.Info(strings.Join(execCtx.EnvVars.EnvStrings(), ", "))

	err := t.runDependencies(ctx, execCtx)
	if err != nil {
		return err
	}

	logger.Debug("running target")

	for _, step := range t.Steps {
		err := t.Executor.Execute(ctx, execCtx, step)
		if err != nil {
			logger.Fatal(err.Error())
			return err
		}
	}

	return nil
}

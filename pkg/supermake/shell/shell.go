package shell

import (
	"bufio"
	"context"
	"os/exec"
	"strings"
	"sync"
)

// Split a shell command into an executable and its arguments.
func SplitCommand(command string) (string, []string) {
	args := strings.Split(command, " ")
	if len(args) < 1 {
		return "", []string{}
	}

	return args[0], args[1:]
}

func StreamShellCommand(ctx context.Context, command string, variables []string) (<-chan string, <-chan error, error) {
	command, args := SplitCommand(command)

	ctx, cancel := context.WithCancel(ctx)
	cmd := exec.CommandContext(ctx, command, args...)

	outChan := make(chan string)
	errChan := make(chan error)

	// Set env
	cmd.Env = append(cmd.Env, variables...)

	// Stream output
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		cancel()
		return outChan, errChan, err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		cancel()
		return outChan, errChan, err
	}

	err = cmd.Start()
	if err != nil {
		cancel()
		return outChan, errChan, err
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			m := scanner.Text()
			outChan <- m
		}
		wg.Done()
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			m := scanner.Text()
			outChan <- m
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		err := cmd.Wait()
		if err != nil {
			errChan <- err
		}
		cancel()
	}()

	return outChan, errChan, nil
}

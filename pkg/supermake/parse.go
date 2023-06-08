package supermake

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var targetRegex = regexp.MustCompile("^\t*([a-zA-Z0-9-_.@]+): +(.*)$")
var executorRegex = regexp.MustCompile("^\t*@([a-zA-Z0-9-_:.]+)( .*)?$")
var variableDeclarationRegex = regexp.MustCompile("^(export +)?(\\w*?) *([?:]?=) *([^#\n]*)")
var variableRegex = regexp.MustCompile("$\\((.*)\\)")

const COMMENT_CHARS = "#"

type VariableEvaluationType int

const (
	RECURSIVE VariableEvaluationType = iota
	FALLBACK
)

type Variable struct {
	Name           string
	EvaluationType VariableEvaluationType
	Export         bool
	Value          string
}

type SupermakeFile struct {
	Targets   map[string]Target
	Variables map[string]Variable
}

func (s *SupermakeFile) Run(target string) error {
	t, ok := s.Targets[target]
	if !ok {
		return fmt.Errorf("no such target: %s", target)
	}

	return t.Run(context.TODO(), ".")
}

// Get the index of the first non-whitespace character.
// Returns -1 if the line is empty or all whitespace.
func FirstCharacter(line string) int {
	return strings.IndexFunc(line, func(r rune) bool {
		return r != ' ' && r != '\t'
	})
}

func countAndRemoveIndentation(line string) (string, int, error) {
	spaceCounter := 0
	indentation := 0

	i := 0
	for _, char := range line {
		if char == ' ' {
			spaceCounter++
		} else if char == '\t' {
			indentation++
		} else {
			break
		}

		if spaceCounter == 4 {
			indentation++
			spaceCounter = 0
		} else if spaceCounter > 0 && spaceCounter < 4 {
			return "", -1, fmt.Errorf("invalid indentation for line '%s'", line)
		}

		i++
	}

	return line[i:], indentation, nil
}

func stripTrailingComment(line string) string {
	lastIndex := strings.LastIndex(line, "#")
	if lastIndex > -1 {
		return line[lastIndex-1:]
	}
	return line
}

// Build a map of cleaned lines and their indentation levels.
// Removes indentation characters as well as comments.
func cleanAndMapLines(lines []string) ([]string, []int, map[int]int) {
	indentationLevels := make([]int, 0, len(lines))
	filteredLines := make([]string, 0, len(lines))
	sourceMap := make(map[int]int)

	// build indentation level map
	for i, line := range lines {
		line, indentation, err := countAndRemoveIndentation(line)
		if err != nil {
			log.Fatal(err)
		}

		if len(line) < 1 {
			continue
		}

		indentationLevels = append(indentationLevels, indentation)
		filteredLines = append(filteredLines, line)
		sourceMap[len(filteredLines)-1] = i
	}

	return filteredLines, indentationLevels, sourceMap
}

// Find the end line of an indented block.
func findBlockEnd(lines []string, indentationLevels []int, index int) int {
	minIndentation := indentationLevels[index] + 1

	for i := index; i < len(lines); i++ {
		if indentationLevels[i] < minIndentation {
			return i
		}
	}

	return len(lines)
}

func parseTarget(lines []string) (*Target, error) {
	headerParts := strings.SplitN(lines[0], ":", 2)

	// Name & Node
	nameParts := strings.SplitN(headerParts[0], "@", 2)
	name := strings.TrimSpace(nameParts[0])
	node := ""
	if len(nameParts) > 1 {
		node = strings.TrimSpace(nameParts[1])
	}

	// Dependencies
	rawDeps := strings.Split(headerParts[1], " ")
	dependencies := make([]string, 0, len(rawDeps))
	for _, dep := range rawDeps {
		if len(dep) > 0 {
			dependencies = append(dependencies, strings.TrimSpace(dep))
		}
	}

	// Steps & executors
	steps := make([]Runable, 0)
	currentCommands := make([]Executable, 0)
	var currentExecutor Executor = new(LocalEnvironment)

	for _, line := range lines[1:] {
		// If executor line, end the current command group.
		if executorRegex.MatchString(line) {
			if len(currentCommands) > 0 {
				steps = append(steps, &CommandGroup{
					Environment: currentExecutor,
					Steps:       currentCommands,
				})
			}

			groups := executorRegex.FindStringSubmatch(line)
			currentExecutor = &DockerEnvironment{
				Image:      strings.TrimSpace(groups[1]),
				Entrypoint: strings.TrimSpace(groups[2]),
			}
		} else {
			command := &ShellCommand{
				Command: line,
			}
			currentCommands = append(currentCommands, command)
		}
	}

	// Add the remaining command group
	steps = append(steps, &CommandGroup{
		Environment: currentExecutor,
		Steps:       currentCommands,
	})

	return &Target{
		Name:         name,
		Dependencies: dependencies,
		Node:         node,
		Steps:        steps,
	}, nil
}

func evaluateVariable(variable string, variables map[string]*Variable) (string, error) {
	evaluatedVariables := make(map[string]struct{}) // prevent loops

	for {
		v, ok := variables[variable]
		if !ok {
			return "", fmt.Errorf("unknown variable '%s'", variable)
		}

		if _, ok := evaluatedVariables[variable]; ok {
			return "", errors.New("variable evaluation loop")
		}
		evaluatedVariables[variable] = struct{}{}

		if strings.HasPrefix(variable, "$(") && strings.HasSuffix(variable, ")") {
			variable = variable[2 : len(variable)-2]
		} else {
			return v.Value, nil
		}
	}
}

func parseVariable(line string) (*Variable, error) {
	exportSubstr := "export "
	i := 0

	// export?
	export := strings.HasPrefix(line, exportSubstr)
	if export {
		i = len(exportSubstr)
	}

	// identifier
	identifierEnd := strings.IndexRune(line[i:], ' ')
	identifier := line[i:identifierEnd]
	i = identifierEnd

	// evaluation type
	i = strings.IndexAny(line, ":?=")
	if i == -1 {
		return nil, fmt.Errorf("invalid variable declaration '%s'", line)
	}
	var evaluationType VariableEvaluationType
	switch line[i] {
	case ':', '=':
		evaluationType = RECURSIVE
	case '?':
		evaluationType = FALLBACK
	default:
		return nil, fmt.Errorf("invalid variable evaluation type '%b'", line[i])
	}

	i = strings.LastIndex(line[i:], "=")

	return &Variable{
		Name:           identifier,
		EvaluationType: evaluationType,
		Export:         export,
		Value:          strings.TrimSpace(line[i:]),
	}, nil
}

func ParseSupermakeFileV2(path string) (*SupermakeFile, error) {
	dataBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	data := string(dataBytes)
	lines := strings.Split(data, "\n")
	lines, indentationLevels, _ := cleanAndMapLines(lines)
	targetStack := Stack[*Target]{}

	variables := make(map[string]Variable)

	for i, line := range lines {
		// Enter target
		targetEnd := i
		if targetRegex.MatchString(line) {
			if i+1 <= len(indentationLevels) {
				targetEnd = findBlockEnd(lines, indentationLevels, i)
			}

			target, err := parseTarget(lines[i : targetEnd+1])
			if err != nil {
				return nil, err
			}
			targetStack.Push(target)
		} else if variableDeclarationRegex.MatchString(line) {
			variable, err := parseVariable(line)
			if err != nil {
				return nil, err
			}
			variables[variable.Name] = *variable
		}
	}

	return &SupermakeFile{}, nil
}

package supermake

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var targetRegex = regexp.MustCompile(`^([a-zA-Z0-9-/_.@]+):(?: +(.*))?$`)
var executorRegex = regexp.MustCompile(`^\t*@([a-zA-Z0-9-_:.]+)( .*)?$`)
var variableDeclarationRegex = regexp.MustCompile(`^(export +)?(\w+) +(=|:=|\?=) +(.*)$`)
var variableRegex = regexp.MustCompile(`\$\((.*)\)`)

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
	value          string
}

func (v *Variable) Value() string {
	if v.EvaluationType == FALLBACK {
		if v, ok := os.LookupEnv(v.Name); ok {
			return v
		}
	}
	return v.value
}

type SupermakeFile struct {
	Targets   map[string]*Target
	Variables map[string]*Variable
}

type SuperMakeFileParser struct {
	rawLines []string

	lines     []string
	indent    []int
	sourcemap []int
}

func NewParser(lines []string) *SuperMakeFileParser {
	return &SuperMakeFileParser{
		rawLines: lines,
	}
}

// Build a map of cleaned lines and their indentation levels.
// Removes indentation characters as well as comments.
func (p *SuperMakeFileParser) cleanAndMapLines() {
	p.indent = make([]int, 0, len(p.rawLines))
	p.lines = make([]string, 0, len(p.rawLines))
	p.sourcemap = make([]int, 0)

	// build indentation level map
	for i, line := range p.rawLines {
		line, indentation, err := countAndRemoveIndentation(line)
		if err != nil {
			log.Fatal(err)
		}

		line = stripTrailingComment(line)
		line = strings.TrimSpace(line)

		if len(line) < 1 {
			continue
		}

		p.indent = append(p.indent, indentation)
		p.lines = append(p.lines, line)
		p.sourcemap = append(p.sourcemap, i)
	}
}

func (p *SuperMakeFileParser) parseGlobalVariables() (map[string]*Variable, error) {
	variables := make(map[string]*Variable)

	for i, line := range p.lines {
		if variableDeclarationRegex.MatchString(line) {
			variable, err := parseVariable(line, variables)
			if err != nil {
				return nil, fmt.Errorf("error on line %d: %s", p.sourcemap[i]+1, err)
			}
			variables[variable.Name] = variable
		}
	}

	return variables, nil
}

func registerNestedTargets(prefix string, target *Target, targets map[string]*Target) {
	name := getNestedTargetName(prefix, target.Name)
	targets[name] = target

	for _, subTarget := range target.SubTargets {
		key := getNestedTargetName(target.Name, subTarget.Name)
		targets[key] = subTarget
		registerNestedTargets(name, subTarget, targets)
	}
}

func (p *SuperMakeFileParser) parseTargets() (map[string]*Target, error) {
	targets := make(map[string]*Target)

	for i, line := range p.lines {
		if targetRegex.MatchString(line) {
			// If not at the last line
			targetEnd := i
			if i+1 <= len(p.indent) {
				targetEnd = findBlockEnd(p.indent, i)
			}
			target, err := parseTarget(p.lines[i:targetEnd], p.indent[i:targetEnd])
			if err != nil {
				return nil, err
			}
			registerNestedTargets("", target, targets)
		}
	}

	return targets, nil
}

func (p *SuperMakeFileParser) Parse() (*SupermakeFile, error) {
	p.cleanAndMapLines()
	globalVariables, err := p.parseGlobalVariables()
	if err != nil {
		return nil, err
	}

	targets, err := p.parseTargets()
	if err != nil {
		return nil, err
	}

	return &SupermakeFile{
		Targets:   targets,
		Variables: globalVariables,
	}, nil
}

// Find the end line of an indented block.
// Assumes the given line is expected_block_indent - 1.
func findBlockEnd(indent []int, index int) int {
	minIndentation := indent[index] + 1

	for i := index + 1; i < len(indent); i++ {
		if indent[i] < minIndentation {
			return i
		}
	}

	return len(indent)
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
		return line[:lastIndex]
	}
	return line
}

func getNestedTargetName(parent, name string) string {
	if parent != "" {
		return fmt.Sprintf("%s::%s", parent, name)
	}
	return name
}

// Parse a target.
// Assumes the target declaration is the first line.
func parseTarget(lines []string, indentationLevels []int) (*Target, error) {
	headerParts := strings.SplitN(lines[0], ":", 2)

	// Name & Node
	nameParts := strings.SplitN(headerParts[0], "@", 2)
	name := strings.TrimSpace(nameParts[0])
	node := ""
	if len(nameParts) > 1 {
		node = strings.TrimSpace(nameParts[1])
	}

	// Dependencies
	dependencies := make([]string, 0)
	if len(headerParts) > 1 {
		rawDeps := strings.Split(headerParts[1], " ")
		for _, dep := range rawDeps {
			if len(dep) > 0 {
				dependencies = append(dependencies, strings.TrimSpace(dep))
			}
		}
	}

	// Sub Targets
	subTargets := make(map[string]*Target)

	// Steps & executors
	steps := make([]Runable, 0)
	currentCommands := make([]Command, 0)
	var currentExecutor CommandExecutor = new(LocalEnvironment)

	for i := 1; i < len(lines); i++ {
		line := lines[i]

		// Subtargets
		if targetRegex.MatchString(line) {
			blockEnd := findBlockEnd(indentationLevels, i)
			subTarget, err := parseTarget(lines[i:blockEnd], indentationLevels[i:blockEnd])
			if err != nil {
				return nil, err
			}
			subTargets[subTarget.Name] = subTarget
			if len(currentCommands) > 0 {
				steps = append(steps, &CommandGroup{
					Environment: currentExecutor,
					Steps:       currentCommands,
				})
			}
			currentCommands = make([]Command, 0)
			currentExecutor = new(LocalEnvironment)
			steps = append(steps, subTarget)
			i = blockEnd - 1
			// Executors
		} else if executorRegex.MatchString(line) {
			// If commands were defined, save them in a new command group
			if len(currentCommands) > 0 {
				steps = append(steps, &CommandGroup{
					Environment: currentExecutor,
					Steps:       currentCommands,
				})
			}

			groups := executorRegex.FindStringSubmatch(line)

			// Reset commands & executor
			currentCommands = make([]Command, 0)
			currentExecutor = &DockerEnvironment{
				Image:      strings.TrimSpace(groups[1]),
				Entrypoint: strings.TrimSpace(groups[2]),
			}
		} else {
			// Commands
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

	target := &Target{
		Name:         name,
		Dependencies: dependencies,
		Node:         node,
		Steps:        steps,
		SubTargets:   subTargets,
		Parent:       nil,
	}

	for _, t := range subTargets {
		t.Parent = target
	}

	return target, nil
}

func runShellScript(script string, variables map[string]*Variable) ([]byte, error) {
	cmd := exec.Command("sh", "-c", script)
	// Set environment variables
	for _, v := range variables {
		if v.Export {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", v.Name, v.Value()))
		}
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("%s: '%s'", err, strings.TrimRight(stderr.String(), "\n\r"))
	}

	return stdout.Bytes(), nil
}

func evaluateVariables(line string, variables map[string]*Variable) (string, error) {
	substitutions := make(map[string]string)

	for _, match := range variableRegex.FindAllStringSubmatch(line, -1) {
		wholeMatch := match[0]
		rawIdentifier := match[1]
		if len(rawIdentifier) < 1 {
			return "", errors.New("empty variable")
		}

		// Substitute nested variables
		evaluatedIdentifier, err := evaluateVariables(rawIdentifier, variables)
		if err != nil {
			return "", err
		}

		if strings.HasPrefix(evaluatedIdentifier, "shell ") {
			shellString := strings.TrimPrefix(evaluatedIdentifier, "shell ")
			stdout, err := runShellScript(shellString, variables)
			if err != nil {
				return "", fmt.Errorf("variable shell command failed to execute: '%s': %s", shellString, err)
			}
			t := strings.TrimRight(string(stdout), "\n\r")
			substitutions[wholeMatch] = t
		} else {
			if variable, ok := variables[evaluatedIdentifier]; ok {
				substitutions[wholeMatch] = variable.Value()
			} else if variable, ok := os.LookupEnv(evaluatedIdentifier); ok {
				substitutions[wholeMatch] = variable
			} else {
				return "", fmt.Errorf("unknown variable for '%s': '%s'", line, evaluatedIdentifier)
			}
		}

	}

	for placeholder, substitution := range substitutions {
		line = strings.ReplaceAll(line, placeholder, substitution)
	}

	return line, nil
}

func parseVariable(line string, variables map[string]*Variable) (*Variable, error) {
	groups := variableDeclarationRegex.FindStringSubmatch(line)
	export := false

	// Starts with export?
	if groups[1] != "" {
		export = true
	}

	// Identifier
	identifier := groups[2]

	// Evaluation type
	var evaluationType VariableEvaluationType
	switch groups[3] {
	case "=", ":=":
		evaluationType = RECURSIVE
	case "?=":
		evaluationType = FALLBACK
	default:
		return nil, fmt.Errorf("invalid variable evaluation type '%b'", line[3])
	}

	// Value
	value, err := evaluateVariables(groups[4], variables)
	if err != nil {
		return nil, err
	}

	return &Variable{
		Name:           identifier,
		EvaluationType: evaluationType,
		Export:         export,
		value:          value,
	}, nil
}

func ParseSupermakeFileV2(path string) (*SupermakeFile, error) {
	dataBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	data := string(dataBytes)
	lines := strings.Split(data, "\n")

	parser := NewParser(lines)
	return parser.Parse()
}

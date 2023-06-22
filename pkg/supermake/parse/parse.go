package parse

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/KillianMeersman/Supermake/pkg/supermake"
	"github.com/KillianMeersman/Supermake/pkg/supermake/executables"
	"github.com/KillianMeersman/Supermake/pkg/supermake/executors"
)

var targetRegex = regexp.MustCompile(`^([a-zA-Z0-9-/_.@]+):(?: +(.*))?$`)
var executorRegex = regexp.MustCompile(`^\t*@([a-zA-Z0-9-_:.]+)( .*)?$`)
var variableDeclarationRegex = regexp.MustCompile(`^(export +)?(\w+) +(=|:=|\?=) +(.*)$`)
var variableRegex = regexp.MustCompile(`\$\((.*?)\)`)

const COMMENT_CHARS = "#"

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

func (p *SuperMakeFileParser) parseGlobalVariables() (map[string]*supermake.Variable, error) {
	variables := make(map[string]*supermake.Variable)

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

func (p *SuperMakeFileParser) parseTargets(variables map[string]*supermake.Variable) (map[string]executors.Runable, error) {
	targets := make(map[string]executors.Runable)

	for i := 0; i < len(p.lines); i++ {
		line := p.lines[i]

		if targetRegex.MatchString(line) {
			// If not at the last line
			targetEnd := i
			if i+1 <= len(p.indent) {
				targetEnd = findBlockEnd(p.indent, i)
			}
			target, err := parseTarget(p.lines[i:targetEnd], p.indent[i:targetEnd], variables)
			if err != nil {
				return nil, err
			}
			targets[target.Name()] = target
			i = targetEnd - 1
		}
	}

	return targets, nil
}

func (p *SuperMakeFileParser) Parse() (*supermake.SupermakeFile, error) {
	p.cleanAndMapLines()
	globalVariables, err := p.parseGlobalVariables()
	if err != nil {
		return nil, err
	}

	targets, err := p.parseTargets(globalVariables)
	if err != nil {
		return nil, err
	}

	return &supermake.SupermakeFile{
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
func parseTarget(lines []string, indentationLevels []int, variables map[string]*supermake.Variable) (*supermake.Target, error) {
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
	subTargets := make(map[string]*supermake.Target)

	// Steps & executors
	steps := make([]executors.Runable, 0)
	currentCommands := make([]executables.Command, 0)
	var currentExecutor executors.CommandExecutor = new(executors.LocalEnvironment)

	for i := 1; i < len(lines); i++ {
		line := lines[i]
		line, err := evaluateVariables(line, variables)
		if err != nil {
			return nil, err
		}

		// Subtargets
		if targetRegex.MatchString(line) {
			blockEnd := findBlockEnd(indentationLevels, i)
			subTarget, err := parseTarget(lines[i:blockEnd], indentationLevels[i:blockEnd], variables)
			if err != nil {
				return nil, err
			}
			subTargets[subTarget.Name()] = subTarget
			if len(currentCommands) > 0 {
				steps = append(steps, &supermake.CommandGroup{
					Environment: currentExecutor,
					Steps:       currentCommands,
				})
			}
			currentCommands = make([]executables.Command, 0)
			currentExecutor = new(executors.LocalEnvironment)
			steps = append(steps, subTarget)
			i = blockEnd - 1
			// Executors
		} else if executorRegex.MatchString(line) {
			// If commands were defined, save them in a new command group
			if len(currentCommands) > 0 {
				steps = append(steps, &supermake.CommandGroup{
					Environment: currentExecutor,
					Steps:       currentCommands,
				})
			}

			groups := executorRegex.FindStringSubmatch(line)

			// Reset commands & executor
			currentCommands = make([]executables.Command, 0)
			currentExecutor = &executors.DockerEnvironment{
				Image:      strings.TrimSpace(groups[1]),
				Entrypoint: strings.TrimSpace(groups[2]),
			}
		} else {
			// Commands
			command := &executables.ShellCommand{
				Command: line,
			}
			currentCommands = append(currentCommands, command)
		}
	}

	if len(currentCommands) > 0 {
		// Add the remaining command group
		steps = append(steps, &supermake.CommandGroup{
			Environment: currentExecutor,
			Steps:       currentCommands,
		})
	}

	target := supermake.NewTarget(name, node, dependencies, steps, variables, subTargets, nil)

	for _, t := range subTargets {
		t.Parent = target
	}

	return target, nil
}

func runShellScript(script string, variables map[string]*supermake.Variable) ([]byte, error) {
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

func evaluateVariables(line string, variables map[string]*supermake.Variable) (string, error) {
	line, err := ReplaceVariables(line, func(v string) (string, error) {
		if strings.HasPrefix(v, "shell ") {
			// Run shell command
			shellString := strings.TrimPrefix(v, "shell ")
			stdout, err := runShellScript(shellString, variables)
			if err != nil {
				return "", fmt.Errorf("variable shell command failed to execute: '%s': %s", shellString, err)
			}
			t := strings.TrimRight(string(stdout), "\n\r")
			return t, nil
		} else {
			if variable, ok := variables[v]; ok {
				return variable.Value(), nil
			} else if variable, ok := os.LookupEnv(v); ok {
				return variable, nil
			} else {
				return "", fmt.Errorf("unknown variable for '%s': '%s'", line, v)
			}
		}
	})

	if err != nil {
		return "", err
	}

	return line, nil
}

func parseVariable(line string, variables map[string]*supermake.Variable) (*supermake.Variable, error) {
	groups := variableDeclarationRegex.FindStringSubmatch(line)
	export := false

	// Starts with export?
	if groups[1] != "" {
		export = true
	}

	// Identifier
	identifier := groups[2]

	// Evaluation type
	var evaluationType supermake.VariableEvaluationType
	switch groups[3] {
	case "=", ":=":
		evaluationType = supermake.RECURSIVE
	case "?=":
		evaluationType = supermake.FALLBACK
	default:
		return nil, fmt.Errorf("invalid variable evaluation type '%b'", line[3])
	}

	// Value
	value, err := evaluateVariables(groups[4], variables)
	if err != nil {
		return nil, err
	}

	return supermake.NewVariable(identifier, evaluationType, export, value), nil
}

func ParseSupermakeFileV2(path string) (*supermake.SupermakeFile, error) {
	dataBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	data := string(dataBytes)
	lines := strings.Split(data, "\n")

	parser := NewParser(lines)
	return parser.Parse()
}

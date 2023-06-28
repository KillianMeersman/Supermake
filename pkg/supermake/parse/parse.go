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
	"github.com/KillianMeersman/Supermake/pkg/supermake/util"
)

var targetRegex = regexp.MustCompile(`^([a-zA-Z0-9-/_.@]+):(?: +(.*))?$`)
var executorRegex = regexp.MustCompile(`^([a-zA-Z0-9]*?)@([a-zA-Z0-9-_:.]+)( .*)?$`)
var variableDeclarationRegex = regexp.MustCompile(`^(export +)?(\w+) +(=|:=|\?=) +(.*)$`)

const COMMENT_CHARS = "#"

type LineType int

const (
	VARIABLE LineType = iota
	TARGET
	EXECUTOR
	COMMAND
)

type SuperMakeFileParser struct {
	i        int
	rawLines []string

	lines     []string
	indent    []int
	sourcemap []int
	lineTypes []LineType
}

func NewParser(lines []string) *SuperMakeFileParser {
	return &SuperMakeFileParser{
		i:        0,
		rawLines: lines,
	}
}

// Build a map of cleaned lines and their indentation levels and class.
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

		lineType := COMMAND
		if variableDeclarationRegex.MatchString(line) {
			lineType = VARIABLE
		} else if targetRegex.MatchString(line) {
			lineType = TARGET
		} else if executorRegex.MatchString(line) {
			lineType = EXECUTOR
		}

		p.indent = append(p.indent, indentation)
		p.lines = append(p.lines, line)
		p.sourcemap = append(p.sourcemap, i)
		p.lineTypes = append(p.lineTypes, lineType)
	}
}

func (p *SuperMakeFileParser) parseGlobalVariables() (map[string]*supermake.Variable, error) {
	variables := make(map[string]*supermake.Variable)

	for ; p.i < len(p.lines); p.i++ {
		if p.lineTypes[p.i] == VARIABLE {
			variable, err := parseVariable(p.lines[p.i], variables)
			if err != nil {
				return nil, fmt.Errorf("error on line %d: %s", p.sourcemap[p.i]+1, err)
			}
			variables[variable.Name] = variable
		}
	}

	return variables, nil
}

func (p *SuperMakeFileParser) parseTargets(variables supermake.Variables) (supermake.Targets, error) {
	targets := make(map[string]*supermake.Target)

	for p.i < len(p.lines) {
		if p.lineTypes[p.i] == TARGET {
			_, err := p.parseTarget(variables, targets, "", []string{})
			if err != nil {
				return nil, err
			}
		} else {
			p.i++
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

	p.i = 0
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
// As is standard in Go, the returned index is exclusive.
func findIndentedBlockEnd(indent []int, index int) int {
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

func getNestedTargetName(parent string, name string) string {
	if parent == "" {
		return name
	}
	return fmt.Sprintf("%s::%s", parent, name)
}

// Parse a command block, optionally starting with an executor.
func (p *SuperMakeFileParser) parseCommandBlock(variables supermake.Variables, defaultName string) (supermake.CommandExecutor, supermake.Command, string, error) {
	var executor supermake.CommandExecutor = supermake.NewLocalEnvironment()
	commands := make(supermake.CommandGroup, 0, 10)

	name := defaultName

	// Parse custom executor if required
	if p.lineTypes[p.i] == EXECUTOR {
		line, err := evaluateVariables(p.lines[p.i], variables)
		if err != nil {
			return nil, nil, defaultName, err
		}
		groups := executorRegex.FindStringSubmatch(line)

		if groups[1] != "" {
			name = groups[1]
		}

		image := strings.TrimSpace(groups[2])
		entrypoint := strings.Split(strings.TrimSpace(groups[3]), " ")

		if image == "local" {
			executor = &supermake.LocalEnvironment{
				Entrypoint: entrypoint,
			}
		} else {
			executor = &supermake.DockerEnvironment{
				Image:      image,
				Entrypoint: entrypoint,
			}
		}
		p.i++
	}

	for ; p.i < len(p.lines); p.i++ {
		if p.lineTypes[p.i] != COMMAND {
			break
		}

		line, err := evaluateVariables(p.lines[p.i], variables)
		if err != nil {
			return nil, nil, name, err
		}
		commands = append(commands, line)
	}

	return executor, commands, name, nil
}

// Parse a target.
// Assumes the target declaration is the first line.
func (p *SuperMakeFileParser) parseTarget(variables supermake.Variables, targets map[string]*supermake.Target, parent string, deps []string) (*supermake.Target, error) {
	headerParts := strings.SplitN(p.lines[p.i], ":", 2)

	// Name & Node
	nameParts := strings.SplitN(headerParts[0], "@", 2)
	name := strings.TrimSpace(nameParts[0])

	if parent != "" {
		name = fmt.Sprintf("%s::%s", parent, name)
	}

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
				dependencies = append(dependencies, getNestedTargetName(parent, strings.TrimSpace(dep)))
			}
		}
	}
	dependencies = append(dependencies, deps...)

	// Keeps track of consecutive targets, so they do not depend on eachother.
	blockDependencies := make([]string, 0)

	end := findIndentedBlockEnd(p.indent, p.i)
	step := 0

	p.i++
	for p.i < end {
		// Subtargets
		if p.lineTypes[p.i] == TARGET {
			subTarget, err := p.parseTarget(variables, targets, name, dependencies)
			if err != nil {
				return nil, err
			}

			targets[subTarget.FQN()] = subTarget
			blockDependencies = append(blockDependencies, subTarget.FQN())
		} else if p.lineTypes[p.i] == EXECUTOR || p.lineTypes[p.i] == COMMAND {
			dependencies = append(dependencies, blockDependencies...)
			blockDependencies = make([]string, 0)

			executor, commands, blockName, err := p.parseCommandBlock(variables, fmt.Sprintf("%d", step))
			if err != nil {
				return nil, err
			}

			target := supermake.NewTarget(fmt.Sprintf("%s::%s", name, blockName), node, []string{}, executor, []supermake.Command{commands}, variables)
			target.Dependencies = append(target.Dependencies, dependencies...)
			targets[target.FQN()] = target
			dependencies = append(dependencies, target.FQN())
			step++
		} else {
			p.i++
		}
	}
	dependencies = append(dependencies, blockDependencies...)
	p.i = end

	target := supermake.NewTarget(name, node, dependencies, supermake.NewLocalEnvironment(), []supermake.Command{}, variables)
	targets[target.FQN()] = target

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
	return ParseSupermakeString(data)
}

func ParseSupermakeString(data string) (*supermake.SupermakeFile, error) {
	lines := util.SplitLines(data)

	parser := NewParser(lines)
	return parser.Parse()
}

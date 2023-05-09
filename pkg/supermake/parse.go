package supermake

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/KillianMeersman/Supermake/internal/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type SupermakeFile struct {
	Targets map[string]Target
}

func (s *SupermakeFile) Run(target string) error {
	return nil
}

type SuperMakeCustomListener struct {
	*parser.BaseSuperMakeFileListener
}

func NewSuperMakeCustomListener() *SuperMakeCustomListener {
	return new(SuperMakeCustomListener)
}

// EnterSupermakefile is called when entering the supermakefile production.
func (l *SuperMakeCustomListener) EnterSupermakefile(c *parser.SupermakefileContext) {
	fmt.Println("Entering Supermake file")
}

// EnterVariable is called when entering the variable production.
func (l *SuperMakeCustomListener) EnterVariable(c *parser.VariableContext) {
	fmt.Println("Entering variable", c.WORD())
}

// ExitSupermakefile is called when exiting the supermakefile production.
func (l *SuperMakeCustomListener) ExitSupermakefile(c *parser.SupermakefileContext) {
	fmt.Println("Exiting Supermake file")
}

func (l *SuperMakeCustomListener) ExitVariable(c *parser.VariableContext) {
	fmt.Println("Exiting variable")
}

func (l *SuperMakeCustomListener) EnterTarget(c *parser.TargetContext) {
	fmt.Println("Entering target", c.WORD())
}

func (l *SuperMakeCustomListener) EnterExecutor_line(c *parser.Executor_lineContext) {
	fmt.Println("Executor:", c.EXECUTOR())
}

func (l *SuperMakeCustomListener) EnterDependencies(c *parser.DependenciesContext) {
	fmt.Println("Dependencies:", c.AllWORD())
}

func (l *SuperMakeCustomListener) ExitTarget(c *parser.TargetContext) {
	fmt.Println("Exiting target", c.WORD())
}

func (l *SuperMakeCustomListener) EnterCommand_line(c *parser.Command_lineContext) {
	fmt.Println("Entering command line", c.AllWORD())
}

func ParseSupermakeFile(path string) (*SupermakeFile, error) {
	input, _ := antlr.NewFileStream(path)
	lexer := parser.NewSuperMakeFileLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSuperMakeFileParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(NewSuperMakeCustomListener(), p.Supermakefile())

	return &SupermakeFile{}, nil
}

func parseTarget(data string) (*Target, error) {
	return nil, nil
}

// Extract a nested block starting from the beginning of data.
func extractBlock(data string) {

}

func ParseSupermakeFileV2(path string) (*SupermakeFile, error) {
	// Regular expressions
	target_re := regexp.MustCompile("\t*([a-zA-Z0-9_]+):$")
	// executor_re := regexp.MustCompile("@(\\w+)(:\\w+)? (\\w*)?$")

	target_stack := make([]string, 3)
	// targets := make(map[string]*Target)

	dataBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	data := string(dataBytes)
	lines := strings.Split(data, "\n")
	indentationLevels := make([]int, len(lines))

	for i, line := range lines {
		// Find last tab character not preceded by other characters.
		lastIndex := strings.LastIndex(line, "\t")
		for lastIndex > 0 && line[lastIndex-1] != '\t' {
			lastIndex = strings.LastIndex(line[:lastIndex], "\t")
		}
		indentationLevels[i] = lastIndex + 1
	}

	currentIndentationLevel := 0
	for i, line := range lines {
		if target_re.MatchString(line) {
			targetData := strings.Builder{}
			for indentationLevels[i] >= currentIndentationLevel {

			}
			target_stack = append(target_stack, line)
		}
	}

	return &SupermakeFile{}, nil
}

// Code generated from SuperMakeFile.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // SuperMakeFile

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseSuperMakeFileListener is a complete listener for a parse tree produced by SuperMakeFileParser.
type BaseSuperMakeFileListener struct{}

var _ SuperMakeFileListener = &BaseSuperMakeFileListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSuperMakeFileListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSuperMakeFileListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSuperMakeFileListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSuperMakeFileListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSupermakefile is called when production supermakefile is entered.
func (s *BaseSuperMakeFileListener) EnterSupermakefile(ctx *SupermakefileContext) {}

// ExitSupermakefile is called when production supermakefile is exited.
func (s *BaseSuperMakeFileListener) ExitSupermakefile(ctx *SupermakefileContext) {}

// EnterLine is called when production line is entered.
func (s *BaseSuperMakeFileListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaseSuperMakeFileListener) ExitLine(ctx *LineContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseSuperMakeFileListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseSuperMakeFileListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseSuperMakeFileListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseSuperMakeFileListener) ExitVariable(ctx *VariableContext) {}

// EnterTarget is called when production target is entered.
func (s *BaseSuperMakeFileListener) EnterTarget(ctx *TargetContext) {}

// ExitTarget is called when production target is exited.
func (s *BaseSuperMakeFileListener) ExitTarget(ctx *TargetContext) {}

// EnterDependencies is called when production dependencies is entered.
func (s *BaseSuperMakeFileListener) EnterDependencies(ctx *DependenciesContext) {}

// ExitDependencies is called when production dependencies is exited.
func (s *BaseSuperMakeFileListener) ExitDependencies(ctx *DependenciesContext) {}

// EnterRecipe is called when production recipe is entered.
func (s *BaseSuperMakeFileListener) EnterRecipe(ctx *RecipeContext) {}

// ExitRecipe is called when production recipe is exited.
func (s *BaseSuperMakeFileListener) ExitRecipe(ctx *RecipeContext) {}

// EnterExecutor_line is called when production executor_line is entered.
func (s *BaseSuperMakeFileListener) EnterExecutor_line(ctx *Executor_lineContext) {}

// ExitExecutor_line is called when production executor_line is exited.
func (s *BaseSuperMakeFileListener) ExitExecutor_line(ctx *Executor_lineContext) {}

// EnterCommand_line is called when production command_line is entered.
func (s *BaseSuperMakeFileListener) EnterCommand_line(ctx *Command_lineContext) {}

// ExitCommand_line is called when production command_line is exited.
func (s *BaseSuperMakeFileListener) ExitCommand_line(ctx *Command_lineContext) {}

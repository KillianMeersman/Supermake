// Code generated from SuperMakeFile.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // SuperMakeFile

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// SuperMakeFileListener is a complete listener for a parse tree produced by SuperMakeFileParser.
type SuperMakeFileListener interface {
	antlr.ParseTreeListener

	// EnterSupermakefile is called when entering the supermakefile production.
	EnterSupermakefile(c *SupermakefileContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterTarget is called when entering the target production.
	EnterTarget(c *TargetContext)

	// EnterDependencies is called when entering the dependencies production.
	EnterDependencies(c *DependenciesContext)

	// EnterRecipe is called when entering the recipe production.
	EnterRecipe(c *RecipeContext)

	// EnterExecutor_line is called when entering the executor_line production.
	EnterExecutor_line(c *Executor_lineContext)

	// EnterCommand_line is called when entering the command_line production.
	EnterCommand_line(c *Command_lineContext)

	// ExitSupermakefile is called when exiting the supermakefile production.
	ExitSupermakefile(c *SupermakefileContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitTarget is called when exiting the target production.
	ExitTarget(c *TargetContext)

	// ExitDependencies is called when exiting the dependencies production.
	ExitDependencies(c *DependenciesContext)

	// ExitRecipe is called when exiting the recipe production.
	ExitRecipe(c *RecipeContext)

	// ExitExecutor_line is called when exiting the executor_line production.
	ExitExecutor_line(c *Executor_lineContext)

	// ExitCommand_line is called when exiting the command_line production.
	ExitCommand_line(c *Command_lineContext)
}

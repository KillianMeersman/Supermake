// Code generated from /home/killy/Code/Supermake/internal/parser/SuperMakeFile.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 11, 74, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 5, 6, 5, 31, 10, 5, 13, 5, 14, 5, 32, 3, 5, 3, 5,
	3, 6, 6, 6, 38, 10, 6, 13, 6, 14, 6, 39, 3, 7, 3, 7, 3, 7, 7, 7, 45, 10,
	7, 12, 7, 14, 7, 48, 11, 7, 3, 8, 3, 8, 6, 8, 52, 10, 8, 13, 8, 14, 8,
	53, 3, 8, 6, 8, 57, 10, 8, 13, 8, 14, 8, 58, 3, 9, 6, 9, 62, 10, 9, 13,
	9, 14, 9, 63, 3, 10, 3, 10, 7, 10, 68, 10, 10, 12, 10, 14, 10, 71, 11,
	10, 3, 11, 3, 11, 2, 2, 12, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9,
	17, 10, 19, 11, 21, 2, 3, 2, 8, 4, 2, 12, 12, 15, 15, 5, 2, 67, 92, 97,
	97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 6, 2, 50, 60, 67, 92,
	97, 97, 99, 124, 6, 2, 11, 12, 15, 15, 34, 34, 37, 38, 3, 2, 50, 59, 2,
	79, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2,
	2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2,
	2, 2, 2, 19, 3, 2, 2, 2, 3, 23, 3, 2, 2, 2, 5, 25, 3, 2, 2, 2, 7, 27, 3,
	2, 2, 2, 9, 30, 3, 2, 2, 2, 11, 37, 3, 2, 2, 2, 13, 41, 3, 2, 2, 2, 15,
	49, 3, 2, 2, 2, 17, 61, 3, 2, 2, 2, 19, 65, 3, 2, 2, 2, 21, 72, 3, 2, 2,
	2, 23, 24, 7, 63, 2, 2, 24, 4, 3, 2, 2, 2, 25, 26, 7, 60, 2, 2, 26, 6,
	3, 2, 2, 2, 27, 28, 7, 11, 2, 2, 28, 8, 3, 2, 2, 2, 29, 31, 7, 34, 2, 2,
	30, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 32, 33, 3,
	2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 35, 8, 5, 2, 2, 35, 10, 3, 2, 2, 2, 36,
	38, 9, 2, 2, 2, 37, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 37, 3, 2, 2,
	2, 39, 40, 3, 2, 2, 2, 40, 12, 3, 2, 2, 2, 41, 42, 7, 38, 2, 2, 42, 46,
	9, 3, 2, 2, 43, 45, 9, 4, 2, 2, 44, 43, 3, 2, 2, 2, 45, 48, 3, 2, 2, 2,
	46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 14, 3, 2, 2, 2, 48, 46, 3,
	2, 2, 2, 49, 51, 7, 66, 2, 2, 50, 52, 9, 5, 2, 2, 51, 50, 3, 2, 2, 2, 52,
	53, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 56, 3, 2, 2,
	2, 55, 57, 9, 4, 2, 2, 56, 55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 56,
	3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 16, 3, 2, 2, 2, 60, 62, 10, 6, 2, 2,
	61, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 63, 64, 3,
	2, 2, 2, 64, 18, 3, 2, 2, 2, 65, 69, 7, 37, 2, 2, 66, 68, 10, 2, 2, 2,
	67, 66, 3, 2, 2, 2, 68, 71, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69, 70, 3,
	2, 2, 2, 70, 20, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 72, 73, 9, 7, 2, 2, 73,
	22, 3, 2, 2, 2, 10, 2, 32, 39, 46, 53, 58, 63, 69, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='", "':'", "'\t'",
}

var lexerSymbolicNames = []string{
	"", "", "", "INDENT", "SPACES", "NEWLINE", "VAR", "EXECUTOR", "WORD", "COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "INDENT", "SPACES", "NEWLINE", "VAR", "EXECUTOR", "WORD",
	"COMMENT", "DIGIT",
}

type SuperMakeFileLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewSuperMakeFileLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *SuperMakeFileLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewSuperMakeFileLexer(input antlr.CharStream) *SuperMakeFileLexer {
	l := new(SuperMakeFileLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "SuperMakeFile.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SuperMakeFileLexer tokens.
const (
	SuperMakeFileLexerT__0     = 1
	SuperMakeFileLexerT__1     = 2
	SuperMakeFileLexerINDENT   = 3
	SuperMakeFileLexerSPACES   = 4
	SuperMakeFileLexerNEWLINE  = 5
	SuperMakeFileLexerVAR      = 6
	SuperMakeFileLexerEXECUTOR = 7
	SuperMakeFileLexerWORD     = 8
	SuperMakeFileLexerCOMMENT  = 9
)

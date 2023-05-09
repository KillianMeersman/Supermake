// Code generated from /home/killy/Code/Supermake/SuperMakeGrammar by ANTLR 4.9.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 12, 65, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 6, 4, 32, 10, 4, 13, 4, 14, 4, 33, 3,
	5, 6, 5, 37, 10, 5, 13, 5, 14, 5, 38, 3, 6, 6, 6, 42, 10, 6, 13, 6, 14,
	6, 43, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11,
	7, 11, 56, 10, 11, 12, 11, 14, 11, 59, 11, 11, 3, 12, 5, 12, 62, 10, 12,
	3, 12, 3, 12, 2, 2, 13, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17,
	10, 19, 11, 21, 12, 23, 2, 3, 2, 6, 4, 2, 12, 12, 15, 15, 3, 2, 11, 11,
	5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 2,
	68, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2,
	2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2,
	2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 3, 25, 3, 2, 2, 2, 5, 27, 3,
	2, 2, 2, 7, 31, 3, 2, 2, 2, 9, 36, 3, 2, 2, 2, 11, 41, 3, 2, 2, 2, 13,
	45, 3, 2, 2, 2, 15, 47, 3, 2, 2, 2, 17, 49, 3, 2, 2, 2, 19, 51, 3, 2, 2,
	2, 21, 53, 3, 2, 2, 2, 23, 61, 3, 2, 2, 2, 25, 26, 10, 2, 2, 2, 26, 4,
	3, 2, 2, 2, 27, 28, 7, 38, 2, 2, 28, 29, 5, 21, 11, 2, 29, 6, 3, 2, 2,
	2, 30, 32, 10, 2, 2, 2, 31, 30, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 31,
	3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 8, 3, 2, 2, 2, 35, 37, 5, 3, 2, 2,
	36, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 38, 39, 3,
	2, 2, 2, 39, 10, 3, 2, 2, 2, 40, 42, 5, 3, 2, 2, 41, 40, 3, 2, 2, 2, 42,
	43, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 12, 3, 2, 2,
	2, 45, 46, 7, 63, 2, 2, 46, 14, 3, 2, 2, 2, 47, 48, 7, 60, 2, 2, 48, 16,
	3, 2, 2, 2, 49, 50, 7, 37, 2, 2, 50, 18, 3, 2, 2, 2, 51, 52, 9, 3, 2, 2,
	52, 20, 3, 2, 2, 2, 53, 57, 9, 4, 2, 2, 54, 56, 9, 5, 2, 2, 55, 54, 3,
	2, 2, 2, 56, 59, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58,
	22, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 60, 62, 7, 15, 2, 2, 61, 60, 3, 2,
	2, 2, 61, 62, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 64, 7, 12, 2, 2, 64,
	24, 3, 2, 2, 2, 8, 2, 33, 38, 43, 57, 61, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "", "", "", "", "'='", "':'", "'#'",
}

var lexerSymbolicNames = []string{
	"", "CHAR", "VARNAME", "VARVALUE", "TARGET", "DEP", "ASSIGNMENT", "COLON",
	"HASH", "TAB", "ID",
}

var lexerRuleNames = []string{
	"CHAR", "VARNAME", "VARVALUE", "TARGET", "DEP", "ASSIGNMENT", "COLON",
	"HASH", "TAB", "ID", "NEWLINE",
}

type SuperMakeGrammarLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewSuperMakeGrammarLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *SuperMakeGrammarLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewSuperMakeGrammarLexer(input antlr.CharStream) *SuperMakeGrammarLexer {
	l := new(SuperMakeGrammarLexer)
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
	l.GrammarFileName = "SuperMakeGrammar"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SuperMakeGrammarLexer tokens.
const (
	SuperMakeGrammarLexerCHAR       = 1
	SuperMakeGrammarLexerVARNAME    = 2
	SuperMakeGrammarLexerVARVALUE   = 3
	SuperMakeGrammarLexerTARGET     = 4
	SuperMakeGrammarLexerDEP        = 5
	SuperMakeGrammarLexerASSIGNMENT = 6
	SuperMakeGrammarLexerCOLON      = 7
	SuperMakeGrammarLexerHASH       = 8
	SuperMakeGrammarLexerTAB        = 9
	SuperMakeGrammarLexerID         = 10
)

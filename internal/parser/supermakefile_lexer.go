// Code generated from SuperMakeFile.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser

import (
	"fmt"
  "sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter


type SuperMakeFileLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var supermakefilelexerLexerStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  channelNames           []string
  modeNames              []string
  literalNames           []string
  symbolicNames          []string
  ruleNames              []string
  predictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func supermakefilelexerLexerInit() {
  staticData := &supermakefilelexerLexerStaticData
  staticData.channelNames = []string{
    "DEFAULT_TOKEN_CHANNEL", "HIDDEN",
  }
  staticData.modeNames = []string{
    "DEFAULT_MODE",
  }
  staticData.literalNames = []string{
    "", "'='", "':'", "'\\t'",
  }
  staticData.symbolicNames = []string{
    "", "", "", "INDENT", "SPACES", "NEWLINE", "VAR", "EXECUTOR", "WORD", 
    "COMMENT",
  }
  staticData.ruleNames = []string{
    "T__0", "T__1", "INDENT", "SPACES", "NEWLINE", "VAR", "EXECUTOR", "WORD", 
    "COMMENT", "DIGIT",
  }
  staticData.predictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 9, 72, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
	4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 
	0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 4, 3, 29, 8, 3, 11, 3, 12, 3, 30, 
	1, 3, 1, 3, 1, 4, 4, 4, 36, 8, 4, 11, 4, 12, 4, 37, 1, 5, 1, 5, 1, 5, 5, 
	5, 43, 8, 5, 10, 5, 12, 5, 46, 9, 5, 1, 6, 1, 6, 4, 6, 50, 8, 6, 11, 6, 
	12, 6, 51, 1, 6, 4, 6, 55, 8, 6, 11, 6, 12, 6, 56, 1, 7, 4, 7, 60, 8, 7, 
	11, 7, 12, 7, 61, 1, 8, 1, 8, 5, 8, 66, 8, 8, 10, 8, 12, 8, 69, 9, 8, 1, 
	9, 1, 9, 0, 0, 10, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 
	9, 19, 0, 1, 0, 6, 2, 0, 10, 10, 13, 13, 3, 0, 65, 90, 95, 95, 97, 122, 
	4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 4, 0, 48, 58, 65, 90, 95, 95, 97, 
	122, 4, 0, 9, 10, 13, 13, 32, 32, 35, 36, 1, 0, 48, 57, 77, 0, 1, 1, 0, 
	0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 
	0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 
	0, 0, 0, 1, 21, 1, 0, 0, 0, 3, 23, 1, 0, 0, 0, 5, 25, 1, 0, 0, 0, 7, 28, 
	1, 0, 0, 0, 9, 35, 1, 0, 0, 0, 11, 39, 1, 0, 0, 0, 13, 47, 1, 0, 0, 0, 
	15, 59, 1, 0, 0, 0, 17, 63, 1, 0, 0, 0, 19, 70, 1, 0, 0, 0, 21, 22, 5, 
	61, 0, 0, 22, 2, 1, 0, 0, 0, 23, 24, 5, 58, 0, 0, 24, 4, 1, 0, 0, 0, 25, 
	26, 5, 9, 0, 0, 26, 6, 1, 0, 0, 0, 27, 29, 5, 32, 0, 0, 28, 27, 1, 0, 0, 
	0, 29, 30, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 32, 
	1, 0, 0, 0, 32, 33, 6, 3, 0, 0, 33, 8, 1, 0, 0, 0, 34, 36, 7, 0, 0, 0, 
	35, 34, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 38, 1, 
	0, 0, 0, 38, 10, 1, 0, 0, 0, 39, 40, 5, 36, 0, 0, 40, 44, 7, 1, 0, 0, 41, 
	43, 7, 2, 0, 0, 42, 41, 1, 0, 0, 0, 43, 46, 1, 0, 0, 0, 44, 42, 1, 0, 0, 
	0, 44, 45, 1, 0, 0, 0, 45, 12, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 47, 49, 
	5, 64, 0, 0, 48, 50, 7, 3, 0, 0, 49, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 
	51, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 54, 1, 0, 0, 0, 53, 55, 7, 
	2, 0, 0, 54, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 56, 
	57, 1, 0, 0, 0, 57, 14, 1, 0, 0, 0, 58, 60, 8, 4, 0, 0, 59, 58, 1, 0, 0, 
	0, 60, 61, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 16, 
	1, 0, 0, 0, 63, 67, 5, 35, 0, 0, 64, 66, 8, 0, 0, 0, 65, 64, 1, 0, 0, 0, 
	66, 69, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 18, 1, 
	0, 0, 0, 69, 67, 1, 0, 0, 0, 70, 71, 7, 5, 0, 0, 71, 20, 1, 0, 0, 0, 8, 
	0, 30, 37, 44, 51, 56, 61, 67, 1, 6, 0, 0,
}
  deserializer := antlr.NewATNDeserializer(nil)
  staticData.atn = deserializer.Deserialize(staticData.serializedATN)
  atn := staticData.atn
  staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
  decisionToDFA := staticData.decisionToDFA
  for index, state := range atn.DecisionToState {
    decisionToDFA[index] = antlr.NewDFA(state, index)
  }
}

// SuperMakeFileLexerInit initializes any static state used to implement SuperMakeFileLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSuperMakeFileLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SuperMakeFileLexerInit() {
  staticData := &supermakefilelexerLexerStaticData
  staticData.once.Do(supermakefilelexerLexerInit)
}

// NewSuperMakeFileLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSuperMakeFileLexer(input antlr.CharStream) *SuperMakeFileLexer {
  SuperMakeFileLexerInit()
	l := new(SuperMakeFileLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
  staticData := &supermakefilelexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "SuperMakeFile.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SuperMakeFileLexer tokens.
const (
	SuperMakeFileLexerT__0 = 1
	SuperMakeFileLexerT__1 = 2
	SuperMakeFileLexerINDENT = 3
	SuperMakeFileLexerSPACES = 4
	SuperMakeFileLexerNEWLINE = 5
	SuperMakeFileLexerVAR = 6
	SuperMakeFileLexerEXECUTOR = 7
	SuperMakeFileLexerWORD = 8
	SuperMakeFileLexerCOMMENT = 9
)


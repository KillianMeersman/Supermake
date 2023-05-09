// Code generated from /home/killy/Code/Supermake/SuperMakeGrammar by ANTLR 4.9.2. DO NOT EDIT.

package parser // SuperMakeGrammar

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 13, 76, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 6, 2, 22, 10, 2, 13, 2, 14, 2,
	23, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 30, 10, 3, 3, 3, 3, 3, 3, 4, 3, 4, 6,
	4, 36, 10, 4, 13, 4, 14, 4, 37, 3, 5, 6, 5, 41, 10, 5, 13, 5, 14, 5, 42,
	3, 6, 5, 6, 46, 10, 6, 3, 6, 3, 6, 3, 6, 5, 6, 51, 10, 6, 3, 6, 5, 6, 54,
	10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 6, 8, 61, 10, 8, 13, 8, 14, 8, 62,
	3, 9, 3, 9, 7, 9, 67, 10, 9, 12, 9, 14, 9, 70, 11, 9, 3, 9, 3, 9, 3, 10,
	3, 10, 3, 10, 2, 2, 11, 2, 4, 6, 8, 10, 12, 14, 16, 18, 2, 2, 2, 77, 2,
	21, 3, 2, 2, 2, 4, 29, 3, 2, 2, 2, 6, 33, 3, 2, 2, 2, 8, 40, 3, 2, 2, 2,
	10, 45, 3, 2, 2, 2, 12, 55, 3, 2, 2, 2, 14, 60, 3, 2, 2, 2, 16, 64, 3,
	2, 2, 2, 18, 73, 3, 2, 2, 2, 20, 22, 5, 4, 3, 2, 21, 20, 3, 2, 2, 2, 22,
	23, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 3, 3, 2, 2,
	2, 25, 30, 5, 6, 4, 2, 26, 30, 5, 10, 6, 2, 27, 30, 5, 16, 9, 2, 28, 30,
	5, 18, 10, 2, 29, 25, 3, 2, 2, 2, 29, 26, 3, 2, 2, 2, 29, 27, 3, 2, 2,
	2, 29, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 32, 7, 13, 2, 2, 32, 5,
	3, 2, 2, 2, 33, 35, 7, 11, 2, 2, 34, 36, 5, 8, 5, 2, 35, 34, 3, 2, 2, 2,
	36, 37, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 7, 3, 2,
	2, 2, 39, 41, 7, 3, 2, 2, 40, 39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 40,
	3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 9, 3, 2, 2, 2, 44, 46, 5, 12, 7, 2,
	45, 44, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 48, 7,
	6, 2, 2, 48, 50, 7, 9, 2, 2, 49, 51, 5, 14, 8, 2, 50, 49, 3, 2, 2, 2, 50,
	51, 3, 2, 2, 2, 51, 53, 3, 2, 2, 2, 52, 54, 5, 6, 4, 2, 53, 52, 3, 2, 2,
	2, 53, 54, 3, 2, 2, 2, 54, 11, 3, 2, 2, 2, 55, 56, 7, 4, 2, 2, 56, 57,
	7, 8, 2, 2, 57, 58, 7, 5, 2, 2, 58, 13, 3, 2, 2, 2, 59, 61, 7, 7, 2, 2,
	60, 59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 63, 3,
	2, 2, 2, 63, 15, 3, 2, 2, 2, 64, 68, 7, 10, 2, 2, 65, 67, 7, 3, 2, 2, 66,
	65, 3, 2, 2, 2, 67, 70, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 69, 3, 2, 2,
	2, 69, 71, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 71, 72, 7, 13, 2, 2, 72, 17,
	3, 2, 2, 2, 73, 74, 7, 13, 2, 2, 74, 19, 3, 2, 2, 2, 11, 23, 29, 37, 42,
	45, 50, 53, 62, 68,
}
var literalNames = []string{
	"", "", "", "", "", "", "'='", "':'", "'#'",
}
var symbolicNames = []string{
	"", "CHAR", "VARNAME", "VARVALUE", "TARGET", "DEP", "ASSIGNMENT", "COLON",
	"HASH", "TAB", "ID", "NEWLINE",
}

var ruleNames = []string{
	"makefile", "line", "recipe", "recipeLine", "rule", "variable", "dependencies",
	"comment", "blank",
}

type SuperMakeGrammarParser struct {
	*antlr.BaseParser
}

// NewSuperMakeGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *SuperMakeGrammarParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewSuperMakeGrammarParser(input antlr.TokenStream) *SuperMakeGrammarParser {
	this := new(SuperMakeGrammarParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "SuperMakeGrammar"

	return this
}

// SuperMakeGrammarParser tokens.
const (
	SuperMakeGrammarParserEOF        = antlr.TokenEOF
	SuperMakeGrammarParserCHAR       = 1
	SuperMakeGrammarParserVARNAME    = 2
	SuperMakeGrammarParserVARVALUE   = 3
	SuperMakeGrammarParserTARGET     = 4
	SuperMakeGrammarParserDEP        = 5
	SuperMakeGrammarParserASSIGNMENT = 6
	SuperMakeGrammarParserCOLON      = 7
	SuperMakeGrammarParserHASH       = 8
	SuperMakeGrammarParserTAB        = 9
	SuperMakeGrammarParserID         = 10
	SuperMakeGrammarParserNEWLINE    = 11
)

// SuperMakeGrammarParser rules.
const (
	SuperMakeGrammarParserRULE_makefile     = 0
	SuperMakeGrammarParserRULE_line         = 1
	SuperMakeGrammarParserRULE_recipe       = 2
	SuperMakeGrammarParserRULE_recipeLine   = 3
	SuperMakeGrammarParserRULE_rule         = 4
	SuperMakeGrammarParserRULE_variable     = 5
	SuperMakeGrammarParserRULE_dependencies = 6
	SuperMakeGrammarParserRULE_comment      = 7
	SuperMakeGrammarParserRULE_blank        = 8
)

// IMakefileContext is an interface to support dynamic dispatch.
type IMakefileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMakefileContext differentiates from other interfaces.
	IsMakefileContext()
}

type MakefileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMakefileContext() *MakefileContext {
	var p = new(MakefileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SuperMakeGrammarParserRULE_makefile
	return p
}

func (*MakefileContext) IsMakefileContext() {}

func NewMakefileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MakefileContext {
	var p = new(MakefileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SuperMakeGrammarParserRULE_makefile

	return p
}

func (s *MakefileContext) GetParser() antlr.Parser { return s.parser }

func (s *MakefileContext) AllLine() []ILineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineContext)(nil)).Elem())
	var tst = make([]ILineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineContext)
		}
	}

	return tst
}

func (s *MakefileContext) Line(i int) ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *MakefileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MakefileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SuperMakeGrammarParser) Makefile() (localctx IMakefileContext) {
	localctx = NewMakefileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SuperMakeGrammarParserRULE_makefile)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SuperMakeGrammarParserVARNAME)|(1<<SuperMakeGrammarParserTARGET)|(1<<SuperMakeGrammarParserHASH)|(1<<SuperMakeGrammarParserTAB)|(1<<SuperMakeGrammarParserNEWLINE))) != 0) {
		{
			p.SetState(18)
			p.Line()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SuperMakeGrammarParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SuperMakeGrammarParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(SuperMakeGrammarParserNEWLINE, 0)
}

func (s *LineContext) Recipe() IRecipeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecipeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRecipeContext)
}

func (s *LineContext) Rule() IRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuleContext)
}

func (s *LineContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *LineContext) Blank() IBlankContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlankContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlankContext)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SuperMakeGrammarParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SuperMakeGrammarParserRULE_line)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SuperMakeGrammarParserTAB:
		{
			p.SetState(23)
			p.Recipe()
		}

	case SuperMakeGrammarParserVARNAME, SuperMakeGrammarParserTARGET:
		{
			p.SetState(24)
			p.Rule()
		}

	case SuperMakeGrammarParserHASH:
		{
			p.SetState(25)
			p.Comment()
		}

	case SuperMakeGrammarParserNEWLINE:
		{
			p.SetState(26)
			p.Blank()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(29)
		p.Match(SuperMakeGrammarParserNEWLINE)
	}

	return localctx
}

// IRecipeContext is an interface to support dynamic dispatch.
type IRecipeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRecipeContext differentiates from other interfaces.
	IsRecipeContext()
}

type RecipeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecipeContext() *RecipeContext {
	var p = new(RecipeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SuperMakeGrammarParserRULE_recipe
	return p
}

func (*RecipeContext) IsRecipeContext() {}

func NewRecipeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecipeContext {
	var p = new(RecipeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SuperMakeGrammarParserRULE_recipe

	return p
}

func (s *RecipeContext) GetParser() antlr.Parser { return s.parser }

func (s *RecipeContext) TAB() antlr.TerminalNode {
	return s.GetToken(SuperMakeGrammarParserTAB, 0)
}

func (s *RecipeContext) AllRecipeLine() []IRecipeLineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRecipeLineContext)(nil)).Elem())
	var tst = make([]IRecipeLineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRecipeLineContext)
		}
	}

	return tst
}

func (s *RecipeContext) RecipeLine(i int) IRecipeLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecipeLineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRecipeLineContext)
}

func (s *RecipeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecipeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SuperMakeGrammarParser) Recipe() (localctx IRecipeContext) {
	localctx = NewRecipeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SuperMakeGrammarParserRULE_recipe)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.Match(SuperMakeGrammarParserTAB)
	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SuperMakeGrammarParserCHAR {
		{
			p.SetState(32)
			p.RecipeLine()
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRecipeLineContext is an interface to support dynamic dispatch.
type IRecipeLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRecipeLineContext differentiates from other interfaces.
	IsRecipeLineContext()
}

type RecipeLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecipeLineContext() *RecipeLineContext {
	var p = new(RecipeLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SuperMakeGrammarParserRULE_recipeLine
	return p
}

func (*RecipeLineContext) IsRecipeLineContext() {}

func NewRecipeLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecipeLineContext {
	var p = new(RecipeLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SuperMakeGrammarParserRULE_recipeLine

	return p
}

func (s *RecipeLineContext) GetParser() antlr.Parser { return s.parser }

func (s *RecipeLineContext) AllCHAR() []antlr.TerminalNode {
	return s.GetTokens(SuperMakeGrammarParserCHAR)
}

func (s *RecipeLineContext) CHAR(i int) antlr.TerminalNode {
	return s.GetToken(SuperMakeGrammarParserCHAR, i)
}

func (s *RecipeLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecipeLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SuperMakeGrammarParser) RecipeLine() (localctx IRecipeLineContext) {
	localctx = NewRecipeLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SuperMakeGrammarParserRULE_recipeLine)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(37)
				p.Match(SuperMakeGrammarParserCHAR)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IRuleContext is an interface to support dynamic dispatch.
type IRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleContext differentiates from other interfaces.
	IsRuleContext()
}

type RuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleContext() *RuleContext {
	var p = new(RuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SuperMakeGrammarParserRULE_rule
	return p
}

func (*RuleContext) IsRuleContext() {}

func NewRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleContext {
	var p = new(RuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SuperMakeGrammarParserRULE_rule

	return p
}

func (s *RuleContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleContext) TARGET() antlr.TerminalNode {
	return s.GetToken(SuperMakeGrammarParserTARGET, 0)
}

func (s *RuleContext) COLON() antlr.TerminalNode {
	return s.GetToken(SuperMakeGrammarParserCOLON, 0)
}

func (s *RuleContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *RuleContext) Dependencies() IDependenciesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDependenciesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDependenciesContext)
}

func (s *RuleContext) Recipe() IRecipeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecipeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRecipeContext)
}

func (s *RuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SuperMakeGrammarParser) Rule() (localctx IRuleContext) {
	localctx = NewRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SuperMakeGrammarParserRULE_rule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SuperMakeGrammarParserVARNAME {
		{
			p.SetState(42)
			p.Variable()
		}

	}
	{
		p.SetState(45)
		p.Match(SuperMakeGrammarParserTARGET)
	}
	{
		p.SetState(46)
		p.Match(SuperMakeGrammarParserCOLON)
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SuperMakeGrammarParserDEP {
		{
			p.SetState(47)
			p.Dependencies()
		}

	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SuperMakeGrammarParserTAB {
		{
			p.SetState(50)
			p.Recipe()
		}

	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SuperMakeGrammarParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SuperMakeGrammarParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) VARNAME() antlr.TerminalNode {
	return s.GetToken(SuperMakeGrammarParserVARNAME, 0)
}

func (s *VariableContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(SuperMakeGrammarParserASSIGNMENT, 0)
}

func (s *VariableContext) VARVALUE() antlr.TerminalNode {
	return s.GetToken(SuperMakeGrammarParserVARVALUE, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SuperMakeGrammarParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SuperMakeGrammarParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(SuperMakeGrammarParserVARNAME)
	}
	{
		p.SetState(54)
		p.Match(SuperMakeGrammarParserASSIGNMENT)
	}
	{
		p.SetState(55)
		p.Match(SuperMakeGrammarParserVARVALUE)
	}

	return localctx
}

// IDependenciesContext is an interface to support dynamic dispatch.
type IDependenciesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDependenciesContext differentiates from other interfaces.
	IsDependenciesContext()
}

type DependenciesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDependenciesContext() *DependenciesContext {
	var p = new(DependenciesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SuperMakeGrammarParserRULE_dependencies
	return p
}

func (*DependenciesContext) IsDependenciesContext() {}

func NewDependenciesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DependenciesContext {
	var p = new(DependenciesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SuperMakeGrammarParserRULE_dependencies

	return p
}

func (s *DependenciesContext) GetParser() antlr.Parser { return s.parser }

func (s *DependenciesContext) AllDEP() []antlr.TerminalNode {
	return s.GetTokens(SuperMakeGrammarParserDEP)
}

func (s *DependenciesContext) DEP(i int) antlr.TerminalNode {
	return s.GetToken(SuperMakeGrammarParserDEP, i)
}

func (s *DependenciesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DependenciesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SuperMakeGrammarParser) Dependencies() (localctx IDependenciesContext) {
	localctx = NewDependenciesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SuperMakeGrammarParserRULE_dependencies)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SuperMakeGrammarParserDEP {
		{
			p.SetState(57)
			p.Match(SuperMakeGrammarParserDEP)
		}

		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SuperMakeGrammarParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SuperMakeGrammarParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) HASH() antlr.TerminalNode {
	return s.GetToken(SuperMakeGrammarParserHASH, 0)
}

func (s *CommentContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(SuperMakeGrammarParserNEWLINE, 0)
}

func (s *CommentContext) AllCHAR() []antlr.TerminalNode {
	return s.GetTokens(SuperMakeGrammarParserCHAR)
}

func (s *CommentContext) CHAR(i int) antlr.TerminalNode {
	return s.GetToken(SuperMakeGrammarParserCHAR, i)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SuperMakeGrammarParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SuperMakeGrammarParserRULE_comment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Match(SuperMakeGrammarParserHASH)
	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SuperMakeGrammarParserCHAR {
		{
			p.SetState(63)
			p.Match(SuperMakeGrammarParserCHAR)
		}

		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(69)
		p.Match(SuperMakeGrammarParserNEWLINE)
	}

	return localctx
}

// IBlankContext is an interface to support dynamic dispatch.
type IBlankContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlankContext differentiates from other interfaces.
	IsBlankContext()
}

type BlankContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlankContext() *BlankContext {
	var p = new(BlankContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SuperMakeGrammarParserRULE_blank
	return p
}

func (*BlankContext) IsBlankContext() {}

func NewBlankContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlankContext {
	var p = new(BlankContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SuperMakeGrammarParserRULE_blank

	return p
}

func (s *BlankContext) GetParser() antlr.Parser { return s.parser }

func (s *BlankContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(SuperMakeGrammarParserNEWLINE, 0)
}

func (s *BlankContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlankContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SuperMakeGrammarParser) Blank() (localctx IBlankContext) {
	localctx = NewBlankContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SuperMakeGrammarParserRULE_blank)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Match(SuperMakeGrammarParserNEWLINE)
	}

	return localctx
}

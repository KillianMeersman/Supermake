// Code generated from /home/killy/Code/Supermake/internal/parser/SuperMakeFile.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // SuperMakeFile

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 11, 90, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 7, 2, 23, 10, 2, 12, 2,
	14, 2, 26, 11, 2, 3, 3, 3, 3, 5, 3, 30, 10, 3, 3, 3, 3, 3, 3, 4, 3, 4,
	5, 4, 36, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 5, 6, 45, 10,
	6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 7, 7, 53, 10, 7, 12, 7, 14, 7, 56,
	11, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 63, 10, 8, 7, 8, 65, 10, 8,
	12, 8, 14, 8, 68, 11, 8, 3, 9, 3, 9, 3, 9, 3, 10, 6, 10, 74, 10, 10, 13,
	10, 14, 10, 75, 3, 10, 7, 10, 79, 10, 10, 12, 10, 14, 10, 82, 11, 10, 6,
	10, 84, 10, 10, 13, 10, 14, 10, 85, 3, 10, 3, 10, 3, 10, 2, 2, 11, 2, 4,
	6, 8, 10, 12, 14, 16, 18, 2, 2, 2, 92, 2, 24, 3, 2, 2, 2, 4, 29, 3, 2,
	2, 2, 6, 35, 3, 2, 2, 2, 8, 37, 3, 2, 2, 2, 10, 41, 3, 2, 2, 2, 12, 49,
	3, 2, 2, 2, 14, 66, 3, 2, 2, 2, 16, 69, 3, 2, 2, 2, 18, 83, 3, 2, 2, 2,
	20, 23, 5, 4, 3, 2, 21, 23, 7, 7, 2, 2, 22, 20, 3, 2, 2, 2, 22, 21, 3,
	2, 2, 2, 23, 26, 3, 2, 2, 2, 24, 22, 3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25,
	3, 3, 2, 2, 2, 26, 24, 3, 2, 2, 2, 27, 30, 5, 14, 8, 2, 28, 30, 5, 6, 4,
	2, 29, 27, 3, 2, 2, 2, 29, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 32,
	7, 7, 2, 2, 32, 5, 3, 2, 2, 2, 33, 36, 5, 8, 5, 2, 34, 36, 5, 10, 6, 2,
	35, 33, 3, 2, 2, 2, 35, 34, 3, 2, 2, 2, 36, 7, 3, 2, 2, 2, 37, 38, 7, 8,
	2, 2, 38, 39, 7, 3, 2, 2, 39, 40, 7, 10, 2, 2, 40, 9, 3, 2, 2, 2, 41, 42,
	7, 10, 2, 2, 42, 44, 7, 4, 2, 2, 43, 45, 5, 12, 7, 2, 44, 43, 3, 2, 2,
	2, 44, 45, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46, 47, 7, 7, 2, 2, 47, 48,
	5, 14, 8, 2, 48, 11, 3, 2, 2, 2, 49, 54, 7, 10, 2, 2, 50, 51, 7, 6, 2,
	2, 51, 53, 7, 10, 2, 2, 52, 50, 3, 2, 2, 2, 53, 56, 3, 2, 2, 2, 54, 52,
	3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 13, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2,
	57, 62, 7, 5, 2, 2, 58, 63, 5, 16, 9, 2, 59, 63, 5, 18, 10, 2, 60, 61,
	7, 5, 2, 2, 61, 63, 5, 10, 6, 2, 62, 58, 3, 2, 2, 2, 62, 59, 3, 2, 2, 2,
	62, 60, 3, 2, 2, 2, 63, 65, 3, 2, 2, 2, 64, 57, 3, 2, 2, 2, 65, 68, 3,
	2, 2, 2, 66, 64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 15, 3, 2, 2, 2, 68,
	66, 3, 2, 2, 2, 69, 70, 7, 9, 2, 2, 70, 71, 7, 7, 2, 2, 71, 17, 3, 2, 2,
	2, 72, 74, 7, 10, 2, 2, 73, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 73,
	3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 80, 3, 2, 2, 2, 77, 79, 7, 6, 2, 2,
	78, 77, 3, 2, 2, 2, 79, 82, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2, 80, 81, 3,
	2, 2, 2, 81, 84, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 83, 73, 3, 2, 2, 2, 84,
	85, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 87, 3, 2, 2,
	2, 87, 88, 7, 7, 2, 2, 88, 19, 3, 2, 2, 2, 13, 22, 24, 29, 35, 44, 54,
	62, 66, 75, 80, 85,
}
var literalNames = []string{
	"", "'='", "':'", "'\t'",
}
var symbolicNames = []string{
	"", "", "", "INDENT", "SPACES", "NEWLINE", "VAR", "EXECUTOR", "WORD", "COMMENT",
}

var ruleNames = []string{
	"supermakefile", "line", "declaration", "variable", "target", "dependencies",
	"recipe", "executor_line", "command_line",
}

type SuperMakeFileParser struct {
	*antlr.BaseParser
}

// NewSuperMakeFileParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *SuperMakeFileParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewSuperMakeFileParser(input antlr.TokenStream) *SuperMakeFileParser {
	this := new(SuperMakeFileParser)
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
	this.GrammarFileName = "SuperMakeFile.g4"

	return this
}

// SuperMakeFileParser tokens.
const (
	SuperMakeFileParserEOF      = antlr.TokenEOF
	SuperMakeFileParserT__0     = 1
	SuperMakeFileParserT__1     = 2
	SuperMakeFileParserINDENT   = 3
	SuperMakeFileParserSPACES   = 4
	SuperMakeFileParserNEWLINE  = 5
	SuperMakeFileParserVAR      = 6
	SuperMakeFileParserEXECUTOR = 7
	SuperMakeFileParserWORD     = 8
	SuperMakeFileParserCOMMENT  = 9
)

// SuperMakeFileParser rules.
const (
	SuperMakeFileParserRULE_supermakefile = 0
	SuperMakeFileParserRULE_line          = 1
	SuperMakeFileParserRULE_declaration   = 2
	SuperMakeFileParserRULE_variable      = 3
	SuperMakeFileParserRULE_target        = 4
	SuperMakeFileParserRULE_dependencies  = 5
	SuperMakeFileParserRULE_recipe        = 6
	SuperMakeFileParserRULE_executor_line = 7
	SuperMakeFileParserRULE_command_line  = 8
)

// ISupermakefileContext is an interface to support dynamic dispatch.
type ISupermakefileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSupermakefileContext differentiates from other interfaces.
	IsSupermakefileContext()
}

type SupermakefileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySupermakefileContext() *SupermakefileContext {
	var p = new(SupermakefileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SuperMakeFileParserRULE_supermakefile
	return p
}

func (*SupermakefileContext) IsSupermakefileContext() {}

func NewSupermakefileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SupermakefileContext {
	var p = new(SupermakefileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SuperMakeFileParserRULE_supermakefile

	return p
}

func (s *SupermakefileContext) GetParser() antlr.Parser { return s.parser }

func (s *SupermakefileContext) AllLine() []ILineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineContext)(nil)).Elem())
	var tst = make([]ILineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineContext)
		}
	}

	return tst
}

func (s *SupermakefileContext) Line(i int) ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *SupermakefileContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(SuperMakeFileParserNEWLINE)
}

func (s *SupermakefileContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(SuperMakeFileParserNEWLINE, i)
}

func (s *SupermakefileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SupermakefileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SuperMakeFileParser) Supermakefile() (localctx ISupermakefileContext) {
	localctx = NewSupermakefileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SuperMakeFileParserRULE_supermakefile)
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
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SuperMakeFileParserINDENT)|(1<<SuperMakeFileParserNEWLINE)|(1<<SuperMakeFileParserVAR)|(1<<SuperMakeFileParserWORD))) != 0 {
		p.SetState(20)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(18)
				p.Line()
			}

		case 2:
			{
				p.SetState(19)
				p.Match(SuperMakeFileParserNEWLINE)
			}

		}

		p.SetState(24)
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
	p.RuleIndex = SuperMakeFileParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SuperMakeFileParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(SuperMakeFileParserNEWLINE, 0)
}

func (s *LineContext) Recipe() IRecipeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecipeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRecipeContext)
}

func (s *LineContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SuperMakeFileParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SuperMakeFileParserRULE_line)

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
	case SuperMakeFileParserINDENT, SuperMakeFileParserNEWLINE:
		{
			p.SetState(25)
			p.Recipe()
		}

	case SuperMakeFileParserVAR, SuperMakeFileParserWORD:
		{
			p.SetState(26)
			p.Declaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(29)
		p.Match(SuperMakeFileParserNEWLINE)
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SuperMakeFileParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SuperMakeFileParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *DeclarationContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SuperMakeFileParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SuperMakeFileParserRULE_declaration)

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

	p.SetState(33)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SuperMakeFileParserVAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(31)
			p.Variable()
		}

	case SuperMakeFileParserWORD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(32)
			p.Target()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = SuperMakeFileParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SuperMakeFileParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) VAR() antlr.TerminalNode {
	return s.GetToken(SuperMakeFileParserVAR, 0)
}

func (s *VariableContext) WORD() antlr.TerminalNode {
	return s.GetToken(SuperMakeFileParserWORD, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SuperMakeFileParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SuperMakeFileParserRULE_variable)

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
		p.SetState(35)
		p.Match(SuperMakeFileParserVAR)
	}
	{
		p.SetState(36)
		p.Match(SuperMakeFileParserT__0)
	}
	{
		p.SetState(37)
		p.Match(SuperMakeFileParserWORD)
	}

	return localctx
}

// ITargetContext is an interface to support dynamic dispatch.
type ITargetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTargetContext differentiates from other interfaces.
	IsTargetContext()
}

type TargetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetContext() *TargetContext {
	var p = new(TargetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SuperMakeFileParserRULE_target
	return p
}

func (*TargetContext) IsTargetContext() {}

func NewTargetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetContext {
	var p = new(TargetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SuperMakeFileParserRULE_target

	return p
}

func (s *TargetContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetContext) WORD() antlr.TerminalNode {
	return s.GetToken(SuperMakeFileParserWORD, 0)
}

func (s *TargetContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(SuperMakeFileParserNEWLINE, 0)
}

func (s *TargetContext) Recipe() IRecipeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecipeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRecipeContext)
}

func (s *TargetContext) Dependencies() IDependenciesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDependenciesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDependenciesContext)
}

func (s *TargetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SuperMakeFileParser) Target() (localctx ITargetContext) {
	localctx = NewTargetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SuperMakeFileParserRULE_target)
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
		p.SetState(39)
		p.Match(SuperMakeFileParserWORD)
	}
	{
		p.SetState(40)
		p.Match(SuperMakeFileParserT__1)
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SuperMakeFileParserWORD {
		{
			p.SetState(41)
			p.Dependencies()
		}

	}
	{
		p.SetState(44)
		p.Match(SuperMakeFileParserNEWLINE)
	}
	{
		p.SetState(45)
		p.Recipe()
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
	p.RuleIndex = SuperMakeFileParserRULE_dependencies
	return p
}

func (*DependenciesContext) IsDependenciesContext() {}

func NewDependenciesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DependenciesContext {
	var p = new(DependenciesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SuperMakeFileParserRULE_dependencies

	return p
}

func (s *DependenciesContext) GetParser() antlr.Parser { return s.parser }

func (s *DependenciesContext) AllWORD() []antlr.TerminalNode {
	return s.GetTokens(SuperMakeFileParserWORD)
}

func (s *DependenciesContext) WORD(i int) antlr.TerminalNode {
	return s.GetToken(SuperMakeFileParserWORD, i)
}

func (s *DependenciesContext) AllSPACES() []antlr.TerminalNode {
	return s.GetTokens(SuperMakeFileParserSPACES)
}

func (s *DependenciesContext) SPACES(i int) antlr.TerminalNode {
	return s.GetToken(SuperMakeFileParserSPACES, i)
}

func (s *DependenciesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DependenciesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SuperMakeFileParser) Dependencies() (localctx IDependenciesContext) {
	localctx = NewDependenciesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SuperMakeFileParserRULE_dependencies)
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
		p.SetState(47)
		p.Match(SuperMakeFileParserWORD)
	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SuperMakeFileParserSPACES {
		{
			p.SetState(48)
			p.Match(SuperMakeFileParserSPACES)
		}
		{
			p.SetState(49)
			p.Match(SuperMakeFileParserWORD)
		}

		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = SuperMakeFileParserRULE_recipe
	return p
}

func (*RecipeContext) IsRecipeContext() {}

func NewRecipeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecipeContext {
	var p = new(RecipeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SuperMakeFileParserRULE_recipe

	return p
}

func (s *RecipeContext) GetParser() antlr.Parser { return s.parser }

func (s *RecipeContext) AllINDENT() []antlr.TerminalNode {
	return s.GetTokens(SuperMakeFileParserINDENT)
}

func (s *RecipeContext) INDENT(i int) antlr.TerminalNode {
	return s.GetToken(SuperMakeFileParserINDENT, i)
}

func (s *RecipeContext) AllExecutor_line() []IExecutor_lineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExecutor_lineContext)(nil)).Elem())
	var tst = make([]IExecutor_lineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExecutor_lineContext)
		}
	}

	return tst
}

func (s *RecipeContext) Executor_line(i int) IExecutor_lineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExecutor_lineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExecutor_lineContext)
}

func (s *RecipeContext) AllCommand_line() []ICommand_lineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommand_lineContext)(nil)).Elem())
	var tst = make([]ICommand_lineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommand_lineContext)
		}
	}

	return tst
}

func (s *RecipeContext) Command_line(i int) ICommand_lineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommand_lineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommand_lineContext)
}

func (s *RecipeContext) AllTarget() []ITargetContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITargetContext)(nil)).Elem())
	var tst = make([]ITargetContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITargetContext)
		}
	}

	return tst
}

func (s *RecipeContext) Target(i int) ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *RecipeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecipeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SuperMakeFileParser) Recipe() (localctx IRecipeContext) {
	localctx = NewRecipeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SuperMakeFileParserRULE_recipe)

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
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(55)
				p.Match(SuperMakeFileParserINDENT)
			}
			p.SetState(60)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case SuperMakeFileParserEXECUTOR:
				{
					p.SetState(56)
					p.Executor_line()
				}

			case SuperMakeFileParserWORD:
				{
					p.SetState(57)
					p.Command_line()
				}

			case SuperMakeFileParserINDENT:
				{
					p.SetState(58)
					p.Match(SuperMakeFileParserINDENT)
				}
				{
					p.SetState(59)
					p.Target()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// IExecutor_lineContext is an interface to support dynamic dispatch.
type IExecutor_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExecutor_lineContext differentiates from other interfaces.
	IsExecutor_lineContext()
}

type Executor_lineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExecutor_lineContext() *Executor_lineContext {
	var p = new(Executor_lineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SuperMakeFileParserRULE_executor_line
	return p
}

func (*Executor_lineContext) IsExecutor_lineContext() {}

func NewExecutor_lineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Executor_lineContext {
	var p = new(Executor_lineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SuperMakeFileParserRULE_executor_line

	return p
}

func (s *Executor_lineContext) GetParser() antlr.Parser { return s.parser }

func (s *Executor_lineContext) EXECUTOR() antlr.TerminalNode {
	return s.GetToken(SuperMakeFileParserEXECUTOR, 0)
}

func (s *Executor_lineContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(SuperMakeFileParserNEWLINE, 0)
}

func (s *Executor_lineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Executor_lineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SuperMakeFileParser) Executor_line() (localctx IExecutor_lineContext) {
	localctx = NewExecutor_lineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SuperMakeFileParserRULE_executor_line)

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
		p.SetState(67)
		p.Match(SuperMakeFileParserEXECUTOR)
	}
	{
		p.SetState(68)
		p.Match(SuperMakeFileParserNEWLINE)
	}

	return localctx
}

// ICommand_lineContext is an interface to support dynamic dispatch.
type ICommand_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommand_lineContext differentiates from other interfaces.
	IsCommand_lineContext()
}

type Command_lineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommand_lineContext() *Command_lineContext {
	var p = new(Command_lineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SuperMakeFileParserRULE_command_line
	return p
}

func (*Command_lineContext) IsCommand_lineContext() {}

func NewCommand_lineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Command_lineContext {
	var p = new(Command_lineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SuperMakeFileParserRULE_command_line

	return p
}

func (s *Command_lineContext) GetParser() antlr.Parser { return s.parser }

func (s *Command_lineContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(SuperMakeFileParserNEWLINE, 0)
}

func (s *Command_lineContext) AllWORD() []antlr.TerminalNode {
	return s.GetTokens(SuperMakeFileParserWORD)
}

func (s *Command_lineContext) WORD(i int) antlr.TerminalNode {
	return s.GetToken(SuperMakeFileParserWORD, i)
}

func (s *Command_lineContext) AllSPACES() []antlr.TerminalNode {
	return s.GetTokens(SuperMakeFileParserSPACES)
}

func (s *Command_lineContext) SPACES(i int) antlr.TerminalNode {
	return s.GetToken(SuperMakeFileParserSPACES, i)
}

func (s *Command_lineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_lineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SuperMakeFileParser) Command_line() (localctx ICommand_lineContext) {
	localctx = NewCommand_lineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SuperMakeFileParserRULE_command_line)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SuperMakeFileParserWORD {
		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(70)
					p.Match(SuperMakeFileParserWORD)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(73)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
		}
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SuperMakeFileParserSPACES {
			{
				p.SetState(75)
				p.Match(SuperMakeFileParserSPACES)
			}

			p.SetState(80)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(85)
		p.Match(SuperMakeFileParserNEWLINE)
	}

	return localctx
}

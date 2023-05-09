// Code generated from SuperMakeFile.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // SuperMakeFile

import (
	"fmt"
	"strconv"
  "sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}


type SuperMakeFileParser struct {
	*antlr.BaseParser
}

var supermakefileParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  literalNames           []string
  symbolicNames          []string
  ruleNames              []string
  predictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func supermakefileParserInit() {
  staticData := &supermakefileParserStaticData
  staticData.literalNames = []string{
    "", "'='", "':'", "'\\t'",
  }
  staticData.symbolicNames = []string{
    "", "", "", "INDENT", "SPACES", "NEWLINE", "VAR", "EXECUTOR", "WORD", 
    "COMMENT",
  }
  staticData.ruleNames = []string{
    "supermakefile", "line", "declaration", "variable", "target", "dependencies", 
    "recipe", "executor_line", "command_line",
  }
  staticData.predictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 9, 88, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4, 
	2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 5, 0, 21, 8, 
	0, 10, 0, 12, 0, 24, 9, 0, 1, 1, 1, 1, 3, 1, 28, 8, 1, 1, 1, 1, 1, 1, 2, 
	1, 2, 3, 2, 34, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 3, 4, 43, 
	8, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 5, 5, 51, 8, 5, 10, 5, 12, 5, 
	54, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 61, 8, 6, 5, 6, 63, 8, 6, 
	10, 6, 12, 6, 66, 9, 6, 1, 7, 1, 7, 1, 7, 1, 8, 4, 8, 72, 8, 8, 11, 8, 
	12, 8, 73, 1, 8, 5, 8, 77, 8, 8, 10, 8, 12, 8, 80, 9, 8, 4, 8, 82, 8, 8, 
	11, 8, 12, 8, 83, 1, 8, 1, 8, 1, 8, 0, 0, 9, 0, 2, 4, 6, 8, 10, 12, 14, 
	16, 0, 0, 90, 0, 22, 1, 0, 0, 0, 2, 27, 1, 0, 0, 0, 4, 33, 1, 0, 0, 0, 
	6, 35, 1, 0, 0, 0, 8, 39, 1, 0, 0, 0, 10, 47, 1, 0, 0, 0, 12, 64, 1, 0, 
	0, 0, 14, 67, 1, 0, 0, 0, 16, 81, 1, 0, 0, 0, 18, 21, 3, 2, 1, 0, 19, 21, 
	5, 5, 0, 0, 20, 18, 1, 0, 0, 0, 20, 19, 1, 0, 0, 0, 21, 24, 1, 0, 0, 0, 
	22, 20, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 1, 1, 0, 0, 0, 24, 22, 1, 0, 
	0, 0, 25, 28, 3, 12, 6, 0, 26, 28, 3, 4, 2, 0, 27, 25, 1, 0, 0, 0, 27, 
	26, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 30, 5, 5, 0, 0, 30, 3, 1, 0, 0, 
	0, 31, 34, 3, 6, 3, 0, 32, 34, 3, 8, 4, 0, 33, 31, 1, 0, 0, 0, 33, 32, 
	1, 0, 0, 0, 34, 5, 1, 0, 0, 0, 35, 36, 5, 6, 0, 0, 36, 37, 5, 1, 0, 0, 
	37, 38, 5, 8, 0, 0, 38, 7, 1, 0, 0, 0, 39, 40, 5, 8, 0, 0, 40, 42, 5, 2, 
	0, 0, 41, 43, 3, 10, 5, 0, 42, 41, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 
	44, 1, 0, 0, 0, 44, 45, 5, 5, 0, 0, 45, 46, 3, 12, 6, 0, 46, 9, 1, 0, 0, 
	0, 47, 52, 5, 8, 0, 0, 48, 49, 5, 4, 0, 0, 49, 51, 5, 8, 0, 0, 50, 48, 
	1, 0, 0, 0, 51, 54, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 
	53, 11, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 55, 60, 5, 3, 0, 0, 56, 61, 3, 
	14, 7, 0, 57, 61, 3, 16, 8, 0, 58, 59, 5, 3, 0, 0, 59, 61, 3, 8, 4, 0, 
	60, 56, 1, 0, 0, 0, 60, 57, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 61, 63, 1, 
	0, 0, 0, 62, 55, 1, 0, 0, 0, 63, 66, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 64, 
	65, 1, 0, 0, 0, 65, 13, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 67, 68, 5, 7, 0, 
	0, 68, 69, 5, 5, 0, 0, 69, 15, 1, 0, 0, 0, 70, 72, 5, 8, 0, 0, 71, 70, 
	1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 
	74, 78, 1, 0, 0, 0, 75, 77, 5, 4, 0, 0, 76, 75, 1, 0, 0, 0, 77, 80, 1, 
	0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 82, 1, 0, 0, 0, 80, 
	78, 1, 0, 0, 0, 81, 71, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 81, 1, 0, 0, 
	0, 83, 84, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 86, 5, 5, 0, 0, 86, 17, 
	1, 0, 0, 0, 11, 20, 22, 27, 33, 42, 52, 60, 64, 73, 78, 83,
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

// SuperMakeFileParserInit initializes any static state used to implement SuperMakeFileParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSuperMakeFileParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SuperMakeFileParserInit() {
  staticData := &supermakefileParserStaticData
  staticData.once.Do(supermakefileParserInit)
}

// NewSuperMakeFileParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSuperMakeFileParser(input antlr.TokenStream) *SuperMakeFileParser {
	SuperMakeFileParserInit()
	this := new(SuperMakeFileParser)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &supermakefileParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "SuperMakeFile.g4"

	return this
}


// SuperMakeFileParser tokens.
const (
	SuperMakeFileParserEOF = antlr.TokenEOF
	SuperMakeFileParserT__0 = 1
	SuperMakeFileParserT__1 = 2
	SuperMakeFileParserINDENT = 3
	SuperMakeFileParserSPACES = 4
	SuperMakeFileParserNEWLINE = 5
	SuperMakeFileParserVAR = 6
	SuperMakeFileParserEXECUTOR = 7
	SuperMakeFileParserWORD = 8
	SuperMakeFileParserCOMMENT = 9
)

// SuperMakeFileParser rules.
const (
	SuperMakeFileParserRULE_supermakefile = 0
	SuperMakeFileParserRULE_line = 1
	SuperMakeFileParserRULE_declaration = 2
	SuperMakeFileParserRULE_variable = 3
	SuperMakeFileParserRULE_target = 4
	SuperMakeFileParserRULE_dependencies = 5
	SuperMakeFileParserRULE_recipe = 6
	SuperMakeFileParserRULE_executor_line = 7
	SuperMakeFileParserRULE_command_line = 8
)

// ISupermakefileContext is an interface to support dynamic dispatch.
type ISupermakefileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLine() []ILineContext
	Line(i int) ILineContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILineContext); ok {
			len++
		}
	}

	tst := make([]ILineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILineContext); ok {
			tst[i] = t.(ILineContext)
			i++
		}
	}

	return tst
}

func (s *SupermakefileContext) Line(i int) ILineContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

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


func (s *SupermakefileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SuperMakeFileListener); ok {
		listenerT.EnterSupermakefile(s)
	}
}

func (s *SupermakefileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SuperMakeFileListener); ok {
		listenerT.ExitSupermakefile(s)
	}
}




func (p *SuperMakeFileParser) Supermakefile() (localctx ISupermakefileContext) {
	this := p
	_ = this

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


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 360) != 0) {
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

	// Getter signatures
	NEWLINE() antlr.TerminalNode
	Recipe() IRecipeContext
	Declaration() IDeclarationContext

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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRecipeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRecipeContext)
}

func (s *LineContext) Declaration() IDeclarationContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

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


func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SuperMakeFileListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SuperMakeFileListener); ok {
		listenerT.ExitLine(s)
	}
}




func (p *SuperMakeFileParser) Line() (localctx ILineContext) {
	this := p
	_ = this

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

	// Getter signatures
	Variable() IVariableContext
	Target() ITargetContext

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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *DeclarationContext) Target() ITargetContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

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


func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SuperMakeFileListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SuperMakeFileListener); ok {
		listenerT.ExitDeclaration(s)
	}
}




func (p *SuperMakeFileParser) Declaration() (localctx IDeclarationContext) {
	this := p
	_ = this

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

	// Getter signatures
	VAR() antlr.TerminalNode
	WORD() antlr.TerminalNode

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


func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SuperMakeFileListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SuperMakeFileListener); ok {
		listenerT.ExitVariable(s)
	}
}




func (p *SuperMakeFileParser) Variable() (localctx IVariableContext) {
	this := p
	_ = this

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

	// Getter signatures
	WORD() antlr.TerminalNode
	NEWLINE() antlr.TerminalNode
	Recipe() IRecipeContext
	Dependencies() IDependenciesContext

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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRecipeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRecipeContext)
}

func (s *TargetContext) Dependencies() IDependenciesContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDependenciesContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

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


func (s *TargetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SuperMakeFileListener); ok {
		listenerT.EnterTarget(s)
	}
}

func (s *TargetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SuperMakeFileListener); ok {
		listenerT.ExitTarget(s)
	}
}




func (p *SuperMakeFileParser) Target() (localctx ITargetContext) {
	this := p
	_ = this

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

	// Getter signatures
	AllWORD() []antlr.TerminalNode
	WORD(i int) antlr.TerminalNode
	AllSPACES() []antlr.TerminalNode
	SPACES(i int) antlr.TerminalNode

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


func (s *DependenciesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SuperMakeFileListener); ok {
		listenerT.EnterDependencies(s)
	}
}

func (s *DependenciesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SuperMakeFileListener); ok {
		listenerT.ExitDependencies(s)
	}
}




func (p *SuperMakeFileParser) Dependencies() (localctx IDependenciesContext) {
	this := p
	_ = this

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

	// Getter signatures
	AllINDENT() []antlr.TerminalNode
	INDENT(i int) antlr.TerminalNode
	AllExecutor_line() []IExecutor_lineContext
	Executor_line(i int) IExecutor_lineContext
	AllCommand_line() []ICommand_lineContext
	Command_line(i int) ICommand_lineContext
	AllTarget() []ITargetContext
	Target(i int) ITargetContext

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExecutor_lineContext); ok {
			len++
		}
	}

	tst := make([]IExecutor_lineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExecutor_lineContext); ok {
			tst[i] = t.(IExecutor_lineContext)
			i++
		}
	}

	return tst
}

func (s *RecipeContext) Executor_line(i int) IExecutor_lineContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExecutor_lineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExecutor_lineContext)
}

func (s *RecipeContext) AllCommand_line() []ICommand_lineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICommand_lineContext); ok {
			len++
		}
	}

	tst := make([]ICommand_lineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICommand_lineContext); ok {
			tst[i] = t.(ICommand_lineContext)
			i++
		}
	}

	return tst
}

func (s *RecipeContext) Command_line(i int) ICommand_lineContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommand_lineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommand_lineContext)
}

func (s *RecipeContext) AllTarget() []ITargetContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITargetContext); ok {
			len++
		}
	}

	tst := make([]ITargetContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITargetContext); ok {
			tst[i] = t.(ITargetContext)
			i++
		}
	}

	return tst
}

func (s *RecipeContext) Target(i int) ITargetContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

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


func (s *RecipeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SuperMakeFileListener); ok {
		listenerT.EnterRecipe(s)
	}
}

func (s *RecipeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SuperMakeFileListener); ok {
		listenerT.ExitRecipe(s)
	}
}




func (p *SuperMakeFileParser) Recipe() (localctx IRecipeContext) {
	this := p
	_ = this

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

	// Getter signatures
	EXECUTOR() antlr.TerminalNode
	NEWLINE() antlr.TerminalNode

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


func (s *Executor_lineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SuperMakeFileListener); ok {
		listenerT.EnterExecutor_line(s)
	}
}

func (s *Executor_lineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SuperMakeFileListener); ok {
		listenerT.ExitExecutor_line(s)
	}
}




func (p *SuperMakeFileParser) Executor_line() (localctx IExecutor_lineContext) {
	this := p
	_ = this

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

	// Getter signatures
	NEWLINE() antlr.TerminalNode
	AllWORD() []antlr.TerminalNode
	WORD(i int) antlr.TerminalNode
	AllSPACES() []antlr.TerminalNode
	SPACES(i int) antlr.TerminalNode

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


func (s *Command_lineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SuperMakeFileListener); ok {
		listenerT.EnterCommand_line(s)
	}
}

func (s *Command_lineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SuperMakeFileListener); ok {
		listenerT.ExitCommand_line(s)
	}
}




func (p *SuperMakeFileParser) Command_line() (localctx ICommand_lineContext) {
	this := p
	_ = this

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



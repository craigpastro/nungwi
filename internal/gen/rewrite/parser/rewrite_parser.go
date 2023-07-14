// Code generated from Rewrite.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Rewrite

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type RewriteParser struct {
	*antlr.BaseParser
}

var RewriteParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rewriteParserInit() {
	staticData := &RewriteParserStaticData
	staticData.LiteralNames = []string{
		"", "'self'", "'computedUserset'", "'('", "')'", "'tupleToUserset'",
		"','", "'union'", "'intersection'", "'exclusion'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "ID", "WS",
	}
	staticData.RuleNames = []string{
		"term", "rewrite",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 11, 42, 2, 0, 7, 0, 2, 1, 7, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 40, 8, 1, 1, 1, 0, 0, 2, 0, 2, 0,
		0, 44, 0, 4, 1, 0, 0, 0, 2, 39, 1, 0, 0, 0, 4, 5, 3, 2, 1, 0, 5, 6, 5,
		0, 0, 1, 6, 1, 1, 0, 0, 0, 7, 40, 5, 1, 0, 0, 8, 9, 5, 2, 0, 0, 9, 10,
		5, 3, 0, 0, 10, 11, 5, 10, 0, 0, 11, 40, 5, 4, 0, 0, 12, 13, 5, 5, 0, 0,
		13, 14, 5, 3, 0, 0, 14, 15, 5, 10, 0, 0, 15, 16, 5, 6, 0, 0, 16, 17, 5,
		10, 0, 0, 17, 40, 5, 4, 0, 0, 18, 19, 5, 7, 0, 0, 19, 20, 5, 3, 0, 0, 20,
		21, 3, 2, 1, 0, 21, 22, 5, 6, 0, 0, 22, 23, 3, 2, 1, 0, 23, 24, 5, 4, 0,
		0, 24, 40, 1, 0, 0, 0, 25, 26, 5, 8, 0, 0, 26, 27, 5, 3, 0, 0, 27, 28,
		3, 2, 1, 0, 28, 29, 5, 6, 0, 0, 29, 30, 3, 2, 1, 0, 30, 31, 5, 4, 0, 0,
		31, 40, 1, 0, 0, 0, 32, 33, 5, 9, 0, 0, 33, 34, 5, 3, 0, 0, 34, 35, 3,
		2, 1, 0, 35, 36, 5, 6, 0, 0, 36, 37, 3, 2, 1, 0, 37, 38, 5, 4, 0, 0, 38,
		40, 1, 0, 0, 0, 39, 7, 1, 0, 0, 0, 39, 8, 1, 0, 0, 0, 39, 12, 1, 0, 0,
		0, 39, 18, 1, 0, 0, 0, 39, 25, 1, 0, 0, 0, 39, 32, 1, 0, 0, 0, 40, 3, 1,
		0, 0, 0, 1, 39,
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

// RewriteParserInit initializes any static state used to implement RewriteParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRewriteParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RewriteParserInit() {
	staticData := &RewriteParserStaticData
	staticData.once.Do(rewriteParserInit)
}

// NewRewriteParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRewriteParser(input antlr.TokenStream) *RewriteParser {
	RewriteParserInit()
	this := new(RewriteParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &RewriteParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Rewrite.g4"

	return this
}

// RewriteParser tokens.
const (
	RewriteParserEOF  = antlr.TokenEOF
	RewriteParserT__0 = 1
	RewriteParserT__1 = 2
	RewriteParserT__2 = 3
	RewriteParserT__3 = 4
	RewriteParserT__4 = 5
	RewriteParserT__5 = 6
	RewriteParserT__6 = 7
	RewriteParserT__7 = 8
	RewriteParserT__8 = 9
	RewriteParserID   = 10
	RewriteParserWS   = 11
)

// RewriteParser rules.
const (
	RewriteParserRULE_term    = 0
	RewriteParserRULE_rewrite = 1
)

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Rewrite() IRewriteContext
	EOF() antlr.TerminalNode

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RewriteParserRULE_term
	return p
}

func InitEmptyTermContext(p *TermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RewriteParserRULE_term
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RewriteParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Rewrite() IRewriteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRewriteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRewriteContext)
}

func (s *TermContext) EOF() antlr.TerminalNode {
	return s.GetToken(RewriteParserEOF, 0)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RewriteListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RewriteListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *RewriteParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RewriteParserRULE_term)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(4)
		p.Rewrite()
	}
	{
		p.SetState(5)
		p.Match(RewriteParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRewriteContext is an interface to support dynamic dispatch.
type IRewriteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllRewrite() []IRewriteContext
	Rewrite(i int) IRewriteContext

	// IsRewriteContext differentiates from other interfaces.
	IsRewriteContext()
}

type RewriteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRewriteContext() *RewriteContext {
	var p = new(RewriteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RewriteParserRULE_rewrite
	return p
}

func InitEmptyRewriteContext(p *RewriteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RewriteParserRULE_rewrite
}

func (*RewriteContext) IsRewriteContext() {}

func NewRewriteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RewriteContext {
	var p = new(RewriteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RewriteParserRULE_rewrite

	return p
}

func (s *RewriteContext) GetParser() antlr.Parser { return s.parser }

func (s *RewriteContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(RewriteParserID)
}

func (s *RewriteContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(RewriteParserID, i)
}

func (s *RewriteContext) AllRewrite() []IRewriteContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRewriteContext); ok {
			len++
		}
	}

	tst := make([]IRewriteContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRewriteContext); ok {
			tst[i] = t.(IRewriteContext)
			i++
		}
	}

	return tst
}

func (s *RewriteContext) Rewrite(i int) IRewriteContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRewriteContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRewriteContext)
}

func (s *RewriteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RewriteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RewriteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RewriteListener); ok {
		listenerT.EnterRewrite(s)
	}
}

func (s *RewriteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RewriteListener); ok {
		listenerT.ExitRewrite(s)
	}
}

func (p *RewriteParser) Rewrite() (localctx IRewriteContext) {
	localctx = NewRewriteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RewriteParserRULE_rewrite)
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RewriteParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(7)
			p.Match(RewriteParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RewriteParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(8)
			p.Match(RewriteParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(9)
			p.Match(RewriteParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(10)
			p.Match(RewriteParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(11)
			p.Match(RewriteParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RewriteParserT__4:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(12)
			p.Match(RewriteParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(13)
			p.Match(RewriteParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(14)
			p.Match(RewriteParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(15)
			p.Match(RewriteParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(16)
			p.Match(RewriteParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(17)
			p.Match(RewriteParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RewriteParserT__6:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(18)
			p.Match(RewriteParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(19)
			p.Match(RewriteParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(20)
			p.Rewrite()
		}
		{
			p.SetState(21)
			p.Match(RewriteParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(22)
			p.Rewrite()
		}
		{
			p.SetState(23)
			p.Match(RewriteParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RewriteParserT__7:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(25)
			p.Match(RewriteParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(26)
			p.Match(RewriteParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(27)
			p.Rewrite()
		}
		{
			p.SetState(28)
			p.Match(RewriteParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(29)
			p.Rewrite()
		}
		{
			p.SetState(30)
			p.Match(RewriteParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RewriteParserT__8:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(32)
			p.Match(RewriteParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(33)
			p.Match(RewriteParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(34)
			p.Rewrite()
		}
		{
			p.SetState(35)
			p.Match(RewriteParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(36)
			p.Rewrite()
		}
		{
			p.SetState(37)
			p.Match(RewriteParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

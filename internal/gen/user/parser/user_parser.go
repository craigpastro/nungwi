// Code generated from User.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // User

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

type UserParser struct {
	*antlr.BaseParser
}

var UserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func userParserInit() {
	staticData := &UserParserStaticData
	staticData.LiteralNames = []string{
		"", "'userset'", "'('", "','", "')'", "'object'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "ID", "WS",
	}
	staticData.RuleNames = []string{
		"term", "user",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 7, 25, 2, 0, 7, 0, 2, 1, 7, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		3, 1, 23, 8, 1, 1, 1, 0, 0, 2, 0, 2, 0, 0, 24, 0, 4, 1, 0, 0, 0, 2, 22,
		1, 0, 0, 0, 4, 5, 3, 2, 1, 0, 5, 6, 5, 0, 0, 1, 6, 1, 1, 0, 0, 0, 7, 8,
		5, 1, 0, 0, 8, 9, 5, 2, 0, 0, 9, 10, 5, 6, 0, 0, 10, 11, 5, 3, 0, 0, 11,
		12, 5, 6, 0, 0, 12, 13, 5, 3, 0, 0, 13, 14, 5, 6, 0, 0, 14, 23, 5, 4, 0,
		0, 15, 16, 5, 5, 0, 0, 16, 17, 5, 2, 0, 0, 17, 18, 5, 6, 0, 0, 18, 19,
		5, 3, 0, 0, 19, 20, 5, 6, 0, 0, 20, 23, 5, 4, 0, 0, 21, 23, 5, 6, 0, 0,
		22, 7, 1, 0, 0, 0, 22, 15, 1, 0, 0, 0, 22, 21, 1, 0, 0, 0, 23, 3, 1, 0,
		0, 0, 1, 22,
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

// UserParserInit initializes any static state used to implement UserParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewUserParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func UserParserInit() {
	staticData := &UserParserStaticData
	staticData.once.Do(userParserInit)
}

// NewUserParser produces a new parser instance for the optional input antlr.TokenStream.
func NewUserParser(input antlr.TokenStream) *UserParser {
	UserParserInit()
	this := new(UserParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &UserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "User.g4"

	return this
}

// UserParser tokens.
const (
	UserParserEOF  = antlr.TokenEOF
	UserParserT__0 = 1
	UserParserT__1 = 2
	UserParserT__2 = 3
	UserParserT__3 = 4
	UserParserT__4 = 5
	UserParserID   = 6
	UserParserWS   = 7
)

// UserParser rules.
const (
	UserParserRULE_term = 0
	UserParserRULE_user = 1
)

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	User() IUserContext
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
	p.RuleIndex = UserParserRULE_term
	return p
}

func InitEmptyTermContext(p *TermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = UserParserRULE_term
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = UserParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) User() IUserContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUserContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUserContext)
}

func (s *TermContext) EOF() antlr.TerminalNode {
	return s.GetToken(UserParserEOF, 0)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UserListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UserListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *UserParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, UserParserRULE_term)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(4)
		p.User()
	}
	{
		p.SetState(5)
		p.Match(UserParserEOF)
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

// IUserContext is an interface to support dynamic dispatch.
type IUserContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsUserContext differentiates from other interfaces.
	IsUserContext()
}

type UserContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUserContext() *UserContext {
	var p = new(UserContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = UserParserRULE_user
	return p
}

func InitEmptyUserContext(p *UserContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = UserParserRULE_user
}

func (*UserContext) IsUserContext() {}

func NewUserContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UserContext {
	var p = new(UserContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = UserParserRULE_user

	return p
}

func (s *UserContext) GetParser() antlr.Parser { return s.parser }

func (s *UserContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(UserParserID)
}

func (s *UserContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(UserParserID, i)
}

func (s *UserContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UserContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UserListener); ok {
		listenerT.EnterUser(s)
	}
}

func (s *UserContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UserListener); ok {
		listenerT.ExitUser(s)
	}
}

func (p *UserParser) User() (localctx IUserContext) {
	localctx = NewUserContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, UserParserRULE_user)
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case UserParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(7)
			p.Match(UserParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(8)
			p.Match(UserParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(9)
			p.Match(UserParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(10)
			p.Match(UserParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(11)
			p.Match(UserParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(12)
			p.Match(UserParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(13)
			p.Match(UserParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(14)
			p.Match(UserParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case UserParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(15)
			p.Match(UserParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(16)
			p.Match(UserParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(17)
			p.Match(UserParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(18)
			p.Match(UserParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(19)
			p.Match(UserParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(20)
			p.Match(UserParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case UserParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(21)
			p.Match(UserParserID)
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

// Code generated from User.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type UserLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var UserLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func userlexerLexerInit() {
	staticData := &UserLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'userset'", "'('", "','", "')'", "'object'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "ID", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "ID", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 7, 48, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 5, 4, 5, 38, 8, 5, 11, 5, 12, 5, 39, 1, 6, 4, 6, 43,
		8, 6, 11, 6, 12, 6, 44, 1, 6, 1, 6, 0, 0, 7, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 7, 1, 0, 2, 5, 0, 45, 45, 48, 57, 65, 90, 95, 95, 97, 122,
		3, 0, 9, 10, 12, 13, 32, 32, 49, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		13, 1, 0, 0, 0, 1, 15, 1, 0, 0, 0, 3, 23, 1, 0, 0, 0, 5, 25, 1, 0, 0, 0,
		7, 27, 1, 0, 0, 0, 9, 29, 1, 0, 0, 0, 11, 37, 1, 0, 0, 0, 13, 42, 1, 0,
		0, 0, 15, 16, 5, 117, 0, 0, 16, 17, 5, 115, 0, 0, 17, 18, 5, 101, 0, 0,
		18, 19, 5, 114, 0, 0, 19, 20, 5, 115, 0, 0, 20, 21, 5, 101, 0, 0, 21, 22,
		5, 116, 0, 0, 22, 2, 1, 0, 0, 0, 23, 24, 5, 40, 0, 0, 24, 4, 1, 0, 0, 0,
		25, 26, 5, 44, 0, 0, 26, 6, 1, 0, 0, 0, 27, 28, 5, 41, 0, 0, 28, 8, 1,
		0, 0, 0, 29, 30, 5, 111, 0, 0, 30, 31, 5, 98, 0, 0, 31, 32, 5, 106, 0,
		0, 32, 33, 5, 101, 0, 0, 33, 34, 5, 99, 0, 0, 34, 35, 5, 116, 0, 0, 35,
		10, 1, 0, 0, 0, 36, 38, 7, 0, 0, 0, 37, 36, 1, 0, 0, 0, 38, 39, 1, 0, 0,
		0, 39, 37, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 12, 1, 0, 0, 0, 41, 43,
		7, 1, 0, 0, 42, 41, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0,
		44, 45, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 47, 6, 6, 0, 0, 47, 14, 1,
		0, 0, 0, 3, 0, 39, 44, 1, 6, 0, 0,
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

// UserLexerInit initializes any static state used to implement UserLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewUserLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func UserLexerInit() {
	staticData := &UserLexerLexerStaticData
	staticData.once.Do(userlexerLexerInit)
}

// NewUserLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewUserLexer(input antlr.CharStream) *UserLexer {
	UserLexerInit()
	l := new(UserLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &UserLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "User.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// UserLexer tokens.
const (
	UserLexerT__0 = 1
	UserLexerT__1 = 2
	UserLexerT__2 = 3
	UserLexerT__3 = 4
	UserLexerT__4 = 5
	UserLexerID   = 6
	UserLexerWS   = 7
)

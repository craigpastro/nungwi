// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

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

type UserLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var userlexerLexerStaticData struct {
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

func userlexerLexerInit() {
	staticData := &userlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'userset'", "'('", "','", "')'", "'object'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "ID", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "ID", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 7, 50, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 5, 1, 5, 5, 5, 39, 8, 5, 10, 5, 12, 5, 42, 9, 5, 1, 6,
		4, 6, 45, 8, 6, 11, 6, 12, 6, 46, 1, 6, 1, 6, 0, 0, 7, 1, 1, 3, 2, 5, 3,
		7, 4, 9, 5, 11, 6, 13, 7, 1, 0, 3, 2, 0, 65, 90, 97, 122, 5, 0, 45, 45,
		48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 12, 13, 32, 32, 51, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 1, 15, 1, 0, 0, 0, 3,
		23, 1, 0, 0, 0, 5, 25, 1, 0, 0, 0, 7, 27, 1, 0, 0, 0, 9, 29, 1, 0, 0, 0,
		11, 36, 1, 0, 0, 0, 13, 44, 1, 0, 0, 0, 15, 16, 5, 117, 0, 0, 16, 17, 5,
		115, 0, 0, 17, 18, 5, 101, 0, 0, 18, 19, 5, 114, 0, 0, 19, 20, 5, 115,
		0, 0, 20, 21, 5, 101, 0, 0, 21, 22, 5, 116, 0, 0, 22, 2, 1, 0, 0, 0, 23,
		24, 5, 40, 0, 0, 24, 4, 1, 0, 0, 0, 25, 26, 5, 44, 0, 0, 26, 6, 1, 0, 0,
		0, 27, 28, 5, 41, 0, 0, 28, 8, 1, 0, 0, 0, 29, 30, 5, 111, 0, 0, 30, 31,
		5, 98, 0, 0, 31, 32, 5, 106, 0, 0, 32, 33, 5, 101, 0, 0, 33, 34, 5, 99,
		0, 0, 34, 35, 5, 116, 0, 0, 35, 10, 1, 0, 0, 0, 36, 40, 7, 0, 0, 0, 37,
		39, 7, 1, 0, 0, 38, 37, 1, 0, 0, 0, 39, 42, 1, 0, 0, 0, 40, 38, 1, 0, 0,
		0, 40, 41, 1, 0, 0, 0, 41, 12, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 43, 45,
		7, 2, 0, 0, 44, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0,
		46, 47, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 49, 6, 6, 0, 0, 49, 14, 1,
		0, 0, 0, 3, 0, 40, 46, 1, 6, 0, 0,
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
	staticData := &userlexerLexerStaticData
	staticData.once.Do(userlexerLexerInit)
}

// NewUserLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewUserLexer(input antlr.CharStream) *UserLexer {
	UserLexerInit()
	l := new(UserLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &userlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
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

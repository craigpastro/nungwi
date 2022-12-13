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

type RewriteLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var rewritelexerLexerStaticData struct {
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

func rewritelexerLexerInit() {
	staticData := &rewritelexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'self'", "'computedUserset'", "'('", "')'", "'tupleToUserset'",
		"','", "'union'", "'intersection'", "'exclusion'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "ID", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"ID", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 11, 108, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 5, 9, 97, 8, 9, 10, 9, 12, 9, 100, 9, 9, 1, 10, 4, 10,
		103, 8, 10, 11, 10, 12, 10, 104, 1, 10, 1, 10, 0, 0, 11, 1, 1, 3, 2, 5,
		3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 1, 0, 3, 2,
		0, 65, 90, 97, 122, 5, 0, 45, 45, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0,
		9, 10, 12, 13, 32, 32, 109, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1,
		0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 1, 23, 1, 0, 0, 0, 3, 28, 1, 0, 0, 0, 5, 44, 1, 0, 0, 0,
		7, 46, 1, 0, 0, 0, 9, 48, 1, 0, 0, 0, 11, 63, 1, 0, 0, 0, 13, 65, 1, 0,
		0, 0, 15, 71, 1, 0, 0, 0, 17, 84, 1, 0, 0, 0, 19, 94, 1, 0, 0, 0, 21, 102,
		1, 0, 0, 0, 23, 24, 5, 115, 0, 0, 24, 25, 5, 101, 0, 0, 25, 26, 5, 108,
		0, 0, 26, 27, 5, 102, 0, 0, 27, 2, 1, 0, 0, 0, 28, 29, 5, 99, 0, 0, 29,
		30, 5, 111, 0, 0, 30, 31, 5, 109, 0, 0, 31, 32, 5, 112, 0, 0, 32, 33, 5,
		117, 0, 0, 33, 34, 5, 116, 0, 0, 34, 35, 5, 101, 0, 0, 35, 36, 5, 100,
		0, 0, 36, 37, 5, 85, 0, 0, 37, 38, 5, 115, 0, 0, 38, 39, 5, 101, 0, 0,
		39, 40, 5, 114, 0, 0, 40, 41, 5, 115, 0, 0, 41, 42, 5, 101, 0, 0, 42, 43,
		5, 116, 0, 0, 43, 4, 1, 0, 0, 0, 44, 45, 5, 40, 0, 0, 45, 6, 1, 0, 0, 0,
		46, 47, 5, 41, 0, 0, 47, 8, 1, 0, 0, 0, 48, 49, 5, 116, 0, 0, 49, 50, 5,
		117, 0, 0, 50, 51, 5, 112, 0, 0, 51, 52, 5, 108, 0, 0, 52, 53, 5, 101,
		0, 0, 53, 54, 5, 84, 0, 0, 54, 55, 5, 111, 0, 0, 55, 56, 5, 85, 0, 0, 56,
		57, 5, 115, 0, 0, 57, 58, 5, 101, 0, 0, 58, 59, 5, 114, 0, 0, 59, 60, 5,
		115, 0, 0, 60, 61, 5, 101, 0, 0, 61, 62, 5, 116, 0, 0, 62, 10, 1, 0, 0,
		0, 63, 64, 5, 44, 0, 0, 64, 12, 1, 0, 0, 0, 65, 66, 5, 117, 0, 0, 66, 67,
		5, 110, 0, 0, 67, 68, 5, 105, 0, 0, 68, 69, 5, 111, 0, 0, 69, 70, 5, 110,
		0, 0, 70, 14, 1, 0, 0, 0, 71, 72, 5, 105, 0, 0, 72, 73, 5, 110, 0, 0, 73,
		74, 5, 116, 0, 0, 74, 75, 5, 101, 0, 0, 75, 76, 5, 114, 0, 0, 76, 77, 5,
		115, 0, 0, 77, 78, 5, 101, 0, 0, 78, 79, 5, 99, 0, 0, 79, 80, 5, 116, 0,
		0, 80, 81, 5, 105, 0, 0, 81, 82, 5, 111, 0, 0, 82, 83, 5, 110, 0, 0, 83,
		16, 1, 0, 0, 0, 84, 85, 5, 101, 0, 0, 85, 86, 5, 120, 0, 0, 86, 87, 5,
		99, 0, 0, 87, 88, 5, 108, 0, 0, 88, 89, 5, 117, 0, 0, 89, 90, 5, 115, 0,
		0, 90, 91, 5, 105, 0, 0, 91, 92, 5, 111, 0, 0, 92, 93, 5, 110, 0, 0, 93,
		18, 1, 0, 0, 0, 94, 98, 7, 0, 0, 0, 95, 97, 7, 1, 0, 0, 96, 95, 1, 0, 0,
		0, 97, 100, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 20,
		1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 101, 103, 7, 2, 0, 0, 102, 101, 1, 0,
		0, 0, 103, 104, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0,
		105, 106, 1, 0, 0, 0, 106, 107, 6, 10, 0, 0, 107, 22, 1, 0, 0, 0, 3, 0,
		98, 104, 1, 6, 0, 0,
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

// RewriteLexerInit initializes any static state used to implement RewriteLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewRewriteLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func RewriteLexerInit() {
	staticData := &rewritelexerLexerStaticData
	staticData.once.Do(rewritelexerLexerInit)
}

// NewRewriteLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewRewriteLexer(input antlr.CharStream) *RewriteLexer {
	RewriteLexerInit()
	l := new(RewriteLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &rewritelexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Rewrite.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// RewriteLexer tokens.
const (
	RewriteLexerT__0 = 1
	RewriteLexerT__1 = 2
	RewriteLexerT__2 = 3
	RewriteLexerT__3 = 4
	RewriteLexerT__4 = 5
	RewriteLexerT__5 = 6
	RewriteLexerT__6 = 7
	RewriteLexerT__7 = 8
	RewriteLexerT__8 = 9
	RewriteLexerID   = 10
	RewriteLexerWS   = 11
)

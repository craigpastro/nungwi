package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	rewrite "github.com/craigpastro/nungwi/internal/gen/rewrite/parser"
	user "github.com/craigpastro/nungwi/internal/gen/user/parser"
)

type errorListener struct {
	*antlr.DefaultErrorListener

	error bool
}

func (r *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	r.error = true
}

func ParseRewrite(s string) bool {
	is := antlr.NewInputStream(s)

	errorListener := &errorListener{}

	lexer := rewrite.NewRewriteLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := rewrite.NewRewriteParser(stream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(errorListener)

	antlr.NewParseTreeWalker().Walk(&rewrite.BaseRewriteListener{}, parser.Term())

	return !errorListener.error
}

func ParseUser(s string) bool {
	is := antlr.NewInputStream(s)

	errorListener := &errorListener{}

	lexer := user.NewUserLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := user.NewUserParser(stream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(errorListener)

	antlr.NewParseTreeWalker().Walk(&user.BaseUserListener{}, parser.Term())

	return !errorListener.error
}

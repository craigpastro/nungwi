// Code generated from Rewrite.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Rewrite

import "github.com/antlr4-go/antlr/v4"

// RewriteListener is a complete listener for a parse tree produced by RewriteParser.
type RewriteListener interface {
	antlr.ParseTreeListener

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterRewrite is called when entering the rewrite production.
	EnterRewrite(c *RewriteContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitRewrite is called when exiting the rewrite production.
	ExitRewrite(c *RewriteContext)
}

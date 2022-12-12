// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // User

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// UserListener is a complete listener for a parse tree produced by UserParser.
type UserListener interface {
	antlr.ParseTreeListener

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterUser is called when entering the user production.
	EnterUser(c *UserContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitUser is called when exiting the user production.
	ExitUser(c *UserContext)
}

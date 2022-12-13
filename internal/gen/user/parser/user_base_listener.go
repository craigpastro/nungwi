// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // User

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseUserListener is a complete listener for a parse tree produced by UserParser.
type BaseUserListener struct{}

var _ UserListener = &BaseUserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseUserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseUserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseUserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseUserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseUserListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseUserListener) ExitTerm(ctx *TermContext) {}

// EnterUser is called when production user is entered.
func (s *BaseUserListener) EnterUser(ctx *UserContext) {}

// ExitUser is called when production user is exited.
func (s *BaseUserListener) ExitUser(ctx *UserContext) {}

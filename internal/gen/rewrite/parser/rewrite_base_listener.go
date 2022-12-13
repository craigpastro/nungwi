// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Rewrite

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseRewriteListener is a complete listener for a parse tree produced by RewriteParser.
type BaseRewriteListener struct{}

var _ RewriteListener = &BaseRewriteListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRewriteListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRewriteListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRewriteListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRewriteListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseRewriteListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseRewriteListener) ExitTerm(ctx *TermContext) {}

// EnterRewrite is called when production rewrite is entered.
func (s *BaseRewriteListener) EnterRewrite(ctx *RewriteContext) {}

// ExitRewrite is called when production rewrite is exited.
func (s *BaseRewriteListener) ExitRewrite(ctx *RewriteContext) {}

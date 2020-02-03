// Code generated from PlayScript.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // PlayScript

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePlayScriptListener is a complete listener for a parse tree produced by PlayScriptParser.
type BasePlayScriptListener struct{}

var _ PlayScriptListener = &BasePlayScriptListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePlayScriptListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePlayScriptListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePlayScriptListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePlayScriptListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BasePlayScriptListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BasePlayScriptListener) ExitStart(ctx *StartContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BasePlayScriptListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BasePlayScriptListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BasePlayScriptListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BasePlayScriptListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BasePlayScriptListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BasePlayScriptListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

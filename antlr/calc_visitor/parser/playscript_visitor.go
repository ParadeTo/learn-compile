// Code generated from PlayScript.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // PlayScript

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by PlayScriptParser.
type PlayScriptVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PlayScriptParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#additiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#multiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}
}

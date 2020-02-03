package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"learn-compile/antlr/calc_visitor/parser"
	"strconv"
)

// visitor 模式
type playVisitor struct {
	parser.BasePlayScriptVisitor
}

func (v *playVisitor) VisitStart(ctx *parser.StartContext) interface{} {
	return v.VisitAdditiveExpression(ctx.AdditiveExpression().(*parser.AdditiveExpressionContext))
}

func (v *playVisitor) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{} {
	if ctx.ADD() != nil || ctx.SUB() != nil {
		left := v.VisitAdditiveExpression(ctx.AdditiveExpression().(*parser.AdditiveExpressionContext))
		right := v.VisitMultiplicativeExpression(ctx.MultiplicativeExpression().(*parser.MultiplicativeExpressionContext))
		if ctx.ADD() != nil {
			return left.(int) + right.(int)
		} else {
			return left.(int) - right.(int)
		}
	} else {
		return v.VisitMultiplicativeExpression(ctx.MultiplicativeExpression().(*parser.MultiplicativeExpressionContext))
	}
}

func (v *playVisitor) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) interface{} {
	if ctx.MUL() != nil || ctx.DIV() != nil {
		left := v.VisitMultiplicativeExpression(ctx.MultiplicativeExpression().(*parser.MultiplicativeExpressionContext))
		right := v.VisitPrimaryExpression(ctx.PrimaryExpression().(*parser.PrimaryExpressionContext))
		if ctx.MUL() != nil {
			return left.(int) * right.(int)
		} else {
			return left.(int) / right.(int)
		}
	} else {
		return v.VisitPrimaryExpression(ctx.PrimaryExpression().(*parser.PrimaryExpressionContext))
	}
}

func (v *playVisitor) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	if ctx.NUMBER() != nil {
		value, _ := strconv.Atoi(ctx.NUMBER().GetText())
		return value
	}
	return 0
}

func main() {
	is := antlr.NewInputStream("1+3*4")

	lexer := parser.NewPlayScriptLexer(is)
	//for {
	//	t := lexer.NextToken()
	//	if t.GetTokenType() == antlr.TokenEOF {
	//		break
	//	}
	//	fmt.Printf("%s (%q)\n",
	//		lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	//}
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewPlayScriptParser(stream)
	tree := p.Start()
	v := &playVisitor{}
	fmt.Println(tree.Accept(v))
}

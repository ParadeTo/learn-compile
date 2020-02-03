package main

// https://blog.gopheracademy.com/advent-2017/parsing-with-antlr4-and-go/
import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"learn-compile/antlr/calc/parser"
	"strconv"
)

type calcListener struct {
	*parser.BaseCalcListener
	stack []int
}

func (l *calcListener) push(i int) {
	l.stack = append(l.stack, i)
}

func (l *calcListener) pop() int {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func (l *calcListener) ExitMulDiv(c *parser.MulDivContext) {
	right, left := l.pop(), l.pop()

	switch c.GetOp().GetTokenType() {
	case parser.CalcParserMUL:
		l.push(left * right)
	case parser.CalcParserDIV:
		l.push(left / right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *calcListener) ExitAddSub(c *parser.AddSubContext) {
	right, left := l.pop(), l.pop()
	switch c.GetOp().GetTokenType() {
	case parser.CalcParserADD:
		l.push(left + right)
	case parser.CalcParserSUB:
		l.push(left - right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *calcListener) ExitNumber(c *parser.NumberContext) {
	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		panic(err.Error())
	}

	l.push(i)
}

// visitor 模式
type calcVisitor struct {
	*parser.BaseCalcVisitor
}

func (v *calcVisitor) VisitAddSub(ctx *parser.AddSubContext) interface{} {
	if ctx.ADD() != nil || ctx.SUB() != nil {
		//left := v.VisitAddSub(ctx.Expression().(*parser.AddSubContext))
		//right :=
	}
	return v.VisitChildren(ctx)
}

func (v *calcVisitor) VisitMulDiv(ctx *parser.AddSubContext) interface{} {
	if ctx.ADD() != nil || ctx.SUB() != nil {

	}
	return v.VisitChildren(ctx)
}

func (v *calcVisitor) VisitNumber(ctx *parser.AddSubContext) interface{} {
	if ctx.ADD() != nil || ctx.SUB() != nil {

	}
	return v.VisitChildren(ctx)
}

func main() {
	// Setup the input
	is := antlr.NewInputStream("1 + 2 * 3")

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)

	// Read all tokens
	//for {
	//	t := lexer.NextToken()
	//	if t.GetTokenType() == antlr.TokenEOF {
	//		break
	//	}
	//	fmt.Printf("%s (%q)\n",
	//		lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	//}

	// Setup the input

	// Create the Lexer
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalcParser(stream)
	tree := p.Start()
	// Finally parse the expression
	//listener := &calcListener{}
	v := &parser.BaseCalcVisitor{}
	//antlr.ParseTreeWalkerDefault.Walk(listener, p.Start())
	//fmt.Println(listener.pop())
	tree.Accept(v)
}

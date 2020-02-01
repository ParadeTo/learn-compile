package main

import (
	"fmt"
	"learn-compile/craft/calculator"
	"learn-compile/craft/lexer"
)

func main() {
	calc := calculator.NewSimpleCalculator()
	lexer := lexer.NewSimpleLexer()
	script := "int age;"
	tokenReader := lexer.Tokenize(script)
	_, node := calc.IntDeclare(tokenReader)
	//calc.PreTraverse(node)
	fmt.Println(calc.PreTraverse(node)[0].GetText())
}

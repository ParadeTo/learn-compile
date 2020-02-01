package main

import (
	"fmt"
	"learn-compile/craft/calculator"
)

func main() {
	calc := calculator.NewSimpleCalculator()
	script := "3+"
	fmt.Println(calc.Evaluate(script))
}

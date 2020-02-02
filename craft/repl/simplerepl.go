package repl

import (
	"errors"
	"fmt"
	"learn-compile/craft/ast"
	"learn-compile/craft/tools"
	"strconv"
)

type SimpleRepl struct {
	variables map[string]int
	verbose   bool
}

func (repl *SimpleRepl) Evaluate(node ast.ASTNode, indent string) int {
	var result int
	var err error
	if repl.verbose {
		fmt.Println(indent + "Calculating: " + string(node.GetType()))
	}

	switch node.GetType() {
	case ast.Programm:
		for _, child := range node.GetChildren() {
			result = repl.Evaluate(child, indent+tools.INDENT)
		}
	case ast.Additive:
		children := node.GetChildren()
		child1 := children[0]
		value1 := repl.Evaluate(child1, indent+tools.INDENT)
		child2 := children[1]
		value2 := repl.Evaluate(child2, indent+tools.INDENT)
		if node.GetText() == "+" {
			result = value1 + value2
		} else {
			result = value1 - value2
		}
	case ast.Multiplicative:
		children := node.GetChildren()
		child1 := children[0]
		value1 := repl.Evaluate(child1, indent+tools.INDENT)
		child2 := children[1]
		value2 := repl.Evaluate(child2, indent+tools.INDENT)
		if node.GetText() == "*" {
			result = value1 * value2
		} else {
			result = value1 / value2
		}
	case ast.IntLiteral:
		result, _ = strconv.Atoi(node.GetText())
	case ast.Identifier:
		varName := node.GetText()
		if value, ok := repl.variables[varName]; ok {
			result = value
		} else {
			err = errors.New("unknown variable: " + varName)
		}
	case ast.AssignmentStmt, ast.IntDeclaration:
		varName := node.GetText()
		if node.GetType() == ast.AssignmentStmt {
			if _, ok := repl.variables[varName]; !ok {
				err = errors.New("unknown variable: " + varName)
			}
		}
		var varValue int
		if len(node.GetChildren()) > 0 {
			child := node.GetChildren()[0]
			result = repl.Evaluate(child, indent+tools.INDENT)
			varValue = result
		}
		repl.variables[varName] = varValue
	default:
	}

	if repl.verbose {
		fmt.Printf(indent+"Result: %d\n", result)
	} else if indent == "" && err == nil {
		fmt.Println(result)
		//if node.GetType() == ast.IntDeclaration || node.GetType() == ast.AssignmentStmt {
		//	fmt.Printf(node.GetText() + ": %d\t", result)
		//} else if node.GetType() == ast.Programm {
		//	fmt.Println(result)
		//}
	}
	return result
}

func NewSimpleRepl(verbose bool) *SimpleRepl {
	return &SimpleRepl{variables: make(map[string]int), verbose: verbose}
}

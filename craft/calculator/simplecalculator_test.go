package calculator

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"learn-compile/craft/ast"
	"learn-compile/craft/lexer"
	"testing"
)

type TestDataIntDeclare struct {
	script   string
	expected []*SimpleASTNode
}

type TestDataEvaluate struct {
	script   string
	expected int
}

func TestSimpleCalculator_IntDeclare(t *testing.T) {
	lexer := lexer.NewSimpleLexer()
	calc := NewSimpleCalculator()
	testData := []TestDataIntDeclare{
		{
			script: "int age;",
			expected: []*SimpleASTNode{
				{nodeType: ast.IntDeclaration, text: "age"},
			},
		},
		{
			script: "int age = 45 + 46;",
			expected: []*SimpleASTNode{
				{nodeType: ast.IntDeclaration, text: "age"},
				{nodeType: ast.Additive, text: "+"},
				{nodeType: ast.IntLiteral, text: "45"},
				{nodeType: ast.IntLiteral, text: "46"},
			},
		},
	}
	for _, data := range testData {
		Convey(fmt.Sprintf("%s", data.script), t, func() {
			reader := lexer.Tokenize(data.script)
			_, node := calc.IntDeclare(reader)
			nodes := calc.PreTraverse(node)
			for i, node := range nodes {
				So(node.GetType(), ShouldEqual, data.expected[i].GetType())
				So(node.GetText(), ShouldEqual, data.expected[i].GetText())
			}
		})
	}
}

func TestSimpleCalculator_Evaluate(t *testing.T) {
	calc := NewSimpleCalculator()
	testData := []TestDataEvaluate{
		{
			script:   "4 + 6 * 10",
			expected: 64,
		},
		{
			script:   "(4 + 6) * 10",
			expected: 100,
		},
		{
			script:   "4 +",
			expected: 0,
		},
	}
	for _, data := range testData {
		Convey(fmt.Sprintf("%s", data.script), t, func() {
			error, result := calc.Evaluate(data.script)
			if error != nil {
				So(error, ShouldBeError)
			} else {
				So(result, ShouldEqual, data.expected)
			}
		})
	}
}

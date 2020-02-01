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

func TestSimpleCalculator_Additive(t *testing.T) {
	lexer := lexer.NewSimpleLexer()
	calc := NewSimpleCalculator()
	testData := []TestDataIntDeclare{
		{
			script: "45 + 46 * 10",
			expected: []*SimpleASTNode{
				{nodeType: ast.Additive, text: "+"},
				{nodeType: ast.IntLiteral, text: "45"},
				{nodeType: ast.Multiplicative, text: "*"},
				{nodeType: ast.IntLiteral, text: "46"},
				{nodeType: ast.IntLiteral, text: "10"},
			},
		},
		{
			script: "(45 + 46) * 10",
			expected: []*SimpleASTNode{
				{nodeType: ast.Multiplicative, text: "*"},
				{nodeType: ast.Additive, text: "+"},
				{nodeType: ast.IntLiteral, text: "45"},
				{nodeType: ast.IntLiteral, text: "46"},
				{nodeType: ast.IntLiteral, text: "10"},
			},
		},
	}
	for _, data := range testData {
		Convey(fmt.Sprintf("%s", data.script), t, func() {
			reader := lexer.Tokenize(data.script)
			_, node := calc.Additive(reader)
			nodes := calc.PreTraverse(node)
			for i, node := range nodes {
				So(node.GetType(), ShouldEqual, data.expected[i].GetType())
				So(node.GetText(), ShouldEqual, data.expected[i].GetText())
			}
		})
	}
}

package calculator

import (
	"fmt"
	"github.com/kataras/iris/core/errors"
	. "github.com/smartystreets/goconvey/convey"
	"learn-compile/craft/ast"
	"learn-compile/craft/lexer"
	"learn-compile/craft/tools"
	"testing"
)

type TestDataIntDeclare struct {
	script   string
	expected []*ast.SimpleASTNode
}

type TestDataEvaluate struct {
	script   string
	expected int
	error    error
}

func TestSimpleCalculator_IntDeclare(t *testing.T) {
	lexer := lexer.NewSimpleLexer()
	calc := NewSimpleCalculator()
	testData := []TestDataIntDeclare{
		{
			script: "int age;",
			expected: []*ast.SimpleASTNode{
				{NodeType: ast.IntDeclaration, Text: "age"},
			},
		},
		{
			script: "int age = 45 + 46;",
			expected: []*ast.SimpleASTNode{
				{NodeType: ast.IntDeclaration, Text: "age"},
				{NodeType: ast.Additive, Text: "+"},
				{NodeType: ast.IntLiteral, Text: "45"},
				{NodeType: ast.IntLiteral, Text: "46"},
			},
		},
	}
	for _, data := range testData {
		Convey(fmt.Sprintf("%s", data.script), t, func() {
			reader := lexer.Tokenize(data.script)
			_, node := calc.IntDeclare(reader)
			nodes := tools.PreTraverse(node)
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
			script:   "2 * 3 * 4",
			expected: 24,
		},
		{
			script:   "4 + 6 + 10",
			expected: 20,
		},
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
			error:    errors.New(""),
		},
	}
	for _, data := range testData {
		Convey(fmt.Sprintf("%s", data.script), t, func() {
			error, result := calc.Evaluate(data.script)
			if data.error != nil {
				So(error, ShouldBeError)
			} else {
				So(result, ShouldEqual, data.expected)
			}
		})
	}
}

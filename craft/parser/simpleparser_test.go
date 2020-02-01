package parser

import (
	"errors"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"learn-compile/craft/ast"
	"learn-compile/craft/tools"
	"testing"
)

type TestDataParse struct {
	script   string
	expected []*ast.SimpleASTNode
	error    error
}

func TestSimpleParser_Parse(t *testing.T) {
	parser := NewSimpleParser()
	testData := []TestDataParse{
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
		{
			script: "int age = 45 + 46; age = 20; age + 10 * 2;",
			expected: []*ast.SimpleASTNode{
				{NodeType: ast.IntDeclaration, Text: "age"},
				{NodeType: ast.Additive, Text: "+"},
				{NodeType: ast.IntLiteral, Text: "45"},
				{NodeType: ast.IntLiteral, Text: "46"},
				{NodeType: ast.AssignmentStmt, Text: "age"},
				{NodeType: ast.IntLiteral, Text: "20"},
				{NodeType: ast.Additive, Text: "+"},
				{NodeType: ast.Identifier, Text: "age"},
				{NodeType: ast.Multiplicative, Text: "*"},
				{NodeType: ast.IntLiteral, Text: "10"},
				{NodeType: ast.IntLiteral, Text: "2"},
			},
		},
		{
			script: "2+3+；",
			error:  errors.New(""),
		},
	}
	for _, data := range testData {
		Convey(fmt.Sprintf("%s", data.script), t, func() {
			error, node := parser.Parse(data.script)

			if data.error != nil {
				So(error, ShouldBeError)
				return
			}

			nodes := tools.PreTraverse(node)
			for i, node := range nodes {
				if i == 0 {
					So(node.GetType(), ShouldEqual, ast.Programm)
				} else {
					So(node.GetType(), ShouldEqual, data.expected[i-1].GetType())
					So(node.GetText(), ShouldEqual, data.expected[i-1].GetText())
				}
			}
		})
	}
}

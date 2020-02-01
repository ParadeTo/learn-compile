package lexer

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"learn-compile/craft/token"
	"testing"
)

type TestDataTokenize struct {
	script   string
	expected []SimpleToken
}

func TestSimpleLexer_Tokenize(t *testing.T) {
	lexer := NewSimpleLexer()
	testData := []TestDataTokenize{
		{
			script: "int age = 45;",
			expected: []SimpleToken{
				{_type: token.Int, text: "int"},
				{_type: token.Identifier, text: "age"},
				{_type: token.Assignment, text: "="},
				{_type: token.IntLiteral, text: "45"},
				{_type: token.SemiColon, text: ";"},
			},
		},
		{
			script: "inta age = 45;",
			expected: []SimpleToken{
				{_type: token.Identifier, text: "inta"},
				{_type: token.Identifier, text: "age"},
				{_type: token.Assignment, text: "="},
				{_type: token.IntLiteral, text: "45"},
				{_type: token.SemiColon, text: ";"},
			},
		},
		{
			script: "in age = 45;",
			expected: []SimpleToken{
				{_type: token.Identifier, text: "in"},
				{_type: token.Identifier, text: "age"},
				{_type: token.Assignment, text: "="},
				{_type: token.IntLiteral, text: "45"},
				{_type: token.SemiColon, text: ";"},
			},
		},
		{
			script: "age >= 45;",
			expected: []SimpleToken{
				{_type: token.Identifier, text: "age"},
				{_type: token.GE, text: ">="},
				{_type: token.IntLiteral, text: "45"},
				{_type: token.SemiColon, text: ";"},
			},
		},
		{
			script: "age > 45;",
			expected: []SimpleToken{
				{_type: token.Identifier, text: "age"},
				{_type: token.GT, text: ">"},
				{_type: token.IntLiteral, text: "45"},
				{_type: token.SemiColon, text: ";"},
			},
		},
		{
			script: "(1 + 2) * 4 / 5 - 6",
			expected: []SimpleToken{
				{_type: token.LeftParen, text: "("},
				{_type: token.IntLiteral, text: "1"},
				{_type: token.Plus, text: "+"},
				{_type: token.IntLiteral, text: "2"},
				{_type: token.RightParen, text: ")"},
				{_type: token.Star, text: "*"},
				{_type: token.IntLiteral, text: "4"},
				{_type: token.Slash, text: "/"},
				{_type: token.IntLiteral, text: "5"},
				{_type: token.Minus, text: "-"},
				{_type: token.IntLiteral, text: "6"},
			},
		},
	}
	for _, data := range testData {
		Convey(fmt.Sprintf("%s", data.script), t, func() {
			reader := lexer.Tokenize(data.script)
			pos := 0
			for {
				token := reader.Read()
				if token == nil {
					break
				}
				So(string(token.GetType()), ShouldEqual, data.expected[pos].GetType())
				pos++
			}
		})
	}
}

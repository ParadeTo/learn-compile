package lexer

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"learn-compile/craft/token"
	"testing"
)

type TestDataTokenize struct {
	script   string
	expected []token.SimpleToken
}

func TestSimpleLexer_Tokenize(t *testing.T) {
	lexer := NewSimpleLexer()
	testData := []TestDataTokenize{
		{
			script: "int age = 45;",
			expected: []token.SimpleToken{
				{Type: token.Int, Text: "int"},
				{Type: token.Identifier, Text: "age"},
				{Type: token.Assignment, Text: "="},
				{Type: token.IntLiteral, Text: "45"},
				{Type: token.SemiColon, Text: ";"},
			},
		},
		{
			script: "inta age = 45;",
			expected: []token.SimpleToken{
				{Type: token.Identifier, Text: "inta"},
				{Type: token.Identifier, Text: "age"},
				{Type: token.Assignment, Text: "="},
				{Type: token.IntLiteral, Text: "45"},
				{Type: token.SemiColon, Text: ";"},
			},
		},
		{
			script: "in age = 45;",
			expected: []token.SimpleToken{
				{Type: token.Identifier, Text: "in"},
				{Type: token.Identifier, Text: "age"},
				{Type: token.Assignment, Text: "="},
				{Type: token.IntLiteral, Text: "45"},
				{Type: token.SemiColon, Text: ";"},
			},
		},
		{
			script: "age >= 45;",
			expected: []token.SimpleToken{
				{Type: token.Identifier, Text: "age"},
				{Type: token.GE, Text: ">="},
				{Type: token.IntLiteral, Text: "45"},
				{Type: token.SemiColon, Text: ";"},
			},
		},
		{
			script: "age > 45;",
			expected: []token.SimpleToken{
				{Type: token.Identifier, Text: "age"},
				{Type: token.GT, Text: ">"},
				{Type: token.IntLiteral, Text: "45"},
				{Type: token.SemiColon, Text: ";"},
			},
		},
		{
			script: "(1 + 2) * 4 / 5 - 6",
			expected: []token.SimpleToken{
				{Type: token.LeftParen, Text: "("},
				{Type: token.IntLiteral, Text: "1"},
				{Type: token.Plus, Text: "+"},
				{Type: token.IntLiteral, Text: "2"},
				{Type: token.RightParen, Text: ")"},
				{Type: token.Star, Text: "*"},
				{Type: token.IntLiteral, Text: "4"},
				{Type: token.Slash, Text: "/"},
				{Type: token.IntLiteral, Text: "5"},
				{Type: token.Minus, Text: "-"},
				{Type: token.IntLiteral, Text: "6"},
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

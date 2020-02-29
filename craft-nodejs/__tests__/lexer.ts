import {TokenType, Token} from '../src/token'
import Lexer from '../src/lexer';

describe('lexer', () => {
  test('tokenize', () => {
    const lexer = new Lexer()

    ;[
      {
        script: 'int age = 45;',
        expected: [
          {type: TokenType.Int, test: 'int'},
          {type: TokenType.Identifier, test: 'age'},
          {type: TokenType.Assignment, test: '='},
          {type: TokenType.IntLiteral, test: '45'},
          {type: TokenType.SemiColon, test: ';'}
        ]
      },
      {
        script: "inta age = 45;",
        expected:[
          {type: TokenType.Identifier, Text: "inta"},
          {type: TokenType.Identifier, Text: "age"},
          {type: TokenType.Assignment, Text: "="},
          {type: TokenType.IntLiteral, Text: "45"},
          {type: TokenType.SemiColon, Text: ";"},
        ],
      },
      {
        script: "in age = 45;",
        expected: [
          {type: TokenType.Identifier, Text: "in"},
          {type: TokenType.Identifier, Text: "age"},
          {type: TokenType.Assignment, Text: "="},
          {type: TokenType.IntLiteral, Text: "45"},
          {type: TokenType.SemiColon, Text: ";"},
        ],
      },
      {
        script: "age >= 45;",
        expected: [
          {type: TokenType.Identifier, Text: "age"},
          {type: TokenType.GE, Text: ">="},
          {type: TokenType.IntLiteral, Text: "45"},
          {type: TokenType.SemiColon, Text: ";"},
        ],
      },
      {
        script: "age > 45;",
        expected: [
          {type: TokenType.Identifier, Text: "age"},
          {type: TokenType.GT, Text: ">"},
          {type: TokenType.IntLiteral, Text: "45"},
          {type: TokenType.SemiColon, Text: ";"},
        ],
      },
      {
        script: "(1 + 2) * 4 / 5 - 6",
        expected: [
          {type: TokenType.LeftParen, Text: "("},
          {type: TokenType.IntLiteral, Text: "1"},
          {type: TokenType.Plus, Text: "+"},
          {type: TokenType.IntLiteral, Text: "2"},
          {type: TokenType.RightParen, Text: ")"},
          {type: TokenType.Star, Text: "*"},
          {type: TokenType.IntLiteral, Text: "4"},
          {type: TokenType.Slash, Text: "/"},
          {type: TokenType.IntLiteral, Text: "5"},
          {type: TokenType.Minus, Text: "-"},
          {type: TokenType.IntLiteral, Text: "6"},
        ],
      },
    ].forEach(({ script, expected}) => {
      const reader = lexer.tokenize(script)
      let token: Token
      let pos = 0
      while (token = reader.read()) {
        expect(token.getType()).toBe(expected[pos++].type)
      }
    })
  })
})

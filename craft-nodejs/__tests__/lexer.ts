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
      }
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

import {TokenType, Token} from '../src/token'
import Parser from '../src/parser'
import { ASTNodeType } from '../src/ast'


describe('parser', () => {
  test('parse', () => {
    const parser = new Parser()

    ;[
      {
        script: 'int age;',
        expected: [
          {nodeType: ASTNodeType.IntDeclaration, text: 'age'},
        ],
        error: false
      },
      {
        script: "int age = 45 + 46;",
        expected:[
          {nodeType: ASTNodeType.IntDeclaration, text: "age"},
          {nodeType: ASTNodeType.Additive, text: "+"},
          {nodeType: ASTNodeType.IntLiteral, text: "45"},
          {nodeType: ASTNodeType.IntLiteral, text: "46"},
        ],
        error: false
      },
      {
        script: "int age = 45 + 46; age = 20; age + 10 * 2;",
        expected: [
          {nodeType: ASTNodeType.IntDeclaration, text: "age"},
          {nodeType: ASTNodeType.Additive, text: "+"},
          {nodeType: ASTNodeType.IntLiteral, text: "45"},
          {nodeType: ASTNodeType.IntLiteral, text: "46"},
          {nodeType: ASTNodeType.AssignmentStmt, text: "age"},
          {nodeType: ASTNodeType.IntLiteral, text: "20"},
          {nodeType: ASTNodeType.Additive, text: "+"},
          {nodeType: ASTNodeType.Identifier, text: "age"},
          {nodeType: ASTNodeType.Multiplicative, text: "*"},
          {nodeType: ASTNodeType.IntLiteral, text: "10"},
          {nodeType: ASTNodeType.IntLiteral, text: "2"},
        ],
      },
      {
        script: "2+3+；",
        error: true,
        expected: []
      }
    ].forEach(({ script, expected, error}) => {
      if (error) {
        expect(() => parser.parse(script)).toThrow()
      } else {
        const node = parser.parse(script)
        const nodes = parser.preTraverse(node)
        nodes.forEach((node, i) => {
          if (i === 0) {
            expect(node.getType()).toBe(ASTNodeType.Programm)
          } else {
            expect(node.getType()).toBe(expected[i - 1].nodeType)
            expect(node.getText()).toBe(expected[i - 1].text)
          }
        })
      }
    })
  })
})

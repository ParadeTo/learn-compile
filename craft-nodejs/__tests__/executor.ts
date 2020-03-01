import Parser from '../src/parser'
import {ASTNodeType} from '../src/ast'
import Executor from '../src/executor'

describe('parser', () => {
  test('parse', () => {
    const parser = new Parser()
    ;[
      {
        script: 'int age;',
        expected: 0
      },
      {
        script: 'int age = 45 + 46; age + 1;',
        expected: 92
      },
      {
        script: 'int age = (1 + 2)*3;',
        expected: 9
      }
    ].forEach(({script, expected}) => {
      const astRoot = parser.parse(script)
      const executor = new Executor()
      expect(executor.execute(astRoot)).toBe(expected)
    })
  })
})

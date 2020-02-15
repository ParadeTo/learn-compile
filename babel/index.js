const { parse } = require('@babel/parser')
const generate = require('@babel/generator').default

const code = 'class Example {}'
const ast = parse(code)
console.log(ast.program)
const output = generate(ast, { sourceMaps: true }, code)
console.log(output)

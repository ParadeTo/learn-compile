const { parse } = require('@babel/parser')
const traverse = require('@babel/traverse').default
const generate = require('@babel/generator').default

const code = `function square(n) {
  {
    return n * n;
  }
}`;

const ast = parse(code);
traverse(ast, {
  BinaryExpression: {
    enter(path) {
      // path.remove()
      console.log(path.isBinaryExpression())
    },
    exit(path) {
      // console.log(path)
      // console.log(path.parent)
      // console.log(path.node)
      // console.log(path.get('left'))
    }
  },
  Identifier: {
    enter(path) {
      if (
        path.node.name === "n"
      ) {
        path.node.name = "x";
      }
    }
  }
});


const output = generate(ast, {  }, code)

console.log(output)
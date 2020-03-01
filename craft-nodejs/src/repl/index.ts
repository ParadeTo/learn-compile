import Parser from "../parser"
import Executor from "../executor"

export default function run() {
  const verbose = true // process.argv[1] === '-v'
  const parser = new Parser()
  const executor = new Executor(verbose)

  verbose && process.stdout.write('verbose mode\n')
  process.stdout.write('\n>')

  const rl = require('readline').createInterface({
    input: process.stdin
  })

  rl.on('line', input => {
    try {
      const astRoot = parser.parse(input)
      process.stdout.write(executor.execute(astRoot).toString())
    } catch (error) {
      process.stdout.write(error.toString())
    } finally {
      process.stdout.write('\n>')
    }
  })
}

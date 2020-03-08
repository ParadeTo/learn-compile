import readline from 'readline'
import Parser from "../parser"
import Executor from "../executor"

export default function run(verbose = false) {
  const parser = new Parser()
  const executor = new Executor(verbose)

  verbose && process.stdout.write('verbose mode\n')
  process.stdout.write('\n>')

  const rl = readline.createInterface({
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

import fs from 'fs'
import Parser from '../parser'
import Executor from '../executor'

export default function interprete(filename: string, verbose = false) {
  const code = fs.readFileSync(filename)
  const parser = new Parser()
  const executor = new Executor(verbose)
  const astRoot = parser.parse(code.toString())
  executor.execute(astRoot)
}
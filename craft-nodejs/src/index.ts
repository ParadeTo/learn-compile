import Parser from './parser'

const parser = new Parser()
const astRoot = parser.parse('inta age = 45 + 46;')
parser.dumpAST(astRoot)

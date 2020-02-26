import Lexer from './lexer'
const lexer = new Lexer()
const reader = lexer.tokenize("int a = 1;")
lexer.dump(reader)
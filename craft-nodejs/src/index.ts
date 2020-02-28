import Lexer from './lexer'
const lexer = new Lexer()
const reader = lexer.tokenize("int age = 43;")
lexer.dump(reader)
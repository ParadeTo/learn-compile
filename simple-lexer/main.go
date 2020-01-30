package main

import "./lexer"

func main() {
	lexer := lexer.NewSimpleLexer()
	script := "int age = 45;"
	tokenReader := lexer.Tokenize(script)
	lexer.Dump(tokenReader)
}
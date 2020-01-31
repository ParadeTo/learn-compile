package craft

import "fmt"

func main() {
	//lexer := lexer.NewSimpleLexer()
	//script := "int age = 45;"
	//tokenReader := lexer.Tokenize(script)
	//lexer.Dump(tokenReader)
	s := "a"
	switch s {
	case "a", "b":
		fmt.Println(1)
	}
}
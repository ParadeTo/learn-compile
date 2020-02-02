package main

import (
	"bufio"
	"flag"
	"fmt"
	"learn-compile/craft/parser"
	repl "learn-compile/craft/repl"
	"os"
)

var v = flag.Bool("v", false, "verbose")

func main() {
	flag.Parse()

	_parser := parser.NewSimpleParser()
	_repl := repl.NewSimpleRepl(*v)

	if *v {
		fmt.Println("verbose mode")
	}
	fmt.Print("\n>")

	f := bufio.NewReader(os.Stdin) //读取输入的内容
	for {
		input, _ := f.ReadString('\n')
		if input == "exit;\n" {
			fmt.Println("Good bye!")
			break
		}

		err, tree := _parser.Parse(input)
		if err != nil {
			fmt.Println(err)
		} else {
			if *v {
				_parser.DumpAST(tree, "")
			}
			_repl.Evaluate(tree, "")
		}

		fmt.Print("\n>")
	}
}

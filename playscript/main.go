package main

import "learn-compile/playscript/playscript"

func main() {
	compiler := playscript.NewPlayScriptCompiler()
	at := compiler.Compile("int foo(int a){ int a=1; int b=a;}")

	//scope := playscript.NewBaseScope()
	//variable := playscript.NewVariable("a", nil, nil)
	//scope.AddSymbol(variable)
	//v := scope.GetVariable("a")
	//fmt.Println(v.ToString())
}

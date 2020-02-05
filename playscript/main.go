package main

import "learn-compile/playscript/playscript"

func main() {
	compiler := playscript.NewPlayScriptCompiler()
	compiler.Compile("int foo(int a){return a+a;} int foo(){}")
	//scope := playscript.NewBaseScope()
	//variable := playscript.NewVariable("a", nil, nil)
	//scope.AddSymbol(variable)
	//v := scope.GetVariable("a")
	//fmt.Println(v.ToString())
}

package main

import "learn-compile/playscript/playscript"

func main() {
	compiler := playscript.NewPlayScriptCompiler()
	compiler.Compile("for(int i = 0; i < 10; i++){45;}")
	//scope := playscript.NewScope()
	//variable := playscript.NewVariable("a", nil, nil)
	//scope.AddSymbol(variable)
	//v := scope.GetVariable("a")
	//fmt.Println(v.ToString())
}

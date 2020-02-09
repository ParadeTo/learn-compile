package main

import "fmt"

func main() {
	//compiler := playscript.NewPlayScriptCompiler()
	//at := compiler.Compile("int foo(){ return 4;}")
	//fmt.Println(at)
	//scope := playscript.NewBaseScope()
	//variable := playscript.NewVariable("a", nil, nil)
	//scope.AddSymbol(variable)
	//v := scope.GetVariable("a")
	//fmt.Println(v.ToString())

	//s1 := tools.NewSet(1,2,3,4)
	//s2 := tools.NewSet(1,2)
	//s1.RemoveSubset(s2)
	//fmt.Println(s1)
	//fmt.Println(s)
	//s.Remove(3)
	//fmt.Println(s)
	//fmt.Println(s.Contains(3))
	//var a interface{}
	//var b interface{}
	//a = 0x12
	b := "abcdef"
	//sa, _ := a.(string)
	//sb, _ := b.(string)
	//sa := fmt.Sprintf("%s", a)
	//sb := fmt.Sprintf("%s", b)
	fmt.Println(b[1 : len(b)-1])
}

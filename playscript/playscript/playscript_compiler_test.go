package playscript

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type TestDataCompile struct {
	script string
	error  string
}

func TestPlayScriptCompiler_Compile(t *testing.T) {
	testData := []TestDataCompile{
		{
			script: "int foo(int a){ int a=1; int b=a;}",
			error:  "Variable or parameter already Declared: a @1:20",
		},
		{
			script: "int foo(){} int foo(){ int a=1; int b=a;}",
			error:  "Function or method already Declared: foo @1:12",
		},
		{
			script: "int foo(){}",
			error:  "return statement expected in function @1:0",
		},
		{
			script: "{return 1;}",
			error:  "return statement not in function body @1:1",
		},
		{
			script: "int a=1; break;",
			error:  "break statement not in loop or switch statements @1:9",
		},
		//{
		//	script: "for(int i = 0; i < 1; i++) {}",
		//	error:  "break statement not in loop or switch statements @1:9",
		//},
		{
			script: "int a; a =1.1;",
			error:  "cannot assign 1.1 of type Float to a of type Integer @1:7",
		},
	}
	for _, data := range testData {
		Convey(fmt.Sprintf("%s", data.script), t, func() {
			compiler := NewPlayScriptCompiler()
			at := compiler.Compile(data.script)
			if data.error != "" {
				So(at.GetFirstCompilationError().ToString(), ShouldEqual, data.error)
			} else {
				So(at.GetFirstCompilationError(), ShouldEqual, nil)
			}
		})
	}
}

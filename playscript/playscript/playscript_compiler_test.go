package playscript

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
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
			script: "void foo(){}",
			error:  "",
		},
		{
			script: "{return 1;}",
			error:  "return statement not in function body @1:1",
		},
		{
			script: "int a=1; break;",
			error:  "break statement not in loop or switch statements @1:9",
		},
		{
			script: "int a; a =1.1;",
			error:  "cannot assign 1.1 of type Float to a of type Integer @1:7",
		},
		{
			script: "for(int i = 0; i < 1; i++)\n { int foo(){break;} }",
			error:  "break statement not in loop or switch statements @2:13",
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

type TestDataExecute struct {
	scriptFilename string
	printlnArr     []interface{}
	result         interface{}
}

func TestPlayScriptCompiler_Execute(t *testing.T) {
	testData := []TestDataExecute{
		{
			scriptFilename: "blockscope.play",
			result:         nil,
			printlnArr:     []interface{}{0, 2, 3, 0},
		},
		{
			scriptFilename: "arithmetic.play",
			result:         6,
			printlnArr:     nil,
		},
		{
			scriptFilename: "function.play",
			result:         nil,
			printlnArr:     []interface{}{"foo(10)=20", "bar(3)=3", "bar(10)=5"},
		},
		{
			scriptFilename: "loop.play",
			result:         nil,
			printlnArr: []interface{}{
				"while loop:",
				"i=1",
				"i=2",
				"i=3",
				"i=4",
				"i=5",
				"i=6",
				"i=7",
				"i=8",
				"i=9",
				"i=10",

				"for loop:",
				"i=0, a=0",
				"i=1, a=1",
				"i=2, a=3",
				"i=3, a=6",
				"i=4, a=10",
				"i=5, a=15",
				"i=6, a=21",
				"i=7, a=28",
				"i=8, a=36",
				"i=9, a=45",

				"while loop, breaks when i > 5:",
				"i=1",
				"i=2",
				"i=3",
				"i=4",
				"i=5",
				"i=6",

				"for loop, breaks when i > 5:",
				"i=0, a=0",
				"i=1, a=1",
				"i=2, a=3",
				"i=3, a=6",
				"i=4, a=10",
				"i=5, a=15",
				"i=6, a=21",

				"for loop within while loop:",
				"i=1, j=0",
				"i=1, j=1",
				"i=1, j=2",
				"i=2, j=0",
				"i=2, j=1",
				"i=2, j=2",
				"i=3, j=0",
				"i=3, j=1",
				"i=3, j=2",
			},
		},
	}
	for _, data := range testData {
		Convey(fmt.Sprintf("%s", data.scriptFilename), t, func() {
			compiler := NewPlayScriptCompiler()
			script, _ := ioutil.ReadFile("./testdata/" + data.scriptFilename)
			at := compiler.Compile(string(script))
			result := compiler.Execute(at, true)
			So(result, ShouldEqual, data.result)
			So(compiler.visitor.printlnArr, ShouldResemble, data.printlnArr)
		})
	}
}

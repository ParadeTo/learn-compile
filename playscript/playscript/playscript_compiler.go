package playscript

import (
	"fmt"
	. "github.com/antlr/antlr4/runtime/Go/antlr"
	"learn-compile/playscript/parser"
)

type PlayScriptCompiler struct {
	at     *AnnotatedTree
	lexer  *parser.PlayScriptLexer
	parser *parser.PlayScriptParser
}

func (compiler *PlayScriptCompiler) Compile(script string) *AnnotatedTree {
	at := NewAnnotatedTree()

	// 词法分析
	lexer := parser.NewPlayScriptLexer(NewInputStream(script))
	stream := NewCommonTokenStream(lexer, TokenDefaultChannel)

	// 语法分析
	_parser := parser.NewPlayScriptParser(stream)
	at.ast = _parser.Prog()

	// 语义分析
	walker := NewParseTreeWalker()

	//多步的语义解析。
	//优点：1.代码更清晰；2.允许使用在声明之前，这在支持面向对象、递归函数等特征时是必须的。
	//pass1：类型和Scope
	pass1 := NewTypeAndScopeScanner(at)
	walker.Walk(pass1, at.ast)

	//pass2：把变量、类继承、函数声明的类型都解析出来。也就是所有声明时用到类型的地方。
	pass2 := NewTypeResolver(at)
	walker.Walk(pass2, at.ast)

	//pass3：消解有的变量应用、函数引用。另外还做了类型的推断。
	pass3 := NewRefResolver(at)
	walker.Walk(pass3, at.ast)

	//pass4：类型检查
	pass4 := NewTypeChecker(at)
	walker.Walk(pass4, at.ast)

	//pass5: 其他语义检查
	pass5 := NewSematicValidator(at)
	walker.Walk(pass5, at.ast)

	//pass6: 闭包分析
	closureAnalyzer := NewClosureAnalyzer(at)
	closureAnalyzer.AnalyzeClosures()

	fmt.Println(at.ast.ToStringTree([]string{"youxingzhi"}, _parser))
	return at
}

func NewPlayScriptCompiler() *PlayScriptCompiler {
	return &PlayScriptCompiler{}
}

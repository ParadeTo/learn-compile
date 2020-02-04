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

	fmt.Println(at.ast.ToStringTree([]string{}, _parser))
	return at
}

func NewPlayScriptCompiler() *PlayScriptCompiler {
	return &PlayScriptCompiler{}
}

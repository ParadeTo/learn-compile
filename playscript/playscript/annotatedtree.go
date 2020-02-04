package playscript

import . "github.com/antlr/antlr4/runtime/Go/antlr"

/**
 * 注释树。
 * 语义分析的结果都放在这里。跟AST的节点建立关联。包括：
 * 1.类型信息，包括基本类型和用户自定义类型；
 * 2.变量和函数调用的消解；
 * 3.作用域Scope。在Scope中包含了该作用域的所有符号。Variable、Function、Class等都是符号。
 */
type AnnotatedTree struct {
	ast        ParseTree
	node2Scope map[ParserRuleContext]IScope
	nameSpace  *NameSpace
}

func NewAnnotatedTree() *AnnotatedTree {
	return &AnnotatedTree{
		node2Scope: make(map[ParserRuleContext]IScope),
	}
}

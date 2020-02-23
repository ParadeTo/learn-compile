package playscript

import (
	"fmt"
	. "github.com/antlr/antlr4/runtime/Go/antlr"
	"learn-compile/playscript/parser"
)

/**
 * 注释树。
 * 语义分析的结果都放在这里。跟AST的节点建立关联。包括：
 * 1.类型信息，包括基本类型和用户自定义类型；
 * 2.变量和函数调用的消解；
 * 3.作用域Scope。在Scope中包含了该作用域的所有符号。Variable、Function、Class等都是符号。
 */
type AnnotatedTree struct {
	ast ParseTree
	// 解析出来的所有类型，包括类和函数，以后还可以包括数组和枚举。类的方法也作为单独的要素放进去。
	types []Type
	// AST节点对应的Scope，如for、函数调用会启动新的Scope
	node2Scope map[Tree]Scope
	// AST节点对应的Symbol
	symbolOfNode map[Tree]Symbol
	// AST节点对应的Type
	typeOfNode map[Tree]Type
	// 在构造函数里引用的 this()
	thisConstructorRef map[Tree]*Function // key 是否要用 *Function
	// 在构造函数里引用的 super()
	superConstructorRef map[Tree]*Function // key 是否要用 *Function
	// 命名空间
	nameSpace *NameSpace
	// 语义分析过程中生成的信息，包括普通信息、警告和错误
	logs []*CompilationLog
}

// 得到第一条编译错误
func (tree *AnnotatedTree) GetFirstCompilationError() *CompilationLog {
	for _, log := range tree.logs {
		if log._type == CompilationLogERROR {
			return log
		}
	}
	return nil
}

// 是否有编译错误
func (tree *AnnotatedTree) HasCompilationError() bool {
	return tree.GetFirstCompilationError() != nil
}

// 通过名称查找Variable 逐级Scope查找
func (tree *AnnotatedTree) LookupVariable(scope Scope, idName string) *Variable {
	variable := scope.GetVariable(idName)
	if variable == nil && scope.GetEnclosingScope() != nil {
		variable = tree.LookupVariable(scope.GetEnclosingScope(), idName)
	}
	return variable
}

// 通过方法的名称。逐级Scope查找
// TODO 通过方法的名称和方法签名查找Function。逐级Scope查找
func (tree *AnnotatedTree) LookupFunction(scope Scope, idName string) *Function {
	function := scope.GetFunction(idName)
	if function == nil && scope.GetEnclosingScope() != nil {
		function = tree.LookupFunction(scope.GetEnclosingScope(), idName)
	}
	return function
}

// 通过名称查找Class。逐级Scope查找
func (tree *AnnotatedTree) LookupClass(scope Scope, idName string) *Class {
	rtn := scope.GetClass(idName)
	if rtn == nil && scope.GetEnclosingScope() != nil {
		rtn = tree.LookupClass(scope.GetEnclosingScope(), idName)
	}
	return rtn
}

// TODO 单纯根据名称并不严密
func (tree *AnnotatedTree) LookupType(idName string) Type {
	var rtn Type
	for _, _type := range tree.types {
		if _type.GetName() == idName {
			rtn = _type
			break
		}
	}
	return rtn
}

func (tree *AnnotatedTree) LookupFunctionVariable(scope Scope, idName string) *Variable {
	variable := scope.GetFunctionVariable(idName)
	if variable == nil && scope.GetEnclosingScope() != nil {
		variable = tree.LookupFunctionVariable(scope.GetEnclosingScope(), idName)
	}
	return variable
}

func (tree *AnnotatedTree) AddType(_type Type) {
	tree.types = append(tree.types, _type)
}

// 查找某节点所在的Scope
// 算法：逐级查找父节点，找到一个对应着Scope的上级节点
func (tree *AnnotatedTree) EnclosingScopeOfNode(node Tree) Scope {
	var rtn Scope
	parent := node.GetParent()
	if parent != nil {
		rtn = tree.node2Scope[parent]
		if rtn == nil {
			rtn = tree.EnclosingScopeOfNode(parent)
		}
	}
	return rtn
}

func (tree *AnnotatedTree) Log(message string, _type int, ctx ParserRuleContext) {
	log := NewCompilationLog()
	log.ctx = ctx
	log.message = message
	log.line = ctx.GetStart().GetLine()
	log.positionInLine = ctx.GetStart().GetColumn()
	log._type = _type
	tree.logs = append(tree.logs, log)
	fmt.Println(log.ToString())
}

func (tree *AnnotatedTree) LogError(message string, ctx ParserRuleContext) {
	tree.Log(message, CompilationLogERROR, ctx)
}

// 包含某节点的函数
func (tree *AnnotatedTree) EnclosingFunctionOfNode(ctx Tree) *Function {
	parent := ctx.GetParent()
	if _, ok := parent.(*parser.FunctionDeclarationContext); ok {
		scope := tree.node2Scope[parent]
		if function, ok := scope.(*Function); ok {
			return function
		}
		return nil
	} else if ctx.GetParent() == nil {
		return nil
	} else {
		return tree.EnclosingFunctionOfNode(ctx.GetParent())
	}
}

// 包含某节点的类
func (tree *AnnotatedTree) EnclosingClassOfNode(ctx Tree) *Class {
	parent := ctx.GetParent()
	if _, ok := parent.(*parser.ClassDeclarationContext); ok {
		scope := tree.node2Scope[parent]
		if class, ok := scope.(*Class); ok {
			return class
		}
		return nil
	} else if ctx.GetParent() == nil {
		return nil
	} else {
		return tree.EnclosingClassOfNode(ctx.GetParent())
	}
}

func NewAnnotatedTree() *AnnotatedTree {
	return &AnnotatedTree{
		node2Scope:   make(map[Tree]Scope),
		symbolOfNode: make(map[Tree]Symbol),
		typeOfNode:   make(map[Tree]Type),
	}
}

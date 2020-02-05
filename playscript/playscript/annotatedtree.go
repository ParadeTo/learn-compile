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
	ast ParseTree
	// 解析出来的所有类型，包括类和函数，以后还可以包括数组和枚举。类的方法也作为单独的要素放进去。
	types []Type
	// AST节点对应的Scope，如for、函数调用会启动新的Scope
	node2Scope map[Tree]Scope
	// AST节点对应的Symbol
	symbolOfNode map[Tree]Symbol
	// AST节点对应的Type
	typeOfNode map[Tree]Type
	// 命名空间
	nameSpace *NameSpace
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

func NewAnnotatedTree() *AnnotatedTree {
	return &AnnotatedTree{
		node2Scope:   make(map[Tree]Scope),
		symbolOfNode: make(map[Tree]Symbol),
		typeOfNode:   make(map[Tree]Type),
	}
}

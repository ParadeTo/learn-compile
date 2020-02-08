package playscript

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"learn-compile/playscript/tools"
)

type ClosureAnalyzer struct {
	at *AnnotatedTree
}

// 看看node1是不是node2的祖先
func (this *ClosureAnalyzer) isAncetor(node1, node2 antlr.Tree) bool {
	if node2.GetParent() == nil {
		return false
	} else if node2.GetParent() == node1 {
		return true
	} else {
		return this.isAncetor(node1, node2.GetParent())
	}
}

// 对所有的函数做闭包分析
// 只做标准函数的分析，不做类的方法的分析
func (this *ClosureAnalyzer) AnalyzeClosures() {
	for _, _type := range this.at.types {
		if function, ok := _type.(*Function); ok {
			set := this.calcClosureVariables(function)
			if set.Size() > 0 {
				function.closureVariables = set
			}
		}
	}
}

// 为某个函数计算闭包变量，也就是它所引用的外部环境变量
// 算法：计算所有的变量引用，去掉内部声明的变量，剩下的就是外部的
func (this *ClosureAnalyzer) calcClosureVariables(function *Function) *tools.Set {
	referred := this.variablesReferredByScope(function)
	declared := this.variablesDeclaredUnderScope(function)
	referred.RemoveSubset(declared)
	return referred
}

// Scope（包括下级Scope）内部的所有变量
func (this *ClosureAnalyzer) variablesReferredByScope(scope Scope) *tools.Set {
	rtn := tools.NewSet()
	scopeNode := scope.GetCtx()

	// 扫描所有的符号引用。这对于大的程序性能不够优化，因为符号表太大
	for node, symbol := range this.at.symbolOfNode {
		if variable, ok := symbol.(*Variable); ok && this.isAncetor(scopeNode, node) {
			rtn.Add(variable)
		}
	}
	return rtn
}

// 在一个Scope（及）下级Scope中声明的所有变量的集合
func (this *ClosureAnalyzer) variablesDeclaredUnderScope(scope Scope) *tools.Set {
	rtn := tools.NewSet()
	for _, symbol := range scope.GetSymbols() {
		if variable, ok := symbol.(*Variable); ok {
			rtn.Add(variable)
		} else if scope, ok := symbol.(Scope); ok {
			rtn.Add(this.variablesDeclaredUnderScope(scope))
		}
	}
	return rtn
}

func NewClosureAnalyzer(at *AnnotatedTree) *ClosureAnalyzer {
	return &ClosureAnalyzer{at}
}

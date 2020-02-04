package playscript

import "github.com/antlr/antlr4/runtime/Go/antlr"

// 变量
type Variable struct {
	*Symbol
	// 变量类型
	_type Type

	// 作为参数的变量的属性
	// 默认值
	defaultValue interface{}
	// 是否允许多次重复，这是一个创新的参数机制
	multiplicity int
}

func (v *Variable) ToString() string {
	return "Variable: " + v.name + " -> " + v._type.ToString()
}

func NewVariable(name string, enclosingScope IScope, ctx antlr.ParserRuleContext) *Variable {
	variable := &Variable{Symbol: &Symbol{name: name, ctx: ctx}}
	variable.SetEnclosingScope(enclosingScope)
	return variable
}

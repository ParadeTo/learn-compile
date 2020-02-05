package playscript

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Function -> BaseScope -> BaseSymbol
type Function struct {
	*BaseScope
	// 参数
	parameters []*Variable
	// 参数类型
	paramTypes []Type
	// 返回值
	returnType Type
	// 闭包变量，即函数里面所引用的外部环境变量
}

func (f *Function) AddParamType(_type Type) {
	f.paramTypes = append(f.paramTypes, _type)
}

func (f *Function) AddParameter(variable *Variable) {
	f.parameters = append(f.parameters, variable)
}

func (f *Function) IsType(_type Type) bool {
	if functionType, ok := _type.(FunctionType); ok {
		return isType(f, functionType)
	}
	return false
}

func (f *Function) GetReturnType() Type {
	return f.returnType
}

func (f *Function) GetParamTypes() []Type {
	for _, param := range f.parameters {
		f.paramTypes = append(f.paramTypes, param._type)
	}
	return f.paramTypes
}

func (f *Function) MatchParameterTypes(paramTypes []Type) bool {
	if len(f.parameters) != len(paramTypes) {
		return false
	}

	match := true
	for i := 0; i < len(paramTypes); i++ {
		_var := f.parameters[i]
		_type := paramTypes[i]
		if !_var._type.IsType(_type) {
			match = false
			break
		}
	}

	return match
}

func (f *Function) ToString() string {
	return "Function: " + f.name
}

// 该函数是不是类的方法 TODO
func (f *Function) IsMethod() bool {
	return true
}

// 该函数是不是类的构造函数 TODO
func (f *Function) IsConstructor() bool {
	return false
}

func NewFunction(name string, enclosingScope Scope, ctx antlr.ParserRuleContext) *Function {
	function := &Function{BaseScope: NewBaseScope()}
	function.name = name
	function.SetEnclosingScope(enclosingScope)
	function.SetCtx(ctx)
	return function
}

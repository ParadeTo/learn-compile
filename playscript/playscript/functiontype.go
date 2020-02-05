package playscript

import (
	"strconv"
)

type FunctionType interface {
	Type
	GetReturnType() Type
	AddParamType(_type Type)
	GetParamTypes() []Type
	MatchParameterTypes(paramTypes []Type) bool
}

// 对于未命名的类型，自动赋予名字
var nameIndex = 1

/**
 * 工具性方法，比较type1是否是type2。
 * 规则：
 * 1.type1的返回值跟type2相等，或者是它的子类型。
 * @param type1
 * @param type2
 * @return
 */
func isType(type1, type2 FunctionType) bool {
	if type1 == type2 {
		return true
	}

	if !type1.GetReturnType().IsType(type2.GetReturnType()) {
		return false
	}

	paramTypes1 := type1.GetParamTypes()
	paramTypes2 := type2.GetParamTypes()

	if len(paramTypes1) != len(paramTypes2) {
		return false
	}

	for i := 0; i < len(paramTypes1); i++ {
		if !paramTypes1[i].IsType(paramTypes2[i]) {
			return false
		}
	}

	return true
}

type DefaultFunctionType struct {
	name           string
	enclosingScope Scope
	returnType     Type
	paramTypes     []Type
}

func (ft *DefaultFunctionType) AddParamType(_type Type) {
	ft.paramTypes = append(ft.paramTypes, _type)
}

func (ft *DefaultFunctionType) GetName() string {
	return ft.name
}

func (ft *DefaultFunctionType) GetEnclosingScope() Scope {
	return ft.enclosingScope
}

func (ft *DefaultFunctionType) IsType(_type Type) bool {
	if functionType, ok := _type.(FunctionType); ok {
		return isType(ft, functionType)
	}
	return false
}

func (ft *DefaultFunctionType) ToString() string {
	return "FunctionType"
}

func (ft *DefaultFunctionType) GetReturnType() Type {
	return ft.returnType
}

func (ft *DefaultFunctionType) GetParamTypes() []Type {
	return ft.paramTypes
}

// 检查该函数是否匹配所需的参数
func (ft *DefaultFunctionType) MatchParameterTypes(paramTypes []Type) bool {
	if len(ft.paramTypes) != len(paramTypes) {
		return false
	}

	match := true
	for i := 0; i < len(paramTypes); i++ {
		type1 := ft.paramTypes[i]
		type2 := paramTypes[i]
		if !type1.IsType(type2) {
			match = false
			break
		}
	}

	return match
}

func NewDefaultFunctionType() *DefaultFunctionType {
	f := &DefaultFunctionType{}
	f.name = "FunctionType: " + strconv.Itoa(nameIndex)
	nameIndex++
	return f
}

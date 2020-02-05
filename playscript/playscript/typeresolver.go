package playscript

import (
	"fmt"
	. "learn-compile/playscript/parser"
)

/**
 * 第二遍扫描。把变量、类继承、函数声明的类型都解析出来。
 * 也就是所有用到Type的地方。
 * 扫描是为了填充 symbolOfNode 的信息
 */
type TypeResolver struct {
	*BasePlayScriptListener
	at *AnnotatedTree
}

// 设置所声明的变量的类型，单个变量和多个变量都整合在 variableDeclarators 语句中
func (tr *TypeResolver) ExitVariableDeclarators(ctx *VariableDeclaratorsContext) {
	_type := tr.at.typeOfNode[ctx.TypeType()]
	for _, child := range ctx.AllVariableDeclarator() {
		symbol := tr.at.symbolOfNode[child]
		if variable, ok := symbol.(*Variable); ok {
			variable._type = _type
		}
	}
}

// 把所有的变量声明加入符号表
func (tr *TypeResolver) EnterVariableDeclaratorId(ctx *VariableDeclaratorIdContext) {
	idName := ctx.IDENTIFIER().GetText()
	// 所属的作用域
	scope := tr.at.EnclosingScopeOfNode(ctx)
	variable := NewVariable(idName, scope, ctx)

	// 查看变量是否已经定义过
	if scope.GetVariable(idName) != nil {
		// TODO 记录log
	}

	scope.AddSymbol(variable)
	tr.at.symbolOfNode[ctx] = variable
}

// 设置函数的返回值类型，并查重
func (tr *TypeResolver) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {
	scope := tr.at.node2Scope[ctx]
	typeCtx := ctx.TypeTypeOrVoid()
	if typeCtx != nil {
		if function, ok := scope.(*Function); ok {
			function.returnType = tr.at.typeOfNode[typeCtx]
		}
	}
	// 函数查重，暂时只检查名称
	// TODO，检查参数
	parentScope := tr.at.EnclosingScopeOfNode(ctx)
	found := parentScope.GetFunction(scope.GetName())
	if found != nil && found != scope {
		// TODO log
		fmt.Println("Function or method already Declared: " + ctx.IDENTIFIER().GetText())
	}
}

// 设置函数的参数的类型，这些参数已经在EnterVariableDeclaratorId中作为变量声明了，现在设置它们的类型
func (tr *TypeResolver) ExitFormalParameter(ctx *FormalParameterContext) {
	_type := tr.at.typeOfNode[ctx.TypeType()]
	symbol := tr.at.symbolOfNode[ctx.VariableDeclaratorId()]
	if variable, ok := symbol.(*Variable); ok {
		variable._type = _type
		scope := tr.at.EnclosingScopeOfNode(ctx)
		if function, ok := scope.(*Function); ok {
			function.AddParameter(variable)
		}
	}
}

// 函数的返回类型
func (tr *TypeResolver) ExitTypeTypeOrVoid(ctx *TypeTypeOrVoidContext) {
	if ctx.VOID() != nil {
		tr.at.typeOfNode[ctx] = GetVoidType()
	} else if ctx.TypeType() != nil {
		tr.at.typeOfNode[ctx] = tr.at.typeOfNode[ctx.TypeType()]
	}
}

// 函数的返回类型之返回非 void
// TypeType 包括的子节点有:
//  ClassOrInterfaceType: MyClass foo(){}
//  FunctionType: function int() foo(){}
//  PrimitiveType: int foo(){}
func (tr *TypeResolver) ExitTypeType(ctx *TypeTypeContext) {
	var _type Type
	if ctx.ClassOrInterfaceType() != nil {
		_type = tr.at.typeOfNode[ctx.ClassOrInterfaceType()]
	} else if ctx.FunctionType() != nil {
		_type = tr.at.typeOfNode[ctx.FunctionType()]
	} else if ctx.PrimitiveType() != nil {
		_type = tr.at.typeOfNode[ctx.PrimitiveType()]
	}
	tr.at.typeOfNode[ctx] = _type
}

//func (tr *TypeResolver) EnterClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) {
//	if ctx.AllIDENTIFIER() != nil {
//		// class 的闭包 scope
//		scope := tr.at.EnclosingScopeOfNode(ctx)
//		idName := ctx.GetText()
//	}
//}

// 比如 function int(string, int, int) foo() {}
// functionType: function int(string, int, int)
// TypeTypeOrVoid: int
// typeList: string, int, int
func (tr *TypeResolver) ExitFunctionType(ctx *FunctionTypeContext) {
	functionType := NewDefaultFunctionType()
	tr.at.AddType(functionType)
	tr.at.typeOfNode[ctx] = functionType
	functionType.returnType = tr.at.typeOfNode[ctx.TypeTypeOrVoid()]
	// 参数的类型
	itcl := ctx.TypeList()
	if itcl != nil {
		//rct := itcl.GetRuleContext()
		if tcl, ok := itcl.(*TypeListContext); ok {
			for _, ttc := range tcl.AllTypeType() {
				_type := tr.at.typeOfNode[ttc]
				functionType.AddParamType(_type)
			}
		}
	}
}

func (tr *TypeResolver) ExitPrimitiveType(ctx *PrimitiveTypeContext) {
	var _type Type
	if ctx.BOOLEAN() != nil {
		_type = Boolean
	} else if ctx.INT() != nil {
		_type = Integer
	} else if ctx.LONG() != nil {
		_type = Long
	} else if ctx.FLOAT() != nil {
		_type = Float
	} else if ctx.DOUBLE() != nil {
		_type = Double
	} else if ctx.BYTE() != nil {
		_type = Byte
	} else if ctx.SHORT() != nil {
		_type = Short
	} else if ctx.CHAR() != nil {
		_type = Char
	} else if ctx.STRING() != nil {
		_type = String
	}
	tr.at.typeOfNode[ctx] = _type
}

func NewTypeResolver(at *AnnotatedTree) *TypeResolver {
	return &TypeResolver{at: at}
}

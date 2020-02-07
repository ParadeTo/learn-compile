package playscript

import (
	"fmt"
	. "learn-compile/playscript/parser"
)

/**
 * 语义解析的第三步：引用消解和类型推断
 * 1.解析所有的本地变量引用、函数调用和类成员引用。
 * 2.类型推断：从下而上推断表达式的类型。
 * 这两件事要放在一起做，因为：
 * (1)对于变量，只有做了消解，才能推断出类型来。
 * (2)对于FunctionCall，只有把参数（表达式)的类型都推断出来，才能匹配到正确的函数（方法)。
 * (3)表达式里包含FunctionCall,所以要推导表达式的类型，必须知道是哪个Function，从而才能得到返回值。
 */
type RefResolver struct {
	*BasePlayScriptListener
	at *AnnotatedTree
	// TODO
}

func (this *RefResolver) ExitPrimary(ctx *PrimaryContext) {
	scope := this.at.EnclosingScopeOfNode(ctx)
	var _type Type
	// 标识符
	if ctx.IDENTIFIER() != nil {
		idName := ctx.IDENTIFIER().GetText()
		variable := this.at.LookupVariable(scope, idName)
		if variable == nil {
			// 看看是不是函数，因为函数可以作为值来传递。这个时候，函数重名没法区分。
			// 因为普通Scope中的函数是不可以重名的，所以这应该是没有问题的。  //TODO 但如果允许重名，那就不行了。
			// TODO 注意，查找function的时候，可能会把类的方法包含进去
			function := this.at.LookupFunction(scope, idName)
			if function != nil {
				this.at.symbolOfNode[ctx] = function
				_type = function
			} else {
				fmt.Println("unknown variable or function: " + idName)
			}
		} else {
			this.at.symbolOfNode[ctx] = variable
			_type = variable._type
		}
	} else if ctx.Literal() != nil { // 字面量，字面量的类型在 ExitLiteral 中完成
		_type = this.at.typeOfNode[ctx.Literal()]
	} else if ctx.Expression() != nil { // 括号的表达式
		_type = this.at.typeOfNode[ctx.Expression()]
	}
	this.at.typeOfNode[ctx] = _type
}

func (this *RefResolver) ExitFunctionCall(ctx *FunctionCallContext) {
	// TODO this
	// TODO super

	// TODO 临时代码，支持println
	if ctx.IDENTIFIER().GetText() == "println" {
		return
	}

	idName := ctx.IDENTIFIER().GetText()
	found := false

	// TODO 看看是不是点符号表达式调用的，调用的是类的方法

	scope := this.at.EnclosingScopeOfNode(ctx)
	// 从当前 scope 逐级查找函数(或方法)
	if !found {
		function := this.at.LookupFunction(scope, idName)
		if function != nil {
			found = true
			this.at.symbolOfNode[ctx] = function
			this.at.typeOfNode[ctx] = function.returnType
		}
	}

	if !found {
		// TODO 看看是不是类的构造函数，用相同的名称查找一个class

		// 看看是不是一个函数型的变量
		variable := this.at.LookupFunctionVariable(scope, idName)
		if variable != nil {
			if functionType, ok := variable._type.(FunctionType); ok {
				found = true
				this.at.symbolOfNode[ctx] = variable
				this.at.typeOfNode[ctx] = functionType
			}
		}
	}
}

func (this *RefResolver) ExitExpression(ctx *ExpressionContext) {
	var _type Type

	// TODO class
	if ctx.GetBop() != nil && ctx.GetBop().GetTokenType() == PlayScriptParserDOT {
		// 这是个左递归，要不断的把左边的节点的计算结果存到node2Symbol，所以要在exitExpression里操作
		// symbol := this.at.symbolOfNode[ctx.Expression(0)]
	} else if ctx.Primary() != nil {
		symbol := this.at.symbolOfNode[ctx.Primary()]
		this.at.symbolOfNode[ctx] = symbol
	}

	// 类型推断和综合
	// 比如 a = 1; 对应 expression = expression -> expresion = primary -> expresion = literal
	if ctx.Primary() != nil {
		fmt.Println(ctx.GetText())
		fmt.Printf("%p\n", ctx)
		//变量引用冒泡： 如果下级是一个变量，往上冒泡传递，以便在点符号表达式中使用
		fmt.Println(this.at.typeOfNode[ctx.Primary()])
		_type = this.at.typeOfNode[ctx.Primary()]
	} else if ctx.FunctionCall() != nil {
		_type = this.at.typeOfNode[ctx.FunctionCall()]
	} else if ctx.GetBop() != nil && len(ctx.AllExpression()) >= 2 {
		type1 := this.at.typeOfNode[ctx.Expression(0)]
		type2 := this.at.typeOfNode[ctx.Expression(1)]
		switch ctx.GetBop().GetTokenType() {
		case PlayScriptParserADD:
			if type1 == String || type2 == String {
				_type = String
			} else if _, ok1 := type1.(*PrimitiveType); ok1 {
				if _, ok2 := type2.(*PrimitiveType); ok2 {
					_type = GetUpperType(type1, type2)
				}
			} else {
				this.at.LogError("operand should be PrimitiveType for additive and multiplicative operation", ctx)
			}
		case PlayScriptParserSUB, PlayScriptParserMUL, PlayScriptParserDIV:
			if _, ok1 := type1.(*PrimitiveType); ok1 {
				if _, ok2 := type2.(*PrimitiveType); ok2 {
					_type = GetUpperType(type1, type2)
				}
			} else {
				this.at.LogError("operand should be PrimitiveType for additive and multiplicative operation", ctx)
			}
		case PlayScriptParserEQUAL, PlayScriptParserNOTEQUAL,
			PlayScriptParserLE, PlayScriptParserLT, PlayScriptParserGE,
			PlayScriptParserGT, PlayScriptParserAND, PlayScriptParserOR,
			PlayScriptParserBANG:
			_type = Boolean
		case PlayScriptParserASSIGN, PlayScriptParserADD_ASSIGN,
			PlayScriptParserSUB_ASSIGN, PlayScriptParserMUL_ASSIGN,
			PlayScriptParserDIV_ASSIGN, PlayScriptParserAND_ASSIGN,
			PlayScriptParserOR_ASSIGN, PlayScriptParserXOR_ASSIGN,
			PlayScriptParserMOD_ASSIGN, PlayScriptParserLSHIFT_ASSIGN,
			PlayScriptParserRSHIFT_ASSIGN, PlayScriptParserURSHIFT_ASSIGN:
			_type = type1
		}
	}
	this.at.typeOfNode[ctx] = _type
}

// 变量初始化部分做类型推断
// 比如 int a = ***
// *** 这一部分即是
func (this *RefResolver) ExitVariableInitializer(ctx *VariableInitializerContext) {
	if ctx.Expression() != nil {
		this.at.typeOfNode[ctx] = this.at.typeOfNode[ctx.Expression()]
	}
}

//根据字面量来推断类型
func (this *RefResolver) ExitLiteral(ctx *LiteralContext) {
	var _type *PrimitiveType
	if ctx.BOOL_LITERAL() != nil {
		_type = Boolean
	} else if ctx.CHAR_LITERAL() != nil {
		_type = Char
	} else if ctx.NULL_LITERAL() != nil {
		_type = Null
	} else if ctx.STRING_LITERAL() != nil {
		_type = String
	} else if ctx.IntegerLiteral() != nil {
		_type = Integer
	} else if ctx.FloatLiteral() != nil {
		_type = Float
	}
	// TODO 这里存的是 literal 的 ctx
	// 但是在 TypeChecker 里面是拿 literal 的父 ctx 来查询 type 的，会返回 nil
	this.at.typeOfNode[ctx] = _type
}

// TODO 在结束扫描之前，把this()和super()构造函数消解掉
//func (this *RefResolver) ExitProg(ctx *ProgContext) {
//}

func NewRefResolver(at *AnnotatedTree) *RefResolver {
	return &RefResolver{at: at}
}

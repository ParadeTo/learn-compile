package playscript

import (
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
				// 比如
				// int a(){};
				// b = a // PrimaryContexts 引用了 a
				this.at.symbolOfNode[ctx] = function
				_type = function
			}
		}
	}
	// TODO 字面量
}

func NewRefResolver(at *AnnotatedTree) *RefResolver {
	return &RefResolver{at: at}
}

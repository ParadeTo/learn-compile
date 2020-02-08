package playscript

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	. "learn-compile/playscript/parser"
)

/**
 * 进行一些语义检查，包括：
 * 01.break 只能出现在循环语句中，或case语句中；
 *
 * 02.return语句
 * 02-01 函数声明了返回值，就一定要有return语句。除非返回值类型是void。
 * 02-02 类的构造函数里如果用到return，不能带返回值。
 * 02-03 return语句只能出现在函数里。
 * 02-04 返回值类型检查 -> (在TypeChecker里做）
 *
 * 03.左值
 * 03-01 标注左值（不标注就是右值)；
 * 03-02 检查表达式能否生成合格的左值。
 *
 * 04.类的声明不能在函数里（TODO 未来应该也可以，只不过对生存期有要求）
 *
 * 05.super()和this()，只能是构造函数中的第一句。  这个在RefResolver中实现了。
 *
 * 06.
 */
type SematicValidator struct {
	*BasePlayScriptListener
	at *AnnotatedTree
}

// break只能出现在循环语句或switch-case语句里
func (v *SematicValidator) checkBreak(ctx antlr.Tree) bool {
	parent := ctx.GetParent()
	if statementContext, ok := parent.(StatementContext); ok {
		return statementContext.FOR() != nil || statementContext.WHILE() != nil
	} else if _, ok := parent.(SwitchBlockStatementGroupContext); ok {
		return true
	} else if _, ok := parent.(FunctionDeclarationContext); ok {
		return false
	} else if parent == nil {
		return false
	} else {
		return v.checkBreak(parent)
	}
}

func (v *SematicValidator) hasReturnStatement(ctx antlr.Tree) bool {
	rtn := false
	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChild(i)
		if statementContext, ok := child.(*StatementContext); ok {
			if statementContext.RETURN() != nil {
				rtn = true
				break
			}
		} else {
			_, ok1 := child.(*FunctionDeclarationContext)
			_, ok2 := child.(*ClassDeclarationContext)
			// 从其他子语句中找 return，但是 function 和 class 中的不算
			if !(ok1 || ok2) {
				rtn = v.hasReturnStatement(child)
				if rtn {
					break
				}
			}
		}
	}
	return rtn
}

func (v *SematicValidator) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {
	//02-01 函数定义了返回值，就一定要有相应的return语句
	//TODO 更完善的是要进行控制流计算，不是仅仅有一个return语句就行了
	if ctx.TypeTypeOrVoid() != nil {
		if !v.hasReturnStatement(ctx) {
			returnType := v.at.typeOfNode[ctx.TypeTypeOrVoid()]
			if returnType != GetVoidType() {
				v.at.LogError("return statement expected in function", ctx)
			}
		}
	}
}

func (v *SematicValidator) ExitStatement(ctx *StatementContext) {
	//02
	if ctx.RETURN() != nil {
		// 02-03 return语句只能出现在函数里
		function := v.at.EnclosingFunctionOfNode(ctx)
		if function == nil {
			v.at.LogError("return statement not in function body", ctx)
		}
		//else if function.isConstructor() && ctx.expression()!=null
	} else if ctx.BREAK() != nil {
		if !v.checkBreak(ctx) {
			v.at.LogError("break statement not in loop or switch statements", ctx)
		}
	}
}

func NewSematicValidator(at *AnnotatedTree) *SematicValidator {
	return &SematicValidator{at: at}
}

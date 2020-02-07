package playscript

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	. "learn-compile/playscript/parser"
)

/**
 * 类型检查。
 * 主要检查:
 * 1.赋值表达式；
 * 2.变量初始化；
 * 3.表达式里的一些运算，比如加减乘除，是否类型匹配；
 * 4.返回值的类型；
 */
type TypeChecker struct {
	*BasePlayScriptListener
	at *AnnotatedTree
}

// 检查类型是不是数值型的
func (this *TypeChecker) checkNumericOperand(_type Type, exp *ExpressionContext, operand IExpressionContext) {
	if !IsNumeric(_type) {
		this.at.LogError("operand for arithmetic operation should be numeric : "+operand.GetText(), exp)
	}
}

// 检查类型是不是布尔型的
func (this *TypeChecker) checkBooleanOperand(_type Type, exp *ExpressionContext, operand IExpressionContext) {
	if _type != Boolean {
		this.at.LogError("operand for logical operation should be boolean : "+operand.GetText(), exp)
	}
}

/**
 * 看一个类型能否赋值成另一个类型，比如：
 * (1) 整型可以转成浮点型；
 * (2) 子类的对象可以赋给父类;
 * (3) 函数赋值，要求签名是一致的。
 * @param from
 * @param to
 * @return
 */
func (this *TypeChecker) checkNumericAssign(from, to Type) bool {
	canAssign := false
	if to == Double {
		canAssign = IsNumeric(from)
	} else if to == Float {
		canAssign = from == Byte ||
			from == Short ||
			from == Integer ||
			from == Long ||
			from == Float
	} else if to == Long {
		canAssign = from == Byte ||
			from == Short ||
			from == Integer ||
			from == Long
	} else if to == Integer {
		canAssign = from == Byte ||
			from == Short ||
			from == Integer
	} else if to == Short {
		canAssign = from == Byte ||
			from == Short
	} else if to == Byte {
		canAssign = from == Byte
	}
	return canAssign
}

func (this *TypeChecker) checkAssign(type1, type2 Type, ctx antlr.ParserRuleContext, operand1, operand2 antlr.ParserRuleContext) {
	if IsNumeric(type2) {
		if !this.checkNumericAssign(type2, type1) {
			this.at.LogError("cannot assign "+operand2.GetText()+" of type "+type2.ToString()+
				" to "+operand1.GetText()+" of type "+type1.ToString(), ctx)
		}
	}
	// 检查类的兼容性
	// 检查函数的兼容性
}

func (this *TypeChecker) ExitVariableDeclarator(ctx *VariableDeclaratorContext) {
	if ctx.VariableInitializer() != nil {
		symbol := this.at.symbolOfNode[ctx.VariableDeclaratorId()]
		if variable, ok := symbol.(*Variable); ok {
			type1 := variable._type
			type2 := this.at.typeOfNode[ctx.VariableInitializer()]
			this.checkAssign(type1, type2, ctx, ctx.VariableDeclaratorId(), ctx.VariableInitializer())
		}
	}
}

func (this *TypeChecker) ExitExpression(ctx *ExpressionContext) {
	if ctx.GetBop() != nil && len(ctx.AllExpression()) >= 2 {
		exp1 := ctx.Expression(0)
		exp2 := ctx.Expression(1)
		fmt.Printf("%p\n", exp1)
		fmt.Printf("%p\n", exp2)
		fmt.Println(this.at.typeOfNode)
		type1 := this.at.typeOfNode[exp1]
		type2 := this.at.typeOfNode[exp2]
		switch ctx.GetBop().GetTokenType() {
		case PlayScriptParserADD:
			//字符串能够跟任何对象做 + 运算，所以这里只需要检查两个操作数都不是字符串的情况
			if type1 != String && type2 != String {
				this.checkNumericOperand(type1, ctx, ctx.Expression(0))
				this.checkNumericOperand(type2, ctx, ctx.Expression(1))
			}
		case PlayScriptParserSUB, PlayScriptParserMUL, PlayScriptParserDIV,
			PlayScriptParserLE, PlayScriptParserLT, PlayScriptParserGE,
			PlayScriptParserGT:
			fmt.Println(ctx.GetText())
			this.checkNumericOperand(type1, ctx, ctx.Expression(0))
			this.checkNumericOperand(type2, ctx, ctx.Expression(1))
		case PlayScriptParserEQUAL, PlayScriptParserNOTEQUAL:
		case PlayScriptParserAND, PlayScriptParserOR:
			this.checkBooleanOperand(type1, ctx, ctx.Expression(0))
			this.checkBooleanOperand(type2, ctx, ctx.Expression(1))
		case PlayScriptParserASSIGN:
			this.checkAssign(type1, type2, ctx, ctx.Expression(0), ctx.Expression(1))
		case PlayScriptParserADD_ASSIGN, PlayScriptParserSUB_ASSIGN, PlayScriptParserMUL_ASSIGN,
			PlayScriptParserDIV_ASSIGN, PlayScriptParserAND_ASSIGN,
			PlayScriptParserOR_ASSIGN, PlayScriptParserXOR_ASSIGN,
			PlayScriptParserMOD_ASSIGN, PlayScriptParserLSHIFT_ASSIGN,
			PlayScriptParserRSHIFT_ASSIGN, PlayScriptParserURSHIFT_ASSIGN:
			if IsNumeric(type2) {
				if !this.checkNumericAssign(type2, type1) {
					this.at.LogError(
						fmt.Sprintf("can not assign %s of type %s to %s of type %s",
							ctx.Expression(1).GetText(), type2.ToString(), ctx.Expression(0), type1.ToString()), ctx)
				}
			} else {
				this.at.LogError("operand "+ctx.Expression(1).GetText()+" should be numeric.", ctx)
			}
		}
	}
	//TODO 对各种一元运算做类型检查，比如NOT操作
}

func NewTypeChecker(at *AnnotatedTree) *TypeChecker {
	return &TypeChecker{at: at}
}

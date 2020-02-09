package playscript

import (
	"fmt"
	. "learn-compile/playscript/parser"
	"strconv"
)

type ASTEvaluator struct {
	*BasePlayScriptVisitor
	at    *AnnotatedTree
	stack []*StackFrame
}

/**
 * 栈桢入栈。
 * 其中最重要的任务，是要保证栈桢的parentFrame设置正确。否则，
 * (1)随着栈的变深，查找变量的性能会降低；
 * (2)甚至有可能找错栈桢，比如在递归(直接或间接)的场景下。
 */
func (this *ASTEvaluator) pushStack(frame *StackFrame) {
	if len(this.stack) > 0 {
		//从栈顶到栈底依次查找
		for i := len(this.stack) - 1; i > 0; i-- {
			f := this.stack[i]

			/*
			   如果新加入的栈桢，跟某个已有的栈桢的enclosingScope是一样的，那么这俩的parentFrame也一样。
			   因为它们原本就是同一级的嘛。
			   比如：
			   void foo(){};
			   void bar(foo());

			   或者：
			   void foo();
			   if (...){
			       foo();
			   }
			*/
			if f.scope.GetEnclosingScope() == frame.scope.GetEnclosingScope() {
				frame.parentFrame = f.parentFrame
				break
			} else if f.scope == frame.scope.GetEnclosingScope() {
				/*
				   如果新加入的栈桢，是某个已有的栈桢的下一级，那么就把把这个父子关系建立起来。比如：
				   void foo(){
				       if (...){  //把这个块往栈桢里加的时候，就符合这个条件。
				       }
				   }
				   再比如,下面的例子:
				   class MyClass{
				       void foo();
				   }
				   MyClass c = MyClass();  //先加Class的栈桢，里面有类的属性，包括父类的
				   c.foo();                //再加foo()的栈桢
				*/
				frame.parentFrame = f
				break
			} else if functionObject, ok := frame.object.(*FunctionObject); ok {
				/*
				   这是针对函数可能是一等公民的情况。这个时候，函数运行时的作用域，与声明时的作用域会不一致。
				   我在这里设计了一个“receiver”的机制，意思是这个函数是被哪个变量接收了。要按照这个receiver的作用域来判断。
				*/
				if functionObject.receiver != nil && functionObject.receiver.GetEnclosingScope() == f.scope {
					frame.parentFrame = f
					break
				}
			}
		}

		// 否则把栈顶作为 parentFrame
		if frame.parentFrame == nil {
			frame.parentFrame = this.stack[len(this.stack)-1]
		}
	}

	this.stack = append(this.stack, frame)

	// 打印 stackFrame
}

func (this *ASTEvaluator) dumpStackFrame() {
	fmt.Println("\nStack Frames --------")
	for _, frame := range this.stack {
		fmt.Println(frame.ToString())
	}
	fmt.Println("-----------------------")

}

func (this *ASTEvaluator) popStack() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *ASTEvaluator) GetLValue(variable *Variable) LValue {
	f := this.stack[len(this.stack)-1]
	var valueContainer PlayObject
	for {
		if f == nil {
			break
		}
		if f.scope.ContainsSymbol(variable) {
			valueContainer = f.object
			break
		}
		f = f.parentFrame
	}

	//通过正常的作用域找不到，就从闭包里找
	//原理：PlayObject中可能有一些变量，其作用域跟StackFrame.scope是不同的
	if valueContainer == nil {
		f = this.stack[len(this.stack)-1]
		for {
			if f == nil {
				break
			}
			if f.Contains(variable) {
				valueContainer = f.object
				break
			}
			f = f.parentFrame
		}
	}

	lvalue := NewMyLValue(valueContainer, variable)
	return lvalue
}

//---- visit ----
func (this *ASTEvaluator) VisitBlock(ctx *BlockContext) interface{} {
	scope := this.at.node2Scope[ctx]
	if scope != nil { // 有些block是不对应scope的，比如函数底下的block.
		if blockScope, ok := scope.(*BlockScope); ok {
			frame := NewFrameForBlockScope(blockScope)
			this.pushStack(frame)
		}
	}

	//rtn :=
}

func (this *ASTEvaluator) VisitBlockStatements(ctx *BlockStatementsContext) interface{} {
	var rtn interface{}
	for _, child := range ctx.AllBlockStatement() {
		rtn
	}
}

func (this *ASTEvaluator) VisitBlockStatement(ctx *BlockStatementContext) interface{} {
	var rtn interface{}
	if ctx.VariableDeclarators() != nil {

	} else if ctx.Statement() != nil {

	}
	return rtn
}

func (this *ASTEvaluator) VisitStatement(ctx *StatementContext) interface{} {
	var rtn interface{}
	if ctx.GetStatementExpression() != nil {
		//if statementExpressionCtx, ok := ctx.GetStatementExpression().(*ExpressionContext)
		//rtn = this.VisitExpression()
	} else if ctx.IF() != nil {
		//condition :=
	}
}

func (this *ASTEvaluator) VisitParExpression(ctx *ParExpressionContext) interface{} {
	if 
	return this.VisitExpression()
}

func (this *ASTEvaluator) VisitVariableDeclarators(ctx *VariableDeclaratorsContext) interface{} {
}

func (this *ASTEvaluator) VisitVariableDeclarator(ctx *VariableDeclaratorContext) interface{} {
}

func (this *ASTEvaluator) VisitVariableInitializer(ctx *VariableInitializerContext) interface{} {
}

func (this *ASTEvaluator) VisitExpression(ctx *ExpressionContext) interface{} {
	var rtn interface{}
	if ctx.GetBop() != nil && len(ctx.AllExpression()) >= 2 {
		exp1, ok1 := ctx.Expression(0).(*ExpressionContext)
		exp2, ok2 := ctx.Expression(1).(*ExpressionContext)
		if !(ok1 && ok2) {
			fmt.Println("VisitExpression exp1 or exp2 cannot convert")
		}
		left := this.VisitExpression(exp1)
		right := this.VisitExpression(exp2)
		leftObject := left
		rightObject := right

		if value, ok := left.(LValue); ok {
			leftObject = value.GetValue()
		}

		if value, ok := right.(LValue); ok {
			right = value.GetValue()
		}

		// 本节点期待的数据类型
		// 在RefResolver中推断出来的
		_type := this.at.typeOfNode[ctx]

		// 左右两个子节点的类型
		type1 := this.at.typeOfNode[exp1]
		type2 := this.at.typeOfNode[exp2]

		switch ctx.GetBop().GetTokenType() {
		case PlayScriptParserADD:
			rtn = this.add(leftObject, rightObject, _type)
		case PlayScriptParserSUB:
			rtn = this.minue(leftObject, rightObject, _type)
		case PlayScriptParserMUL:
			rtn = this.mul(leftObject, rightObject, _type)
		case PlayScriptParserDIV:
			rtn = this.div(leftObject, rightObject, _type)
		case PlayScriptParserEQUAL:
			rtn = this.EQ(leftObject, rightObject, GetUpperType(type1, type2))
		case PlayScriptParserNOTEQUAL:
			rtn = !this.EQ(leftObject, rightObject, GetUpperType(type1, type2))
		case PlayScriptParserLE:
			rtn = this.LE(leftObject, rightObject, GetUpperType(type1, type2))
		case PlayScriptParserLT:
			rtn = this.LT(leftObject, rightObject, GetUpperType(type1, type2))
		case PlayScriptParserGE:
			rtn = this.GE(leftObject, rightObject, GetUpperType(type1, type2))
		case PlayScriptParserGT:
			rtn = this.GT(leftObject, rightObject, GetUpperType(type1, type2))
		case PlayScriptParserAND:
			b1, _ := leftObject.(bool)
			b2, _ := rightObject.(bool)
			rtn = b1 && b2
		case PlayScriptParserOR:
			b1, _ := leftObject.(bool)
			b2, _ := rightObject.(bool)
			rtn = b1 || b2
		case PlayScriptParserASSIGN:
			if left, ok := leftObject.(LValue); ok {
				left.SetValue(rightObject)
				rtn = right
			} else {
				fmt.Println("Unsupported feature during assignment")
			}
		}
	} else if ctx.GetBop() != nil && ctx.GetBop().GetTokenType() == PlayScriptParserDOT {
		// 此语法是左递归的，算法体现这一点
		// TODO class
	} else if ctx.Primary() != nil {
		if primaryCtx, ok := ctx.Primary().(*PrimaryContext); ok {
			rtn = this.VisitPrimary(primaryCtx)
		}
	} else if ctx.GetPostfix() != nil { // 后缀运算，如 i++ 或 i--
		var value interface{}
		var lValue LValue
		var _type Type
		if expCtx, ok := ctx.Expression(0).(*ExpressionContext); ok {
			value = this.VisitExpression(expCtx)
			_type = this.at.typeOfNode[expCtx]
		}
		if lv, ok := value.(LValue); ok {
			lValue = lv
			value = lValue.GetValue()
		}
		switch ctx.GetPostfix().GetTokenType() {
		case PlayScriptParserINC:
			// short 呢
			if _type == Integer {
				lValue.SetValue(value.(int32) + 1)
			} else {
				lValue.SetValue(value.(int64) + 1)
			}
			rtn = value
		case PlayScriptParserDEC:
			if _type == Integer {
				lValue.SetValue(value.(int32) - 1)
			} else {
				lValue.SetValue(value.(int64) - 1)
			}
			rtn = value
		}
	} else if ctx.GetPrefix() != nil { // 前缀操作，如 ++i 或 --i
		var value interface{}
		var lValue LValue
		var _type Type
		if expCtx, ok := ctx.Expression(0).(*ExpressionContext); ok {
			value = this.VisitExpression(expCtx)
			_type = this.at.typeOfNode[expCtx]
		}
		if lv, ok := value.(LValue); ok {
			lValue = lv
			value = lValue.GetValue()
		}
		switch ctx.GetPostfix().GetTokenType() {
		case PlayScriptParserINC:
			// short 呢
			if _type == Integer {
				rtn = value.(int32) + 1
			} else {
				rtn = value.(int64) + 1
			}
			lValue.SetValue(rtn)
		case PlayScriptParserDEC:
			if _type == Integer {
				rtn = value.(int32) - 1
			} else {
				rtn = value.(int64) - 1
			}
			lValue.SetValue(rtn)
		}
	} else if ctx.FunctionCall() != nil {
		if functionCallContext, ok := ctx.FunctionCall().(*FunctionCallContext); ok {
			rtn = this.VisitFunctionCall(functionCallContext)
		}
	}

	return rtn
}

func (this *ASTEvaluator) VisitPrimary(ctx *PrimaryContext) interface{} {
	var rtn interface{}
	// 字面量
	if ctx.Literal() != nil {
		if literalCtx, ok := ctx.Literal().(*LiteralContext); ok {
			rtn = this.VisitLiteral(literalCtx)
		}
	} else if ctx.IDENTIFIER() != nil { // 变量
		symbol := this.at.symbolOfNode[ctx]
		if variable, ok := symbol.(*Variable); ok {
			rtn = this.GetLValue(variable)
		} else if function, ok := symbol.(*Function); ok {
			obj := NewFunctionObject(function)
			rtn = obj
		}
	} else if ctx.Expression() != nil { // 括号括起来的表达式
		if expressionCtx, ok := ctx.Expression().(*ExpressionContext); ok {
			rtn = this.VisitExpression(expressionCtx)
		}
	} else if ctx.THIS() != nil { // this

	} else if ctx.SUPER() != nil { // super

	}
	return rtn
}

func (this *ASTEvaluator) VisitLiteral(ctx *LiteralContext) interface{} {
	var rtn interface{}
	// 整数
	if ctx.IntegerLiteral() != nil {
		ctx, ok := ctx.IntegerLiteral().(*IntegerLiteralContext)
		if ok {
			rtn = this.VisitIntegerLiteral(ctx)
		}
	} else if ctx.FloatLiteral() != nil { 	// 浮点数
		ctx, ok := ctx.FloatLiteral().(*FloatLiteralContext)
		if ok {
			rtn = this.VisitFloatLiteral(ctx)
		}
	} else if ctx.BOOL_LITERAL() != nil { // 布尔值
		if ctx.BOOL_LITERAL().GetText() == "true" {
			rtn = true
		} else {
			rtn = false
		}
	} else if ctx.STRING_LITERAL() != nil { // 字符串
		strWithQuotationMark := ctx.STRING_LITERAL().GetText()
		rtn = strWithQuotationMark[1:len(strWithQuotationMark)-1]
	} else if ctx.CHAR_LITERAL() != nil { // 单个字符
		//rtn = ctx.CHAR_LITERAL().GetText()
	} else if ctx.NULL_LITERAL() != nil { // null 字面量
		rtn = GetNullObject()
	}
	return rtn
}

func (this *ASTEvaluator) VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{} {
	var rtn interface{}
	decimalLiteralCtx := ctx.DECIMAL_LITERAL()
	if decimalLiteralCtx != nil {
		if decimal, error := strconv.Atoi(decimalLiteralCtx.GetText()); error == nil {
			rtn = decimal
		}
	}
	return rtn
}

func (this *ASTEvaluator) VisitFloatLiteral(ctx *FloatLiteralContext) interface{} {
	var rtn interface{}
	if float, error := strconv.ParseFloat(ctx.GetText(), 32); error == nil {
		rtn = float32(float)
	}
	return rtn
}

// -----各种运算-----
func (this *ASTEvaluator) add(obj1, obj2 interface{}, targetType Type) interface{} {
	var rtn interface{}
	if targetType == String {
		str1, ok1 := obj1.(string)
		str2, ok2 := obj2.(string)
		if ok1 && ok2 {
			rtn = str1 + str2
		}
	} else if targetType == Integer {
		int1, ok1 := obj1.(int32)
		int2, ok2 := obj2.(int32)
		if ok1 && ok2 {
			rtn = int1 + int2
		}
	} else if targetType == Short {
		int1, ok1 := obj1.(int16)
		int2, ok2 := obj2.(int16)
		if ok1 && ok2 {
			rtn = int1 + int2
		}
	} else if targetType == Long {
		int1, ok1 := obj1.(int64)
		int2, ok2 := obj2.(int64)
		if ok1 && ok2 {
			rtn = int1 + int2
		}
	} else if targetType == Float {
		float1, ok1 := obj1.(float32)
		float2, ok2 := obj2.(float32)
		if ok1 && ok2 {
			rtn = float1 + float2
		}
	} else if targetType == Double {
		double1, ok1 := obj1.(float64)
		double2, ok2 := obj2.(float64)
		if ok1 && ok2 {
			rtn = double1 + double2
		}
	} else {
		fmt.Println("unsupported add operation")
	}

	return rtn
}

func (this *ASTEvaluator) minue(obj1, obj2 interface{}, targetType Type) interface{} {
	var rtn interface{}
	if targetType == Integer {
		int1, ok1 := obj1.(int32)
		int2, ok2 := obj2.(int32)
		if ok1 && ok2 {
			rtn = int1 - int2
		}
	} else if targetType == Short {
		int1, ok1 := obj1.(int16)
		int2, ok2 := obj2.(int16)
		if ok1 && ok2 {
			rtn = int1 - int2
		}
	} else if targetType == Long {
		int1, ok1 := obj1.(int64)
		int2, ok2 := obj2.(int64)
		if ok1 && ok2 {
			rtn = int1 - int2
		}
	} else if targetType == Float {
		float1, ok1 := obj1.(float32)
		float2, ok2 := obj2.(float32)
		if ok1 && ok2 {
			rtn = float1 - float2
		}
	} else if targetType == Double {
		double1, ok1 := obj1.(float64)
		double2, ok2 := obj2.(float64)
		if ok1 && ok2 {
			rtn = double1 - double2
		}
	} else {
		fmt.Println("unsupported minus operation")
	}

	return rtn
}

func (this *ASTEvaluator) mul(obj1, obj2 interface{}, targetType Type) interface{} {
	var rtn interface{}
	if targetType == Integer {
		int1, ok1 := obj1.(int32)
		int2, ok2 := obj2.(int32)
		if ok1 && ok2 {
			rtn = int1 * int2
		}
	} else if targetType == Short {
		int1, ok1 := obj1.(int16)
		int2, ok2 := obj2.(int16)
		if ok1 && ok2 {
			rtn = int1 * int2
		}
	} else if targetType == Long {
		int1, ok1 := obj1.(int64)
		int2, ok2 := obj2.(int64)
		if ok1 && ok2 {
			rtn = int1 * int2
		}
	} else if targetType == Float {
		float1, ok1 := obj1.(float32)
		float2, ok2 := obj2.(float32)
		if ok1 && ok2 {
			rtn = float1 * float2
		}
	} else if targetType == Double {
		double1, ok1 := obj1.(float64)
		double2, ok2 := obj2.(float64)
		if ok1 && ok2 {
			rtn = double1 * double2
		}
	} else {
		fmt.Println("unsupported mul operation")
	}

	return rtn
}

func (this *ASTEvaluator) div(obj1, obj2 interface{}, targetType Type) interface{} {
	var rtn interface{}
	if targetType == Integer {
		int1, ok1 := obj1.(int32)
		int2, ok2 := obj2.(int32)
		if ok1 && ok2 {
			rtn = int1 / int2
		}
	} else if targetType == Short {
		int1, ok1 := obj1.(int16)
		int2, ok2 := obj2.(int16)
		if ok1 && ok2 {
			rtn = int1 / int2
		}
	} else if targetType == Long {
		int1, ok1 := obj1.(int64)
		int2, ok2 := obj2.(int64)
		if ok1 && ok2 {
			rtn = int1 / int2
		}
	} else if targetType == Float {
		float1, ok1 := obj1.(float32)
		float2, ok2 := obj2.(float32)
		if ok1 && ok2 {
			rtn = float1 / float2
		}
	} else if targetType == Double {
		double1, ok1 := obj1.(float64)
		double2, ok2 := obj2.(float64)
		if ok1 && ok2 {
			rtn = double1 / double2
		}
	} else {
		fmt.Println("unsupported div operation")
	}

	return rtn
}

func (this *ASTEvaluator) EQ(obj1, obj2 interface{}, targetType Type) bool {
	var rtn bool
	if targetType == Integer {
		int1, ok1 := obj1.(int32)
		int2, ok2 := obj2.(int32)
		if ok1 && ok2 {
			rtn = int1 == int2
		}
	} else if targetType == Short {
		int1, ok1 := obj1.(int16)
		int2, ok2 := obj2.(int16)
		if ok1 && ok2 {
			rtn = int1 == int2
		}
	} else if targetType == Long {
		int1, ok1 := obj1.(int64)
		int2, ok2 := obj2.(int64)
		if ok1 && ok2 {
			rtn = int1 == int2
		}
	} else if targetType == Float {
		float1, ok1 := obj1.(float32)
		float2, ok2 := obj2.(float32)
		if ok1 && ok2 {
			rtn = float1 == float2
		}
	} else if targetType == Double {
		double1, ok1 := obj1.(float64)
		double2, ok2 := obj2.(float64)
		if ok1 && ok2 {
			rtn = double1 == double2
		}
	} else {
		rtn = obj1 == obj2
	}

	return rtn
}

func (this *ASTEvaluator) GE(obj1, obj2 interface{}, targetType Type) bool {
	rtn := false
	if targetType == Integer {
		int1, ok1 := obj1.(int32)
		int2, ok2 := obj2.(int32)
		if ok1 && ok2 {
			rtn = int1 >= int2
		}
	} else if targetType == Short {
		int1, ok1 := obj1.(int16)
		int2, ok2 := obj2.(int16)
		if ok1 && ok2 {
			rtn = int1 >= int2
		}
	} else if targetType == Long {
		int1, ok1 := obj1.(int64)
		int2, ok2 := obj2.(int64)
		if ok1 && ok2 {
			rtn = int1 >= int2
		}
	} else if targetType == Float {
		float1, ok1 := obj1.(float32)
		float2, ok2 := obj2.(float32)
		if ok1 && ok2 {
			rtn = float1 >= float2
		}
	} else if targetType == Double {
		double1, ok1 := obj1.(float64)
		double2, ok2 := obj2.(float64)
		if ok1 && ok2 {
			rtn = double1 >= double2
		}
	}

	return rtn
}

func (this *ASTEvaluator) GT(obj1, obj2 interface{}, targetType Type) bool {
	rtn := false
	if targetType == Integer {
		int1, ok1 := obj1.(int32)
		int2, ok2 := obj2.(int32)
		if ok1 && ok2 {
			rtn = int1 > int2
		}
	} else if targetType == Short {
		int1, ok1 := obj1.(int16)
		int2, ok2 := obj2.(int16)
		if ok1 && ok2 {
			rtn = int1 > int2
		}
	} else if targetType == Long {
		int1, ok1 := obj1.(int64)
		int2, ok2 := obj2.(int64)
		if ok1 && ok2 {
			rtn = int1 > int2
		}
	} else if targetType == Float {
		float1, ok1 := obj1.(float32)
		float2, ok2 := obj2.(float32)
		if ok1 && ok2 {
			rtn = float1 > float2
		}
	} else if targetType == Double {
		double1, ok1 := obj1.(float64)
		double2, ok2 := obj2.(float64)
		if ok1 && ok2 {
			rtn = double1 > double2
		}
	}

	return rtn
}

func (this *ASTEvaluator) LE(obj1, obj2 interface{}, targetType Type) bool {
	rtn := false
	if targetType == Integer {
		int1, ok1 := obj1.(int32)
		int2, ok2 := obj2.(int32)
		if ok1 && ok2 {
			rtn = int1 <= int2
		}
	} else if targetType == Short {
		int1, ok1 := obj1.(int16)
		int2, ok2 := obj2.(int16)
		if ok1 && ok2 {
			rtn = int1 <= int2
		}
	} else if targetType == Long {
		int1, ok1 := obj1.(int64)
		int2, ok2 := obj2.(int64)
		if ok1 && ok2 {
			rtn = int1 <= int2
		}
	} else if targetType == Float {
		float1, ok1 := obj1.(float32)
		float2, ok2 := obj2.(float32)
		if ok1 && ok2 {
			rtn = float1 <= float2
		}
	} else if targetType == Double {
		double1, ok1 := obj1.(float64)
		double2, ok2 := obj2.(float64)
		if ok1 && ok2 {
			rtn = double1 <= double2
		}
	}

	return rtn
}

func (this *ASTEvaluator) LT(obj1, obj2 interface{}, targetType Type) bool {
	rtn := false
	if targetType == Integer {
		int1, ok1 := obj1.(int32)
		int2, ok2 := obj2.(int32)
		if ok1 && ok2 {
			rtn = int1 < int2
		}
	} else if targetType == Short {
		int1, ok1 := obj1.(int16)
		int2, ok2 := obj2.(int16)
		if ok1 && ok2 {
			rtn = int1 < int2
		}
	} else if targetType == Long {
		int1, ok1 := obj1.(int64)
		int2, ok2 := obj2.(int64)
		if ok1 && ok2 {
			rtn = int1 < int2
		}
	} else if targetType == Float {
		float1, ok1 := obj1.(float32)
		float2, ok2 := obj2.(float32)
		if ok1 && ok2 {
			rtn = float1 < float2
		}
	} else if targetType == Double {
		double1, ok1 := obj1.(float64)
		double2, ok2 := obj2.(float64)
		if ok1 && ok2 {
			rtn = double1 < double2
		}
	}

	return rtn
}

func NewASTEvaluator(at *AnnotatedTree) *ASTEvaluator {
	return &ASTEvaluator{at: at}
}

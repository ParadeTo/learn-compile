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
	// 是否保存打印的信息
	cachePrintln bool
	printlnArr   []interface{}
}

// 执行一个函数的方法体。需要先设置参数值，然后再执行代码。
func (this *ASTEvaluator) functionCall(functionObject *FunctionObject, paramValues []interface{}) interface{} {
	var rtn interface{}

	// 添加函数栈帧
	functionFrame := NewFrameForFunction(functionObject)
	this.pushStack(functionFrame)

	// 给参数赋值，这些值进入 functionFrame
	functionCtx := functionObject.function.ctx.(*FunctionDeclarationContext)
	formalParametersContext := functionCtx.FormalParameters().(*FormalParametersContext)
	fmt.Println(formalParametersContext)
	fmt.Println(formalParametersContext.FormalParameterList())
	if formalParametersContext.FormalParameterList() != nil {
		formalParameterListCtx := formalParametersContext.FormalParameterList().(*FormalParameterListContext)
		for i, iCtx := range formalParameterListCtx.AllFormalParameter() {
			paramCtx := iCtx.(*FormalParameterContext)
			lvalue := this.VisitVariableDeclaratorId(paramCtx.VariableDeclaratorId().(*VariableDeclaratorIdContext)).(LValue)
			lvalue.SetValue(paramValues[i])
		}
	}

	// 调用函数（方法）体
	rtn = this.VisitFunctionDeclaration(functionCtx)

	this.popStack()

	if returnObject, ok := rtn.(*ReturnObject); ok {
		rtn = returnObject.returnValue
	}
	return rtn
}

func (this *ASTEvaluator) calcParamValues(ctx *FunctionCallContext) []interface{} {
	var paramValues []interface{}
	if ctx.ExpressionList() != nil {
		for _, exp := range ctx.ExpressionList().(*ExpressionListContext).AllExpression() {
			value := this.VisitExpression(exp.(*ExpressionContext))
			if lvalue, ok := value.(LValue); ok {
				value = lvalue.GetValue()
			}
			paramValues = append(paramValues, value)
		}
	}
	return paramValues
}

/**
 * 根据函数调用的上下文，返回一个FunctionObject。
 * 对于函数类型的变量，这个functionObject是存在变量里的；
 * 对于普通的函数调用，此时创建一个。
 * @param ctx
 * @return
 */
func (this *ASTEvaluator) getFunctionObject(ctx *FunctionCallContext) *FunctionObject {
	if ctx.IDENTIFIER() == nil {
		return nil
	}

	var function *Function
	var functionObject *FunctionObject

	symbol := this.at.symbolOfNode[ctx]

	// 函数类型的变量
	if variable, ok := symbol.(*Variable); ok {
		lvalue := this.GetLValue(variable)
		value := lvalue.GetValue()
		if _functionObject, ok := value.(*FunctionObject); ok {
			functionObject = _functionObject
			function = functionObject.function
		}
	} else if _function, ok := symbol.(*Function); ok { // 普通函数
		function = _function
	} else {
		functionName := ctx.IDENTIFIER().GetText()
		this.at.LogError("unable to find function or function variable "+functionName, ctx)
		return nil
	}

	if functionObject == nil {
		functionObject = NewFunctionObject(function)
	}

	return functionObject
}

func (this *ASTEvaluator) println(ctx *FunctionCallContext) {
	expList := ctx.ExpressionList()
	if expList != nil {
		if expListCtx, ok := expList.(*ExpressionListContext); ok {
			value := this.VisitExpressionList(expListCtx)
			if lvalue, ok := value.(LValue); ok {
				//fmt.Printf("%p %d\n", lvalue.GetVariable(), lvalue.GetValue())
				value = lvalue.GetValue()
			}
			if this.cachePrintln {
				this.printlnArr = append(this.printlnArr, value)
			}
			fmt.Println(value)
		}
	} else {
		fmt.Println()
	}
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
	//fmt.Printf("variable=%p\n", variable)
	for {
		if f == nil {
			break
		}
		//fmt.Println(f.scope.ToString())
		//for _, symbol := range f.scope.GetSymbols() {
		//	if a, ok := symbol.(*Variable); ok {
		//		fmt.Printf("%p %+v %+v\n", a, a, a.GetName())
		//	}
		//	if a, ok := symbol.(*BlockScope); ok {
		//		fmt.Println(a.GetName())
		//	}
		//}
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

func (this *ASTEvaluator) convertToBool(value interface{}) bool {
	if string, ok := value.(string); ok {
		return string != ""
	} else if i, ok := value.(int); ok {
		return i != 0
	} else if float32, ok := value.(float32); ok {
		return float32 != 0
	} else if bool, ok := value.(bool); ok {
		return bool
	} else {
		return value != nil
	}
}

/**
* 为闭包获取环境变量的值
* @param function 闭包所关联的函数。这个函数会访问一些环境变量。
* @param valueContainer 存放环境变量的值的容器
 */
func (this *ASTEvaluator) getClosureValuesForFunction(function *Function, valueContainer PlayObject) {
	if function.closureVariables != nil {
		function.closureVariables.ForEach(func(item interface{}) {
			if _var, ok := item.(*Variable); ok {
				lvalue := this.GetLValue(_var) // 现在还可以从栈里取，退出函数以后就不行了
				value := lvalue.GetValue()
				valueContainer.SetValue(_var, value)
			}
		})
	}
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

	rtn := this.VisitBlockStatements(ctx.BlockStatements().(*BlockStatementsContext))

	if scope != nil {
		this.popStack()
	}

	return rtn
}

func (this *ASTEvaluator) VisitBlockStatements(ctx *BlockStatementsContext) interface{} {
	var rtn interface{}
	for _, child := range ctx.AllBlockStatement() {
		rtn = this.VisitBlockStatement(child.(*BlockStatementContext))
		//如果返回的是break，那么不执行下面的语句, todo 表示怀疑
		if _, ok := rtn.(*BreakObject); ok {
			break
		} else if _, ok := rtn.(*ReturnObject); ok { //碰到return，退出函数
			// TODO 要能层层退出一个个block，弹出一个栈桢
			break
		}
	}
	return rtn
}

func (this *ASTEvaluator) VisitBlockStatement(ctx *BlockStatementContext) interface{} {
	var rtn interface{}
	// 这里没有处理 functionDeclaration 和 classDeclaration
	// 因为这里是求值阶段，这两个语句不会有返回值
	if ctx.VariableDeclarators() != nil {
		rtn = this.VisitVariableDeclarators(ctx.VariableDeclarators().(*VariableDeclaratorsContext))
	} else if ctx.Statement() != nil {
		rtn = this.VisitStatement(ctx.Statement().(*StatementContext))
	}
	return rtn
}

func (this *ASTEvaluator) VisitStatement(ctx *StatementContext) interface{} {
	var rtn interface{}
	if ctx.GetStatementExpression() != nil {
		statementExpressionCtx := ctx.GetStatementExpression().(*ExpressionContext)
		rtn = this.VisitExpression(statementExpressionCtx)
	} else if ctx.IF() != nil {
		condition := this.convertToBool(this.VisitParExpression(ctx.ParExpression().(*ParExpressionContext)))
		if condition == true {
			rtn = this.VisitStatement(ctx.Statement(0).(*StatementContext))
		} else if ctx.ELSE() != nil {
			rtn = this.VisitStatement(ctx.Statement(1).(*StatementContext))
		}
	} else if ctx.WHILE() != nil {
		parExpressionCtx := ctx.ParExpression().(*ParExpressionContext)
		if parExpressionCtx.Expression() != nil && ctx.Statement(0) != nil {
			for {
				//每次循环都要计算一下循环条件
				condition := true
				value := this.VisitExpression(parExpressionCtx.Expression().(*ExpressionContext))
				if lvalue, ok := value.(LValue); ok {
					condition = this.convertToBool(lvalue.GetValue())
				} else {
					condition = this.convertToBool(value)
				}

				if condition {
					rtn = this.VisitStatement(ctx.Statement(0).(*StatementContext))
					if _, ok := rtn.(*BreakObject); ok {
						rtn = nil
						break
					} else if _, ok := rtn.(*ReturnObject); ok {
						break
					}
				} else {
					break
				}
			}
		}
	} else if ctx.FOR() != nil {
		scope := this.at.node2Scope[ctx].(*BlockScope)
		frame := NewFrameForBlockScope(scope)
		this.pushStack(frame)

		forControl := ctx.ForControl().(*ForControlContext)
		if forControl.EnhancedForControl() != nil {

		} else {
			// 初始化部分执行一次
			if forControl.ForInit() != nil {
				rtn = this.VisitForInit(forControl.ForInit().(*ForInitContext))
			}

			for {
				// 如果没有条件判断部分，意味着一直循环
				condition := true
				if forControl.Expression() != nil {
					value := this.VisitExpression(forControl.Expression().(*ExpressionContext))
					if lvalue, ok := value.(LValue); ok {
						condition = this.convertToBool(lvalue.GetValue())
					} else {
						condition = this.convertToBool(value)
					}
				}

				if condition {
					// 执行for的语句体
					rtn = this.VisitStatement(ctx.Statement(0).(*StatementContext))
					if _, ok := rtn.(*BreakObject); ok {
						rtn = nil
						break
					} else if _, ok := rtn.(*ReturnObject); ok {
						break
					}

					// 执行forUpdate，通常是“i++”这样的语句。这个执行顺序不能出错。
					if forControl.GetForUpdate() != nil {
						this.VisitExpressionList(forControl.GetForUpdate().(*ExpressionListContext))
					}
				} else {
					break
				}
			}
		}
		this.popStack()
	} else if ctx.GetBlockLabel() != nil { // block
		rtn = this.VisitBlock(ctx.GetBlockLabel().(*BlockContext))
	} else if ctx.BREAK() != nil {
		rtn = GetBreakObject()
	} else if ctx.RETURN() != nil {
		if ctx.Expression() != nil {
			rtn = this.VisitExpression(ctx.Expression().(*ExpressionContext))
			//return语句应该不需要左值   //TODO 取左值的场景需要优化，目前都是取左值。
			if lvalue, ok := rtn.(LValue); ok {
				rtn = lvalue.GetValue()
			}
			// 把闭包涉及的环境变量都打包带走
			if functionObject, ok := rtn.(*FunctionObject); ok {
				this.getClosureValuesForFunction(functionObject.function, functionObject)
			} else if false { // TODO class

			}
		}

		//把真实的返回值封装在一个ReturnObject对象里，告诉visitBlockStatements停止执行下面的语句
		rtn = NewReturnObject(rtn)
	}
	return rtn
}

func (this *ASTEvaluator) VisitParExpression(ctx *ParExpressionContext) interface{} {
	if expCtx, ok := ctx.Expression().(*ExpressionContext); ok {
		return this.VisitExpression(expCtx)
	}
	return nil
}

func (this *ASTEvaluator) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	var rtn interface{}
	for _, child := range ctx.AllExpression() {
		rtn = this.VisitExpression(child.(*ExpressionContext))
	}
	return rtn
}

func (this *ASTEvaluator) VisitForInit(ctx *ForInitContext) interface{} {
	var rtn interface{}
	if ctx.VariableDeclarators() != nil {
		rtn = this.VisitVariableDeclarators(ctx.VariableDeclarators().(*VariableDeclaratorsContext))
	} else if ctx.ExpressionList() != nil {
		rtn = this.VisitExpressionList(ctx.ExpressionList().(*ExpressionListContext))
	}
	return rtn
}

func (this *ASTEvaluator) VisitVariableDeclarators(ctx *VariableDeclaratorsContext) interface{} {
	var rtn interface{}
	for _, child := range ctx.AllVariableDeclarator() {
		if variableDeclaratorContext, ok := child.(*VariableDeclaratorContext); ok {
			rtn = this.VisitVariableDeclarator(variableDeclaratorContext)
		}
	}
	return rtn
}

func (this *ASTEvaluator) VisitVariableDeclarator(ctx *VariableDeclaratorContext) interface{} {
	var rtn interface{}
	value := this.VisitVariableDeclaratorId(ctx.VariableDeclaratorId().(*VariableDeclaratorIdContext))
	if ctx.VariableInitializer() != nil {
		rtn = this.VisitVariableInitializer(ctx.VariableInitializer().(*VariableInitializerContext))
		if lvalue, ok := rtn.(LValue); ok {
			rtn = lvalue.GetValue()
		}
		value.(LValue).SetValue(rtn)
		//fmt.Printf("%p %d\n", value.(LValue).GetVariable(), value.(LValue).GetValue())
	}
	return rtn
}

func (this *ASTEvaluator) VisitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) interface{} {
	var rtn interface{}
	symbol := this.at.symbolOfNode[ctx]
	if variable, ok := symbol.(*Variable); ok {
		rtn = this.GetLValue(variable)
	}
	return rtn
}

func (this *ASTEvaluator) VisitVariableInitializer(ctx *VariableInitializerContext) interface{} {
	var rtn interface{}
	if ctx.Expression() != nil {
		if expCtx, ok := ctx.Expression().(*ExpressionContext); ok {
			//fmt.Println(expCtx.GetText())
			rtn = this.VisitExpression(expCtx)
		}
	}
	return rtn
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
			rightObject = value.GetValue()
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
			if left, ok := left.(LValue); ok {
				left.SetValue(rightObject)
				//fmt.Printf("%p %d\n", left.GetVariable(), left.GetValue())
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
			// 暂时不支持 long 和 short
			if _type == Integer {
				lValue.SetValue(value.(int) + 1)
			} else {
				lValue.SetValue(value.(int) + 1)
			}
			rtn = value
		case PlayScriptParserDEC:
			if _type == Integer {
				lValue.SetValue(value.(int) - 1)
			} else {
				lValue.SetValue(value.(int) - 1)
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
				rtn = value.(int) + 1
			} else {
				rtn = value.(int) + 1
			}
			lValue.SetValue(rtn)
		case PlayScriptParserDEC:
			if _type == Integer {
				rtn = value.(int) - 1
			} else {
				rtn = value.(int) - 1
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

func (this *ASTEvaluator) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	// TODO 暂时不考虑 this 和 super

	var rtn interface{}
	//这是调用时的名称，不一定是真正的函数名，还可能是函数尅性的变量名
	functionName := ctx.IDENTIFIER().GetText()

	if functionName == "println" {
		this.println(ctx)
		return rtn
	}

	//在上下文中查找出函数，并根据需要创建FunctionObject
	functionObject := this.getFunctionObject(ctx)
	//function := functionObject.function

	// todo 如果是对象的构造方法，则按照对象方法调用去执行，并返回所创建出的对象。

	// 计算参数值
	paramValues := this.calcParamValues(ctx)

	fmt.Println("\n>>FunctionCall: " + ctx.GetText())

	// 计算
	rtn = this.functionCall(functionObject, paramValues)
	return rtn
}

func (this *ASTEvaluator) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return this.VisitFunctionBody(ctx.FunctionBody().(*FunctionBodyContext))
}

func (this *ASTEvaluator) VisitFunctionBody(ctx *FunctionBodyContext) interface{} {
	var rtn interface{}
	if ctx.Block() != nil {
		rtn = this.VisitBlock(ctx.Block().(*BlockContext))
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
	} else if ctx.FloatLiteral() != nil { // 浮点数
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
		rtn = strWithQuotationMark[1 : len(strWithQuotationMark)-1]
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

func (this *ASTEvaluator) VisitProg(ctx *ProgContext) interface{} {
	var rtn interface{}
	scope := this.at.node2Scope[ctx]
	//if blockScope, ok := scope.(*BlockScope); ok {
	this.pushStack(NewFrame(scope))
	rtn = this.VisitBlockStatements(ctx.BlockStatements().(*BlockStatementsContext))
	this.popStack()
	//}
	return rtn
}

// -----各种运算-----
func (this *ASTEvaluator) add(obj1, obj2 interface{}, targetType Type) interface{} {
	var rtn interface{}
	// 暂时不支持真正的 long 和 short 以及 double
	if targetType == String {
		str1, ok1 := obj1.(string)
		str2, ok2 := obj2.(string)
		if ok1 && ok2 {
			rtn = str1 + str2
		}
	} else if targetType == Integer {
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
		if ok1 && ok2 {
			rtn = int1 + int2
		}
	} else if targetType == Short {
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
		if ok1 && ok2 {
			rtn = int1 + int2
		}
	} else if targetType == Long {
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
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
		double1, ok1 := obj1.(float32)
		double2, ok2 := obj2.(float32)
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
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
		if ok1 && ok2 {
			rtn = int1 - int2
		}
	} else if targetType == Short {
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
		if ok1 && ok2 {
			rtn = int1 - int2
		}
	} else if targetType == Long {
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
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
		double1, ok1 := obj1.(float32)
		double2, ok2 := obj2.(float32)
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
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
		if ok1 && ok2 {
			rtn = int1 * int2
		}
	} else if targetType == Short {
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
		if ok1 && ok2 {
			rtn = int1 * int2
		}
	} else if targetType == Long {
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
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
		double1, ok1 := obj1.(float32)
		double2, ok2 := obj2.(float32)
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
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
		if ok1 && ok2 {
			rtn = int1 / int2
		}
	} else if targetType == Short {
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
		if ok1 && ok2 {
			rtn = int1 / int2
		}
	} else if targetType == Long {
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
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
		double1, ok1 := obj1.(float32)
		double2, ok2 := obj2.(float32)
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
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
		if ok1 && ok2 {
			rtn = int1 == int2
		}
	} else if targetType == Short {
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
		if ok1 && ok2 {
			rtn = int1 == int2
		}
	} else if targetType == Long {
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
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
		double1, ok1 := obj1.(float32)
		double2, ok2 := obj2.(float32)
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
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
		if ok1 && ok2 {
			rtn = int1 >= int2
		}
	} else if targetType == Short {
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
		if ok1 && ok2 {
			rtn = int1 >= int2
		}
	} else if targetType == Long {
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
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
		double1, ok1 := obj1.(float32)
		double2, ok2 := obj2.(float32)
		if ok1 && ok2 {
			rtn = double1 >= double2
		}
	}

	return rtn
}

func (this *ASTEvaluator) GT(obj1, obj2 interface{}, targetType Type) bool {
	rtn := false
	if targetType == Integer {
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
		if ok1 && ok2 {
			rtn = int1 > int2
		}
	} else if targetType == Short {
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
		if ok1 && ok2 {
			rtn = int1 > int2
		}
	} else if targetType == Long {
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
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
		double1, ok1 := obj1.(float32)
		double2, ok2 := obj2.(float32)
		if ok1 && ok2 {
			rtn = double1 > double2
		}
	}

	return rtn
}

func (this *ASTEvaluator) LE(obj1, obj2 interface{}, targetType Type) bool {
	rtn := false
	if targetType == Integer {
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
		if ok1 && ok2 {
			rtn = int1 <= int2
		}
	} else if targetType == Short {
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
		if ok1 && ok2 {
			rtn = int1 <= int2
		}
	} else if targetType == Long {
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
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
		double1, ok1 := obj1.(float32)
		double2, ok2 := obj2.(float32)
		if ok1 && ok2 {
			rtn = double1 <= double2
		}
	}

	return rtn
}

func (this *ASTEvaluator) LT(obj1, obj2 interface{}, targetType Type) bool {
	rtn := false
	if targetType == Integer {
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
		if ok1 && ok2 {
			rtn = int1 < int2
		}
	} else if targetType == Short {
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
		if ok1 && ok2 {
			rtn = int1 < int2
		}
	} else if targetType == Long {
		int1, ok1 := obj1.(int)
		int2, ok2 := obj2.(int)
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
		double1, ok1 := obj1.(float32)
		double2, ok2 := obj2.(float32)
		if ok1 && ok2 {
			rtn = double1 < double2
		}
	}

	return rtn
}

func NewASTEvaluator(at *AnnotatedTree, cachePrintln bool) *ASTEvaluator {
	return &ASTEvaluator{at: at, cachePrintln: cachePrintln}
}

package playscript

// 存放一个函数运行时的本地变量的值，包括参数的值
type FunctionObject struct {
	*BasePlayObject
	function *Function
	receiver *Variable
}

func (functionObject *FunctionObject) SetFunction(function *Function) {
	functionObject.function = function
}

func NewFunctionObject(function *Function) *FunctionObject {
	return &FunctionObject{function: function}
}

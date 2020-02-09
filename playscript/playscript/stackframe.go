package playscript

type StackFrame struct {
	// 该 frame 对应的 scope
	scope Scope
	//
	parentFrame *StackFrame
	// 实际存放变量的地方
	object PlayObject
}

// 本栈帧里有没有包含某个变量
func (this *StackFrame) Contains(variable *Variable) bool {
	if this.object != nil && this.object.GetFields() != nil {
		_, ok := this.object.GetFields()[variable]
		return ok
	}
	return false
}

func (this *StackFrame) ToString() string {
	rtn := this.scope.ToString()
	if this.parentFrame != nil {
		rtn += " -> " + this.parentFrame.ToString()
	}
	return rtn
}

func NewFrameForBlockScope(scope *BlockScope) *StackFrame {
	stackFrame := &StackFrame{}
	stackFrame.scope = scope
	stackFrame.object = NewBasePlayObject()
	return stackFrame
}

func NewFrameForFunction(object *FunctionObject) *StackFrame {
	stackFrame := &StackFrame{}
	stackFrame.scope = object.function
	stackFrame.object = object
	return stackFrame
}

package playscript

// 对栈中的值的引用
type LValue interface {
	GetValue() interface{}
	SetValue(interface{})
	GetVariable() *Variable
	GetValueContainer() PlayObject
}

type MyLValue struct {
	variable       *Variable
	valueContainer PlayObject
}

func (this *MyLValue) GetValue() interface{} {
	// TODO 对于this或super关键字，直接返回这个对象，应该是ClassObject
	return this.valueContainer.GetValue(this.variable)
}

func (this *MyLValue) SetValue(value interface{}) {
	this.valueContainer.SetValue(this.variable, value)
}

func (this *MyLValue) GetVariable() *Variable {
	return this.variable
}

func (this *MyLValue) GetValueContainer() PlayObject {
	return this.valueContainer
}

func NewMyLValue(valueContainer PlayObject, variable *Variable) *MyLValue {
	return &MyLValue{valueContainer: valueContainer, variable: variable}
}

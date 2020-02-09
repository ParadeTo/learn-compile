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

func (MyLValue) GetValue() interface{} {
	panic("implement me")
}

func (MyLValue) SetValue(interface{}) {
	panic("implement me")
}

func (MyLValue) GetVariable() *Variable {
	panic("implement me")
}

func (MyLValue) GetValueContainer() PlayObject {
	panic("implement me")
}

func NewMyLValue(valueContainer PlayObject, variable *Variable) *MyLValue {
	return &MyLValue{valueContainer: valueContainer, variable: variable}
}

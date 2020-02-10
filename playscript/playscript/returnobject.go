package playscript

type ReturnObject struct {
	returnValue interface{}
}

func (breakObject *ReturnObject) ToString() string {
	return "ReturnObject"
}

func NewReturnObject(value interface{}) *ReturnObject {
	return &ReturnObject{returnValue: value}
}

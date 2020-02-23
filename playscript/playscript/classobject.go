package playscript

type ClassObject struct {
	*BasePlayObject
	_type *Class
}

func NewClassObject() *ClassObject {
	return &ClassObject{}
}

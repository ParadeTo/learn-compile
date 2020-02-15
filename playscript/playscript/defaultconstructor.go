package playscript

type DefaultConstructor struct {
	*Function
}

func (constructor *DefaultConstructor) Class() *Class {
	return constructor.GetEnclosingScope().(*Class)
}

func NewDefaultConstructor(name string, theClase *Class) *DefaultConstructor {
	return &DefaultConstructor{Function: NewFunction(name, theClase, nil)}
}

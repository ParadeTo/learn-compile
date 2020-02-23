package playscript

type DefaultConstructor Function

func (constructor *DefaultConstructor) Class() *Class {
	return constructor.GetEnclosingScope().(*Class)
}

func NewDefaultConstructor(name string, class *Class) *DefaultConstructor {
	return (*DefaultConstructor)(NewFunction(name, class, nil))
}

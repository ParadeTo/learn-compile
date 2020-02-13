package playscript

type Class struct {
	*BaseScope
	// 父类
	parentClass *Class
	// this
	thisRef *This
	// super
}

func (c *Class) IsType(_type Type) bool {
	panic("implement me")
}

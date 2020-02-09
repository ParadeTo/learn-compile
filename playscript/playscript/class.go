package playscript

type Class struct {
	*BaseScope
}

func (c *Class) IsType(_type Type) bool {
	panic("implement me")
}

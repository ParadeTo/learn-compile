package playscript

import "github.com/antlr/antlr4/runtime/Go/antlr"

var rootClass = NewClass("Object", nil)

type Class struct {
	*BaseScope
	// 父类
	parentClass *Class
	// this
	thisRef *This
	// super
	superRef *Super

	defaultConstructor *DefaultConstructor
}

// 当自身是目标类型的子类型的时候，返回true
func (c *Class) IsType(_type Type) bool {
	if c == _type {
		return true
	}

	if class, ok := _type.(*Class); ok {
		return class.isAncestor(c)
	}

	return false
}

func (c *Class) getDefaultConstructor() *DefaultConstructor {
	if c.defaultConstructor == nil {
		c.defaultConstructor = NewDefaultConstructor(c.name, c)
	}
	return c.defaultConstructor
}

// 本类型是不是另一个类型的祖先类型
func (c *Class) isAncestor(theClass *Class) bool {
	if theClass.GetParentClass() != nil {
		if theClass.GetParentClass() == c {
			return true
		} else {
			return c.isAncestor(theClass.GetParentClass())
		}
	}
	return false
}

func (c *Class) GetParentClass() *Class {
	return c.parentClass
}

func (c *Class) SetParentClass(theClass *Class) {
	c.parentClass = theClass
	c.superRef = NewSuper(theClass, c.ctx)
	c.superRef._type = theClass
}

func (c *Class) GetThis() *This {
	return c.thisRef
}

func (c *Class) GetSuper() *Super {
	return c.superRef
}

func (c *Class) ToString() string {
	return "Class " + c.name
}

// 是否包含某个Variable，包括自身及父类
func (c *Class) GetVariable(name string) *Variable {
	rtn := c.BaseScope.GetVariable(name)
	if rtn == nil && c.parentClass != nil {
		rtn = c.parentClass.GetVariable(name)
	}
	return rtn
}

// 是否包含某个Class
func (c *Class) GetClass(name string) *Class {
	rtn := c.BaseScope.GetClass(name)
	if rtn == nil && c.parentClass != nil {
		rtn = c.parentClass.GetClass(name)
	}
	return rtn
}

func (c *Class) FindConstructor(paramTypes []Type) *Function {
	return c.BaseScope.GetFunctionByNameAndParams(c.name, paramTypes)
}

// 在自身及父类中找到某个方法
// 不支持函数重载
func (c *Class) GetFunction(name string) *Function {
	rtn := c.BaseScope.GetFunction(name)
	if rtn == nil && c.parentClass != nil {
		rtn = c.parentClass.GetFunction(name)
	}
	return rtn
}

// 函数类型的成员变量
func (c *Class) GetFunctionVariable(name string) *Variable {
	rtn := c.BaseScope.GetFunctionVariable(name)
	if rtn == nil && c.parentClass != nil {
		rtn = c.parentClass.GetFunctionVariable(name)
	}
	return rtn
}

// 是否包含某个Symbol。这时候要把父类的成员考虑进来
func (c *Class) ContainsSymbol(symbol Symbol) bool {
	if symbol == c.thisRef || symbol == c.superRef {
		return true
	}

	rtn := false

	for _, _symbol := range c.symbols {
		if _symbol == symbol {
			rtn = true
			break
		}
	}

	if !rtn && c.parentClass != nil {
		rtn = c.parentClass.ContainsSymbol(symbol)
	}

	return rtn
}

func NewClass(name string, ctx antlr.Tree) *Class {
	class := &Class{BaseScope: NewBaseScope()}
	class.name = name
	class.ctx = ctx
	class.thisRef = NewThis(class, ctx)
	class.thisRef._type = class
	return class
}

package playscript

import "github.com/antlr/antlr4/runtime/Go/antlr"

type Super struct {
	*Variable
}

func (super *Super) Class() *Class {
	return super.GetEnclosingScope().(*Class)
}

func NewSuper(class *Class, ctx antlr.Tree) *Super {
	return &Super{Variable: NewVariable("super", class, ctx)}
}

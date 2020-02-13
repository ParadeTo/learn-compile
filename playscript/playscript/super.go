package playscript

import "github.com/antlr/antlr4/runtime/Go/antlr"

type This struct {
	*Variable
}

func (this *This) Class() *Class {
	return this.GetEnclosingScope().(*Class)
}

func NewThis(class *Class, ctx antlr.Tree) *This {
	return &This{Variable: NewVariable("this", class, ctx)}
}

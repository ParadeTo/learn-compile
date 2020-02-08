package playscript

import . "github.com/antlr/antlr4/runtime/Go/antlr"

type Symbol interface {
	GetName() string
	GetEnclosingScope() Scope
	SetEnclosingScope(scope Scope)
	SetCtx(ctx Tree)
	GetCtx() Tree
}

// 编译过程中产生的变量、函数、类、块，都被称作符号
type BaseSymbol struct {
	name           string
	visibility     int   // 可见性，如 public 还是 private
	enclosingScope Scope // 所属作用域
	ctx            Tree
}

func (s *BaseSymbol) GetName() string {
	return s.name
}

func (s *BaseSymbol) GetEnclosingScope() Scope {
	return s.enclosingScope
}

func (s *BaseSymbol) SetEnclosingScope(scope Scope) {
	s.enclosingScope = scope
}

func (s *BaseSymbol) SetCtx(ctx Tree) {
	s.ctx = ctx
}

func (s *BaseSymbol) GetCtx() Tree {
	return s.ctx
}

func NewBaseSymbol() *BaseSymbol {
	return &BaseSymbol{}
}

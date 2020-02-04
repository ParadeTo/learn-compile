package playscript

import . "github.com/antlr/antlr4/runtime/Go/antlr"

type ISymbol interface {
	GetName() string
	GetEnclosingScope() IScope
	SetEnclosingScope(scope IScope)
	SetCtx(ctx ParserRuleContext)
}

// 编译过程中产生的变量、函数、类、块，都被称作符号
type Symbol struct {
	name           string
	visibility     int    // 可见性，如 public 还是 private
	enclosingScope IScope // 所属作用域
	ctx            ParserRuleContext
}

func (s *Symbol) GetName() string {
	return s.name
}

func (s *Symbol) GetEnclosingScope() IScope {
	return s.enclosingScope
}

func (s *Symbol) SetEnclosingScope(scope IScope) {
	s.enclosingScope = scope
}

func (s *Symbol) SetCtx(ctx ParserRuleContext) {
	s.ctx = ctx
}

func NewSymbol() *Symbol {
	return &Symbol{}
}

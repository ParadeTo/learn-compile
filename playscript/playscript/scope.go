package playscript

type IScope interface {
	ISymbol
	AddSymbol(symbol ISymbol)
	GetVariable(name string) *Variable
	ToString() string
}

type Scope struct {
	*Symbol
	// 该Scope中的成员，包括变量、方法、类等
	symbols []ISymbol
}

// 向scope中添加符号，同时设置好该符号的enclosingScope
func (s *Scope) AddSymbol(symbol ISymbol) {
	s.symbols = append(s.symbols, symbol)
	// 设置该符号的作用域
	symbol.SetEnclosingScope(s)
}

func (s *Scope) GetVariable(name string) *Variable {
	for _, symbol := range s.symbols {
		if variable, ok := symbol.(*Variable); ok && variable.GetName() == name {
			return variable
		}
	}
	return nil
}

func (s *Scope) ToString() string {
	return "Scope: " + s.name
}

func NewScope() *Scope {
	return &Scope{Symbol: NewSymbol()}
}

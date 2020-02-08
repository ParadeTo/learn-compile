package playscript

type Scope interface {
	Symbol
	AddSymbol(symbol Symbol)
	GetSymbols() []Symbol
	GetVariable(name string) *Variable
	// TODO，支持函数重载，需要传入
	//GetFunction(name string, paramTypes []Type) *Function
	GetFunction(name string) *Function
	GetFunctionVariable(name string) *Variable
	ToString() string
}

type BaseScope struct {
	*BaseSymbol
	// 该Scope中的成员，包括变量、方法、类等
	symbols []Symbol
}

// 向scope中添加符号，同时设置好该符号的enclosingScope
func (s *BaseScope) AddSymbol(symbol Symbol) {
	s.symbols = append(s.symbols, symbol)
	// 设置该符号的作用域
	symbol.SetEnclosingScope(s)
}

func (s *BaseScope) GetSymbols() []Symbol {
	return s.symbols
}

func (s *BaseScope) GetVariable(name string) *Variable {
	for _, symbol := range s.symbols {
		if variable, ok := symbol.(*Variable); ok && variable.GetName() == name {
			return variable
		}
	}
	return nil
}

func (s *BaseScope) GetFunction(name string) *Function {
	for _, symbol := range s.symbols {
		if function, ok := symbol.(*Function); ok && function.GetName() == name {
			return function
		}
	}
	return nil
}

func (s *BaseScope) GetFunctionVariable(name string) *Variable {
	for _, symbol := range s.symbols {
		if variable, ok := symbol.(*Variable); ok {
			if _, ok := variable._type.(FunctionType); ok && variable.GetName() == name {
				return variable
			}
		}
	}
	return nil
}

func (s *BaseScope) ToString() string {
	return "BaseScope: " + s.name
}

func NewBaseScope() *BaseScope {
	return &BaseScope{BaseSymbol: NewBaseSymbol()}
}

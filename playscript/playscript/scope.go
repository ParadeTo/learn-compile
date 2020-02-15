package playscript

type Scope interface {
	Symbol
	ContainsSymbol(symbol Symbol) bool
	AddSymbol(symbol Symbol)
	GetSymbols() []Symbol
	GetVariable(name string) *Variable
	GetClass(name string) *Class
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

func (s *BaseScope) ContainsSymbol(symbol Symbol) bool {
	for _, _symbol := range s.symbols {
		if _symbol == symbol {
			return true
		}
	}
	return false
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

func (s *BaseScope) GetClass(name string) *Class {
	for _, symbol := range s.symbols {
		if class, ok := symbol.(*Class); ok && class.GetName() == name {
			return class
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

// 仅供构造函数重载用
func (s *BaseScope) GetFunctionByNameAndParams(name string, paramTypes []Type) *Function {
	var rtn *Function
	for _, symbol := range s.symbols {
		if function, ok := symbol.(*Function); ok && function.GetName() == name {
			if function.MatchParameterTypes(paramTypes) {
				rtn = function
				break
			}
		}
	}
	return rtn
}

// 获取一个函数类型的变量
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

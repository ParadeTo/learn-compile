package playscript

type PrimitiveType struct {
	name string
}

func (p *PrimitiveType) GetName() string {
	return p.name
}

func (p *PrimitiveType) GetEnclosingScope() Scope {
	return nil
}

func (p *PrimitiveType) IsType(_type Type) bool {
	return p == _type
}

func (p *PrimitiveType) ToString() string {
	return p.name
}

// 计算两个类型中比较“高”的一级，比如int和long相加，要取long
func GetUpperType(type1, type2 Type) *PrimitiveType {
	var _type *PrimitiveType
	if type1 == String || type2 == String {
		_type = String
	} else if type1 == Double || type2 == Double {
		_type = Double
	} else if type1 == Float || type2 == Float {
		_type = Float
	} else if type1 == Long || type2 == Long {
		_type = Long
	} else if type1 == Integer || type2 == Integer {
		_type = Integer
	} else if type1 == Short || type2 == Short {
		_type = Short
	} else {
		_type = Byte
	}
	return _type
}

func IsNumeric(_type Type) bool {
	return _type == Byte ||
		_type == Short ||
		_type == Integer ||
		_type == Long ||
		_type == Float ||
		_type == Double
}

func NewPrimitiveType(name string) *PrimitiveType {
	return &PrimitiveType{name}
}

var (
	Integer = NewPrimitiveType("Integer")
	Long    = NewPrimitiveType("Long")
	Float   = NewPrimitiveType("Float")
	Double  = NewPrimitiveType("Double")
	Boolean = NewPrimitiveType("Boolean")
	Byte    = NewPrimitiveType("Byte")
	Char    = NewPrimitiveType("Char")
	Short   = NewPrimitiveType("Short")
	String  = NewPrimitiveType("String")
	Null    = NewPrimitiveType("Null")
)

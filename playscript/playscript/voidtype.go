package playscript

var voidType *VoidType

type VoidType struct{}

func (t VoidType) GetName() string {
	return "void"
}

// 闭包函数的返回值不会是 void
func (t VoidType) GetEnclosingScope() Scope {
	return nil
}

func (t VoidType) IsType(_type Type) bool {
	return t == _type
}

func (t VoidType) ToString() string {
	return "void"
}

func GetVoidType() *VoidType {
	if voidType == nil {
		voidType = &VoidType{}
	}
	return voidType
}

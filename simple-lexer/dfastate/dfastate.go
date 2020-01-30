package dfastate

// 确定性有限状态机
type DfaState uint8

const (
	Initial DfaState = iota
	Id
	Id_int1 // int 的 i
	Id_int2 // int 的 n
	Id_int3 // int 的 t
	IntLiteral
	GT
	GE
	Assignment // = 赋值
)

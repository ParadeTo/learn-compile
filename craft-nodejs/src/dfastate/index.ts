// 确定有限状态机

export enum DfaState {
  Initial, // 初始状态
	Id, // 标识符
	Id_int1, // int 的 i
	Id_int2, // int 的 n
	Id_int3, // int 的 t
	IntLiteral, // 整形字面量
	GT,
	GE,
	Assignment, // = 赋值

	Plus,  // +
	Minus, // -
	Star,  // *
	Slash, // /

	SemiColon,  // ;
	LeftParen,  // (
	RightParen, // )
}
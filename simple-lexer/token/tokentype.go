package token

type TokenType string

const (
	Plus  TokenType = "Plus"
	Minus TokenType = "Minus" // -
	Star  TokenType = "Star"  // *
	Slash TokenType = "Slash" // /

	GE TokenType = "GE" // >=
	GT TokenType = "GT" // >
	EQ TokenType = "EQ" // ==
	LE TokenType = "LE" // <=
	LT TokenType = "LT" // <

	SemiColon  TokenType = "SemiColon"  // ;
	LeftParen  TokenType = "LeftParen"  // (
	RightParen TokenType = "RightParen" // )

	Assignment TokenType = "Assignment" // =

	If   TokenType = "If"
	Else TokenType = "Else"

	Int TokenType = "Int"

	Identifier TokenType = "Identifier" //标识符

	IntLiteral    TokenType = "IntLiteral"    //整型字面量
	StringLiteral TokenType = "StringLiteral" //字符串字面量
)

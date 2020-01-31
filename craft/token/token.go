package token

type Token interface {
	GetType() TokenType
	GetText() string
}
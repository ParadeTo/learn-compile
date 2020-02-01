package token

type Token interface {
	GetType() TokenType
	GetText() string
}

type SimpleToken struct {
	Type TokenType
	Text string
}

func (token SimpleToken) GetType() TokenType {
	return token.Type
}

func (token SimpleToken) GetText() string {
	return token.Text
}

package lexer

import (
	"fmt"
	"learn-compile/craft/dfastate"
	"learn-compile/craft/token"
)

type SimpleToken struct {
	_type token.TokenType
	text  string
}

func (token SimpleToken) GetType() token.TokenType {
	return token._type
}

func (token SimpleToken) GetText() string {
	return token.text
}

type SimpleTokenReader struct {
	tokens []token.Token
	pos    int
}

func (reader *SimpleTokenReader) Read() token.Token {
	if reader.pos < len(reader.tokens) {
		p := reader.pos
		reader.pos++
		return reader.tokens[p]
	}
	return nil
}

func (reader *SimpleTokenReader) Peek() token.Token {
	if reader.pos < len(reader.tokens) {
		return reader.tokens[reader.pos]
	}
	return nil
}

func (reader *SimpleTokenReader) UnRead() {
	if reader.pos > 0 {
		reader.pos--
	}
}

func (reader *SimpleTokenReader) GetPosition() int {
	return reader.pos
}

func (reader *SimpleTokenReader) SetPosition(pos int) {
	if pos >= 0 && pos < len(reader.tokens) {
		reader.pos = pos
	}
}

func NewSimpleTokenReader(tokens []token.Token) *SimpleTokenReader {
	return &SimpleTokenReader{tokens: tokens, pos: 0}
}

type SimpleLexer struct {
	tokenText string // 临时保存token的文本
	tokens    []token.Token
	token     SimpleToken
}

func (lexer *SimpleLexer) isAlpha(ch string) bool {
	return ch >= "a" && ch <= "z" || ch >= "A" && ch <= "Z"
}

func (lexer *SimpleLexer) isDigit(ch string) bool {
	return ch >= "0" && ch <= "9"
}

func (lexer *SimpleLexer) isBlank(ch string) bool {
	return ch == " " || ch == "\t" || ch == "\n"
}

func (lexer *SimpleLexer) initToken(ch string) dfastate.DfaState {
	// 结束当前的 token
	if len(lexer.tokenText) > 0 {
		lexer.token.text = lexer.tokenText
		lexer.tokens = append(lexer.tokens, lexer.token)
	}

	// 开启下一个 token
	lexer.tokenText = ""
	lexer.token = SimpleToken{text: ""}
	newState := dfastate.Initial
	if lexer.isAlpha(ch) {
		if ch == "i" {
			newState = dfastate.Id_int1
		} else {
			newState = dfastate.Id
		}
		lexer.token._type = token.Identifier
	} else if lexer.isDigit(ch) {
		newState = dfastate.IntLiteral
		lexer.token._type = token.IntLiteral
	} else if ch == ">" {
		newState = dfastate.GT
		lexer.token._type = token.GT
	} else if ch == "=" {
		newState = dfastate.Assignment
		lexer.token._type = token.Assignment
	} else if ch == "+" {
		newState = dfastate.Plus
		lexer.token._type = token.Plus
	} else if ch == "-" {
		newState = dfastate.Minus
		lexer.token._type = token.Minus
	} else if ch == "*" {
		newState = dfastate.Star
		lexer.token._type = token.Star
	} else if ch == "/" {
		newState = dfastate.Slash
		lexer.token._type = token.Slash
	} else if ch == ";" {
		newState = dfastate.SemiColon
		lexer.token._type = token.SemiColon
	} else if ch == "(" {
		newState = dfastate.LeftParen
		lexer.token._type = token.LeftParen
	} else if ch == ")" {
		newState = dfastate.RightParen
		lexer.token._type = token.RightParen
	}
	if newState != dfastate.Initial {
		lexer.tokenText += ch
	}
	return newState
}

func (lexer *SimpleLexer) Tokenize(code string) *SimpleTokenReader {
	lexer.tokenText = ""
	lexer.token = SimpleToken{text: ""}
	lexer.tokens = nil

	state := dfastate.Initial
	for _, ascii := range code {
		ch := string(ascii)
		switch state {
		case dfastate.Initial:
			state = lexer.initToken(ch)
		case dfastate.Id:
			if lexer.isAlpha(ch) || lexer.isDigit(ch) {
				lexer.tokenText += ch
			} else {
				state = lexer.initToken(ch)
			}
		case dfastate.GT:
			if ch == "=" {
				lexer.token._type = token.GE
				state = dfastate.GE
				lexer.tokenText += ch
			} else {
				state = lexer.initToken(ch)
			}
		case dfastate.GE:
			state = lexer.initToken(ch)
		case dfastate.Id_int1:
			if ch == "n" {
				state = dfastate.Id_int2
				lexer.tokenText += ch
			} else if lexer.isDigit(ch) || lexer.isAlpha(ch) {
				state = dfastate.Id
				lexer.tokenText += ch
			} else {
				state = lexer.initToken(ch)
			}
		case dfastate.Id_int2:
			if ch == "t" {
				state = dfastate.Id_int3
				lexer.tokenText += ch
			} else if lexer.isDigit(ch) || lexer.isAlpha(ch) {
				state = dfastate.Id
				lexer.tokenText += ch
			} else {
				state = lexer.initToken(ch)
			}
		case dfastate.Id_int3:
			if lexer.isBlank(ch) {
				lexer.token._type = token.Int
				state = lexer.initToken(ch)
			} else {
				state = dfastate.Id
				lexer.tokenText += ch
			}
		case dfastate.Plus, dfastate.Minus, dfastate.Star, dfastate.Slash, dfastate.Assignment,
			dfastate.LeftParen, dfastate.RightParen, dfastate.SemiColon:
			state = lexer.initToken(ch)
		case dfastate.IntLiteral:
			if lexer.isDigit(ch) {
				lexer.tokenText += ch
			} else {
				state = lexer.initToken(ch)
			}
		default:
		}
	}

	if len(lexer.tokenText) > 0 {
		lexer.initToken("")
	}

	return NewSimpleTokenReader(lexer.tokens)
}

func (lexer *SimpleLexer) Dump(reader *SimpleTokenReader) {
	fmt.Println("type\ttext")
	for {
		token := reader.Read()
		if token == nil {
			break
		}
		fmt.Println(string(token.GetType()) + "\t" + token.GetText())
	}
}

func NewSimpleLexer() *SimpleLexer {
	return &SimpleLexer{token: SimpleToken{}}
}

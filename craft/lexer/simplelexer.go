package lexer

import (
	"fmt"
	"learn-compile/craft/dfastate"
	"learn-compile/craft/token"
)

type SimpleLexer struct {
	tokenText string // 临时保存token的文本
	tokens    []token.Token
	token     token.SimpleToken
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
	if len(lexer.token.Text) > 0 {
		// lexer.token.Text = lexer.token.Text
		lexer.tokens = append(lexer.tokens, lexer.token)
	}

	// 开启下一个 token
	lexer.token = token.SimpleToken{Text: ""}
	newState := dfastate.Initial
	if lexer.isAlpha(ch) {
		if ch == "i" {
			newState = dfastate.Id_int1
		} else {
			newState = dfastate.Id
		}
		lexer.token.Type = token.Identifier
	} else if lexer.isDigit(ch) {
		newState = dfastate.IntLiteral
		lexer.token.Type = token.IntLiteral
	} else if ch == ">" {
		newState = dfastate.GT
		lexer.token.Type = token.GT
	} else if ch == "=" {
		newState = dfastate.Assignment
		lexer.token.Type = token.Assignment
	} else if ch == "+" {
		newState = dfastate.Plus
		lexer.token.Type = token.Plus
	} else if ch == "-" {
		newState = dfastate.Minus
		lexer.token.Type = token.Minus
	} else if ch == "*" {
		newState = dfastate.Star
		lexer.token.Type = token.Star
	} else if ch == "/" {
		newState = dfastate.Slash
		lexer.token.Type = token.Slash
	} else if ch == ";" {
		newState = dfastate.SemiColon
		lexer.token.Type = token.SemiColon
	} else if ch == "(" {
		newState = dfastate.LeftParen
		lexer.token.Type = token.LeftParen
	} else if ch == ")" {
		newState = dfastate.RightParen
		lexer.token.Type = token.RightParen
	}
	if newState != dfastate.Initial {
		lexer.token.Text += ch
	}
	return newState
}

func (lexer *SimpleLexer) Tokenize(code string) *token.SimpleTokenReader {
	lexer.token = token.SimpleToken{Text: ""}
	lexer.tokens = nil

	state := dfastate.Initial
	for _, ascii := range code {
		ch := string(ascii)
		switch state {
		case dfastate.Initial:
			state = lexer.initToken(ch)
		case dfastate.Id:
			if lexer.isAlpha(ch) || lexer.isDigit(ch) {
				lexer.token.Text += ch
			} else {
				state = lexer.initToken(ch)
			}
		case dfastate.GT:
			if ch == "=" {
				lexer.token.Type = token.GE
				state = dfastate.GE
				lexer.token.Text += ch
			} else {
				state = lexer.initToken(ch)
			}
		case dfastate.GE:
			state = lexer.initToken(ch)
		case dfastate.Id_int1:
			if ch == "n" {
				state = dfastate.Id_int2
				lexer.token.Text += ch
			} else if lexer.isDigit(ch) || lexer.isAlpha(ch) {
				state = dfastate.Id
				lexer.token.Text += ch
			} else {
				state = lexer.initToken(ch)
			}
		case dfastate.Id_int2:
			if ch == "t" {
				state = dfastate.Id_int3
				lexer.token.Text += ch
			} else if lexer.isDigit(ch) || lexer.isAlpha(ch) {
				state = dfastate.Id
				lexer.token.Text += ch
			} else {
				state = lexer.initToken(ch)
			}
		case dfastate.Id_int3:
			if lexer.isBlank(ch) {
				lexer.token.Type = token.Int
				state = lexer.initToken(ch)
			} else {
				state = dfastate.Id
				lexer.token.Text += ch
			}
		case dfastate.Plus, dfastate.Minus, dfastate.Star, dfastate.Slash, dfastate.Assignment,
			dfastate.LeftParen, dfastate.RightParen, dfastate.SemiColon:
			state = lexer.initToken(ch)
		case dfastate.IntLiteral:
			if lexer.isDigit(ch) {
				lexer.token.Text += ch
			} else {
				state = lexer.initToken(ch)
			}
		default:
		}
	}

	if len(lexer.token.Text) > 0 {
		lexer.initToken("")
	}

	return token.NewSimpleTokenReader(lexer.tokens)
}

func (lexer *SimpleLexer) Dump(reader *token.SimpleTokenReader) {
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
	return &SimpleLexer{token: token.SimpleToken{}}
}

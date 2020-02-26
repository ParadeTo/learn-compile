import { DfaState } from "../dfastate"
import { Token, TokenType, TokenReader } from "../token"

class SimpleToken implements Token {
  text: string = ''
  type!: TokenType

  constructor(text: string) {
    this.text = text
  }

  getType(): TokenType {
    return this.type
  }

  getText(): string {
    return this.text
  }
}

class SimpleTokenReader implements TokenReader {
  tokens: Token[] = []
  pos: number = 0

  constructor(tokens: Token[]) {
    this.tokens = tokens
  }

  read(): Token {
    let token: Token
    if (this.pos < this.tokens.length) {
      token = this.tokens[this.pos++]
    }
    return token!
  }

  peek(): Token {
    let token: Token
    if (this.pos < this.tokens.length) {
      token = this.tokens[this.pos]
    }
    return token!
  }

  unRead() {
    if (this.pos > 0) {
      this.pos--
    }
  }

  getPosition(): number {
    return this.pos
  }

  setPosition(pos: number) {
    this.pos = pos
  }
}

// 词法分析器
export default class Lexer {
  // 保存所有的token
  tokens: Token[] = []
  // 当前正在处理的token
  token!: SimpleToken

  isAlpha(ch: string): boolean {
    return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
  }

  isDigit(ch: string): boolean {
    return ch >= '0' && ch <= '9'
  }

  isBlank(ch: string): boolean {
    return ch == ' ' || ch == '\t' || ch == '\n'
  }

  initToken(ch: string): DfaState {
    // 结束当前的 token
    if (this.token.text.length > 0) {
      this.tokens.push(this.token)
    }

    // 开启下一个 token
    this.token = new SimpleToken('')
    let newState = DfaState.Initial
    if (this.isAlpha(ch)) {
      if (ch == 'i') {
        newState = DfaState.Id_int1
      } else {
        newState = DfaState.Id
      }
      this.token.type = TokenType.Identifier
    } else if (this.isDigit(ch)) {
      newState = DfaState.IntLiteral
      this.token.type = TokenType.IntLiteral
    } else if (ch == '>') {
      newState = DfaState.GT
      this.token.type = TokenType.GT
    } else if (ch == '=') {
      newState = DfaState.Assignment
      this.token.type = TokenType.Assignment
    } else if (ch == '+') {
      newState = DfaState.Plus
      this.token.type = TokenType.Plus
    } else if (ch == '-') {
      newState = DfaState.Minus
      this.token.type = TokenType.Minus
    } else if (ch == '*') {
      newState = DfaState.Star
      this.token.type = TokenType.Star
    } else if (ch == '/') {
      newState = DfaState.Slash
      this.token.type = TokenType.Slash
    } else if (ch == ';') {
      newState = DfaState.SemiColon
      this.token.type = TokenType.SemiColon
    } else if (ch == '(') {
      newState = DfaState.LeftParen
      this.token.type = TokenType.LeftParen
    } else if (ch == ')') {
      newState = DfaState.RightParen
      this.token.type = TokenType.RightParen
    }

    if (newState != DfaState.Initial) {
      this.token.text += ch
    }

    return newState
  }

  tokenize(code: string): SimpleTokenReader {
    // 每次解析前先初始化数据
    this.token = new SimpleToken('')
    this.tokens = []

    let state = DfaState.Initial
    for (let i = 0; i < code.length; i++) {
      const ch = code[i]
      switch (state) {
        case DfaState.Initial:
          state = this.initToken(ch)
          break
        case DfaState.Id:
          // 标识符只能含有字母和数字
          if (this.isAlpha(ch) || this.isDigit(ch)) {
            this.token.text += ch
          } else {
            // 否则重新开始
            state = this.initToken(ch)
          }
          break
        case DfaState.GT:
          if (ch === '=') {
            this.token.type = TokenType.GE
            state = DfaState.GE
            this.token.text += ch
          } else {
            state = this.initToken(ch)
          }
          break
        case DfaState.Id_int1:
          if (ch === 'n') {
            state = DfaState.Id_int2
            this.token.text += ch
          } else if (this.isDigit(ch) || this.isAlpha(ch)) {
            state = DfaState.Id
            this.token.text += ch
          } else {
            state = this.initToken(ch)
          }
          break
        case DfaState.Id_int2:
          if (ch === 't') {
            state = DfaState.Id_int3
            this.token.text += ch
          } else if (this.isDigit(ch) || this.isAlpha(ch)) {
            state = DfaState.Id
            this.token.text += ch
          } else {
            state = this.initToken(ch)
          }
          break
        case DfaState.Id_int3:
          if (this.isBlank(ch)) {
            this.token.type = TokenType.Int
            state = this.initToken(ch)
          } else if (this.isDigit(ch) || this.isAlpha(ch)) {
            state = DfaState.Id
            this.token.text += ch
          } else {
            state = this.initToken(ch)
          }
          break
        case DfaState.IntLiteral:
          if (this.isDigit(ch)) {
            this.token.text += ch
          } else {
            state = this.initToken(ch)
          }
          break
        case DfaState.GE:
        case DfaState.Plus:
        case DfaState.Minus:
        case DfaState.Star:
        case DfaState.Slash:
        case DfaState.Assignment:
        case DfaState.LeftParen:
        case DfaState.RightParen:
        case DfaState.SemiColon:
          state = this.initToken(ch)
        default:
          break
      }
    }

    // 完成最后一个token
    if (this.token.text.length > 0) {
      this.initToken('')
    }

    return new SimpleTokenReader(this.tokens)
  }

  dump(reader: SimpleTokenReader) {
    console.log("type\ttext")
    let token
    while (token = reader.read()) {
      console.log(token.getType() + "\t" + token.getText())
    }
  }
}

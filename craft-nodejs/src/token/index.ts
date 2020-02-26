export enum TokenType {
  Plus = 'Plus', // +
  Minus = 'Minus', // -
  Star = 'Star', // *
  Slash = 'Slash', // /

  GE = 'GE', // >=
  GT = 'GT', // >
  EQ = 'EQ', // ==
  LE = 'LE', // <=
  LT = 'LT', // <

  SemiColon = 'SemiColon', // ;
  LeftParen = 'LeftParen', // (
  RightParen = 'RightParen', // )

  Assignment = 'Assignment', // =

  Int = 'Int', // int
  Identifier = 'Identifier', // 标识符, 比如变量名称

  IntLiteral = 'IntLiteral' // 整型字面量
}

export interface Token {
  getType(): TokenType
  getText(): string
}

export interface TokenReader {
  /**
   * 返回Token流中下一个Token，并从流中取出。 如果流已经为空，返回null;
   */
  read(): Token

  /**
   * 返回Token流中下一个Token，但不从流中取出。 如果流已经为空，返回null;
   */
  peek(): Token

  /**
   * Token流回退一步。恢复原来的Token。
   */
  unRead()

  /**
   * 获取Token流当前的读取位置。
   * @return
   */
  getPosition(): number

  /**
   * 设置Token流当前的读取位置
   * @param position
   */
  setPosition(pos: number)
}

package token

type TokenReader interface {
	/**
	 * 返回Token流中下一个Token，并从流中取出。 如果流已经为空，返回null;
	 */
	Read() Token

	/**
	 * 返回Token流中下一个Token，但不从流中取出。 如果流已经为空，返回null;
	 */
	Peek() Token

	/**
	 * Token流回退一步。恢复原来的Token。
	 */
	UnRead()

	/**
	 * 获取Token流当前的读取位置。
	 * @return
	 */
	GetPosition() int

	/**
	 * 设置Token流当前的读取位置
	 * @param position
	 */
	SetPosition(pos int)
}

type SimpleTokenReader struct {
	tokens []Token
	pos    int
}

func (reader *SimpleTokenReader) Read() Token {
	if reader.pos < len(reader.tokens) {
		p := reader.pos
		reader.pos++
		return reader.tokens[p]
	}
	return nil
}

func (reader *SimpleTokenReader) Peek() Token {
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

func NewSimpleTokenReader(tokens []Token) *SimpleTokenReader {
	return &SimpleTokenReader{tokens: tokens, pos: 0}
}

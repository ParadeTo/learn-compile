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

package playscript

// 给 block 编号的数字
var index = 1

type BlockScope struct {
	*Scope
}

func (b *BlockScope) ToString() string {
	return "Block: " + b.name
}

func NewBlockScope() *BlockScope {
	blockScope := &BlockScope{Scope: NewScope()}
	index++
	return blockScope
}

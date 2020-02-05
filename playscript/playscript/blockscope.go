package playscript

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strconv"
)

// 给 block 编号的数字
var index = 1

type BlockScope struct {
	*BaseScope
}

func (b *BlockScope) ToString() string {
	return "Block: " + b.name
}

func NewBlockScope() *BlockScope {
	blockScope := &BlockScope{BaseScope: NewBaseScope()}
	blockScope.name = "block" + strconv.Itoa(index)
	index++
	return blockScope
}

func NewBlockScopeWithParams(enclosingScope Scope, ctx antlr.ParserRuleContext) *BlockScope {
	blockScope := &BlockScope{BaseScope: NewBaseScope()}
	blockScope.name = "block" + strconv.Itoa(index)
	blockScope.SetEnclosingScope(enclosingScope)
	blockScope.SetCtx(ctx)
	index++
	return blockScope
}

package playscript

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type NameSpace struct {
	*BlockScope
	parent        *NameSpace
	subNameSpaces []*NameSpace
	name          string
}

func (nameSpace *NameSpace) GetName() string {
	return nameSpace.name
}

func (nameSpace *NameSpace) GetSubNameSpaces() []*NameSpace {
	return nameSpace.subNameSpaces
}

func (nameSpace *NameSpace) AddSubNameSpace(child *NameSpace) {
	child.parent = nameSpace
	nameSpace.subNameSpaces = append(nameSpace.subNameSpaces, child)
}

func (nameSpace *NameSpace) RemoveSubNameSpace(child *NameSpace) {
	child.parent = nil
	for i, _item := range nameSpace.subNameSpaces {
		if _item == child {
			index = i
			break
		}
	}
	if index > -1 {
		nameSpace.subNameSpaces = append(
			nameSpace.subNameSpaces[:index], nameSpace.subNameSpaces[index+1:]...)
	}
}

func NewNameSpace(name string, enclosingScope IScope, ctx antlr.ParserRuleContext) *NameSpace {
	nameSpace := &NameSpace{BlockScope: NewBlockScopeWithParams(enclosingScope, ctx)}
	nameSpace.name = name
	nameSpace.SetEnclosingScope(enclosingScope)
	nameSpace.ctx = ctx
	return nameSpace
}

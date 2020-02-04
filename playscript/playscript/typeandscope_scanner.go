package playscript

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	. "learn-compile/playscript/parser"
)

type TypeAndScopeScanner struct {
	*BasePlayScriptListener
	at         *AnnotatedTree
	scopeStack []IScope
}

func (scanner *TypeAndScopeScanner) pushScope(scope IScope, ctx antlr.ParserRuleContext) IScope {
	scanner.at.node2Scope[ctx] = scope
	scope.SetCtx(ctx)
	scanner.scopeStack = append(scanner.scopeStack, scope)
	return scope
}

// 当前的Scope
func (scanner *TypeAndScopeScanner) currentScope() IScope {
	_len := len(scanner.scopeStack)
	if _len > 0 {
		return scanner.scopeStack[_len-1]
	} else {
		return nil
	}
}

func (scanner *TypeAndScopeScanner) EnterProg(ctx *ProgContext) {
	nameSpace := NewNameSpace("", scanner.currentScope(), ctx)
	scanner.at.nameSpace = nameSpace
	scanner.pushScope(nameSpace, ctx)
}

// 块级作用域
func (scanner *TypeAndScopeScanner) EnterBlock(ctx *BlockContext) {
	// 函数体的 {} 形成的块级作用域不需要额外再建立了
	// 直接用函数的作用域就好了
	if _, ok := ctx.GetParent().(*FunctionBodyContext); !ok {
		fmt.Println("aaa")
	}
}

func NewTypeAndScopeScanner(at *AnnotatedTree) *TypeAndScopeScanner {
	return &TypeAndScopeScanner{at: at}
}

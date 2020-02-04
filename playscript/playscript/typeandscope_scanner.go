package playscript

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	. "learn-compile/playscript/parser"
)

/**
 * 第一遍扫描：识别出所有类型（包括类和函数)，以及Scope。（但函数的参数信息要等到下一个阶段才会添加进去。）
 * 扫描是为了填充 subNameSpaces 的信息
 */
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

func (scanner *TypeAndScopeScanner) popScope() {
	l := len(scanner.scopeStack)
	scanner.scopeStack = scanner.scopeStack[:l-1]
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
	// 程序最顶层的作用域叫 NameSpace
	nameSpace := NewNameSpace("", scanner.currentScope(), ctx)
	scanner.at.nameSpace = nameSpace
	scanner.pushScope(nameSpace, ctx)
}

func (scanner *TypeAndScopeScanner) ExitProg(ctx *ProgContext) {
	scanner.popScope()
}

// 块级作用域
func (scanner *TypeAndScopeScanner) EnterBlock(ctx *BlockContext) {
	// 函数体的 {} 形成的块级作用域不需要额外再建立了
	// 直接用函数的作用域就好了
	if _, ok := ctx.GetParent().(*FunctionBodyContext); !ok {
		// currentScope 是 blockScope 的父 scope
		currentScope := scanner.currentScope()
		blockScope := NewBlockScopeWithParams(currentScope, ctx)
		currentScope.AddSymbol(blockScope)
		scanner.pushScope(blockScope, ctx)
	}
}

func (scanner *TypeAndScopeScanner) ExitBlock(ctx *BlockContext) {
	if _, ok := ctx.GetParent().(*FunctionBodyContext); !ok {
		scanner.popScope()
	}
}

// for 循环语句
func (scanner *TypeAndScopeScanner) EnterStatement(ctx *StatementContext) {
	// 为for的()建立额外的Scope
	if ctx.FOR() != nil {
		currentScope := scanner.currentScope()
		scope := NewBlockScopeWithParams(currentScope, ctx)
		currentScope.AddSymbol(scope)
		scanner.pushScope(scope, ctx)
	}
}

func (scanner *TypeAndScopeScanner) ExitStatement(ctx *StatementContext) {
	if ctx.FOR() != nil {
		scanner.popScope()
	}
}

// 函数
func (scanner *TypeAndScopeScanner) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {
	//idName := ctx.IDENTIFIER().GetText()
	//注意：目前funtion的信息并不完整，参数要等到 TypeResolver 中去确定。
	//Function function = new Function(idName, currentScope(), ctx);
}

func NewTypeAndScopeScanner(at *AnnotatedTree) *TypeAndScopeScanner {
	return &TypeAndScopeScanner{at: at}
}

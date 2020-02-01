package parser

import (
	"errors"
	"learn-compile/craft/ast"
	"learn-compile/craft/lexer"
	"learn-compile/craft/token"
)

/**
 * 一个简单的语法解析器。
 * 能够解析简单的表达式、变量声明和初始化语句、赋值语句。
 * 它支持的语法规则为：
 *
 * programm -> intDeclare | expressionStatement | assignmentStatement
 * intDeclare -> 'int' Id ( = additive) ';'
 * expressionStatement -> additive ';'
 * assignmentStatement -> Id = expressionStatement
 * additive -> multiplicative ( (+ | -) multiplicative)*
 * multiplicative -> primary ( (* | /) primary)*
 * primary -> IntLiteral | Id | (additive)
 */
type SimpleParser struct {
}

/**
 * 加/减法表达式，左结合
 * Additive: Multiplicative (Plus/Minus Multiplicative)*
 */
func (parser SimpleParser) Additive(reader token.TokenReader) (error, *ast.SimpleASTNode) {
	_, child1 := parser.Multiplicative(reader)
	node := child1
	if child1 != nil {
		for {
			_token := reader.Peek()
			if _token != nil && (_token.GetType() == token.Plus || _token.GetType() == token.Minus) {
				_token = reader.Read()
				_, child2 := parser.Multiplicative(reader)
				if child2 != nil {
					node = ast.NewSimpleASTNode(ast.Additive, _token.GetText())
					node.AddChild(child1)
					node.AddChild(child2)
					child1 = node
				} else {
					return errors.New("invalid additive expression, expecting the right part"), nil
				}
			} else {
				break
			}
		}
	}
	return nil, node
}

/**
 * 乘/除法表达式，左结合版
 * Multiplicative: Primary (Star/Slash Primary)*
 */
func (parser SimpleParser) Multiplicative(reader token.TokenReader) (error, *ast.SimpleASTNode) {
	_, child1 := parser.Primary(reader)
	node := child1
	if child1 != nil {
		for {
			_token := reader.Peek()
			if _token != nil && (_token.GetType() == token.Star || _token.GetType() == token.Slash) {
				_token = reader.Read()
				_, child2 := parser.Primary(reader)
				if child2 != nil {
					node = ast.NewSimpleASTNode(ast.Multiplicative, _token.GetText())
					node.AddChild(child1)
					node.AddChild(child2)
					child1 = node
				} else {
					return errors.New("invalid multiplicative expression, expecting the right part"), nil
				}
			} else {
				break
			}
		}
	}
	return nil, node
}

/**
 * 基础表达式，这里做了简化，没有构造一个 Primary 类型的节点，而是直接返回了它的子节点，因为只有一个子节点
 */
func (parser SimpleParser) Primary(reader token.TokenReader) (error, *ast.SimpleASTNode) {
	var node *ast.SimpleASTNode
	_token := reader.Peek()
	if _token != nil {
		if _token.GetType() == token.IntLiteral { // 整形字面量
			_token = reader.Read()
			node = ast.NewSimpleASTNode(ast.IntLiteral, _token.GetText())
		} else if _token.GetType() == token.Identifier { // 标识符
			_token = reader.Read()
			node = ast.NewSimpleASTNode(ast.Identifier, _token.GetText())
		} else if _token.GetType() == token.LeftParen {
			reader.Read() // 消耗掉 (
			_, node = parser.Additive(reader)
			if node != nil {
				_token = reader.Peek()
				if _token != nil && _token.GetType() == token.RightParen {
					reader.Read()
				} else {
					return errors.New("expecting right parenthesis"), nil
				}
			} else {
				return errors.New("expecting an additive expression inside parenthesis"), nil
			}
		}
	}
	return nil, node
}

/**
 * 整型变量声明语句，如：
 * int a;
 * int b = 2*3;
 */
func (parser SimpleParser) IntDeclare(reader token.TokenReader) (error, *ast.SimpleASTNode) {
	var node *ast.SimpleASTNode
	_token := reader.Peek()
	if _token != nil && _token.GetType() == token.Int {
		reader.Read()                                    // 消耗掉 int
		if reader.Peek().GetType() == token.Identifier { //是否为标识符
			_token = reader.Read()
			//创建当前节点，并把变量名记到AST节点的文本值中，这里新建一个变量子节点也是可以的
			node = ast.NewSimpleASTNode(ast.IntDeclaration, _token.GetText())
			_token = reader.Peek()                                     // 预读
			if _token != nil && _token.GetType() == token.Assignment { // 没有 else，因为声明语句可以不赋值
				reader.Read()
				err, child := parser.Additive(reader)
				if err != nil {
					return errors.New("invalid variable initialization, expecting an expression"), nil
				}
				node.AddChild(child)
			}
		} else {
			return errors.New("variable name expected"), nil
		}

		if node != nil {
			_token = reader.Peek()
			if _token != nil && _token.GetType() == token.SemiColon {
				reader.Read()
			} else {
				return errors.New("invalid statement, expecting semicolon"), nil
			}
		}
	}
	return nil, node
}

/**
 * 表达式语句
 */
func (parser SimpleParser) ExpressionStatement(reader token.TokenReader) (error, *ast.SimpleASTNode) {
	pos := reader.GetPosition()
	error, node := parser.Additive(reader)
	if node != nil {
		_token := reader.Peek()
		if _token != nil && _token.GetType() == token.SemiColon {
			reader.Read()
		} else {
			node = nil
			reader.SetPosition(pos) // 回溯
		}
	}
	return error, node //直接返回子节点，简化了AST。
}

// 赋值语句, 如 age = 10*2
func (parser SimpleParser) AssignmentStatement(reader token.TokenReader) (error, *ast.SimpleASTNode) {
	var node *ast.SimpleASTNode
	_token := reader.Peek() // 预读，看是否为标识符
	if _token != nil && _token.GetType() == token.Identifier {
		_token = reader.Read()
		node = ast.NewSimpleASTNode(ast.AssignmentStmt, _token.GetText())
		_token = reader.Peek() // 预读，看是否为等号
		if _token != nil && _token.GetType() == token.Assignment {
			reader.Read()
			_, child := parser.Additive(reader)
			if child == nil { // 等号右面没有一个合法的表达式
				return errors.New("invalid assignment statement, expecting an expression"), nil
			} else {
				node.AddChild(child)
				_token = reader.Peek() // 看是不是分号
				if _token != nil && _token.GetType() == token.SemiColon {
					reader.Read() // 消耗掉这个分号
				} else {
					return errors.New("invalid statement, expecting semicolon"), nil
				}
			}
		} else {
			reader.UnRead() // 回溯
			node = nil
		}
	}
	return nil, node
}

func (parser SimpleParser) Prog(reader token.TokenReader) (error, *ast.SimpleASTNode) {
	node := ast.NewSimpleASTNode(ast.Programm, "pwc")

	for {
		if reader.Peek() == nil {
			break
		}

		_, child := parser.IntDeclare(reader)
		if child == nil {
			_, child = parser.ExpressionStatement(reader)
		}
		if child == nil {
			_, child = parser.AssignmentStatement(reader)
		}

		if child != nil {
			node.AddChild(child)
		} else {
			return errors.New("unknown statement"), nil
		}
	}

	return nil, node
}

func (parser SimpleParser) Parse(script string) (error, *ast.SimpleASTNode) {
	lexer := lexer.NewSimpleLexer()
	reader := lexer.Tokenize(script)
	return parser.Prog(reader)
}

func NewSimpleParser() *SimpleParser {
	return &SimpleParser{}
}

package calculator

import (
	"github.com/kataras/iris/core/errors"
	"learn-compile/craft/ast"
	"learn-compile/craft/token"
)

type SimpleASTNode struct {
	parent   *SimpleASTNode
	children []ast.ASTNode
	nodeType ast.ASTNodeType
	text     string
}

func (astNode *SimpleASTNode) GetParent() ast.ASTNode {
	return astNode.parent
}

func (astNode *SimpleASTNode) GetChildren() []ast.ASTNode {
	return astNode.children
}

func (astNode *SimpleASTNode) GetType() ast.ASTNodeType {
	return astNode.nodeType
}

func (astNode *SimpleASTNode) GetText() string {
	return astNode.text
}

func (astNode *SimpleASTNode) AddChild(child *SimpleASTNode) {
	astNode.children = append(astNode.children, child)
	child.parent = astNode
}

func NewSimpleASTNode(nodeType ast.ASTNodeType, text string) *SimpleASTNode {
	return &SimpleASTNode{nodeType: nodeType, text: text}
}

type SimpleCalculator struct {
}

/**
 * 加/减法表达式
 * Additive: Multiplicative | Multiplicative Plus/Minus Additive
 */
func (calc SimpleCalculator) Additive(reader token.TokenReader) (error, *SimpleASTNode) {
	_, child1 := calc.Multiplicative(reader)
	node := child1
	_token := reader.Peek()
	if child1 != nil && _token != nil {
		if _token.GetType() == token.Plus || _token.GetType() == token.Minus {
			_token = reader.Read() // 读取运算符
			_, child2 := calc.Additive(reader)
			if child2 != nil {
				node = NewSimpleASTNode(ast.Additive, _token.GetText())
				node.AddChild(child1)
				node.AddChild(child2)
			} else {
				return errors.New("invalid additive expression, expecting the right part"), nil
			}
		}
	}
	return nil, node
}

/**
 * 乘/除法表达式
 * Multiplicative: Primary | Multiplicative Star/Slash Primary // 这样会陷入死循环
 * Multiplicative: Primary | Primary Star/Slash Multiplicative
 */
func (calc SimpleCalculator) Multiplicative(reader token.TokenReader) (error, *SimpleASTNode) {
	_, child1 := calc.Primary(reader)
	node := child1
	_token := reader.Peek()
	if child1 != nil && _token != nil {
		if _token.GetType() == token.Star || _token.GetType() == token.Slash {
			_token = reader.Read() // 读取运算符
			_, child2 := calc.Multiplicative(reader)
			if child2 != nil {
				node = NewSimpleASTNode(ast.Multiplicative, _token.GetText())
				node.AddChild(child1)
				node.AddChild(child2)
			} else {
				return errors.New("invalid multiplicative expression, expecting the right part"), nil
			}
		}
	}
	return nil, node
}

/**
 * 基础表达式，这里做了简化，没有构造一个 Primary 类型的节点，而是直接返回了它的子节点，因为只有一个子节点
 */
func (calc SimpleCalculator) Primary(reader token.TokenReader) (error, *SimpleASTNode) {
	var node *SimpleASTNode
	_token := reader.Peek()
	if _token != nil {
		if _token.GetType() == token.IntLiteral { // 整形字面量
			_token = reader.Read()
			node = NewSimpleASTNode(ast.IntLiteral, _token.GetText())
		} else if _token.GetType() == token.Identifier { // 标识符
			_token = reader.Read()
			node = NewSimpleASTNode(ast.Identifier, _token.GetText())
		} else if _token.GetType() == token.LeftParen {
			reader.Read() // 消耗掉 (
			_, node = calc.Additive(reader)
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
func (calc SimpleCalculator) IntDeclare(reader token.TokenReader) (error, *SimpleASTNode) {
	var node *SimpleASTNode
	_token := reader.Peek()
	if _token != nil && _token.GetType() == token.Int {
		reader.Read()                                    // 消耗掉 int
		if reader.Peek().GetType() == token.Identifier { //是否为标识符
			_token = reader.Read()
			//创建当前节点，并把变量名记到AST节点的文本值中，这里新建一个变量子节点也是可以的
			node = NewSimpleASTNode(ast.IntDeclaration, _token.GetText())
			_token = reader.Peek()                                     // 预读
			if _token != nil && _token.GetType() == token.Assignment { // 没有 else，因为声明语句可以不赋值
				reader.Read()
				err, child := calc.Additive(reader)
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

// 前序遍历
func (calc SimpleCalculator) PreTraverse(node *SimpleASTNode) []*SimpleASTNode {
	var nodes []*SimpleASTNode
	var childrenNodes []*SimpleASTNode
	nodes = append(nodes, node)

	for _, child := range node.children {
		childrenNodes = append(childrenNodes, calc.PreTraverse(child.(*SimpleASTNode))...)
	}

	for _, childNode := range childrenNodes {
		nodes = append(nodes, childNode)
	}
	return nodes
}

func NewSimpleCalculator() *SimpleCalculator {
	return &SimpleCalculator{}
}

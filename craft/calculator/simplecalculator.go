package calculator

import (
	"fmt"
	"github.com/kataras/iris/core/errors"
	"learn-compile/craft/ast"
	"learn-compile/craft/lexer"
	"learn-compile/craft/token"
	"strconv"
)

type SimpleCalculator struct {
}

/**
 * 加/减法表达式，为了解决左递归无限循环的问题，这里的表达式写成了右递归，造成的问题是计算是又结合的
 * Additive: Multiplicative | Multiplicative Plus/Minus Additive
 */
//func (calc SimpleCalculator) Additive(reader token.TokenReader) (error, *ast.SimpleASTNode) {
//	_, child1 := calc.Multiplicative(reader)
//	node := child1
//	_token := reader.Peek()
//	if child1 != nil && _token != nil {
//		if _token.GetType() == token.Plus || _token.GetType() == token.Minus {
//			_token = reader.Read() // 读取运算符
//			_, child2 := calc.Additive(reader)
//			if child2 != nil {
//				node = Newast.SimpleASTNode(ast.Additive, _token.GetText())
//				node.AddChild(child1)
//				node.AddChild(child2)
//			} else {
//				return errors.New("invalid additive expression, expecting the right part"), nil
//			}
//		}
//	}
//	return nil, node
//}

/**
 * 加/减法表达式，左结合
 * Additive: Multiplicative (Plus/Minus Multiplicative)*
 */
func (calc SimpleCalculator) Additive(reader token.TokenReader) (error, *ast.SimpleASTNode) {
	_, child1 := calc.Multiplicative(reader)
	node := child1
	if child1 != nil {
		for {
			_token := reader.Peek()
			if _token != nil && (_token.GetType() == token.Plus || _token.GetType() == token.Minus) {
				_token = reader.Read()
				_, child2 := calc.Multiplicative(reader)
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
 * 乘/除法表达式
 * Multiplicative: Primary | Multiplicative Star/Slash Primary // 这样会陷入死循环
 * Multiplicative: Primary | Primary Star/Slash Multiplicative
 */
//func (calc SimpleCalculator) Multiplicative(reader token.TokenReader) (error, *ast.SimpleASTNode) {
//	_, child1 := calc.Primary(reader)
//	node := child1
//	_token := reader.Peek()
//	if child1 != nil && _token != nil {
//		if _token.GetType() == token.Star || _token.GetType() == token.Slash {
//			_token = reader.Read() // 读取运算符
//			_, child2 := calc.Multiplicative(reader)
//			if child2 != nil {
//				node = Newast.SimpleASTNode(ast.Multiplicative, _token.GetText())
//				node.AddChild(child1)
//				node.AddChild(child2)
//			} else {
//				return errors.New("invalid multiplicative expression, expecting the right part"), nil
//			}
//		}
//	}
//	return nil, node
//}

/**
 * 乘/除法表达式，左结合版
 * Multiplicative: Primary (Star/Slash Primary)*
 */
func (calc SimpleCalculator) Multiplicative(reader token.TokenReader) (error, *ast.SimpleASTNode) {
	_, child1 := calc.Primary(reader)
	node := child1
	if child1 != nil {
		for {
			_token := reader.Peek()
			if _token != nil && (_token.GetType() == token.Star || _token.GetType() == token.Slash) {
				_token = reader.Read()
				_, child2 := calc.Primary(reader)
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
func (calc SimpleCalculator) Primary(reader token.TokenReader) (error, *ast.SimpleASTNode) {
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
func (calc SimpleCalculator) IntDeclare(reader token.TokenReader) (error, *ast.SimpleASTNode) {
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

func (calc SimpleCalculator) Prog(reader token.TokenReader) (error, *ast.SimpleASTNode) {
	node := ast.NewSimpleASTNode(ast.Programm, "Calculator")
	error, child := calc.Additive(reader)
	if child != nil {
		node.AddChild(child)
	}
	return error, node
}

func (calc SimpleCalculator) parse(script string) (error, *ast.SimpleASTNode) {
	lexer := lexer.NewSimpleLexer()
	reader := lexer.Tokenize(script)
	return calc.Prog(reader)
}

func (calc SimpleCalculator) Evaluate(script string) (error, int) {
	error, tree := calc.parse(script)
	if error != nil {
		return error, 0
	}
	return error, calc.evaluate(tree, "")
}

/**
 * 求值，并打印求值过程
 */
func (calc SimpleCalculator) evaluate(node ast.ASTNode, indent string) int {
	result := 0
	fmt.Println(indent + "Calculating: " + string(node.GetType()))
	switch node.GetType() {
	case ast.Programm:
		for _, child := range node.GetChildren() {
			result = calc.evaluate(child, indent+"--")
		}
	case ast.Additive:
		children := node.GetChildren()
		child1 := children[0]
		value1 := calc.evaluate(child1, indent+"--")
		child2 := children[1]
		value2 := calc.evaluate(child2, indent+"--")
		if node.GetText() == "+" {
			result = value1 + value2
		} else {
			result = value1 - value2
		}
	case ast.Multiplicative:
		children := node.GetChildren()
		child1 := children[0]
		value1 := calc.evaluate(child1, indent+"--")
		child2 := children[1]
		value2 := calc.evaluate(child2, indent+"--")
		if node.GetText() == "*" {
			result = value1 * value2
		} else {
			result = value1 / value2
		}
	case ast.IntLiteral:
		result, _ = strconv.Atoi(node.GetText())
	}
	fmt.Printf(indent+"Result: %d\n", result)
	return result
}

func NewSimpleCalculator() *SimpleCalculator {
	return &SimpleCalculator{}
}

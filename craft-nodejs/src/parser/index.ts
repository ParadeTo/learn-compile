import {ASTNode, ASTNodeType} from '../ast'
import {TokenReader, TokenType, Token} from '../token'
import Lexer from '../lexer'

/**
 * 一个简单的语法解析器。
 * 能够解析简单的表达式(四则运算)、变量声明和初始化语句、赋值语句。
 * 它支持的语法规则为：
 *
 * programm -> intDeclare | expressionStatement | assignmentStatement
 * intDeclare -> 'int' Id ( = additive)? ';'
 * expressionStatement -> additive ';'
 * assignmentStatement -> Id = expressionStatement
 * additive -> multiplicative ( (+ | -) multiplicative )*
 * multiplicative -> primary ( (* | /) primary )*
 * primary -> IntLiteral | Id | (additive)
 */
export default class Parser {
  /**
   * 基础表达式，这里做了简化，没有构造一个 Primary 类型的节点，而是直接返回了它的子节点，因为只有一个子节点
   */
  primary(reader: TokenReader): SimpleASTNode {
    let node: SimpleASTNode
    let token = reader.peek()
    if (token) {
      // 整型字面量
      if (token.getType() === TokenType.IntLiteral) {
        node = new SimpleASTNode(
          ASTNodeType.IntLiteral,
          reader.read().getText()
        )
      }
      // 标识符
      else if (token.getType() === TokenType.Identifier) {
        node = new SimpleASTNode(
          ASTNodeType.Identifier,
          reader.read().getText()
        )
      }
      // 括号括起来的加法表达式
      else if (token.getType() === TokenType.LeftParen) {
        reader.read() // 消耗掉 (
        node = this.additive(reader)
        if (node) {
          token = reader.peek()
          if (token && token.getType() === TokenType.RightParen) {
            reader.read()
          } else {
            throw new Error('expecting right parenthesis')
          }
        } else {
          throw new Error('expecting an additive expression inside parenthesis')
        }
      }
    }
    return node!
  }

  /**
   * 乘/除法表达式
   * multiplicative: primary ( (+ | -) primary )*
   */
  multiplicative(reader: TokenReader): SimpleASTNode {
    let child1 = this.primary(reader)
    let node = child1
    let token
    if (child1) {
      while ((token = reader.peek())) {
        if (
          token &&
          (token.getType() === TokenType.Star ||
            token.getType() === TokenType.Slash)
        ) {
          token = reader.read()
          const child2 = this.primary(reader)
          if (child2) {
            node = new SimpleASTNode(
              ASTNodeType.Multiplicative,
              token.getText()
            )
            node.addChild(child1)
            node.addChild(child2)
            // 新生成的结点作为左边的节点
            child1 = node
          } else {
            throw new Error(
              'invalid multiplicative expression, expecting the right part'
            )
          }
        } else {
          break
        }
      }
    }
    return node
  }

  /**
   * 加/减法表达式
   * Additive: multiplicative ( (+ | -) multiplicative )*
   */
  additive(reader: TokenReader): SimpleASTNode {
    let child1 = this.multiplicative(reader)
    let node = child1
    let token
    if (child1) {
      while ((token = reader.peek())) {
        if (
          token &&
          (token.getType() === TokenType.Plus ||
            token.getType() === TokenType.Minus)
        ) {
          token = reader.read()
          const child2 = this.multiplicative(reader)
          if (child2) {
            node = new SimpleASTNode(ASTNodeType.Additive, token.getText())
            node.addChild(child1)
            node.addChild(child2)
            // 新生成的结点作为左边的节点
            child1 = node
          } else {
            throw new Error(
              'invalid additive expression, expecting the right part'
            )
          }
        } else {
          break
        }
      }
    }
    return node
  }

  /**
   * 整型变量声明语句，如：
   * int a;
   * int b = 2*3;
   */
  intDeclare(reader: TokenReader): SimpleASTNode {
      let node: SimpleASTNode
      let token = reader.peek()
      if (token && token.getType() === TokenType.Int) {
        // 消耗掉 int，我们已经知道是整型变量声明语句了，int 这个 token 已经不需要了
        reader.read()
        // int 后面必须接标识符
        if (reader.peek().getType() === TokenType.Identifier) {
          token = reader.read()
          //创建当前节点，并把变量名记到AST节点的文本值中，这里新建一个变量子节点也是可以的
          node = new SimpleASTNode(ASTNodeType.IntDeclaration, token.getText())
          token = reader.peek()
          // 声明语句后还有赋值操作
          if (token && token.getType() === TokenType.Assignment) {
            reader.read() // 消耗掉 =
            const child = this.additive(reader)
            if (!child) {
              throw new Error(
                'invalid variable initialization, expecting an expression'
              )
            }
            node.addChild(child)
          }
          // 这里不需要 else 因为声明语句后可以没有赋值操作
        }
        // 否则报错
        else {
          throw new Error('variable name expected')
        }

        if (node) {
          token = reader.peek()
          if (token && token.getType() === TokenType.SemiColon) {
            reader.read()
          } else {
            throw new Error('invalid statement, expecting semicolon')
          }
        }
      }
      return node!
  }

  // 表达式
  expressionStatement(reader: TokenReader): SimpleASTNode {
    const pos = reader.getPosition()
    let node = this.additive(reader)
    let token
    if (node) {
      token = reader.peek()
      if (token && token.getType() === TokenType.SemiColon) {
        reader.read()
      } else { // 解析失败，回溯
        node = null!
        reader.setPosition(pos)
      }
    }
    // 直接返回子节点，简化 AST
    return node
  }

  // 赋值语句，如 age = 20;
  assignmentStatement(reader: TokenReader): SimpleASTNode {
    let node: SimpleASTNode
    let token = reader.peek()
    if (token && token.getType() === TokenType.Identifier) {
      token = reader.read()
      node = new SimpleASTNode(ASTNodeType.AssignmentStmt, token.getText())
      token = reader.peek()
      if (token && token.getType() === TokenType.Assignment) {
        reader.read()
        const child = this.additive(reader)
        if (!child) {// 等号右面没有一个合法的表达式
          throw new Error('invalid assignment statement, expecting an expression')
        } else {
          node.addChild(child)
          token = reader.peek() // 看是不是分号
          if (token && token.getType() == TokenType.SemiColon) {
            reader.read() // 消耗掉这个分号
          } else {
            throw new Error('invalid statement, expecting semicolon')
          }
        }
      } else {
        reader.unRead() // 回溯
        node = null!
      }
    }
    return node!
  }

  // 入口
  prog(reader: TokenReader): SimpleASTNode {
    // 根节点
    const node = new SimpleASTNode(ASTNodeType.Programm, 'simple script')
    let child
    while (reader.peek()) {
      // 依次尝试三种语句
      child = this.intDeclare(reader)
      if (!child) {
        child = this.expressionStatement(reader)
      }
      if (!child) {
        child = this.assignmentStatement(reader)
      }

      if (child) {
        node.addChild(child)
      } else {
        throw new Error('unknow statement')
      }
    }
    return node
  }

  parse(script: string): SimpleASTNode {
    const lexer = new Lexer()
    return this.prog(lexer.tokenize(script))
  }

  preTraverse(node: ASTNode): ASTNode[] {
    const nodes: ASTNode[] = []
    const childrenNodes: ASTNode[] = []
    nodes.push(node)
    node.getChildren().forEach(child => {
      childrenNodes.push(...this.preTraverse(child))
    })
    childrenNodes.forEach(node => {
      nodes.push(node)
    })
    return nodes
  }

  dumpAST(node: ASTNode, indent: string = '\t') {
    console.log(indent + node.getType() + " " + node.getText())
    node.getChildren().forEach(child => {
      this.dumpAST(child, indent+'\t')
    });
  }
}

class SimpleASTNode implements ASTNode {
  parent!: SimpleASTNode
  nodeType!: ASTNodeType
  children: ASTNode[] = []
  text: string = ''

  constructor(nodeType: ASTNodeType, text: string) {
    this.nodeType = nodeType
    this.text = text
  }

  getParent(): ASTNode {
    return this.parent
  }

  getChildren(): ASTNode[] {
    return this.children
  }

  getType(): ASTNodeType {
    return this.nodeType
  }

  getText(): string {
    return this.text
  }

  addChild(child: SimpleASTNode) {
    this.children.push(child)
    child.parent = this
  }
}

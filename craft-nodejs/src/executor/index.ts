import {ASTNode, ASTNodeType} from '../ast'

export default class Executor {
  verbose: boolean = false
  variables: {[k: string]: number} = {}

  constructor(verbose: boolean = false) {
    this.verbose = verbose
  }

  execute(node: ASTNode, indent: string = '\t'): number {
    let result: any = 0
    if (this.verbose) {
      console.log(indent + 'Executing: ' + node.getType())
    }

    // 因为解析 expressionStatement 的时候直接返回的是子节点 additive，所以这里没有 expressionStatement
    switch (node.getType()) {
      case ASTNodeType.PrintCall:
        result = void 0
        console.log(this.execute(node.getChildren()[0], indent + '\t'))
      case ASTNodeType.Programm:
        node.getChildren().forEach(child => {
          result = this.execute(child, indent + '\t')
        })
        break
      case ASTNodeType.Additive: {
        const [child1, child2] = node.getChildren()
        const value1 = this.execute(child1, indent + '\t')
        const value2 = this.execute(child2, indent + '\t')
        if (node.getText() === '+') {
          result = value1 + value2
        } else {
          result = value1 - value2
        }
        break
      }
      case ASTNodeType.Multiplicative:
        const [child1, child2] = node.getChildren()
        const value1 = this.execute(child1, indent + '\t')
        const value2 = this.execute(child2, indent + '\t')
        if (node.getText() === '*') {
          result = value1 * value2
        } else {
          result = value1 / value2
        }
        break
      case ASTNodeType.IntLiteral:
        result = +node.getText()
        break
      case ASTNodeType.Identifier: {
        const varName = node.getText()
        const value = this.variables[varName]
        if (value !== undefined) {
          result = this.variables[varName]
        } else {
          throw new Error('unknown variable: ' + varName)
        }
        break
      }
      case ASTNodeType.AssignmentStmt:
      case ASTNodeType.IntDeclaration:
        const varName = node.getText()
        if (
          node.getType() === ASTNodeType.AssignmentStmt &&
          !this.variables[varName]
        ) {
          throw new Error('unknown variable: ' + varName)
        }

        let varValue = 0 // 变量的默认值为 0
        const children = node.getChildren()
        // 有赋值的语句
        if (children.length > 0) {
          const child = children[0]
          result = this.execute(child, indent + '\t')
          varValue = result
        }
        this.variables[varName] = varValue
        break
      default:
        break
    }

    if (this.verbose) {
      console.log(indent + 'Result: ' + result)
    }

    return result
  }
}

package tools

import "learn-compile/craft/ast"

const INDENT = "\t"

// 前序遍历
func PreTraverse(node *ast.SimpleASTNode) []*ast.SimpleASTNode {
	var nodes []*ast.SimpleASTNode
	var childrenNodes []*ast.SimpleASTNode
	nodes = append(nodes, node)

	for _, child := range node.GetChildren() {
		childrenNodes = append(childrenNodes, PreTraverse(child.(*ast.SimpleASTNode))...)
	}

	for _, childNode := range childrenNodes {
		nodes = append(nodes, childNode)
	}
	return nodes
}

package ast

type ASTNode interface {
	GetParent() ASTNode
	GetChildren() []ASTNode
	GetType() ASTNodeType
	GetText() string
}

type SimpleASTNode struct {
	parent   *SimpleASTNode
	children []ASTNode
	NodeType ASTNodeType
	Text     string
}

func (astNode *SimpleASTNode) GetParent() ASTNode {
	return astNode.parent
}

func (astNode *SimpleASTNode) GetChildren() []ASTNode {
	return astNode.children
}

func (astNode *SimpleASTNode) GetType() ASTNodeType {
	return astNode.NodeType
}

func (astNode *SimpleASTNode) GetText() string {
	return astNode.Text
}

func (astNode *SimpleASTNode) AddChild(child *SimpleASTNode) {
	astNode.children = append(astNode.children, child)
	child.parent = astNode
}

func NewSimpleASTNode(nodeType ASTNodeType, text string) *SimpleASTNode {
	return &SimpleASTNode{NodeType: nodeType, Text: text}
}

package ast

type ASTNode interface {
	GetParent() ASTNode
	GetChildren() []ASTNode
	GetType() ASTNodeType
	GetText() string
}

// Code generated from PlayScript.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // PlayScript

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePlayScriptListener is a complete listener for a parse tree produced by PlayScriptParser.
type BasePlayScriptListener struct{}

var _ PlayScriptListener = &BasePlayScriptListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePlayScriptListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePlayScriptListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePlayScriptListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePlayScriptListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *BasePlayScriptListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *BasePlayScriptListener) ExitClassDeclaration(ctx *ClassDeclarationContext) {}

// EnterClassBody is called when production classBody is entered.
func (s *BasePlayScriptListener) EnterClassBody(ctx *ClassBodyContext) {}

// ExitClassBody is called when production classBody is exited.
func (s *BasePlayScriptListener) ExitClassBody(ctx *ClassBodyContext) {}

// EnterClassBodyDeclaration is called when production classBodyDeclaration is entered.
func (s *BasePlayScriptListener) EnterClassBodyDeclaration(ctx *ClassBodyDeclarationContext) {}

// ExitClassBodyDeclaration is called when production classBodyDeclaration is exited.
func (s *BasePlayScriptListener) ExitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) {}

// EnterMemberDeclaration is called when production memberDeclaration is entered.
func (s *BasePlayScriptListener) EnterMemberDeclaration(ctx *MemberDeclarationContext) {}

// ExitMemberDeclaration is called when production memberDeclaration is exited.
func (s *BasePlayScriptListener) ExitMemberDeclaration(ctx *MemberDeclarationContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BasePlayScriptListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BasePlayScriptListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterFunctionBody is called when production functionBody is entered.
func (s *BasePlayScriptListener) EnterFunctionBody(ctx *FunctionBodyContext) {}

// ExitFunctionBody is called when production functionBody is exited.
func (s *BasePlayScriptListener) ExitFunctionBody(ctx *FunctionBodyContext) {}

// EnterTypeTypeOrVoid is called when production typeTypeOrVoid is entered.
func (s *BasePlayScriptListener) EnterTypeTypeOrVoid(ctx *TypeTypeOrVoidContext) {}

// ExitTypeTypeOrVoid is called when production typeTypeOrVoid is exited.
func (s *BasePlayScriptListener) ExitTypeTypeOrVoid(ctx *TypeTypeOrVoidContext) {}

// EnterQualifiedNameList is called when production qualifiedNameList is entered.
func (s *BasePlayScriptListener) EnterQualifiedNameList(ctx *QualifiedNameListContext) {}

// ExitQualifiedNameList is called when production qualifiedNameList is exited.
func (s *BasePlayScriptListener) ExitQualifiedNameList(ctx *QualifiedNameListContext) {}

// EnterFormalParameters is called when production formalParameters is entered.
func (s *BasePlayScriptListener) EnterFormalParameters(ctx *FormalParametersContext) {}

// ExitFormalParameters is called when production formalParameters is exited.
func (s *BasePlayScriptListener) ExitFormalParameters(ctx *FormalParametersContext) {}

// EnterFormalParameterList is called when production formalParameterList is entered.
func (s *BasePlayScriptListener) EnterFormalParameterList(ctx *FormalParameterListContext) {}

// ExitFormalParameterList is called when production formalParameterList is exited.
func (s *BasePlayScriptListener) ExitFormalParameterList(ctx *FormalParameterListContext) {}

// EnterFormalParameter is called when production formalParameter is entered.
func (s *BasePlayScriptListener) EnterFormalParameter(ctx *FormalParameterContext) {}

// ExitFormalParameter is called when production formalParameter is exited.
func (s *BasePlayScriptListener) ExitFormalParameter(ctx *FormalParameterContext) {}

// EnterLastFormalParameter is called when production lastFormalParameter is entered.
func (s *BasePlayScriptListener) EnterLastFormalParameter(ctx *LastFormalParameterContext) {}

// ExitLastFormalParameter is called when production lastFormalParameter is exited.
func (s *BasePlayScriptListener) ExitLastFormalParameter(ctx *LastFormalParameterContext) {}

// EnterVariableModifier is called when production variableModifier is entered.
func (s *BasePlayScriptListener) EnterVariableModifier(ctx *VariableModifierContext) {}

// ExitVariableModifier is called when production variableModifier is exited.
func (s *BasePlayScriptListener) ExitVariableModifier(ctx *VariableModifierContext) {}

// EnterQualifiedName is called when production qualifiedName is entered.
func (s *BasePlayScriptListener) EnterQualifiedName(ctx *QualifiedNameContext) {}

// ExitQualifiedName is called when production qualifiedName is exited.
func (s *BasePlayScriptListener) ExitQualifiedName(ctx *QualifiedNameContext) {}

// EnterFieldDeclaration is called when production fieldDeclaration is entered.
func (s *BasePlayScriptListener) EnterFieldDeclaration(ctx *FieldDeclarationContext) {}

// ExitFieldDeclaration is called when production fieldDeclaration is exited.
func (s *BasePlayScriptListener) ExitFieldDeclaration(ctx *FieldDeclarationContext) {}

// EnterConstructorDeclaration is called when production constructorDeclaration is entered.
func (s *BasePlayScriptListener) EnterConstructorDeclaration(ctx *ConstructorDeclarationContext) {}

// ExitConstructorDeclaration is called when production constructorDeclaration is exited.
func (s *BasePlayScriptListener) ExitConstructorDeclaration(ctx *ConstructorDeclarationContext) {}

// EnterVariableDeclarators is called when production variableDeclarators is entered.
func (s *BasePlayScriptListener) EnterVariableDeclarators(ctx *VariableDeclaratorsContext) {}

// ExitVariableDeclarators is called when production variableDeclarators is exited.
func (s *BasePlayScriptListener) ExitVariableDeclarators(ctx *VariableDeclaratorsContext) {}

// EnterVariableDeclarator is called when production variableDeclarator is entered.
func (s *BasePlayScriptListener) EnterVariableDeclarator(ctx *VariableDeclaratorContext) {}

// ExitVariableDeclarator is called when production variableDeclarator is exited.
func (s *BasePlayScriptListener) ExitVariableDeclarator(ctx *VariableDeclaratorContext) {}

// EnterVariableDeclaratorId is called when production variableDeclaratorId is entered.
func (s *BasePlayScriptListener) EnterVariableDeclaratorId(ctx *VariableDeclaratorIdContext) {}

// ExitVariableDeclaratorId is called when production variableDeclaratorId is exited.
func (s *BasePlayScriptListener) ExitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) {}

// EnterVariableInitializer is called when production variableInitializer is entered.
func (s *BasePlayScriptListener) EnterVariableInitializer(ctx *VariableInitializerContext) {}

// ExitVariableInitializer is called when production variableInitializer is exited.
func (s *BasePlayScriptListener) ExitVariableInitializer(ctx *VariableInitializerContext) {}

// EnterArrayInitializer is called when production arrayInitializer is entered.
func (s *BasePlayScriptListener) EnterArrayInitializer(ctx *ArrayInitializerContext) {}

// ExitArrayInitializer is called when production arrayInitializer is exited.
func (s *BasePlayScriptListener) ExitArrayInitializer(ctx *ArrayInitializerContext) {}

// EnterClassOrInterfaceType is called when production classOrInterfaceType is entered.
func (s *BasePlayScriptListener) EnterClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) {}

// ExitClassOrInterfaceType is called when production classOrInterfaceType is exited.
func (s *BasePlayScriptListener) ExitClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) {}

// EnterTypeArgument is called when production typeArgument is entered.
func (s *BasePlayScriptListener) EnterTypeArgument(ctx *TypeArgumentContext) {}

// ExitTypeArgument is called when production typeArgument is exited.
func (s *BasePlayScriptListener) ExitTypeArgument(ctx *TypeArgumentContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BasePlayScriptListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BasePlayScriptListener) ExitLiteral(ctx *LiteralContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *BasePlayScriptListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *BasePlayScriptListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterFloatLiteral is called when production floatLiteral is entered.
func (s *BasePlayScriptListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production floatLiteral is exited.
func (s *BasePlayScriptListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterProg is called when production prog is entered.
func (s *BasePlayScriptListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BasePlayScriptListener) ExitProg(ctx *ProgContext) {}

// EnterBlock is called when production block is entered.
func (s *BasePlayScriptListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BasePlayScriptListener) ExitBlock(ctx *BlockContext) {}

// EnterBlockStatements is called when production blockStatements is entered.
func (s *BasePlayScriptListener) EnterBlockStatements(ctx *BlockStatementsContext) {}

// ExitBlockStatements is called when production blockStatements is exited.
func (s *BasePlayScriptListener) ExitBlockStatements(ctx *BlockStatementsContext) {}

// EnterBlockStatement is called when production blockStatement is entered.
func (s *BasePlayScriptListener) EnterBlockStatement(ctx *BlockStatementContext) {}

// ExitBlockStatement is called when production blockStatement is exited.
func (s *BasePlayScriptListener) ExitBlockStatement(ctx *BlockStatementContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasePlayScriptListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasePlayScriptListener) ExitStatement(ctx *StatementContext) {}

// EnterSwitchBlockStatementGroup is called when production switchBlockStatementGroup is entered.
func (s *BasePlayScriptListener) EnterSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) {
}

// ExitSwitchBlockStatementGroup is called when production switchBlockStatementGroup is exited.
func (s *BasePlayScriptListener) ExitSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) {
}

// EnterSwitchLabel is called when production switchLabel is entered.
func (s *BasePlayScriptListener) EnterSwitchLabel(ctx *SwitchLabelContext) {}

// ExitSwitchLabel is called when production switchLabel is exited.
func (s *BasePlayScriptListener) ExitSwitchLabel(ctx *SwitchLabelContext) {}

// EnterForControl is called when production forControl is entered.
func (s *BasePlayScriptListener) EnterForControl(ctx *ForControlContext) {}

// ExitForControl is called when production forControl is exited.
func (s *BasePlayScriptListener) ExitForControl(ctx *ForControlContext) {}

// EnterForInit is called when production forInit is entered.
func (s *BasePlayScriptListener) EnterForInit(ctx *ForInitContext) {}

// ExitForInit is called when production forInit is exited.
func (s *BasePlayScriptListener) ExitForInit(ctx *ForInitContext) {}

// EnterEnhancedForControl is called when production enhancedForControl is entered.
func (s *BasePlayScriptListener) EnterEnhancedForControl(ctx *EnhancedForControlContext) {}

// ExitEnhancedForControl is called when production enhancedForControl is exited.
func (s *BasePlayScriptListener) ExitEnhancedForControl(ctx *EnhancedForControlContext) {}

// EnterParExpression is called when production parExpression is entered.
func (s *BasePlayScriptListener) EnterParExpression(ctx *ParExpressionContext) {}

// ExitParExpression is called when production parExpression is exited.
func (s *BasePlayScriptListener) ExitParExpression(ctx *ParExpressionContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BasePlayScriptListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BasePlayScriptListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BasePlayScriptListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BasePlayScriptListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePlayScriptListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePlayScriptListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BasePlayScriptListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BasePlayScriptListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterTypeList is called when production typeList is entered.
func (s *BasePlayScriptListener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *BasePlayScriptListener) ExitTypeList(ctx *TypeListContext) {}

// EnterTypeType is called when production typeType is entered.
func (s *BasePlayScriptListener) EnterTypeType(ctx *TypeTypeContext) {}

// ExitTypeType is called when production typeType is exited.
func (s *BasePlayScriptListener) ExitTypeType(ctx *TypeTypeContext) {}

// EnterFunctionType is called when production functionType is entered.
func (s *BasePlayScriptListener) EnterFunctionType(ctx *FunctionTypeContext) {}

// ExitFunctionType is called when production functionType is exited.
func (s *BasePlayScriptListener) ExitFunctionType(ctx *FunctionTypeContext) {}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *BasePlayScriptListener) EnterPrimitiveType(ctx *PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *BasePlayScriptListener) ExitPrimitiveType(ctx *PrimitiveTypeContext) {}

// EnterCreator is called when production creator is entered.
func (s *BasePlayScriptListener) EnterCreator(ctx *CreatorContext) {}

// ExitCreator is called when production creator is exited.
func (s *BasePlayScriptListener) ExitCreator(ctx *CreatorContext) {}

// EnterSuperSuffix is called when production superSuffix is entered.
func (s *BasePlayScriptListener) EnterSuperSuffix(ctx *SuperSuffixContext) {}

// ExitSuperSuffix is called when production superSuffix is exited.
func (s *BasePlayScriptListener) ExitSuperSuffix(ctx *SuperSuffixContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BasePlayScriptListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BasePlayScriptListener) ExitArguments(ctx *ArgumentsContext) {}

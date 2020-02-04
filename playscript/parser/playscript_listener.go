// Code generated from PlayScript.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // PlayScript

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PlayScriptListener is a complete listener for a parse tree produced by PlayScriptParser.
type PlayScriptListener interface {
	antlr.ParseTreeListener

	// EnterClassDeclaration is called when entering the classDeclaration production.
	EnterClassDeclaration(c *ClassDeclarationContext)

	// EnterClassBody is called when entering the classBody production.
	EnterClassBody(c *ClassBodyContext)

	// EnterClassBodyDeclaration is called when entering the classBodyDeclaration production.
	EnterClassBodyDeclaration(c *ClassBodyDeclarationContext)

	// EnterMemberDeclaration is called when entering the memberDeclaration production.
	EnterMemberDeclaration(c *MemberDeclarationContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterFunctionBody is called when entering the functionBody production.
	EnterFunctionBody(c *FunctionBodyContext)

	// EnterTypeTypeOrVoid is called when entering the typeTypeOrVoid production.
	EnterTypeTypeOrVoid(c *TypeTypeOrVoidContext)

	// EnterQualifiedNameList is called when entering the qualifiedNameList production.
	EnterQualifiedNameList(c *QualifiedNameListContext)

	// EnterFormalParameters is called when entering the formalParameters production.
	EnterFormalParameters(c *FormalParametersContext)

	// EnterFormalParameterList is called when entering the formalParameterList production.
	EnterFormalParameterList(c *FormalParameterListContext)

	// EnterFormalParameter is called when entering the formalParameter production.
	EnterFormalParameter(c *FormalParameterContext)

	// EnterLastFormalParameter is called when entering the lastFormalParameter production.
	EnterLastFormalParameter(c *LastFormalParameterContext)

	// EnterVariableModifier is called when entering the variableModifier production.
	EnterVariableModifier(c *VariableModifierContext)

	// EnterQualifiedName is called when entering the qualifiedName production.
	EnterQualifiedName(c *QualifiedNameContext)

	// EnterFieldDeclaration is called when entering the fieldDeclaration production.
	EnterFieldDeclaration(c *FieldDeclarationContext)

	// EnterConstructorDeclaration is called when entering the constructorDeclaration production.
	EnterConstructorDeclaration(c *ConstructorDeclarationContext)

	// EnterVariableDeclarators is called when entering the variableDeclarators production.
	EnterVariableDeclarators(c *VariableDeclaratorsContext)

	// EnterVariableDeclarator is called when entering the variableDeclarator production.
	EnterVariableDeclarator(c *VariableDeclaratorContext)

	// EnterVariableDeclaratorId is called when entering the variableDeclaratorId production.
	EnterVariableDeclaratorId(c *VariableDeclaratorIdContext)

	// EnterVariableInitializer is called when entering the variableInitializer production.
	EnterVariableInitializer(c *VariableInitializerContext)

	// EnterArrayInitializer is called when entering the arrayInitializer production.
	EnterArrayInitializer(c *ArrayInitializerContext)

	// EnterClassOrInterfaceType is called when entering the classOrInterfaceType production.
	EnterClassOrInterfaceType(c *ClassOrInterfaceTypeContext)

	// EnterTypeArgument is called when entering the typeArgument production.
	EnterTypeArgument(c *TypeArgumentContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterIntegerLiteral is called when entering the integerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// EnterFloatLiteral is called when entering the floatLiteral production.
	EnterFloatLiteral(c *FloatLiteralContext)

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterBlockStatements is called when entering the blockStatements production.
	EnterBlockStatements(c *BlockStatementsContext)

	// EnterBlockStatement is called when entering the blockStatement production.
	EnterBlockStatement(c *BlockStatementContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterSwitchBlockStatementGroup is called when entering the switchBlockStatementGroup production.
	EnterSwitchBlockStatementGroup(c *SwitchBlockStatementGroupContext)

	// EnterSwitchLabel is called when entering the switchLabel production.
	EnterSwitchLabel(c *SwitchLabelContext)

	// EnterForControl is called when entering the forControl production.
	EnterForControl(c *ForControlContext)

	// EnterForInit is called when entering the forInit production.
	EnterForInit(c *ForInitContext)

	// EnterEnhancedForControl is called when entering the enhancedForControl production.
	EnterEnhancedForControl(c *EnhancedForControlContext)

	// EnterParExpression is called when entering the parExpression production.
	EnterParExpression(c *ParExpressionContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterTypeList is called when entering the typeList production.
	EnterTypeList(c *TypeListContext)

	// EnterTypeType is called when entering the typeType production.
	EnterTypeType(c *TypeTypeContext)

	// EnterFunctionType is called when entering the functionType production.
	EnterFunctionType(c *FunctionTypeContext)

	// EnterPrimitiveType is called when entering the primitiveType production.
	EnterPrimitiveType(c *PrimitiveTypeContext)

	// EnterCreator is called when entering the creator production.
	EnterCreator(c *CreatorContext)

	// EnterSuperSuffix is called when entering the superSuffix production.
	EnterSuperSuffix(c *SuperSuffixContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// ExitClassDeclaration is called when exiting the classDeclaration production.
	ExitClassDeclaration(c *ClassDeclarationContext)

	// ExitClassBody is called when exiting the classBody production.
	ExitClassBody(c *ClassBodyContext)

	// ExitClassBodyDeclaration is called when exiting the classBodyDeclaration production.
	ExitClassBodyDeclaration(c *ClassBodyDeclarationContext)

	// ExitMemberDeclaration is called when exiting the memberDeclaration production.
	ExitMemberDeclaration(c *MemberDeclarationContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitFunctionBody is called when exiting the functionBody production.
	ExitFunctionBody(c *FunctionBodyContext)

	// ExitTypeTypeOrVoid is called when exiting the typeTypeOrVoid production.
	ExitTypeTypeOrVoid(c *TypeTypeOrVoidContext)

	// ExitQualifiedNameList is called when exiting the qualifiedNameList production.
	ExitQualifiedNameList(c *QualifiedNameListContext)

	// ExitFormalParameters is called when exiting the formalParameters production.
	ExitFormalParameters(c *FormalParametersContext)

	// ExitFormalParameterList is called when exiting the formalParameterList production.
	ExitFormalParameterList(c *FormalParameterListContext)

	// ExitFormalParameter is called when exiting the formalParameter production.
	ExitFormalParameter(c *FormalParameterContext)

	// ExitLastFormalParameter is called when exiting the lastFormalParameter production.
	ExitLastFormalParameter(c *LastFormalParameterContext)

	// ExitVariableModifier is called when exiting the variableModifier production.
	ExitVariableModifier(c *VariableModifierContext)

	// ExitQualifiedName is called when exiting the qualifiedName production.
	ExitQualifiedName(c *QualifiedNameContext)

	// ExitFieldDeclaration is called when exiting the fieldDeclaration production.
	ExitFieldDeclaration(c *FieldDeclarationContext)

	// ExitConstructorDeclaration is called when exiting the constructorDeclaration production.
	ExitConstructorDeclaration(c *ConstructorDeclarationContext)

	// ExitVariableDeclarators is called when exiting the variableDeclarators production.
	ExitVariableDeclarators(c *VariableDeclaratorsContext)

	// ExitVariableDeclarator is called when exiting the variableDeclarator production.
	ExitVariableDeclarator(c *VariableDeclaratorContext)

	// ExitVariableDeclaratorId is called when exiting the variableDeclaratorId production.
	ExitVariableDeclaratorId(c *VariableDeclaratorIdContext)

	// ExitVariableInitializer is called when exiting the variableInitializer production.
	ExitVariableInitializer(c *VariableInitializerContext)

	// ExitArrayInitializer is called when exiting the arrayInitializer production.
	ExitArrayInitializer(c *ArrayInitializerContext)

	// ExitClassOrInterfaceType is called when exiting the classOrInterfaceType production.
	ExitClassOrInterfaceType(c *ClassOrInterfaceTypeContext)

	// ExitTypeArgument is called when exiting the typeArgument production.
	ExitTypeArgument(c *TypeArgumentContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitIntegerLiteral is called when exiting the integerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)

	// ExitFloatLiteral is called when exiting the floatLiteral production.
	ExitFloatLiteral(c *FloatLiteralContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitBlockStatements is called when exiting the blockStatements production.
	ExitBlockStatements(c *BlockStatementsContext)

	// ExitBlockStatement is called when exiting the blockStatement production.
	ExitBlockStatement(c *BlockStatementContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitSwitchBlockStatementGroup is called when exiting the switchBlockStatementGroup production.
	ExitSwitchBlockStatementGroup(c *SwitchBlockStatementGroupContext)

	// ExitSwitchLabel is called when exiting the switchLabel production.
	ExitSwitchLabel(c *SwitchLabelContext)

	// ExitForControl is called when exiting the forControl production.
	ExitForControl(c *ForControlContext)

	// ExitForInit is called when exiting the forInit production.
	ExitForInit(c *ForInitContext)

	// ExitEnhancedForControl is called when exiting the enhancedForControl production.
	ExitEnhancedForControl(c *EnhancedForControlContext)

	// ExitParExpression is called when exiting the parExpression production.
	ExitParExpression(c *ParExpressionContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitTypeList is called when exiting the typeList production.
	ExitTypeList(c *TypeListContext)

	// ExitTypeType is called when exiting the typeType production.
	ExitTypeType(c *TypeTypeContext)

	// ExitFunctionType is called when exiting the functionType production.
	ExitFunctionType(c *FunctionTypeContext)

	// ExitPrimitiveType is called when exiting the primitiveType production.
	ExitPrimitiveType(c *PrimitiveTypeContext)

	// ExitCreator is called when exiting the creator production.
	ExitCreator(c *CreatorContext)

	// ExitSuperSuffix is called when exiting the superSuffix production.
	ExitSuperSuffix(c *SuperSuffixContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)
}

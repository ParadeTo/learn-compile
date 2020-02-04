// Code generated from PlayScript.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // PlayScript

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by PlayScriptParser.
type PlayScriptVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PlayScriptParser#classDeclaration.
	VisitClassDeclaration(ctx *ClassDeclarationContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#classBody.
	VisitClassBody(ctx *ClassBodyContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#classBodyDeclaration.
	VisitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#memberDeclaration.
	VisitMemberDeclaration(ctx *MemberDeclarationContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#functionBody.
	VisitFunctionBody(ctx *FunctionBodyContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#typeTypeOrVoid.
	VisitTypeTypeOrVoid(ctx *TypeTypeOrVoidContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#qualifiedNameList.
	VisitQualifiedNameList(ctx *QualifiedNameListContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#formalParameters.
	VisitFormalParameters(ctx *FormalParametersContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#formalParameterList.
	VisitFormalParameterList(ctx *FormalParameterListContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#formalParameter.
	VisitFormalParameter(ctx *FormalParameterContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#lastFormalParameter.
	VisitLastFormalParameter(ctx *LastFormalParameterContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#variableModifier.
	VisitVariableModifier(ctx *VariableModifierContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#qualifiedName.
	VisitQualifiedName(ctx *QualifiedNameContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#fieldDeclaration.
	VisitFieldDeclaration(ctx *FieldDeclarationContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#constructorDeclaration.
	VisitConstructorDeclaration(ctx *ConstructorDeclarationContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#variableDeclarators.
	VisitVariableDeclarators(ctx *VariableDeclaratorsContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#variableDeclarator.
	VisitVariableDeclarator(ctx *VariableDeclaratorContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#variableDeclaratorId.
	VisitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#variableInitializer.
	VisitVariableInitializer(ctx *VariableInitializerContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#arrayInitializer.
	VisitArrayInitializer(ctx *ArrayInitializerContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#classOrInterfaceType.
	VisitClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#typeArgument.
	VisitTypeArgument(ctx *TypeArgumentContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#integerLiteral.
	VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#floatLiteral.
	VisitFloatLiteral(ctx *FloatLiteralContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#blockStatements.
	VisitBlockStatements(ctx *BlockStatementsContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#blockStatement.
	VisitBlockStatement(ctx *BlockStatementContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#switchBlockStatementGroup.
	VisitSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#switchLabel.
	VisitSwitchLabel(ctx *SwitchLabelContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#forControl.
	VisitForControl(ctx *ForControlContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#forInit.
	VisitForInit(ctx *ForInitContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#enhancedForControl.
	VisitEnhancedForControl(ctx *EnhancedForControlContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#parExpression.
	VisitParExpression(ctx *ParExpressionContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#typeList.
	VisitTypeList(ctx *TypeListContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#typeType.
	VisitTypeType(ctx *TypeTypeContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#functionType.
	VisitFunctionType(ctx *FunctionTypeContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#primitiveType.
	VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#creator.
	VisitCreator(ctx *CreatorContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#superSuffix.
	VisitSuperSuffix(ctx *SuperSuffixContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}
}

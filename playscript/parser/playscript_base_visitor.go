// Code generated from PlayScript.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // PlayScript

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasePlayScriptVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePlayScriptVisitor) VisitClassDeclaration(ctx *ClassDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitClassBody(ctx *ClassBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitMemberDeclaration(ctx *MemberDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitFunctionBody(ctx *FunctionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitTypeTypeOrVoid(ctx *TypeTypeOrVoidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitQualifiedNameList(ctx *QualifiedNameListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitFormalParameters(ctx *FormalParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitFormalParameterList(ctx *FormalParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitFormalParameter(ctx *FormalParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitLastFormalParameter(ctx *LastFormalParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitVariableModifier(ctx *VariableModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitQualifiedName(ctx *QualifiedNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitFieldDeclaration(ctx *FieldDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitConstructorDeclaration(ctx *ConstructorDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitVariableDeclarators(ctx *VariableDeclaratorsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitVariableDeclarator(ctx *VariableDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitVariableInitializer(ctx *VariableInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitArrayInitializer(ctx *ArrayInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitTypeArgument(ctx *TypeArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitFloatLiteral(ctx *FloatLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitBlockStatements(ctx *BlockStatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitBlockStatement(ctx *BlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitSwitchLabel(ctx *SwitchLabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitForControl(ctx *ForControlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitForInit(ctx *ForInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitEnhancedForControl(ctx *EnhancedForControlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitParExpression(ctx *ParExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitTypeList(ctx *TypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitTypeType(ctx *TypeTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitFunctionType(ctx *FunctionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitCreator(ctx *CreatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitSuperSuffix(ctx *SuperSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

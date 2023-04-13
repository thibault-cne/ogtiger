package ast

import "ogtiger/parser"

type OperationNegation struct {
	Expr Ast
	Ctx  parser.IOperationNegationContext
}

func (e *OperationNegation) Display() string {
	return " negation"
}

func (e *OperationNegation) Draw(prefix string) {
	// TODO: Draw the AST
}

func (l *AstCreatorListener) OperationNegationEnter(ctx parser.IOperationNegationContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *AstCreatorListener) OperationNegationExit(ctx parser.IOperationNegationContext) {
	operationNegation := &OperationNegation{
		Ctx: ctx,
	}

	operationNegation.Expr = l.PopAst()

	l.PushAst(operationNegation)
}

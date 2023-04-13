package ast

import "ogtiger/parser"

type OperationNegation struct {
	Expr Ast
}

func (e OperationNegation) Display() string {
	return " negation"
}

func (l *AstCreatorListener) OperationNegationEnter(ctx parser.IOperationNegationContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *AstCreatorListener) OperationNegationExit(ctx parser.IOperationNegationContext) {
	operationNegation := &OperationNegation{}

	operationNegation.Expr = l.PopAst()

	l.PushAst(operationNegation)
}

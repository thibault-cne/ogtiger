package ast

import "ogtiger/parser"

type OperationNegation struct {
	Expr Ast
}

func (e OperationNegation) Display() string {
	return " operationNegation"
}

func  (l *AstCreatorListener) OperationNegationEnter(ctx parser.IOperationNegationContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *AstCreatorListener) OperationNegationExit(ctx parser.IOperationNegationContext) {
	operationNegation := &OperationNegation{}

	// Get the first expr
	expr := l.AstStack[len(l.AstStack)-1]

	// Pop the last element of the stack
	l.AstStack = l.AstStack[:len(l.AstStack)-1]

	operationNegation.Expr = expr

	// Push the new element on the stack
	l.AstStack = append(l.AstStack, operationNegation)
}
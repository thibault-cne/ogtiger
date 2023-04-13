package ast

import "ogtiger/parser"

type OperationOu struct {
	Left Ast
	Right []Ast 
}

func (e OperationOu) Display() string {
	return " operationOu"
}

func  (l *AstCreatorListener) OperationOuEnter(ctx parser.IOperationOuContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func  (l *AstCreatorListener) OperationOuExit(ctx parser.IOperationOuContext) {
	exprOu := &OperationOu{}

	if ctx.GetChildCount() == 1 {
		return
	}

	// Get the first exprEt
	exprEt := l.AstStack[len(l.AstStack)-1]

	// Pop the last element of the stack
	l.AstStack = l.AstStack[:len(l.AstStack)-1]

	exprOu.Left = exprEt

	// Get the other exprEt
	for i := 0; i < (ctx.GetChildCount() - 1) / 2; i++ {
		exprOu.Right = append(exprOu.Right, l.AstStack[len(l.AstStack)-1])

		// Pop the last element of the stack
		l.AstStack = l.AstStack[:len(l.AstStack)-1]
	}

	// Push the new element on the stack
	l.AstStack = append(l.AstStack, exprOu)
}
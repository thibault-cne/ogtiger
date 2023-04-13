package ast

import "ogtiger/parser"

type OperationEt struct {
	Left  Ast
	Right []Ast
}

func (e OperationEt) Display() string {
	return " et"
}

func (l *AstCreatorListener) OperationEtEnter(ctx parser.IOperationEtContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *AstCreatorListener) OperationEtExit(ctx parser.IOperationEtContext) {
	operationEt := &OperationEt{}

	if ctx.GetChildCount() == 1 {
		return
	}

	// Get the first exprEt
	operationCompar := l.AstStack[len(l.AstStack)-1]

	// Pop the last element of the stack
	l.AstStack = l.AstStack[:len(l.AstStack)-1]

	operationEt.Left = operationCompar

	// Get the other exprEt
	for i := 0; i < (ctx.GetChildCount()-1)/2; i++ {
		operationEt.Right = append(operationEt.Right, l.AstStack[len(l.AstStack)-1])

		// Pop the last element of the stack
		l.AstStack = l.AstStack[:len(l.AstStack)-1]
	}

	// Push the new element on the stack
	l.AstStack = append(l.AstStack, operationEt)
}

package ast

import "ogtiger/parser"

type OperationOu struct {
	Left  Ast
	Right []Ast
}

func (e OperationOu) Display() string {
	return " ou"
}

func (l *AstCreatorListener) OperationOuEnter(ctx parser.IOperationOuContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *AstCreatorListener) OperationOuExit(ctx parser.IOperationOuContext) {
	operationOu := &OperationOu{}

	if ctx.GetChildCount() == 1 {
		return
	}

	operationOu.Left = l.PopAst()

	// Get the other exprEt
	for i := 0; i < (ctx.GetChildCount()-1)/2; i++ {
		operationOu.Right = append(operationOu.Right, l.PopAst())
	}

	l.PushAst(operationOu)
}

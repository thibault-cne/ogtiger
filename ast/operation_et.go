package ast

import "ogtiger/parser"

type OperationEt struct {
	Left  Ast
	Right []Ast
	Ctx   parser.IOperationEtContext
}

func (e *OperationEt) Display() string {
	return " et"
}

func (e *OperationEt) Draw(prefix string) {
	// TODO: Draw the AST
}

func (l *AstCreatorListener) OperationEtEnter(ctx parser.IOperationEtContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *AstCreatorListener) OperationEtExit(ctx parser.IOperationEtContext) {
	operationEt := &OperationEt{
		Ctx: ctx,
	}

	if ctx.GetChildCount() == 1 {
		return
	}

	operationEt.Left = l.PopAst()

	// Get the other exprEt
	for i := 0; i < len(ctx.AllOperationComparaison())-1; i++ {
		operationEt.Right = append(operationEt.Right, l.PopAst())
	}

	l.PushAst(operationEt)
}

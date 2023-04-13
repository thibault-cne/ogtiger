package ast

import "ogtiger/parser"

type OperationTantQue struct {
	Cond  Ast
	Block Ast
}

func (e OperationTantQue) Display() string {
	return " si"
}

func (l *AstCreatorListener) OperationTantQueEnter(ctx parser.IOperationTantqueContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *AstCreatorListener) OperationTantQueExit(ctx parser.IOperationTantqueContext) {
	oT := &OperationTantQue{}

	oT.Cond = l.PopAst()
	oT.Block = l.PopAst()

	l.PushAst(oT)
}

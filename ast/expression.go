package ast

import (
	"ogtiger/parser"
)

type Expression struct {
	Left  Ast
	Right Ast
}

func (e *Expression) Display() string {
	return " expression"
}

func (l *AstCreatorListener) ExprEnter(ctx parser.ExpressionContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) ExprExit(ctx parser.ExpressionContext) {
	// Get back the last element of the stack
	expr := &Expression{}

	if ctx.GetChildCount() == 1 {
		return
	}

	expr.Left = l.PopAst()
	expr.Right = l.PopAst()

	l.PushAst(expr)
}

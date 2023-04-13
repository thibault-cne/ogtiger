package ast

import (
	"ogtiger/parser"
)

type Expression struct {
	Left  Ast
	Right Ast
	Ctx   parser.IExpressionContext
}

func (e *Expression) Display() string {
	return " expression"
}

func (l *AstCreatorListener) ExprEnter(ctx parser.IExpressionContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) ExprExit(ctx parser.IExpressionContext) {
	expr := &Expression{
		Ctx: ctx,
	}

	if ctx.GetChildCount() == 1 {
		return
	}

	expr.Left = l.PopAst()
	expr.Right = l.PopAst()

	l.PushAst(expr)
}

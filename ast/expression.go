package ast

import (
	"ogtiger/parser"

	"github.com/goccy/go-graphviz"
)

type Expression struct {
	Left  Ast
	Right Ast
	Ctx   parser.IExpressionContext
}

func (e *Expression) Display() string {
	return " expression"
}

func (e *Expression) Draw(prefix string, g *graphviz.Graphviz) {
	// TODO: Draw the AST
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

package ast

import (
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type Expression struct {
	Left  Ast
	Right Ast
	Ctx   parser.IExpressionContext
	Type  ttype.TigerType
}

func (e *Expression) ReturnType() ttype.TigerType {
	return e.Type
}

func (e *Expression) Display() string {
	return " expression"
}

func (e *Expression) Draw(prefix string, g *cgraph.Graph) {
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

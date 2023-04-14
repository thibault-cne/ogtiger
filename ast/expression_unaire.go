package ast

import (
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type ExpressionUnaire struct {
	Expr Ast
	Ctx  parser.IExpressionUnaireContext
	Type ttype.TigerType
}

func (e *ExpressionUnaire) ReturnType() ttype.TigerType {
	return e.Type
}

func (e *ExpressionUnaire) Display() string {
	return " expressionUnaire"
}

func (e *ExpressionUnaire) Draw(prefix string, g *cgraph.Graph) {
	// TODO: Draw the AST
}

func (l *AstCreatorListener) ExpressionUnaireEnter(ctx parser.IExpressionUnaireContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *AstCreatorListener) ExpressionUnaireExit(ctx parser.IExpressionUnaireContext) {
	expressionUnaire := &ExpressionUnaire{
		Ctx: ctx,
	}

	if ctx.GetChildCount() == 1 {
		return
	}

	// Get the first expr
	// Pop the last element of the stack
	expr := l.PopAst()

	expressionUnaire.Expr = expr

	// Push the new element on the stack
	l.PushAst(expressionUnaire)
}

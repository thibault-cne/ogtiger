package ast

import (
	"ogtiger/parser"

	"github.com/goccy/go-graphviz"
)

type Integer struct {
	Valeur string
	Ctx    parser.EntierContext
}

func (e *Integer) Display() string {
	return " int"
}

func (e *Integer) Draw(prefix string, g *graphviz.Graphviz) {
	// TODO: Draw the AST
}

func (l *AstCreatorListener) IntegerEnter(ctx parser.EntierContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) IntegerExit(ctx parser.EntierContext) {
	// Get back the last element of the stack
	it := &Integer{
		Ctx: ctx,
	}

	it.Valeur = ctx.GetText()

	l.PushAst(it)
}

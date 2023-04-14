package ast

import (
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type Integer struct {
	Valeur string
	Ctx    parser.EntierContext
	Type   ttype.TigerType
}

func (e *Integer) ReturnType() ttype.TigerType {
	return e.Type
}

func (e *Integer) Display() string {
	return " int"
}

func (e *Integer) Draw(prefix string, g *cgraph.Graph) {
	// TODO: Draw the AST
}

func (l *AstCreatorListener) IntegerEnter(ctx parser.EntierContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) IntegerExit(ctx parser.EntierContext) {
	// Get back the last element of the stack
	it := &Integer{
		Ctx:  ctx,
		Type: ttype.NewTigerType(ttype.Int),
	}

	it.Valeur = ctx.GetText()

	l.PushAst(it)
}

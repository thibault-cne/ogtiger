package ast

import (
	"fmt"
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type Integer struct {
	Valeur string
	Ctx    parser.EntierContext
	Type   *ttype.TigerType
}

func (e *Integer) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *Integer) Draw(g *cgraph.Graph) *cgraph.Node {
	id := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(id)
	node.SetLabel(e.Valeur)

	return node
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

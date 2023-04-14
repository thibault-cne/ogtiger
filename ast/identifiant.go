package ast

import (
	"fmt"
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type Identifiant struct {
	Id   string
	Ctx  parser.IIdentifiantContext
	Type ttype.TigerType
}

func (e *Identifiant) ReturnType() ttype.TigerType {
	return e.Type
}

func (e *Identifiant) Draw(g *cgraph.Graph) *cgraph.Node {
	id := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(id)
	node.SetLabel(e.Id)

	return node
}

func (l *AstCreatorListener) IdentifiantEnter(ctx parser.IIdentifiantContext) {
	// Nothing to do
}

func (l *AstCreatorListener) IdentifiantExit(ctx parser.IIdentifiantContext) {
	identifiant := &Identifiant{
		Id:  ctx.GetText(),
		Ctx: ctx,
	}

	l.PushAst(identifiant)
}

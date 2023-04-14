package ast

import (
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type Nil struct {
	Ctx  parser.NilContext
	Type ttype.TigerType
}

func (e *Nil) ReturnType() ttype.TigerType {
	return e.Type
}

func (e *Nil) Display() string {
	return " nil"
}

func (e *Nil) Draw(g *cgraph.Graph) *cgraph.Node {
	node, _ := g.CreateNode("Nil")

	return node
}

func (l *AstCreatorListener) NilEnter(ctx parser.NilContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) NilExit(ctx parser.NilContext) {
	// Get back the last element of the stack
	it := &Nil{
		Ctx:  ctx,
		Type: ttype.NewTigerType(ttype.AnyRecord),
	}

	l.PushAst(it)
}

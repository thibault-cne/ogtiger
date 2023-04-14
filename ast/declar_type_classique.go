package ast

import (
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type DeclarationTypeClassique struct {
	Identifiant Ast
	TType       Ast
	Ctx         parser.DeclarationTypeClassiqueContext
	Type        ttype.TigerType
}

func (e *DeclarationTypeClassique) ReturnType() ttype.TigerType {
	return e.Type
}

func (e *DeclarationTypeClassique) Display() string {
	return " alias"
}

func (e *DeclarationTypeClassique) Draw(g *cgraph.Graph) *cgraph.Node {
	node, _ := g.CreateNode("DeclarationTypeClassique")
	id := e.Identifiant.Draw(g)
	g.CreateEdge("Id", node, id)

	typeNode := e.Type.Draw(g)
	g.CreateEdge("Type", node, typeNode)

	return node
}

func (l *AstCreatorListener) DeclarationTypeClassiqueEnter(ctx parser.DeclarationTypeClassiqueContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) DeclarationTypeClassiqueExit(ctx parser.DeclarationTypeClassiqueContext) {
	// Get back the last element of the stack
	declType := &DeclarationTypeClassique{
		Ctx: ctx,
	}

	declType.Identifiant = l.PopAst()
	declType.TType = l.PopAst()

	l.PushAst(declType)
}

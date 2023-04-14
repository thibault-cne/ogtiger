package ast

import (
	"ogtiger/parser"

	"github.com/goccy/go-graphviz/cgraph"
)

type DeclarationTypeClassique struct {
	Identifiant Ast
	Type        Ast
	Ctx         parser.DeclarationTypeClassiqueContext
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
	declType.Type = l.PopAst()

	l.PushAst(declType)
}

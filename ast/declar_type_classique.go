package ast

import (
	"fmt"
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

func (e *DeclarationTypeClassique) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("DeclarationTypeClassique")

	id := e.Identifiant.Draw(g)
	g.CreateEdge("Id", node, id)

	typeNode := e.TType.Draw(g)
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

	declType.TType = l.PopAst()
	declType.Identifiant = l.PopAst()

	l.PushAst(declType)
}

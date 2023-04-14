package ast

import (
	"ogtiger/parser"

	"github.com/goccy/go-graphviz/cgraph"
)

type DeclarationArrayType struct {
	Identifiant Ast
	Type        Ast
	Ctx         parser.DeclarationArrayTypeContext
}

func (e *DeclarationArrayType) Display() string {
	return " array_type"
}

func (e *DeclarationArrayType) Draw(g *cgraph.Graph) *cgraph.Node {
	node, _ := g.CreateNode("DeclarationArrayType")
	id := e.Identifiant.Draw(g)
	g.CreateEdge("Id", node, id)

	typeNode := e.Type.Draw(g)
	g.CreateEdge("Type", node, typeNode)

	return node
}

func (l *AstCreatorListener) DeclarationArrayTypeEnter(ctx parser.DeclarationArrayTypeContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) DeclarationArrayTypeExit(ctx parser.DeclarationArrayTypeContext) {
	// Get back the last element of the stack
	declArrType := &DeclarationArrayType{
		Ctx: ctx,
	}

	declArrType.Identifiant = l.PopAst()
	declArrType.Type = l.PopAst()

	l.PushAst(declArrType)
}

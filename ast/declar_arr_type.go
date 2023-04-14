package ast

import (
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type DeclarationArrayType struct {
	Identifiant Ast
	AType       Ast
	Ctx         parser.DeclarationArrayTypeContext
	Type        ttype.TigerType
}

func (e *DeclarationArrayType) ReturnType() ttype.TigerType {
	return e.Type
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
	declArrType.AType = l.PopAst()

	l.PushAst(declArrType)
}

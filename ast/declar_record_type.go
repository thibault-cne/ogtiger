package ast

import (
	"ogtiger/parser"

	"github.com/goccy/go-graphviz/cgraph"
)

type DeclarationRecordType struct {
	Identifiant Ast
	Type        Ast
	Ctx         parser.DeclarationRecordTypeContext
}

func (e *DeclarationRecordType) Display() string {
	return " record_type"
}

func (e *DeclarationRecordType) Draw(g *cgraph.Graph) *cgraph.Node {
	node, _ := g.CreateNode("DeclarationRecordType")
	id := e.Identifiant.Draw(g)
	g.CreateEdge("Id", node, id)

	typeNode := e.Type.Draw(g)
	g.CreateEdge("Type", node, typeNode)

	return node
}

func (l *AstCreatorListener) DeclarationRecordTypeEnter(ctx parser.DeclarationRecordTypeContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) DeclarationRecordTypeExit(ctx parser.DeclarationRecordTypeContext) {
	// Get back the last element of the stack
	declRecordType := &DeclarationRecordType{
		Ctx: ctx,
	}

	declRecordType.Identifiant = l.PopAst()
	declRecordType.Type = l.PopAst()

	l.PushAst(declRecordType)
}

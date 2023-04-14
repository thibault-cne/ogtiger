package ast

import (
	"fmt"
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type DeclarationRecordType struct {
	Identifiant Ast
	RType       Ast
	Ctx         parser.DeclarationRecordTypeContext
	Type        ttype.TigerType
}

func (e *DeclarationRecordType) ReturnType() ttype.TigerType {
	return e.Type
}

func (e *DeclarationRecordType) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("DeclarationRecordType")

	id := e.Identifiant.Draw(g)
	g.CreateEdge("Id", node, id)

	typeNode := e.RType.Draw(g)
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

	declRecordType.RType = l.PopAst()
	declRecordType.Identifiant = l.PopAst()

	l.PushAst(declRecordType)
}

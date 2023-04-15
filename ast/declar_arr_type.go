package ast

import (
	"fmt"
	"ogtiger/logger"
	"ogtiger/parser"
	"ogtiger/slt"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type DeclarationArrayType struct {
	Identifiant Ast
	AType       Ast
	Ctx         parser.DeclarationArrayTypeContext
	Type        *ttype.TigerType
}

func (e *DeclarationArrayType) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *DeclarationArrayType) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("DeclarationArrayType")

	id := e.Identifiant.Draw(g)
	g.CreateEdge("Id", node, id)

	typeNode := e.AType.Draw(g)
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

	declArrType.AType = l.PopAst()
	declArrType.Identifiant = l.PopAst()

	// Verify that the type is not already defined
	if _, err := l.Slt.GetSymbol(declArrType.Identifiant.(*Identifiant).Id); err == nil {
		l.Logger.NewSemanticError(logger.ErrorIdIsAlreadyDefinedInScope, &ctx, declArrType.Identifiant.(*Identifiant).Id)
	}

	t := &slt.Symbol{
		Name: declArrType.Identifiant.(*Identifiant).Id,
		Type: ttype.NewArrayType(declArrType.AType.ReturnType()),
	}
	l.Slt.AddSymbol(declArrType.Identifiant.(*Identifiant).Id, t)

	// Add the type to the node
	declArrType.Type = t.Type

	l.PushAst(declArrType)
}

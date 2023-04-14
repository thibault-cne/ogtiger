package ast

import (
	"fmt"
	"ogtiger/logger"
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type DeclarationRecordType struct {
	Identifiant Ast
	Champs      []Ast
	Ctx         parser.DeclarationRecordTypeContext
	Type        *ttype.TigerType
}

func (e *DeclarationRecordType) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *DeclarationRecordType) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("DeclarationRecordType")

	id := e.Identifiant.Draw(g)
	g.CreateEdge("Id", node, id)

	cnodeId := fmt.Sprintf("N%pc", e.Champs)
	cnode, _ := g.CreateNode(cnodeId)
	cnode.SetLabel("Champs")

	for _, champ := range e.Champs {
		typeNode := champ.Draw(g)
		g.CreateEdge("Champ", cnode, typeNode)
	}

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

	for range ctx.AllDeclarationChamp() {
		declRecordType.Champs = append(declRecordType.Champs, l.PopAst())
	}

	declRecordType.Identifiant = l.PopAst()

	// Verify that the type is not already defined
	if _, err := l.Slt.GetSymbol(declRecordType.Identifiant.(*Identifiant).Id); err == nil {
		l.Logger.NewSemanticError(logger.ErrorIdIsAlreadyDefinedInScope, &ctx, declRecordType.Identifiant.(*Identifiant).Id)
	}

	typeFields := []*ttype.RecordField{}
	for _, champ := range declRecordType.Champs {
		if champ.(*DeclarationChamp).Right.ReturnType() == nil {
			l.Logger.NewSemanticError(logger.ErrorTypeIsNotDefined, &ctx, champ.(*DeclarationChamp).Right.(*Identifiant).Id)
		}
		typeFields = append(typeFields, &ttype.RecordField{
			Name: champ.(*DeclarationChamp).Left.(*Identifiant).Id,
			Type: champ.(*DeclarationChamp).Right.ReturnType(),
		})
	}

	l.Slt.CreateSymbol(declRecordType.Identifiant.(*Identifiant).Id, ttype.NewRecordType(typeFields))

	l.PushAst(declRecordType)
}

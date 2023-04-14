package ast

import (
	"fmt"
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type DeclarationChamp struct {
	Left  Ast
	Right Ast
	Ctx   parser.IDeclarationChampContext
	Type  *ttype.TigerType
}

func (e *DeclarationChamp) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *DeclarationChamp) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("DeclarationChamp")

	left := e.Left.Draw(g)
	g.CreateEdge("Left", node, left)

	right := e.Right.Draw(g)
	g.CreateEdge("Rigth", node, right)

	return node
}

func (l *AstCreatorListener) DeclarationChampEnter(ctx parser.IDeclarationChampContext) {
	// Nothing to do
}

func (l *AstCreatorListener) DeclarationChampExit(ctx parser.IDeclarationChampContext) {
	declarationChamp := &DeclarationChamp{
		Ctx: ctx,
	}

	declarationChamp.Right = l.PopAst()
	declarationChamp.Left = l.PopAst()

	l.PushAst(declarationChamp)
}

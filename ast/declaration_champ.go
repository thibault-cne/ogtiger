package ast

import (
	"ogtiger/parser"

	"github.com/goccy/go-graphviz/cgraph"
)

type DeclarationChamp struct {
	Left  Ast
	Rigth Ast
	Ctx   parser.IDeclarationChampContext
}

func (e *DeclarationChamp) Display() string {
	return " declarationChamp"
}

func (e *DeclarationChamp) Draw(g *cgraph.Graph) *cgraph.Node {
	node, _ := g.CreateNode("DeclarationChamp")
	left := e.Left.Draw(g)
	g.CreateEdge("Left", node, left)

	right := e.Rigth.Draw(g)
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

	declarationChamp.Left = l.PopAst()
	declarationChamp.Rigth = l.PopAst()

	l.PushAst(declarationChamp)
}

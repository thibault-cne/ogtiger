package ast

import (
	"fmt"
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type DeclarationValeur struct {
	Id    Ast
	VType Ast
	Expr  Ast
	Ctx   parser.IDeclarationValeurContext
	Type  ttype.TigerType
}

func (e *DeclarationValeur) ReturnType() ttype.TigerType {
	return e.Type
}

func (e *DeclarationValeur) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("DeclarationValeur")

	id := e.Id.Draw(g)
	g.CreateEdge("Id", node, id)

	if e.VType != nil {
		typeNode := e.VType.Draw(g)
		g.CreateEdge("Type", node, typeNode)
	}

	expr := e.Expr.Draw(g)
	g.CreateEdge("Expr", node, expr)

	return node
}

func (l *AstCreatorListener) DeclarationValeurEnter(ctx parser.IDeclarationValeurContext) {
	// Nothing to do
}

func (l *AstCreatorListener) DeclarationValeurExit(ctx parser.IDeclarationValeurContext) {
	declarationValeur := &DeclarationValeur{
		Ctx: ctx,
	}

	declarationValeur.Expr = l.PopAst()

	if len(ctx.AllIdentifiant()) > 1 {
		declarationValeur.VType = l.PopAst()
	}

	declarationValeur.Id = l.PopAst()

	l.PushAst(declarationValeur)
}

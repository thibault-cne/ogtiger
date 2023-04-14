package ast

import (
	"fmt"
	"ogtiger/logger"
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type DeclarationValeur struct {
	Id    Ast
	VType Ast
	Expr  Ast
	Ctx   parser.IDeclarationValeurContext
	Type  *ttype.TigerType
}

func (e *DeclarationValeur) ReturnType() *ttype.TigerType {
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
	declarationValeur.Type = declarationValeur.Expr.ReturnType()

	if len(ctx.AllIdentifiant()) > 1 {
		declarationValeur.VType = l.PopAst()
		declarationValeur.Type = declarationValeur.VType.ReturnType()

		// Verify that the type exists
		if _, err := l.Slt.GetSymbol(declarationValeur.VType.(*Identifiant).Id); err != nil {
			l.Logger.NewSemanticError(logger.ErrorTypeIsNotDefined, ctx, declarationValeur.VType.(*Identifiant).Id)
		}
	}

	declarationValeur.Id = l.PopAst()
	fmt.Printf("\n\n%#+v\n\n", declarationValeur.Id)
	fmt.Printf("\n\n%#+v\n\n", declarationValeur.Expr)
	fmt.Printf("\n\n%#+v\n\n", l.AstStack[len(l.AstStack)-2])
	fmt.Printf("\n\n%#+v\n\n", declarationValeur.VType)
	if _, err := l.Slt.GetSymbolInScoope(declarationValeur.Id.(*Identifiant).Id); err == nil {
		l.Logger.NewSemanticError(logger.ErrorIdIsAlreadyDefinedInScope, ctx, declarationValeur.Id.(*Identifiant).Id)
	}

	l.PushAst(declarationValeur)
}

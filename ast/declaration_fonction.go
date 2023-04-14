package ast

import (
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type DeclarationFontion struct {
	Id    Ast
	FType Ast
	Args  []Ast
	Expr  Ast
	Ctx   parser.IDeclarationFonctionContext
	Type  ttype.TigerType
}

func (e *DeclarationFontion) ReturnType() ttype.TigerType {
	return e.Type
}

func (e *DeclarationFontion) Display() string {
	return " declarationFontion"
}

func (e *DeclarationFontion) Draw(g *cgraph.Graph) *cgraph.Node {
	node, _ := g.CreateNode("DeclarationFontion")
	id := e.Id.Draw(g)
	g.CreateEdge("Id", node, id)

	if e.FType != nil {
		typeNode := e.FType.Draw(g)
		g.CreateEdge("Type", node, typeNode)
	}

	for _, arg := range e.Args {
		argNode := arg.Draw(g)
		g.CreateEdge("Arg", node, argNode)
	}

	exprNode := e.Expr.Draw(g)
	g.CreateEdge("Expr", node, exprNode)

	return node
}

func (l *AstCreatorListener) DeclarationFontionEnter(ctx parser.IDeclarationFonctionContext) {
	// Nothing to do
}

func (l *AstCreatorListener) DeclarationFontionExit(ctx parser.IDeclarationFonctionContext) {
	declarationFontion := &DeclarationFontion{Ctx: ctx}

	declarationFontion.Id = l.PopAst()

	// Pop all args
	for i := 0; i < len(ctx.AllDeclarationChamp()); i++ {
		declarationFontion.Args = append(declarationFontion.Args, l.PopAst())
	}

	// Pop the type if it exists
	if len(ctx.AllIdentifiant()) > 1 {
		declarationFontion.FType = l.PopAst()
	}

	declarationFontion.Expr = l.PopAst()

	l.PushAst(declarationFontion)
}

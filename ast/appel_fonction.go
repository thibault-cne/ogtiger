package ast

import (
	"fmt"
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type AppelFonction struct {
	Identifiant Ast
	Args        []Ast
	Ctx         parser.AppelFonctionContext
	Type        ttype.TigerType
}

func (e *AppelFonction) ReturnType() ttype.TigerType {
	return e.Type
}

func (e *AppelFonction) Display() string {
	return " appel"
}

func (a *AppelFonction) Draw(g *cgraph.Graph) *cgraph.Node {
	node, _ := g.CreateNode("AppelFonction")
	id := a.Identifiant.Draw(g)
	g.CreateEdge("Id", node, id)

	for i, arg := range a.Args {
		argNode := arg.Draw(g)
		g.CreateEdge(fmt.Sprintf("arg %d", i), node, argNode)
	}

	return node
}

func (l *AstCreatorListener) AppelFonctionEnter(ctx parser.AppelFonctionContext) {
	// Nothing to do
}

func (l *AstCreatorListener) AppelFonctionExit(ctx parser.AppelFonctionContext) {
	appelFonction := &AppelFonction{
		Ctx: ctx,
	}

	// Get the identifiant
	expr := l.PopAst()

	appelFonction.Identifiant = expr

	// Get the args
	for i := 2; i < ctx.GetChildCount()-1; i += 2 {
		// Pop the last element of the stack
		// Add it to the args
		appelFonction.Args = append(appelFonction.Args, l.PopAst())
	}

	// Push the new element on the stack
	l.PushAst(appelFonction)
}

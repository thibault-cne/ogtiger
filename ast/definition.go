package ast

import (
	"fmt"
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type Definition struct {
	Declarations []Ast
	Expressions  []Ast
	Ctx          parser.IDefinitionContext
	Type         *ttype.TigerType
}

func (e *Definition) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *Definition) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("Definition")

	for _, declaration := range e.Declarations {
		declarationNode := declaration.Draw(g)
		g.CreateEdge("", node, declarationNode)
	}

	for _, expression := range e.Expressions {
		expressionNode := expression.Draw(g)
		g.CreateEdge("", node, expressionNode)
	}

	return node
}

func (l *AstCreatorListener) DefinitionEnter(ctx parser.IDefinitionContext) {
	// Create a TDS for the definition
	l.Slt = l.Slt.CreateRegion()
}

func (l *AstCreatorListener) DefinitionExit(ctx parser.IDefinitionContext) {
	expr := &Definition{
		Ctx: ctx,
	}

	for range ctx.AllExpression() {
		// Prepend the expressions to the list
		expr.Expressions = append([]Ast{ l.PopAst() }, expr.Expressions...)
	}


	for range ctx.AllDeclaration() {
		// Prepend the declarations to the list
		d := l.PopAst()
		expr.Declarations = append([]Ast{ d }, expr.Declarations...)

		// Add the declaration to the list of symbols
		switch d := d.(type) {
		case *DeclarationValeur:
			l.Slt.CreateSymbol(d.Id.(*Identifiant).Id, d.ReturnType())
		}
	}

	// Pop the TDS
	l.Slt = l.Slt.Parent

	l.PushAst(expr)
}

package ast

import (
	"ogtiger/parser"

	"github.com/goccy/go-graphviz/cgraph"
)

type Definition struct {
	Declarations []Ast
	Expressions  []Ast
	Ctx          parser.IDefinitionContext
}

func (e *Definition) Display() string {
	return " letin"
}

func (e *Definition) Draw(g *cgraph.Graph) *cgraph.Node {
	node, _ := g.CreateNode("Let in")
	def, _ := g.CreateNode("Def")
	expr, _ := g.CreateNode("Expr")

	for _, declaration := range e.Declarations {
		declarationNode := declaration.Draw(g)
		g.CreateEdge("", def, declarationNode)
	}

	for _, expression := range e.Expressions {
		expressionNode := expression.Draw(g)
		g.CreateEdge("", expr, expressionNode)
	}

	g.CreateEdge("Def", node, def)
	g.CreateEdge("Expr", node, expr)

	return node
}

func (l *AstCreatorListener) DefinitionEnter(ctx parser.IDefinitionContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) DefinitionExit(ctx parser.IDefinitionContext) {
	expr := &Definition{
		Ctx: ctx,
	}

	for range ctx.AllDeclaration() {
		expr.Declarations = append(expr.Declarations, l.PopAst())
	}

	for range ctx.AllExpression() {
		expr.Expressions = append(expr.Expressions, l.PopAst())
	}

	l.PushAst(expr)
}

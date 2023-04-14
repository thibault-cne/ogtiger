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
	// l.AstStack = append(l.AstStack, &Expr{})
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
		expr.Declarations = append([]Ast{ l.PopAst() }, expr.Declarations...)
	}

	l.PushAst(expr)
}

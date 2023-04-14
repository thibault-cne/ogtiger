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
	Type         ttype.TigerType
}

func (e *Definition) ReturnType() ttype.TigerType {
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
		expr.Expressions = append(expr.Expressions, l.PopAst())
	}

	// Reverse the order of the expressions
	for i, j := 0, len(expr.Expressions)-1; i < j; i, j = i+1, j-1 {
		expr.Expressions[i], expr.Expressions[j] = expr.Expressions[j], expr.Expressions[i]
	}

	for range ctx.AllDeclaration() {
		expr.Declarations = append(expr.Declarations, l.PopAst())
	}

	// Reverse the order of the declarations
	for i, j := 0, len(expr.Declarations)-1; i < j; i, j = i+1, j-1 {
		expr.Declarations[i], expr.Declarations[j] = expr.Declarations[j], expr.Declarations[i]
	}

	l.PushAst(expr)
}

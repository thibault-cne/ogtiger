package ast

import (
	"fmt"
	"ogtiger/logger"
	"ogtiger/parser"
	"ogtiger/slt"
	"ogtiger/ttype"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/goccy/go-graphviz/cgraph"
)

type Definition struct {
	Declarations []Ast
	Expressions  []Ast
	Slt          *slt.SymbolTable
	Ctx          parser.IDefinitionContext
	Type         *ttype.TigerType
}

func (e *Definition) VisitSemControl(slt *slt.SymbolTable, L *logger.StepLogger) antlr.ParserRuleContext {
	for _, declaration := range e.Declarations {
		declaration.VisitSemControl(e.Slt, L)
	}

	for _, expression := range e.Expressions {
		expression.VisitSemControl(e.Slt, L)
	}

	return e.Ctx
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

	expr.Type = expr.Expressions[len(expr.Expressions)-1].ReturnType()

	for range ctx.AllDeclaration() {
		// Prepend the declarations to the list
		expr.Declarations = append([]Ast{ l.PopAst() }, expr.Declarations...)
	}

	// Pop the TDS
	expr.Slt = l.Slt
	l.Slt = l.Slt.Parent

	l.PushAst(expr)
}

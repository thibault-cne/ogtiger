package ast

import (
	"fmt"
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type ExpressionUnaire struct {
	Expr Ast
	Ctx  parser.IExpressionUnaireContext
	Type ttype.TigerType
}

func (e *ExpressionUnaire) ReturnType() ttype.TigerType {
	return e.Type
}

func (e *ExpressionUnaire) Display() string {
	return " expressionUnaire"
}

func (e *ExpressionUnaire) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("ExpressionUnaire")

	expr := e.Expr.Draw(g)
	g.CreateEdge("Expr", node, expr)

	return node
}

func (l *AstCreatorListener) ExpressionUnaireEnter(ctx parser.IExpressionUnaireContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *AstCreatorListener) ExpressionUnaireExit(ctx parser.IExpressionUnaireContext) {
	// Nothing to do
}

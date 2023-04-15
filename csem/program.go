package ast

import (
	"fmt"
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type Program struct {
	Expr Ast
	Ctx  parser.IProgramContext
	Type *ttype.TigerType
}

func (e *Program) ReturnType() *ttype.TigerType {
	return e.Type
}

func (e *Program) Draw(g *cgraph.Graph) *cgraph.Node {
	nodeId := fmt.Sprintf("N%p", e)
	node, _ := g.CreateNode(nodeId)
	node.SetLabel("Program")

	expr := e.Expr.Draw(g)
	g.CreateEdge("Expr", node, expr)

	return node
}

func (l *SemControlListener) ProgramEnter(ctx parser.IProgramContext) {
	// l.AstStack = append(l.AstStack, &Program{})
}

func (l *SemControlListener) ProgramExit(ctx parser.IProgramContext) {
	prog := &Program{
		Ctx:  ctx,
		Type: ttype.NewTigerType(ttype.NoReturn),
	}

	prog.Expr = l.PopAst()

	// Push the new element on the stack
	l.PushAst(prog)
}

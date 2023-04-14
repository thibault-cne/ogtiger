package ast

import (
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type OperationTantQue struct {
	Cond  Ast
	Block Ast
	Ctx   parser.IOperationTantqueContext
	Type  ttype.TigerType
}

func (e *OperationTantQue) ReturnType() ttype.TigerType {
	return e.Type
}

func (e *OperationTantQue) Display() string {
	return " tantque"
}

func (e *OperationTantQue) Draw(prefix string, g *cgraph.Graph) {
	// TODO: Draw the AST
}

func (l *AstCreatorListener) OperationTantQueEnter(ctx parser.IOperationTantqueContext) {
	// l.AstStack = append(l.AstStack, &ExprOu{})
}

func (l *AstCreatorListener) OperationTantQueExit(ctx parser.IOperationTantqueContext) {
	oT := &OperationTantQue{
		Ctx: ctx,
	}

	oT.Cond = l.PopAst()
	oT.Block = l.PopAst()

	l.PushAst(oT)
}

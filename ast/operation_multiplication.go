package ast

import (
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/goccy/go-graphviz/cgraph"
)

type OperationMultiplication struct {
	Left  Ast
	Right []*OperationMultiplicationFD
	Ctx   parser.IOperationMultiplicationContext
	Type  ttype.TigerType
}

func (e *OperationMultiplication) ReturnType() ttype.TigerType {
	return e.Type
}

type OperationMultiplicationFD struct {
	Op    string
	Right Ast
}

func (e *OperationMultiplication) Display() string {
	return " multiplication"
}

func (e *OperationMultiplication) Draw(prefix string, g *cgraph.Graph) {
	// TODO: Draw the AST
}

func (l *AstCreatorListener) OperationMultiplicationEnter(ctx parser.IOperationMultiplicationContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) OperationMultiplicationExit(ctx parser.IOperationMultiplicationContext) {
	// Get back the last element of the stack
	opMultiplication := &OperationMultiplication{
		Ctx: ctx,
	}

	if ctx.GetChildCount() == 1 {
		return
	}

	opMultiplication.Left = l.PopAst()

	// Get minus and plus and term number
	for i := 0; i < (ctx.GetChildCount()-1)/2; i++ {
		right := &OperationMultiplicationFD{}

		right.Op = ctx.GetChild(2*i + 1).(*antlr.TerminalNodeImpl).GetText()
		right.Right = l.PopAst()

		opMultiplication.Right = append(opMultiplication.Right, right)
	}

	l.PushAst(opMultiplication)
}

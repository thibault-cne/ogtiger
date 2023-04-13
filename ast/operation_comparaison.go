package ast

import (
	"ogtiger/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type OperationComparaison struct {
	Left  Ast
	Right []*OperationComparaisonFD
	Ctx   parser.IOperationComparaisonContext
}

type OperationComparaisonFD struct {
	Op    string
	Right Ast
}

func (e *OperationComparaison) Display() string {
	return " comparaison"
}

func (l *AstCreatorListener) OperationComparaisonEnter(ctx parser.IOperationComparaisonContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) OperationComparaisonExit(ctx parser.IOperationComparaisonContext) {
	// Get back the last element of the stack
	opCompar := &OperationComparaison{
		Ctx: ctx,
	}

	if ctx.GetChildCount() == 1 {
		return
	}

	opCompar.Left = l.PopAst()

	// Get minus and plus and term number
	for i := 0; i < (ctx.GetChildCount()-1)/2; i++ {
		right := &OperationComparaisonFD{}

		right.Op = ctx.GetChild(2*i + 1).(*antlr.TerminalNodeImpl).GetText()
		right.Right = l.PopAst()

		opCompar.Right = append(opCompar.Right, right)
	}

	l.PushAst(opCompar)
}

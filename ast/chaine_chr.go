package ast

import (
	"ogtiger/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type ChaineChr struct {
	Valeur string
	Ctx    parser.ChaineChrContext
}

func (e *ChaineChr) Display() string {
	return " alias"
}

func (l *AstCreatorListener) ChaineChrEnter(ctx parser.ChaineChrContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) ChaineChrExit(ctx parser.ChaineChrContext) {
	// Get back the last element of the stack
	chainChr := &ChaineChr{
		Ctx: ctx,
	}

	chainChr.Valeur = ctx.GetChild(0).(*antlr.TerminalNodeImpl).GetText()

	l.PushAst(chainChr)
}

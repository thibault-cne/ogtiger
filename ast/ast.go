package ast

import (
	"ogtiger/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type AstCreatorListener struct {
	AstStack []Ast 
	*parser.BasetigerVisitor
}

type Ast interface {}

// Example
func (ast *AstCreatorListener) VisitTerminal(node antlr.TerminalNode) {
	// fmt.Printf("%v\n", node.GetText())
}

func (ast *AstCreatorListener) VisitErrorNode(node antlr.ErrorNode) {
	// fmt.Printf("%v\n", node.GetText())
}

func (s *AstCreatorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	switch c := ctx.(type) {
	case parser.IFactorContext:
		s.FactorEnter(c)
	case parser.IExprContext:
		s.ExprEnter(c)
	case parser.ITermContext:
		s.TermEnter(c)
	default:
		break
	}
}

func (s *AstCreatorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	switch c := ctx.(type) {
	case parser.IFactorContext:
		s.FactorExit(c)
	case parser.IExprContext:
		s.ExprExit(c)
	case parser.ITermContext:
		s.TermExit(c)
	default:
		break
	}
}

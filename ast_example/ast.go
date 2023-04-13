package ast

import (
	"fmt"
	"ogtiger/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type AstCreatorListener struct {
	AstStack []Ast 
	*parser.BasetigerVisitor
}

type Ast interface {
	Display() string
}

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

func (s * AstCreatorListener) DisplayAST() {
	fmt.Printf("\nAST\n")
	display(s.AstStack[0], "", true)
}

func display(a Ast, prefix string, isLast bool) {
	if a == nil {
		return
	}

	if isLast {
		fmt.Printf("%s└───%s\n", prefix, a.Display())
		prefix += "    "
	} else {
		fmt.Printf("%s├───%s\n", prefix, a.Display())
		prefix += "    "
	}

	switch c := a.(type) {
	case *Expr:
		if (len(c.Right) == 0) {
			display(c.Left, prefix, true)
		} else {
			display(c.Left, prefix, false)
		}

		for i, r := range c.Right {
			if i == len(c.Right) - 1 {
				display(r.Right, prefix, true)
			} else {
				display(r.Right, prefix, false)
			}
		}
	case *Term:
		if (len(c.Right) == 0) {
			display(c.Left, prefix, true)
		} else {
			display(c.Left, prefix, false)
		}

		for i, r := range c.Right {
			if i == len(c.Right) - 1 {
				display(r.Right, prefix, true)
			} else {
				display(r.Right, prefix, false)
			}
		}
	case *Factor:
		if (c.Expr != nil) {
			display(c.Expr, prefix, true)
		}
	default:
		break
	}
}

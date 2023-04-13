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
	// TODO: FILL
	default:
		break
	}
}

func (s *AstCreatorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	switch c := ctx.(type) {
	// TODO: FILL
	default:
		break
	}
}

func (s *AstCreatorListener) DisplayAST() {
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
	// TODO: FILL
	default:
		break
	}
}

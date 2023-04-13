package ast

import (
	"fmt"
	"ogtiger/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type AstCreatorListener struct {
	Output string
	*parser.BasetigerVisitor
}

// Example
func (ast *AstCreatorListener) VisitTerminal(node antlr.TerminalNode) {
	fmt.Printf("%v\n", node.GetText())
}

func (ast *AstCreatorListener) VisitErrorNode(node antlr.ErrorNode) {
	// fmt.Printf("%v\n", node.GetText())
}

func (s *AstCreatorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	// fmt.Printf("%v\n", ctx.GetText())
}

func (s *AstCreatorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	// fmt.Printf("%v\n", ctx.GetText())
}
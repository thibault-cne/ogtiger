package main

import (
	"ogtiger/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type echoListener struct {
	*parser.BasetigerVisitor
}

// Example
func (s *echoListener) VisitTerminal(node antlr.TerminalNode) {
	//fmt.Printf("%v\n", node.GetText())
}

func (s *echoListener) VisitErrorNode(node antlr.ErrorNode) {
	// fmt.Printf("%v\n", node.GetText())
}

func (s *echoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	// fmt.Printf("%v\n", ctx.GetText())
}

func (s *echoListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	// fmt.Printf("%v\n", ctx.GetText())
}

// Parse the input expression and build the AST
func parse(input string) {
	// Create an input stream from the input string
	inputStream := antlr.NewInputStream(input)

	// Create a lexer from the input stream
	lexer := parser.NewtigerLexer(inputStream)

	// Create a token stream from the lexer
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create a parser from the token stream
	p := parser.NewtigerParser(tokenStream)

	// Create a listener to build the AST
	listener := &echoListener{}

	// Walk the parse tree using the listener to build the AST
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Expr())
}

func main() {
	input := "2 * ((3 + 4)"
	parse(input)
}

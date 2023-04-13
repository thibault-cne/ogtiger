package main

import (
	"ogtiger/ast"
	"ogtiger/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

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
	listener := &ast.AstCreatorListener{}

	// Walk the parse tree using the listener to build the AST
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Expr())
}

func main() {
	input := "2 * (3 + 4)"
	parse(input)
}

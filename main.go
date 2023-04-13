package main

import (
	"fmt"
	"ogtiger/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func main() {
	// Create a new input stream from your source code
	inputStream := antlr.NewInputStream("2 * (3 + 4)")

	// Create a lexer that reads from the input stream
	lexer := parser.NewtigerLexer(inputStream)

	// Create a token stream that reads from the lexer
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)

	// Create a parser that reads from the token stream
	TigerParser := parser.NewtigerParser(tokenStream)

	// Call the top-level rule of your grammar to start parsing
	result := TigerParser.Program()

	// Print the result
	fmt.Printf("%#v\n", result)
}

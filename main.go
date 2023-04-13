package main

import (
	"ogtiger/ast"
	"ogtiger/logger"
	"ogtiger/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type emptyErrorListener struct {
	*antlr.DefaultErrorListener
}

func (l *emptyErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	log.LogErrorAfterf("SyntaxError at line %d:%d => %s", line, column, msg)
	failed = true
}

var failed = false
var log = logger.NewStepLogger(2)

// Parse the input expression and build the AST
func parse(input string) {
	emptyErrorListener := &emptyErrorListener{}

	// Create an input stream from the input string
	inputStream := antlr.NewInputStream(input)

	// Create a lexer from the input stream
	log.Log("Lexing the input expression")

	lexer := parser.NewtigerLexer(inputStream)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(emptyErrorListener)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	log.Log("Lexing complete")
	log.Step()

	if failed {
		return
	}

	// Create a parser from the token stream
	log.Log("Parsing the input expression")
	p := parser.NewtigerParser(tokenStream)
	p.RemoveErrorListeners()
	p.AddErrorListener(emptyErrorListener)
	listener := &ast.AstCreatorListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Program())

	if failed {
		return
	}

	log.Log("Parsing complete")
	log.Step()

	listener.DisplayAST()
}

func main() {
	input := "(1 + ((2 * 3))) / 4 * 5"
	parse(input)
}

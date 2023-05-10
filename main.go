package main

import (
	"ogtiger/asm"
	"ogtiger/ast"
	"ogtiger/logger"
	"ogtiger/options"
	"ogtiger/parser"
	"os"

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
func parse(input string, options *options.Options) {
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
	listener := ast.NewAstCreatorListener(log)
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Program())

	if failed {
		return
	}

	log.Log("Parsing complete")
	log.Step()

	prog := listener.GetProgram()
	prog.VisitSemControl(listener.Slt, log)

	if failed {
		return
	}

	log.Log("Semantic Controls complete")

	log.DisplaySemanticErrors()

	if options.Ast != "" {
		listener.DisplayAST(options.Ast)
	}
	if options.Slt != "" {
		listener.Slt.DisplaySLT(options.Slt)
	}
	if options.Asm != "" {
		writer := asm.NewAssemblyWriter()
		listener.AstStack[0].EnterAsm(writer, listener.Slt)
		writer.WriteToFile(options.Ast)
	}
}

func main() {
	options, err := options.Parse(os.Args)

	if err != nil {
		log.LogErrorf("Error while parsing options: %s\n", err)
		return
	}

	data, err := os.ReadFile(options.File)
	if err != nil {
		log.LogErrorf("Error while reading file: %s\n", err)
		return
	}

	// Create the output folder
	if _, err := os.Stat("output"); os.IsNotExist(err) {
		os.Mkdir("output", 0731)
	}

	parse(string(data), options)
}

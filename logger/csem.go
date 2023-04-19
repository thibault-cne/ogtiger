package logger

import (
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type SemError string

const (
	ErrorIdIsAlreadyDefinedInScope 	SemError = "identifier %v is already defined in this scope"
	ErrorTypeIsNotDefined          	SemError = "type %v is not defined"
	ErrorFieldIsAlreadyDefined     	SemError = "field %v is already defined"
	ErrorFunctionNotDefined			SemError = "Function %v is not defined"
	ErrorWrongNumberOfArgs			SemError = "Wrong number of arguments for function %v. Expected %v arguments, got %v instead"
	ErrorWrongTypeOfArgs			SemError = "Wrong type of arguments for function %v. Expected %v for argument %v, got %v instead"
	ErrorVarIsNotAnArray			SemError = "Variable %v is not an array"
	ErrorFieldDoesNotExist			SemError = "Field %v does not exist in type %v"
	ErrorArrayIndexIsNotAnInteger	SemError = "Array index %v is not an integer"
	ErrorArraySizeIsNotAnInteger	SemError = "Array size %v is not an integer"
	ErrorWrongDefaultValueArray		SemError = "Wrong default value for array of type %v. Expected %v, got %v instead"
	ErrorWrongTypesExpression		SemError = "Wrong types for expression. Left is %v, and right is %v"
)

type SemanticError struct {
	Err SemError
	Ctx antlr.ParserRuleContext
}

func NewSemanticError(err SemError, ctx antlr.ParserRuleContext, args ...interface{}) *SemanticError {
	if len(args) == 0 {
		return &SemanticError{
			Err: err,
			Ctx: ctx,
		}
	}
	return &SemanticError{
		Err: SemError(fmt.Sprintf(string(err), args...)),
		Ctx: ctx,
	}
}

func (e SemanticError) Error() string {
	// System.out.println("[\033[31;1merror\033[0m] : " + error + " at line " + ctx.getStart().getLine() + " column " + ctx.getStart().getCharPositionInLine());
	// String program = ctx.getStart().getInputStream().toString();
	// String currentLine = program.split("\n")[ctx.getStart().getLine() - 1];
	// System.out.println("   "+currentLine);
	// System.out.print("   ");
	// for (int i = 0; i < ctx.getStart().getCharPositionInLine()+offset; i++) {
	// 	System.out.print(" ");
	// }
	// System.out.println("^");

	var out string
	out = fmt.Sprintf("[\033[31;1merror\033[0m] : %v at line %v column %v\n", e.Err, e.Ctx.GetStart().GetLine(), e.Ctx.GetStart().GetColumn())
	program := e.Ctx.GetStart().GetInputStream()
	line := strings.Split(fmt.Sprintf("%s", program), "\n")[e.Ctx.GetStart().GetLine()-1]
	out += fmt.Sprintf("%v", line)
	out += "\n"
	for i := 0; i < e.Ctx.GetStart().GetColumn(); i++ {
		out += " "
	}
	out += "^"
	return out
}

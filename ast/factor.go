package ast

import (
	"ogtiger/parser"
	"strconv"
)

type Factor struct {
	Expr *Expr
	Digit int
}

func  (l *AstCreatorListener) FactorEnter(ctx parser.IFactorContext) {
	// l.AstStack = append(l.AstStack, &Factor{})
}

func  (l *AstCreatorListener) FactorExit(ctx parser.IFactorContext) {
	// Get back the last element of the stack
	factor := &Factor{}

	// Pop the last element of the stack
	// l.AstStack = l.AstStack[:len(l.AstStack)-1]

	// Fill the element
	if len(ctx.AllDIGIT()) > 0 {
		digits := ctx.AllDIGIT()
		temp := ""

		for _, digit := range digits {
			temp += digit.GetText()
		}

		factor.Digit, _ = strconv.Atoi(temp)
	} else {
		factor.Expr = l.AstStack[len(l.AstStack)-1].(*Expr)

		// Pop the last element of the stack
		l.AstStack = l.AstStack[:len(l.AstStack)-1]
	}

	// Push the new element on the stack
	l.AstStack = append(l.AstStack, factor)
}
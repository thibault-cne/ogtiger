package ast

import "ogtiger/parser"

type InstanciationType struct {
	Identifiant Ast
	Identifiants []Ast
	Expressions []Ast
}

func (i InstanciationType) Display() string {
	return " instanciationType"
}

func  (l *AstCreatorListener) InstanciationTypeEnter(ctx parser.InstanciationTypeContext) {
	// Nothing to do
}

func  (l *AstCreatorListener) InstanciationTypeExit(ctx parser.InstanciationTypeContext) {
	instanciationType := &InstanciationType{}

	// Get the identifiant
	identifiant := l.PopAst()

	instanciationType.Identifiant = identifiant

	// Get the identifiants count
	identifiantsCount := len(ctx.AllIdentifiant())

	// Get the identifiants and the expressions
	for i := 0; i < identifiantsCount; i++ {
		// Pop the last element of the stack
		// Add it to the identifiants
		instanciationType.Identifiants = append(instanciationType.Identifiants, l.PopAst())

		// Pop the last element of the stack
		// Add it to the expressions
		instanciationType.Expressions = append(instanciationType.Expressions, l.PopAst())
	}

	// Push the new element on the stack
	l.PushAst(instanciationType)
}
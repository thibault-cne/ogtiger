package ast

import (
	"ogtiger/parser"
	"ogtiger/ttype"

	"github.com/goccy/go-graphviz/cgraph"
)

type DeclarationRecordType struct {
	Identifiant Ast
	RType       Ast
	Ctx         parser.DeclarationRecordTypeContext
	Type        ttype.TigerType
}

func (e *DeclarationRecordType) ReturnType() ttype.TigerType {
	return e.Type
}

func (e *DeclarationRecordType) Display() string {
	return " record_type"
}

func (e *DeclarationRecordType) Draw(prefix string, g *cgraph.Graph) {
	// TODO: Draw the AST
}

func (l *AstCreatorListener) DeclarationRecordTypeEnter(ctx parser.DeclarationRecordTypeContext) {
	// l.AstStack = append(l.AstStack, &Expr{})
}

func (l *AstCreatorListener) DeclarationRecordTypeExit(ctx parser.DeclarationRecordTypeContext) {
	// Get back the last element of the stack
	declRecordType := &DeclarationRecordType{
		Ctx: ctx,
	}

	declRecordType.Identifiant = l.PopAst()
	declRecordType.RType = l.PopAst()

	l.PushAst(declRecordType)
}

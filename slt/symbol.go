package slt

import "ogtiger/ttype"

type Symbol struct {
	Name   string
	Type   *ttype.TigerType
	
	// For variables and parameters
	Offset int

	// For functions
	SymbolTable *SymbolTable
}

package slt

import (
	"fmt"
	"ogtiger/ttype"
)

var REGION_CPT = 0

type SymbolTable struct {
	Table  		map[string]*Symbol
	Region 		int
	Scope  		int
	
	Parent 		*SymbolTable
	Children  	[]*SymbolTable

	// Offset to handle the stack
	Offset 		int
}

func NewSymbolTable() *SymbolTable {
	slt := &SymbolTable{
		Table:  make(map[string]*Symbol),
		Region: 0,
		Scope:  0,

		Offset: -4,
	}

	// Add the built-in functions and types
	print := &Symbol{
		Name: "print",
		Type: ttype.NewFunctionType(ttype.NewTigerType(ttype.NoReturn), []*ttype.FunctionParameter{}),
		SymbolTable: &SymbolTable{
			Table:  make(map[string]*Symbol),
			Region: REGION_CPT + 1,
			Scope:  1,
			Offset: -4,
		},
	}
	flush := &Symbol{
		Name: "flush",
		Type: ttype.NewFunctionType(ttype.NewTigerType(ttype.NoReturn), []*ttype.FunctionParameter{}),
		SymbolTable: &SymbolTable{
			Table:  make(map[string]*Symbol),
			Region: REGION_CPT + 2,
			Scope:  1,
			Offset: -4,
		},
	}
	getchar := &Symbol{
		Name: "getchar",
		Type: ttype.NewFunctionType(ttype.NewTigerType(ttype.String), []*ttype.FunctionParameter{}),
		SymbolTable: &SymbolTable{
			Table:  make(map[string]*Symbol),
			Region: REGION_CPT + 3,
			Scope:  1,
			Offset: -4,
		},
	}
	REGION_CPT += 3
	slt.AddSymbol("print", print)
	slt.AddSymbol("flush", flush)
	slt.AddSymbol("getchar", getchar)

	slt.AddSymbol("int", &Symbol{Name: "int", Type: ttype.NewTigerType(ttype.Int)})
	slt.AddSymbol("string", &Symbol{Name: "string", Type: ttype.NewTigerType(ttype.String)})
	return slt
}

func (s *SymbolTable) CreateRegion() *SymbolTable {
	c := &SymbolTable{
		Table:  make(map[string]*Symbol),
		Region: REGION_CPT + 1,
		Scope:  s.Scope + 1,

		Offset: -4,
		
		Parent: s,
		Children: make([]*SymbolTable, 0),
	}

	REGION_CPT++
	s.Children = append(s.Children, c)

	return c
}

func (s *SymbolTable) CreateSymbol(name string, t *ttype.TigerType) {
	s.Table[name] = &Symbol{
		Name: name,
		Type: t,
		Offset: s.Offset,
	}

	s.Offset -= 4
}

func (s *SymbolTable) AddSymbol(name string, symbol *Symbol) {
	s.Table[name] = symbol
}

func (s *SymbolTable) GetSymbol(name string) (*Symbol, error) {
	if val, ok := s.Table[name]; ok {
		return val, nil
	}

	if s.Parent != nil {
		return s.Parent.GetSymbol(name)
	}

	return nil, fmt.Errorf("Symbol %s not found", name)
}

func (s *SymbolTable) GetSymbolInScoope(name string) (*Symbol, error) {
	if val, ok := s.Table[name]; ok {
		return val, nil
	}

	return nil, fmt.Errorf("Symbol %s not found", name)
}

func (s *SymbolTable) GetSymbolType(name string) (*ttype.TigerType, error) {
	if val, ok := s.Table[name]; ok {
		return val.Type, nil
	}

	if s.Parent != nil {
		return s.Parent.GetSymbolType(name)
	}

	return nil, fmt.Errorf("Symbol %s not found", name)
}

func (s *SymbolTable) Display() {
	indent := ""
	visitedRegions := make(map[int]bool)
	vars := false

	for i := 0; i < s.Scope; i++ {
		indent += "  "
	}

	fmt.Printf("%sScope %d (Region %d)\n", indent, s.Scope, s.Region)

	for k, v := range s.Table {
		if v.Type.ID != ttype.Function {
			vars = true
			if v.Offset == 0 {
				fmt.Printf("%s  %s: %s (type)\n", indent, k, v.Type.String())
			} else {
				fmt.Printf("%s  %s: %s (offset %d)\n", indent, k, v.Type.String(), v.Offset)
			}
		}
	}


	if vars {
		fmt.Println() 
	}

	for k, v := range s.Table {
		if v.Type.ID == ttype.Function {
			fmt.Printf("%sfunc %s: %s\n", indent, k, v.Type.String())
			v.SymbolTable.Display()
			visitedRegions[v.SymbolTable.Region] = true
			fmt.Println()
		}
	}

	for _, child := range s.Children {
		if _, ok := visitedRegions[child.Region]; !ok {
			child.Display()
		}
	}
}
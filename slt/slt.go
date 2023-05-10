package slt

import (
	"fmt"
	"io/ioutil"
	"ogtiger/ttype"
	"os"
	"strings"

	"github.com/emicklei/dot"
)

var REGION_CPT = 0

type SymbolTable struct {
	Table  map[string]*Symbol
	Region int
	Scope  int

	Parent   *SymbolTable
	Children []*SymbolTable

	// Offset to handle the stack
	Offset int
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
		Type: ttype.NewFunctionType(ttype.NewTigerType(ttype.NoReturn), []*ttype.FunctionParameter{
			{
				Name: "x",
				Type: ttype.NewTigerType(ttype.Any),
			},
		}),
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

	print.SymbolTable.CreateSymbol("x", ttype.NewTigerType(ttype.Any))
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

		Parent:   s,
		Children: make([]*SymbolTable, 0),
	}

	REGION_CPT++
	s.Children = append(s.Children, c)

	return c
}

func (s *SymbolTable) CreateSymbol(name string, t *ttype.TigerType) {
	s.Table[name] = &Symbol{
		Name:   name,
		Type:   t,
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

func (s *SymbolTable) Draw(g *dot.Graph, edges *string) dot.Node {
	indent := ""
	visitedRegions := make(map[int]bool)

	for i := 0; i < s.Scope; i++ {
		indent += "  "
	}

	// label should be "Scope 0 (Region 0)|{name|type|offset}|{name|type|offset}"
	label := fmt.Sprintf("{<base> Scope %d (Region %d)}", s.Scope, s.Region)
	label += "|{name|type|offset}"

	for k, v := range s.Table {
		if  v.Type != nil && v.Type.ID != ttype.Function {
			t := strings.ReplaceAll(v.Type.String(), "{", "(")
			t = strings.ReplaceAll(t, "}", ")")
			label += fmt.Sprintf("|{%s|%s|%d}", k, t, v.Offset)
		}
	}

	if len(s.Table) > 0 {
		label += "|"
	}

	var nodeLabel []string
	var nodes []dot.Node

	for k, v := range s.Table {
		if v.Type != nil && v.Type.ID == ttype.Function {
			t := strings.ReplaceAll(v.Type.String(), "{", "(")
			t = strings.ReplaceAll(t, "}", ")")
			label += fmt.Sprintf("|{%s|%s|<%s> 0}", k, t, k)
			to := v.SymbolTable.Draw(g, edges)
			visitedRegions[v.SymbolTable.Region] = true
			nodes = append(nodes, to)
			nodeLabel = append(nodeLabel, k)
		}
	}

	for i, child := range s.Children {
		if _, ok := visitedRegions[child.Region]; !ok {
			label += fmt.Sprintf("|{<slt_%d> Enfant %d}", i, i)
			to := child.Draw(g, edges)
			visitedRegions[child.Region] = true
			nodes = append(nodes, to)
			nodeLabel = append(nodeLabel, fmt.Sprintf("slt_%d", i))
		}
	}

	node := g.Node(fmt.Sprintf("region_%d", s.Region))
	node.Label(label)

	for i, f := range nodes {
		g.EdgeWithPorts(node, f, nodeLabel[i], "base")
	}

	return node
}

func (s *SymbolTable) DisplaySLT(file string) {
	graph := dot.NewGraph(dot.GraphTypeOption{Name: "SLT"})

	graph.Attr("rankdir", "LR")

	var edges string
	s.Draw(graph, &edges)

	// Put the graph in a DOT file.
	path := fmt.Sprintf("output/%s.dot", file)
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	graph.Write(f)

	// Reopen the file to read
	f, err = os.Open(path)
	if err != nil {
		panic(err)
	}

	// Read all
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	// Lines
	lines := strings.Split(string(data), "\n")

	// Replace first line to "digraph SLT {"
	lines[0] = "digraph SLT {"

	// Add a line between 1 and 2
	lines = append(lines[:1], append([]string{"  node [shape=record];"}, lines[1:]...)...)

	// Add edges right before the last line
	lines = append(lines[:len(lines)-2], append([]string{edges}, lines[len(lines)-2:]...)...)

	// Save the file
	ioutil.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
}

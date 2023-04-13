package main

import (
	"fmt"

	"ogtiger/parser"
)

func main() {
	input := []byte{'2', '+', '2'}
	parsed, err := parser.Parse("arithmetic expression", input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+#v\n", parsed)
	}
}

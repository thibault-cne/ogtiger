package options

import "fmt"

type Options struct {
	Ast  string
	File string
	Slt string
	Asm string

	Steps int
}

func Parse(osArgs []string) (*Options, error) {
	var options Options
	options.Steps = 1

	flags, err := parse(osArgs)

	if err != nil {
		return nil, err
	}

	for _, flag := range flags {
		switch flag.Flag {
		case &AST:
			options.Ast = flag.Value
			options.Steps++
		case &File:
			options.File = flag.Value
			options.Steps++
		case &Slt:
			options.Slt = flag.Value
			options.Steps++
		case &ASM:
			options.Asm = flag.Value
			options.Steps++
		}
	}

	if options.File == "" {
		return nil, fmt.Errorf("no file specified")
	}

	return &options, nil
}
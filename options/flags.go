package options

var AST = Flag{
	Short:     'a',
	Long:      "ast",
	Values:     []string{"output.png"},
	TakeValue: Necessary,
}

var File = Flag{
	Short:     'f',
	Long:      "file",
	Values:     []string{"input.tig"},
	TakeValue: Necessary,
}

var Flags = []*Flag{
	&AST, &File,
}
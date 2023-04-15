package options

var AST = Flag{
	Short:     'a',
	Long:      "ast",
	Values:     []string{"output"},
	TakeValue: Necessary,
}

var File = Flag{
	Short:     'f',
	Long:      "file",
	Values:     []string{"input"},
	TakeValue: Necessary,
}

var Slt = Flag{
	Short:     's',
	Long:      "slt",
	Values:     []string{"output"},
	TakeValue: Necessary,
}

var Flags = []*Flag{
	&AST, &File, &Slt,
}
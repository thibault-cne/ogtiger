package options

import "fmt"

// Create a parser to parse flags from the command line
// the flags are
// -f --file : the file to parse
// -a --ast : the file to write the AST to

type TakeValue int
type ShortArg byte
type LongArg string
type Values []string

const (
	Necessary TakeValue = iota
	Optional
	Forbidden
)

type Flag struct {
	Short ShortArg
	Long  LongArg
	Values Values

	TakeValue TakeValue
}

type MatchedFlag struct {
	Flag *Flag
	Value string
}

func parse(osArgs []string) ([]MatchedFlag, error) {
	var flags []MatchedFlag

	for index := 1; index < len(osArgs); index++ {
		arg := osArgs[index]

		if arg[0] == '-' {
			if arg[1] == '-' {
				// Long flag
				longArgName := arg[2:]

				if arg, value := splitOnEqual(longArgName); value != "" {
					// Long flag with value in the same argument
					flag := lookupLong(LongArg(arg))

					if flag == nil {
						return nil, fmt.Errorf("unknown flag %s", arg)
					}

					switch flag.TakeValue {
					case Necessary:
						flags = append(flags, MatchedFlag{flag, value})
					case Optional:
						flags = append(flags, MatchedFlag{flag, value})
					case Forbidden:
						return nil, fmt.Errorf("flag %s does not take a value", arg)
					}
				} else {
					// Long flag without value in the same argument
					flag := lookupLong(LongArg(longArgName))

					if flag == nil {
						return nil, fmt.Errorf("unknown flag %s", longArgName)
					}

					switch flag.TakeValue {
					case Forbidden:
						flags = append(flags, MatchedFlag{flag, ""})
					case Necessary:
						if index+1 < len(osArgs) {
							flags = append(flags, MatchedFlag{flag, osArgs[index+1]})
							index++
						} else {
							return nil, fmt.Errorf("flag %s requires a value", longArgName)
						}
					case Optional:
						if index+1 < len(osArgs) {
							flags = append(flags, MatchedFlag{flag, osArgs[index+1]})
							index++
						} else {
							flags = append(flags, MatchedFlag{flag, ""})
						}
					}
				}
			} else {
				// Short flag
				shortName := arg[1:]

				if arg, value := splitOnEqual(shortName); value != "" {
					// Short flag with value in the same argument
					argWithValue, otherArgs := arg[len(arg)-1], arg[:len(arg)-1]
					
					for _, shortArg := range otherArgs {
						flag := lookupShort(ShortArg(shortArg))

						if flag == nil {
							return nil, fmt.Errorf("unknown flag %c", shortArg)
						}

						switch flag.TakeValue {
						case Necessary:
							return nil, fmt.Errorf("flag %c requires a value", shortArg)
						case Optional:
							flags = append(flags, MatchedFlag{flag, ""})
						case Forbidden:
							flags = append(flags, MatchedFlag{flag, ""})
						}
					}

					flag := lookupShort(ShortArg(argWithValue))

					if flag == nil {
						return nil, fmt.Errorf("unknown flag %c", argWithValue)
					}

					switch flag.TakeValue {
					case Necessary:
						flags = append(flags, MatchedFlag{flag, value})
					case Optional:
						flags = append(flags, MatchedFlag{flag, value})
					case Forbidden:
						return nil, fmt.Errorf("flag %c does not take a value", argWithValue)
					}
				} else {
					b := false

					for i, shortArg := range shortName {
						flag := lookupShort(ShortArg(shortArg))

						if flag == nil {
							return nil, fmt.Errorf("unknown flag %c", shortArg)
						}

						switch flag.TakeValue {
						case Necessary:
							if i < len(shortName) - 1 {
								flags = append(flags, MatchedFlag{flag, string(shortName[i+1:])})
								b = true
							} else if index < len(osArgs) - 1 {
								flags = append(flags, MatchedFlag{flag, osArgs[index + 1]})
								index++
							} else {
								return nil, fmt.Errorf("flag %c requires a value", shortArg)
							}
						case Optional:
							if i < len(shortName) - 1 {
								flags = append(flags, MatchedFlag{flag, string(shortName[i+1:])})
								b = true
							} else if index < len(osArgs) - 1 {
								flags = append(flags, MatchedFlag{flag, osArgs[index+1]})
								index++
							} else {
								flags = append(flags, MatchedFlag{flag, ""})
							}
						case Forbidden:
							flags = append(flags, MatchedFlag{flag, ""})
						}
					}

					if b {
						break
					}
				}
			}
		}
	}
	
	return flags, nil
}

func lookupShort(short ShortArg) *Flag {
	for _, flag := range Flags {
		if flag.Short == short {
			return flag
		}
	}

	return nil
}

func lookupLong(long LongArg) *Flag {
	for _, flag := range Flags {
		if flag.Long == long {
			return flag
		}
	}

	return nil
}

func splitOnEqual(b string) (string, string) {
	for i, c := range b {
		if c == '=' {
			return b[:i], b[i+1:]
		}
	}

	return b, ""
}




package ttype

type TigerType string // Type of a tiger expression

const (
	Integer  TigerType = "integer"
	String   TigerType = "string"
	Record   TigerType = "record"
	Array    TigerType = "array"
	NoReturn TigerType = "noreturn"
)

package lexer

const (
	NUM = iota + 256
	FLOAT
	ID
	TRUE
	FALSE
	EQ  // ==
	NEQ // !=
	LEQ // <=
	GEQ // >=
	INC // ++
	DEC // --
)

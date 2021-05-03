package lexer

type Word interface {
	Token
	GetLexeme() string
}

func NewWord(tag int, lexeme string) Word {
	return &wordImpl{
		tokenImpl: tokenImpl{tag: tag},
		lexeme:    lexeme,
	}
}

type wordImpl struct {
	tokenImpl
	lexeme string
}

func (w *wordImpl) GetLexeme() string {
	return w.lexeme
}

func (w *wordImpl) String() string {
	return w.lexeme
}

var (
	And   = NewWord(AND, "&&")
	Or    = NewWord(OR, "||")
	Eq    = NewWord(EQ, "==")
	Ne    = NewWord(NE, "!=")
	Le    = NewWord(LE, "<=")
	Ge    = NewWord(GE, ">=")
	Minus = NewWord(MINUS, "minus")
	True  = NewWord(TRUE, "true")
	False = NewWord(FALSE, "false")
	Temp  = NewWord(TEMP, "t")
)

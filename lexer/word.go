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

func (n *wordImpl) GetLexeme() string {
	return n.lexeme
}

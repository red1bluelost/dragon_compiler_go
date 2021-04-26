package lexer

type Token interface {
	GetTag() int
}

func NewToken(tag int) Token {
	return &tokenImpl{tag: tag}
}

type tokenImpl struct {
	tag int
}

func (t *tokenImpl) GetTag() int {
	return t.tag
}

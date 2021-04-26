package lexer

type Token interface {
	getTag() int
}

func NewToken(tag int) Token {
	return &tokenImpl{tag: tag}
}

type tokenImpl struct {
	tag int
}

func (t *tokenImpl) getTag() int {
	return t.tag
}

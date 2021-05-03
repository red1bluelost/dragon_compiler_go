package lexer

import "fmt"

type Token interface {
	fmt.Stringer
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

func (t *tokenImpl) String() string {
	return fmt.Sprintf("%c", t.tag)
}

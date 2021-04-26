package lexer

type Num interface {
	Token
	getValue() int
}

func NewNum(tag int, value int) Num {
	return &numImpl{
		tokenImpl: tokenImpl{tag: tag},
		value:     value,
	}
}

type numImpl struct {
	tokenImpl
	value int
}

func (n *numImpl) getValue() int {
	return n.value
}

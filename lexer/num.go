package lexer

type Num interface {
	Token
	GetValue() int
}

func NewNum(value int) Num {
	return &numImpl{
		tokenImpl: tokenImpl{tag: NUM},
		value:     value,
	}
}

type numImpl struct {
	tokenImpl
	value int
}

func (n *numImpl) GetValue() int {
	return n.value
}

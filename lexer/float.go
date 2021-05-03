package lexer

type Float interface {
	Token
	GetValue() float64
}

func NewFloat(value float64) Float {
	return &floatImpl{
		tokenImpl: tokenImpl{tag: REAL},
		value:     value,
	}
}

type floatImpl struct {
	tokenImpl
	value float64
}

func (n *floatImpl) GetValue() float64 {
	return n.value
}

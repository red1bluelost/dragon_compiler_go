package lexer

import "fmt"

type Real interface {
	Token
	GetValue() float64
}

func NewReal(value float64) Real {
	return &realImpl{
		tokenImpl: tokenImpl{tag: REAL},
		value:     value,
	}
}

type realImpl struct {
	tokenImpl
	value float64
}

func (r *realImpl) GetValue() float64 {
	return r.value
}

func (r *realImpl) String() string {
	return fmt.Sprintf("%g", r.value)
}

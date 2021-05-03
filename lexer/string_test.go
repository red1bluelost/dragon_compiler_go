package lexer

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToken_String(t *testing.T) {
	tests := []struct {
		input Token
		want  string
	}{
		{NewToken('+'), "+"},
		{NewToken('-'), "-"},
		{NewToken('!'), "!"},
		{NewToken('.'), "."},
		{NewToken('%'), "%"},
	}
	for _, tt := range tests {
		if got := tt.input.String(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Lexer.Scan() = %v, want %v", got, tt.want)
		}
	}
}

func TestNum_String(t *testing.T) {
	tests := []struct {
		input Token
		want  string
	}{
		{NewNum(3), "3"},
		{NewNum(43), "43"},
		{NewNum(-69), "-69"},
		{NewNum(-42), "-42"},
		{NewNum(100000), "100000"},
	}
	for _, tt := range tests {
		if got := tt.input.String(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Lexer.Scan() = %v, want %v", got, tt.want)
		}
	}
}

func TestReal_String(t *testing.T) {
	tests := []struct {
		input Token
		want  string
	}{
		{NewReal(3.3), fmt.Sprintf("%f", 3.3)},
		{NewReal(4.32), fmt.Sprintf("%f", 4.32)},
		{NewReal(2.913), fmt.Sprintf("%f", 2.913)},
		{NewReal(-420.69), fmt.Sprintf("%f", -420.69)},
		{NewReal(-13.49), fmt.Sprintf("%f", -13.49)},
	}
	for _, tt := range tests {
		if got := tt.input.String(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Lexer.Scan() = %v, want %v", got, tt.want)
		}
	}
}

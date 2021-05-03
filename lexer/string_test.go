package lexer

import (
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
		{NewReal(3.3), "3.3"},
		{NewReal(4.32), "4.32"},
		{NewReal(2.913), "2.913"},
		{NewReal(-420.69), "-420.69"},
		{NewReal(-13.49), "-13.49"},
	}
	for _, tt := range tests {
		if got := tt.input.String(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Lexer.Scan() = %v, want %v", got, tt.want)
		}
	}
}

func TestWord_String(t *testing.T) {
	tests := []struct {
		input Token
		want  string
	}{
		{NewWord(ID, "i"), "i"},
		{NewWord(ID, "retval"), "retval"},
		{NewWord(ID, "foo"), "foo"},
		{And, "&&"},
		{Minus, "minus"},
	}
	for _, tt := range tests {
		if got := tt.input.String(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Lexer.Scan() = %v, want %v", got, tt.want)
		}
	}
}
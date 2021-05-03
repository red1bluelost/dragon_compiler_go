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
		{NewToken('+'),"+"},
		{NewToken('-'),"-"},
		{NewToken('!'),"!"},
		{NewToken('.'),"."},
		{NewToken('%'),"%"},
	}
	for _, tt := range tests {
		if got := tt.input.String(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Lexer.Scan() = %v, want %v", got, tt.want)
		}
	}
}

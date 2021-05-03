package lexer

import (
	"reflect"
	"strings"
	"testing"
)

func TestLexer_Scan(t *testing.T) {
	type test struct {
		want Token
	}
	tester := func(l Lexer, tests []test) {
		for _, tt := range tests {
			if got := l.Scan(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lexer.Scan() = %v, want %v", got, tt.want)
			}
		}
	}
	t.Run("Textbook simple examples", func(t *testing.T) {
		t.Run("Simple input 1", func(t *testing.T) {
			l := NewLexer(strings.NewReader("31 + 28 + 59"))
			tests := []test{
				{NewNum(31)},
				{NewToken('+')},
				{NewNum(28)},
				{NewToken('+')},
				{NewNum(59)},
			}
			tester(l, tests)
		})
		t.Run("Simple input 2", func(t *testing.T) {
			l := NewLexer(strings.NewReader("count = count + increment"))
			tests := []test{
				{NewWord(ID, "count")},
				{NewToken('=')},
				{NewWord(ID, "count")},
				{NewToken('+')},
				{NewWord(ID, "increment")},
			}
			tester(l, tests)
		})
	})
	t.Run("Textbook exercise 2.6.1", func(t *testing.T) {
		t.Run("a) single line comment", func(t *testing.T) {
			s := "31 + 28 // dumb comment\n" +
				" + 59//another for good measure"
			l := NewLexer(strings.NewReader(s))
			tests := []test{
				{NewNum(31)},
				{NewToken('+')},
				{NewNum(28)},
				{NewToken('+')},
				{NewNum(59)},
			}
			tester(l, tests)
		})
		t.Run("b) block comment", func(t *testing.T) {
			s := "count = /* hello */ \n" +
				"count + /* more than\n" +
				"one line*/ increment\n"
			l := NewLexer(strings.NewReader(s))
			tests := []test{
				{NewWord(ID, "count")},
				{NewToken('=')},
				{NewWord(ID, "count")},
				{NewToken('+')},
				{NewWord(ID, "increment")},
			}
			tester(l, tests)
		})
	})
	t.Run("Textbook exercise 2.6.2", func(t *testing.T) {
		t.Run("relational operators", func(t *testing.T) {
			s := "< <= == != >= >"
			l := NewLexer(strings.NewReader(s))
			tests := []test{
				{NewToken('<')},
				{NewWord(LE, "<=")},
				{NewWord(EQ, "==")},
				{NewWord(NE, "!=")},
				{NewWord(GE, ">=")},
				{NewToken('>')},
			}
			tester(l, tests)
		})
	})
	t.Run("Textbook exercise 2.6.3", func(t *testing.T) {
		t.Run("floating point numbers", func(t *testing.T) {
			s := "2. 3.14 .5"
			l := NewLexer(strings.NewReader(s))
			tests := []test{
				{NewFloat(2.)},
				{NewFloat(3.14)},
				{NewFloat(.5)},
			}
			tester(l, tests)
		})
	})
}

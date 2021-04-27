package lexer

import (
	"fmt"
	"io"
	"strconv"
	"unicode"
)

type Lexer interface {
	GetLine() int
	Scan() Token
}

func NewLexer() Lexer {
	l := &lexerImpl{
		line:  1,
		peek:  ' ',
		next:  ' ',
		words: make(map[string]Word),
	}
	l.reserve(NewWord(TRUE, "true"))
	l.reserve(NewWord(FALSE, "false"))
	return l
}

type lexerImpl struct {
	line  int
	peek  byte
	next  byte
	words map[string]Word
}

func (l *lexerImpl) reserve(w Word) {
	l.words[w.GetLexeme()] = w
}

func (l *lexerImpl) GetLine() int {
	return l.line
}

func (l *lexerImpl) Scan() Token {
	if err := l.clearUselessCharacters(); err == io.EOF {
		return nil
	}
	if unicode.IsDigit(rune(l.peek)) {
		return l.handleDigit()
	}
	if unicode.IsLetter(rune(l.peek)) {
		return l.handleWord()
	}
	t := NewToken(int(l.peek))
	l.peek = ' '
	return t
}

// grabNextChar handles pulling in the input characters, no more input returns EOF
func (l *lexerImpl) grabNextChar() (err error) {
	l.peek = l.next
	l.next, err = ReadCharStdio()
	if l.peek == 0 {
		return err
	} else {
		return nil
	}
}

// clearUselessCharacters factors out hte handling of white space and comments
func (l *lexerImpl) clearUselessCharacters() (err error) {
	for {
		switch l.peek {
		case 0:
			return io.EOF
		case '/':
			if l.next == '/' {
				if err = l.handleSingleComment(); err != nil {
					return err
				}
			} else if l.next == '*' {
				if err = l.handleBlockComment(); err != nil {
					return err
				}
			} else {
				return err
			}
		case ' ', '\t', '\r':
		case '\n':
			l.line++
		default:
			return err
		}
		err = l.grabNextChar()
	}
}

// handleSingleComment factors out the process of ignoreing single comments
func (l *lexerImpl) handleSingleComment() error {
	for {
		err := l.grabNextChar()
		if err == io.EOF {
			return err
		} else if l.next == '\n' {
			return nil
		}
	}
}

// handleBlockComment factors out the process of ignoring block comments
func (l *lexerImpl) handleBlockComment() error {
	_ = l.grabNextChar()
	for {
		err := l.grabNextChar()
		if err == io.EOF {
			fmt.Printf("EOF, block comment should be closed.\n")
			return err
		} else if l.peek == '*' && l.next == '/' {
			l.peek, l.next = ' ', ' '
			return nil
		}
	}
}

// handleDigit factors out the handling of a digit in the lexer
func (l *lexerImpl) handleDigit() Token {
	v := 0
	for {
		i, err := strconv.Atoi(string(l.peek))
		if err != nil {
			break
		}
		v = 10*v + i
		_ = l.grabNextChar()
	}
	return NewNum(v)
}

// handleWord factors out the handling of a word in the lexer
func (l *lexerImpl) handleWord() Token {
	buf := make([]byte, 1)
	for {
		buf = append(buf, l.peek)
		if err := l.grabNextChar(); err != nil || !(unicode.IsLetter(rune(l.peek)) || unicode.IsDigit(rune(l.peek))) {
			break
		}
	}
	s := string(buf)
	if w, ok := l.words[s]; ok {
		return w
	} else {
		w = NewWord(ID, s)
		l.reserve(w)
		return w
	}
}

package lexer

import (
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
		words: make(map[string]Word),
	}
	l.reserve(NewWord(TRUE, "true"))
	l.reserve(NewWord(FALSE, "false"))
	return l
}

type lexerImpl struct {
	line  int
	peek  byte
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
	return NewToken(int(l.peek))
}

// clearUselessCharacters factors out hte handling of white space and comments
func (l *lexerImpl) clearUselessCharacters() (err error) {
	for {
		l.peek, err = ReadCharStdio()
		switch l.peek {
		case ' ', '\t', '\r':
			continue
		case '\n':
			l.line++
			continue
		default:
			return err
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
		l.peek, _ = ReadCharStdio()
		if err == io.EOF {
			break
		}
	}
	return NewNum(v)
}

// handleWord factors out the handling of a word in the lexer
func (l *lexerImpl) handleWord() Token {
	buf := make([]byte, 1)
	for {
		buf = append(buf, l.peek)
		var err error
		l.peek, err = ReadCharStdio()
		if err == io.EOF {
			break
		}
		if !(unicode.IsLetter(rune(l.peek)) || unicode.IsDigit(rune(l.peek))) {
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

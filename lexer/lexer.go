package lexer

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"unicode"
)

type Lexer interface {
	GetLine() int
	Scan() Token
}

func NewLexer(reader io.Reader) Lexer {
	l := &lexerImpl{
		line:   1,
		peek:   ' ',
		next:   ' ',
		reader: bufio.NewReader(reader),
		words:  make(map[string]Word),
	}
	l.reserve(NewWord(TRUE, "true"))
	l.reserve(NewWord(FALSE, "false"))
	return l
}

type lexerImpl struct {
	line   int
	peek   byte
	next   byte
	reader *bufio.Reader

	words map[string]Word
}

func (l *lexerImpl) reserve(w Word) {
	l.words[w.GetLexeme()] = w
}

func (l *lexerImpl) GetLine() int {
	return l.line
}

func (l *lexerImpl) Scan() Token {
	defer func() {
		l.peek = ' '
	}()
	if err := l.clearUselessCharacters(); err == io.EOF {
		return nil
	}
	if unicode.IsDigit(rune(l.peek)) {
		return l.handleDigit()
	}
	if unicode.IsLetter(rune(l.peek)) {
		return l.handleWord()
	}
	switch l.peek {
	case '<':
		return l.handleTwoCharToken('=', LE)
	case '>':
		return l.handleTwoCharToken('=', GE)
	case '=':
		return l.handleTwoCharToken('=', EQ)
	case '!':
		return l.handleTwoCharToken('=', NE)
	case '+':
		return l.handleTwoCharToken('+', INC)
	case '-':
		return l.handleTwoCharToken('-', DEC)
	case '.':
		if unicode.IsDigit(rune(l.next)) {
			return l.handleFloat(0)
		}
	}
	return NewToken(int(l.peek))
}

// grabNextChar handles pulling in the input characters, no more input returns EOF
func (l *lexerImpl) grabNextChar() (err error) {
	l.peek = l.next
	l.next, err = l.reader.ReadByte()
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
	if l.peek == '.' {
		return l.handleFloat(v)
	}
	return NewNum(v)
}

// handleFloat factors out the handling of a float in the lexer
func (l *lexerImpl) handleFloat(num int) Token {
	_ = l.grabNextChar() //clear the decimal
	v, d := 0, 1.
	for {
		i, err := strconv.Atoi(string(l.peek))
		if err != nil {
			break
		}
		d *= 0.1
		v = 10*v + i
		_ = l.grabNextChar()
	}
	return NewReal(float64(num) + float64(v)*d)
}

// handleWord factors out the handling of a word in the lexer
func (l *lexerImpl) handleWord() Token {
	buf := make([]byte, 0)
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

func (l *lexerImpl) handleTwoCharToken(second byte, tag int) Token {
	if l.next == second {
		defer func() {
			l.next = ' '
		}()
		return NewWord(tag, string([]byte{l.peek, l.next}))
	}
	return NewToken(int(l.peek))
}

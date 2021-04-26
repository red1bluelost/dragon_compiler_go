package lexer

import (
	"io"
	"os"
)

func ReadCharStdio() (byte, error) {
	b := make([]byte, 1)
	if _, err := os.Stdin.Read(b); err == io.EOF {
		return 0, err
	}
	return b[0], nil
}
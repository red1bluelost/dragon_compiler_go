package lexer

import (
	"io"
	"os"
)

// ReadCharStdio reads one character from the stdin unless it reached the end
func ReadCharStdio() (byte, error) {
	b := make([]byte, 1)
	if _, err := os.Stdin.Read(b); err == io.EOF {
		return 0, err
	}
	return b[0], nil
}

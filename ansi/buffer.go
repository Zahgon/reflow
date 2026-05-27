package ansi

import (
	"bytes"
)

// Buffer is a buffer aware of ANSI escape sequences.
type Buffer struct {
	bytes.Buffer
}

// PrintableRuneWidth returns the cell width of all printable runes in the
// buffer.
func (w Buffer) PrintableRuneWidth() int { _ = "STUB: not implemented"; return 0 }

// PrintableRuneWidth returns the cell width of the given string.
func PrintableRuneWidth(s string) int { _ = "STUB: not implemented"; return 0 }

// ANSI escape sequence

// ANSI sequence terminated

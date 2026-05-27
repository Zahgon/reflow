package indent

import (
	"bytes"
	"io"

	"github.com/muesli/reflow/ansi"
)

type IndentFunc func(w io.Writer)

type Writer struct {
	Indent     uint
	IndentFunc IndentFunc

	ansiWriter *ansi.Writer
	buf        bytes.Buffer
	skipIndent bool
	ansi       bool
}

func NewWriter(indent uint, indentFunc IndentFunc) *Writer { _ = "STUB: not implemented"; return nil }

func NewWriterPipe(forward io.Writer, indent uint, indentFunc IndentFunc) *Writer {
	_ = "STUB: not implemented"
	return nil
}

// Bytes is shorthand for declaring a new default indent-writer instance,
// used to immediately indent a byte slice.
func Bytes(b []byte, indent uint) []byte { _ = "STUB: not implemented"; return nil }

// String is shorthand for declaring a new default indent-writer instance,
// used to immediately indent a string.
func String(s string, indent uint) string { _ = "STUB: not implemented"; return "" }

// Write is used to write content to the indent buffer.
func (w *Writer) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// ANSI escape sequence

// ANSI sequence terminated

// end of current line

// Bytes returns the indented result as a byte slice.
func (w *Writer) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// String returns the indented result as a string.
func (w *Writer) String() string { _ = "STUB: not implemented"; return "" }

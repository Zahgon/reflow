package truncate

import (
	"bytes"
	"io"

	"github.com/muesli/reflow/ansi"
)

type Writer struct {
	width uint
	tail  string

	ansiWriter *ansi.Writer
	buf        bytes.Buffer
	ansi       bool
}

func NewWriter(width uint, tail string) *Writer { _ = "STUB: not implemented"; return nil }

func NewWriterPipe(forward io.Writer, width uint, tail string) *Writer {
	_ = "STUB: not implemented"
	return nil
}

// Bytes is shorthand for declaring a new default truncate-writer instance,
// used to immediately truncate a byte slice.
func Bytes(b []byte, width uint) []byte { _ = "STUB: not implemented"; return nil }

// Bytes is shorthand for declaring a new default truncate-writer instance,
// used to immediately truncate a byte slice. A tail is then added to the
// end of the byte slice.
func BytesWithTail(b []byte, width uint, tail []byte) []byte { _ = "STUB: not implemented"; return nil }

// String is shorthand for declaring a new default truncate-writer instance,
// used to immediately truncate a string.
func String(s string, width uint) string { _ = "STUB: not implemented"; return "" }

// StringWithTail is shorthand for declaring a new default truncate-writer instance,
// used to immediately truncate a string. A tail is then added to the end of the
// string.
func StringWithTail(s string, width uint, tail string) string { _ = "STUB: not implemented"; return "" }

// Write truncates content at the given printable cell width, leaving any
// ansi sequences intact.
func (w *Writer) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// ANSI escape sequence

// ANSI sequence terminated

// Bytes returns the truncated result as a byte slice.
func (w *Writer) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// String returns the truncated result as a string.
func (w *Writer) String() string { _ = "STUB: not implemented"; return "" }

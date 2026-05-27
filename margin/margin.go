package margin

import (
	"bytes"
	"io"

	"github.com/muesli/reflow/indent"
	"github.com/muesli/reflow/padding"
)

type Writer struct {
	buf bytes.Buffer
	pw  *padding.Writer
	iw  *indent.Writer
}

func NewWriter(width uint, margin uint, marginFunc func(io.Writer)) *Writer {
	_ = "STUB: not implemented"
	return nil
}

// Bytes is shorthand for declaring a new default margin-writer instance,
// used to immediately apply a margin to a byte slice.
func Bytes(b []byte, width uint, margin uint) []byte { _ = "STUB: not implemented"; return nil }

// String is shorthand for declaring a new default margin-writer instance,
// used to immediately apply a margin to a string.
func String(s string, width uint, margin uint) string { _ = "STUB: not implemented"; return "" }

func (w *Writer) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Close will finish the margin operation. Always call it before trying to
// retrieve the final result.
func (w *Writer) Close() error { _ = "STUB: not implemented"; return nil }

// Bytes returns the result as a byte slice.
func (w *Writer) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// String returns the result as a string.
func (w *Writer) String() string { _ = "STUB: not implemented"; return "" }

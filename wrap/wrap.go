package wrap

import (
	"bytes"
)

var (
	defaultNewline  = []rune{'\n'}
	defaultTabWidth = 4
)

type Wrap struct {
	Limit         int
	Newline       []rune
	KeepNewlines  bool
	PreserveSpace bool
	TabWidth      int

	buf             *bytes.Buffer
	lineLen         int
	ansi            bool
	forcefulNewline bool
}

// NewWriter returns a new instance of a wrapping writer, initialized with
// default settings.
func NewWriter(limit int) *Wrap { _ = "STUB: not implemented"; return nil }

// Keep whitespaces following a forceful line break. If disabled,
// leading whitespaces in a line are only kept if the line break
// was not forceful, meaning a line break that was already present
// in the input

// Bytes is shorthand for declaring a new default Wrap instance,
// used to immediately wrap a byte slice.
func Bytes(b []byte, limit int) []byte { _ = "STUB: not implemented"; return nil }

func (w *Wrap) addNewLine() { _ = "STUB: not implemented"; return }

// String is shorthand for declaring a new default Wrap instance,
// used to immediately wrap a string.
func String(s string, limit int) string { _ = "STUB: not implemented"; return "" }

func (w *Wrap) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Bytes returns the wrapped result as a byte slice.
func (w *Wrap) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// String returns the wrapped result as a string.
func (w *Wrap) String() string { _ = "STUB: not implemented"; return "" }

func inGroup(a []rune, c rune) bool { _ = "STUB: not implemented"; return false }

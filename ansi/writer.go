package ansi

import (
	"bytes"
	"io"
)

type Writer struct {
	Forward io.Writer

	ansi       bool
	ansiseq    bytes.Buffer
	lastseq    bytes.Buffer
	seqchanged bool
	runeBuf    []byte
}

// Write is used to write content to the ANSI buffer.
func (w *Writer) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// ANSI escape sequence

// ANSI sequence terminated

// reset sequence

// color code

func (w *Writer) writeRune(r rune) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *Writer) LastSequence() string { _ = "STUB: not implemented"; return "" }

func (w *Writer) ResetAnsi() { _ = "STUB: not implemented"; return }

func (w *Writer) RestoreAnsi() { _ = "STUB: not implemented"; return }

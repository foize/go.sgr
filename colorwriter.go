package sgr

import (
	"io"
)

// ForcedNewline is written to the Writer set to ColorWriter when the written bytes didn't end with a newline
var ForcedNewline = []byte(FgWhite + BgRed + `↵ ` + Reset)

// ExistingNewline is written to the Writer set to ColorWriter when the written bytes ended with a newline
var ExistingNewline = []byte(FgWhite + BgGreen + `↵ ` + Reset)

var newlineBytes = []byte("\n")
var resetBytes = []byte(Reset)

// ColorWriter implements io.Writer
type ColorWriter struct {
	w            io.Writer
	colorBytes   []byte
	forceNewline bool
}

// NewColorWriter creates a new ColorWriter.
// example: NewColorWriter(os.Stdout, sgr.FgBlue, true)
func NewColorWriter(w io.Writer, color string, forceNewline bool) *ColorWriter {
	return &ColorWriter{
		w:            w,
		colorBytes:   []byte(color),
		forceNewline: forceNewline,
	}
}

// Write writes the bytes to underlying writer, but adds coloring
func (cw *ColorWriter) Write(p []byte) (n int, err error) {
	// add colors
	cw.w.Write(cw.colorBytes)

	// write tekst/data
	cw.w.Write(p[:len(p)-2])

	// force newline if wanted && required
	if cw.forceNewline {
		if p[len(p)-1] == '\n' {
			cw.w.Write(ExistingNewline)
			cw.w.Write(newlineBytes)
		} else {
			cw.w.Write(p[len(p)-1:])
			cw.w.Write(ForcedNewline)
			cw.w.Write(newlineBytes)
		}
	}
	cw.w.Write(resetBytes)

	//++ TODO: proper return.
	//++ should this return the amount of p bytes that where written? or also the sgr bytes?
	return len(p), nil
}

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
	writer       io.Writer
	colorBytes   []byte
	forceNewline bool
}

// NewColorWriter creates a new ColorWriter.
// example: NewColorWriter(os.Stdout, sgr.FgBlue, true)
func NewColorWriter(w io.Writer, color string, forceNewline bool) *ColorWriter {
	return &ColorWriter{
		writer:       w,
		colorBytes:   []byte(color),
		forceNewline: forceNewline,
	}
}

// Write writes the bytes to underlying writer, but adds coloring
func (cw *ColorWriter) Write(p []byte) (n int, err error) {
	var nAdd int

	// add colors
	_, err = cw.writer.Write(cw.colorBytes)
	if err != nil {
		return
	}

	// write tekst/data
	nAdd, err = cw.writer.Write(p[:len(p)-1])
	if err != nil {
		return
	}
	n += nAdd

	// force newline if wanted && required
	if cw.forceNewline {
		if p[len(p)-1] == '\n' {
			_, err = cw.writer.Write(ExistingNewline)
			if err != nil {
				return
			}

			nAdd, err = cw.writer.Write(p[len(p)-1:])
			if err != nil {
				return
			}
			n += nAdd
		} else {
			nAdd, err = cw.writer.Write(p[len(p)-1:])
			if err != nil {
				return
			}
			n += nAdd

			_, err = cw.writer.Write(ForcedNewline)
			if err != nil {
				return
			}

			_, err = cw.writer.Write(newlineBytes)
			if err != nil {
				return
			}
		}
	}
	_, err = cw.writer.Write(resetBytes)
	if err != nil {
		return
	}

	return
}

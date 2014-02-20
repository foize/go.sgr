package sgr

import (
	"bytes"
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

	var nExtra int
	var nExtraAdd int

	buf := bytes.NewBuffer(make([]byte, 0, len(p)+40))
	defer func() {
		if err != nil {
			return
		}
		n, err = cw.writer.Write(buf.Bytes())
		n = n - nExtra
		if n < 0 {
			n = 0
		}
	}()

	// add colors
	nExtraAdd, err = buf.Write(cw.colorBytes)
	if err != nil {
		return
	}
	nExtra += nExtraAdd

	// force newline if wanted && required
	if !cw.forceNewline {
		// write all bytes
		_, err = buf.Write(p)
		if err != nil {
			return
		}
	} else {
		// write tekst/data except last byte
		_, err = buf.Write(p[:len(p)-1])
		if err != nil {
			return
		}

		// check last byte
		if p[len(p)-1] == '\n' {
			nExtraAdd, err = buf.Write(ExistingNewline)
			if err != nil {
				return
			}
			nExtra += nExtraAdd

			_, err = buf.Write(p[len(p)-1:])
			if err != nil {
				return
			}
		} else {
			_, err = buf.Write(p[len(p)-1:])
			if err != nil {
				return
			}

			nExtraAdd, err = buf.Write(ForcedNewline)
			if err != nil {
				return
			}
			nExtra += nExtraAdd

			nExtraAdd, err = buf.Write(newlineBytes)
			if err != nil {
				return
			}
			nExtra += nExtraAdd
		}
	}
	nExtraAdd, err = buf.Write(resetBytes)
	if err != nil {
		return
	}
	nExtra += nExtraAdd

	return
}

package tcolor

import (
	"bytes"
	"fmt"
)

func Format(parts ...interface{}) (string, error) {
	return format(true, parts)
}

func FormatWithoutReset(parts ...interface{}) (string, error) {
	return format(false, parts)
}

func format(reset bool, parts ...interface{}) (string, error) {
	var b bytes.Buffer

	for i, part := range parts {
		// Just a simple piece of string to add to the result
		str, ok := part.(string)
		if ok {
			b.WriteString(str)
			continue
		}

		// Add a SGR definition
		sgr, ok := part.(Sgr)
		if ok {
			b.Write(CSI)
			b.Write(sgr.Sequence())
			b.WriteByte(SgrEnd)
			continue
		}

		return "", fmt.Errorf("Invalid part at position %d.", i)
	}

	if reset {
		b.Write(Reset.Sequence())
	}
	return b.String(), nil
}

package tcolor

import (
	"bytes"
	"fmt"
)

func Format(parts ...interface{}) (string, error) {
	return format(true, parts...)
}

func FormatWithoutReset(parts ...interface{}) (string, error) {
	return format(false, parts...)
}

func format(reset bool, parts ...interface{}) (string, error) {
	var b bytes.Buffer

	for i, part := range parts {
		// Just a simple piece of string to add to the result
		str, strOk := part.(string)
		if strOk {
			b.WriteString(str)
			continue
		}

		// Add a SGR definition
		sgr, sgrOk := part.(Sgr)
		if sgrOk {
			b.Write(CSI)
			b.Write(sgr.Sequence())
			b.WriteByte(SgrEnd)
			continue
		}

		return "", fmt.Errorf("Invalid part at position %d. Expecting type `string` or `tcolor.Sgr`.", i)
	}

	if reset {
		b.Write(CSI)
		b.Write(Reset.Sequence())
		b.WriteByte(SgrEnd)
	}
	return b.String(), nil
}

package tcolor

import (
	"bytes"
	"fmt"
)

// Same as Format, but panics instead of returning an error.
func MustFormat(parts ...interface{}) string {
	str, err := format(true, parts...)
	if err != nil {
		panic(err)
	}
	return str
}

// Same as FormatWithoutReset, but panics instead of returning an error.
func MustFormatWithoutReset(parts ...interface{}) string {
	str, err := format(false, parts...)
	if err != nil {
		panic(err)
	}
	return str
}

// Formats a string 
func Format(parts ...interface{}) (string, error) {
	return format(true, parts...)
}

func FormatWithoutReset(parts ...interface{}) (string, error) {
	return format(false, parts...)
}

func format(reset bool, parts ...interface{}) (string, error) {
	// Buffer used to build the colored string.
	var b bytes.Buffer

	// Loop through all the given parts
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
			b.Write(sgr.sequence())
			b.WriteByte(SgrEnd)
			continue
		}

		// Invalid type for this part.
		return "", fmt.Errorf("Invalid part at position %d. Expecting type `string` or `tcolor.Sgr`.", i)
	}

	// Reset sequence if needed.
	if reset {
		b.Write(CSI)
		b.Write(Reset.sequence())
		b.WriteByte(SgrEnd)
	}

	// All done, return resulting string
	return b.String(), nil
}
